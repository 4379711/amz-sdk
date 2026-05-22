package auth

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/4379711/amz-sdk/pkg"
)

// mockRoundTripper 按调用顺序返回预设响应,并记录每次被调用时的 Authorization header,
// 用于端到端验证 RoundTrip 的流程(LwA 刷新 + 业务请求 + 兜底重试)。
type mockRoundTripper struct {
	mu    sync.Mutex
	calls []recordedCall
	// handler 允许根据 URL.Path 决定返回内容。
	handler func(reqNum int, req *http.Request) *http.Response
}

type recordedCall struct {
	Path          string
	Authorization string
	AccessToken   string // x-amz-access-token,SP 路径用
	IsLwA         bool
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	m.calls = append(m.calls, recordedCall{
		Path:          req.URL.Path,
		Authorization: req.Header.Get("Authorization"),
		AccessToken:   req.Header.Get("x-amz-access-token"),
		IsLwA:         strings.Contains(req.URL.Path, "/auth/o2/token"),
	})
	n := len(m.calls)
	m.mu.Unlock()
	return m.handler(n, req), nil
}

func (m *mockRoundTripper) CallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.calls)
}

func (m *mockRoundTripper) LwACount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	c := 0
	for _, r := range m.calls {
		if r.IsLwA {
			c++
		}
	}
	return c
}

// captureHooks 安装可记录的 hook,并返回已记录事件的快照访问器以及恢复函数。
func captureHooks(t *testing.T) (getRefreshes func() []pkg.TokenRefreshEvent, getRetries func() []pkg.AuthRetryEvent) {
	t.Helper()
	var mu sync.Mutex
	var refreshes []pkg.TokenRefreshEvent
	var retries []pkg.AuthRetryEvent

	oldRefresh := pkg.OnTokenRefresh
	oldRetry := pkg.OnAuthRetry
	pkg.OnTokenRefresh = func(e pkg.TokenRefreshEvent) {
		mu.Lock()
		refreshes = append(refreshes, e)
		mu.Unlock()
	}
	pkg.OnAuthRetry = func(e pkg.AuthRetryEvent) {
		mu.Lock()
		retries = append(retries, e)
		mu.Unlock()
	}
	t.Cleanup(func() {
		pkg.OnTokenRefresh = oldRefresh
		pkg.OnAuthRetry = oldRetry
	})

	getRefreshes = func() []pkg.TokenRefreshEvent {
		mu.Lock()
		defer mu.Unlock()
		cp := make([]pkg.TokenRefreshEvent, len(refreshes))
		copy(cp, refreshes)
		return cp
	}
	getRetries = func() []pkg.AuthRetryEvent {
		mu.Lock()
		defer mu.Unlock()
		cp := make([]pkg.AuthRetryEvent, len(retries))
		copy(cp, retries)
		return cp
	}
	return getRefreshes, getRetries
}

// stubResp 构造一个可重放 body 的 HTTP Response。
func stubResp(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     http.Header{},
	}
}

// installMockTransport 把 pkg.DefaultClient.Transport 替换为 rt,以拦截 LwA 刷新请求。
// 业务请求走 headerInjector.rt,也可以用同一个 mock 实例。
func installMockTransport(t *testing.T, rt http.RoundTripper) {
	t.Helper()
	oldClientT := pkg.DefaultClient.Transport
	pkg.DefaultClient.Transport = rt
	t.Cleanup(func() { pkg.DefaultClient.Transport = oldClientT })
}

// newTestAuth 生成一个 refresh_token 足够长且唯一的 *AdAuth,
// 以保证和前面的测试不会共享 cache(cache key 就是 refresh_token)。
func newTestAuth(rt string) *AdAuth {
	return &AdAuth{
		App:    &App{ClientID: "cid", ClientSecret: "csec"},
		Seller: &Seller{CountryCode: "US", ProfileId: "ENTITY123"},
		Token:  &Token{RefreshToken: rt},
	}
}

// TestRoundTrip_HappyPath_EmitsStartEnd
// 场景:cache 空 -> 业务请求触发 cache-miss -> LwA 成功 -> 业务请求返回 200。
// 期望:hook 触发 "cache-miss" 的 start+end 事件,不触发 auth retry。
func TestRoundTrip_HappyPath_EmitsStartEnd(t *testing.T) {
	getRefreshes, getRetries := captureHooks(t)

	rt := &mockRoundTripper{
		handler: func(n int, req *http.Request) *http.Response {
			if strings.Contains(req.URL.Path, "/auth/o2/token") {
				// LwA 返回一枚有效 token,过期时间 1h 后
				return stubResp(200, `{"access_token":"ACCESS_HAPPY","token_type":"bearer","expires_in":3600,"refresh_token":"RT_HAPPY_IGNORED"}`)
			}
			return stubResp(200, `{"ok":true}`)
		},
	}
	installMockTransport(t, rt)

	auth := newTestAuth("rt_happy_path_unique_long_suffix_for_test_001")
	h := &headerInjector{auth: auth, rt: rt}

	req, _ := http.NewRequest("GET", "https://advertising-api.amazon.com/v2/profiles", nil)
	resp, err := h.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip err: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("RoundTrip status = %d, want 200", resp.StatusCode)
	}

	rs := getRefreshes()
	if len(rs) != 2 {
		t.Fatalf("want 2 refresh events (start+end), got %d: %+v", len(rs), rs)
	}
	if rs[0].Phase != pkg.EventPhaseStart || rs[0].Reason != "cache-miss" {
		t.Fatalf("event[0] want phase=start reason=cache-miss, got %+v", rs[0])
	}
	if rs[1].Phase != pkg.EventPhaseEnd || !rs[1].Success {
		t.Fatalf("event[1] want phase=end success=true, got %+v", rs[1])
	}
	if rs[1].AccessTokenSuffix != "...SS_HAPPY" {
		t.Fatalf("event[1].AccessTokenSuffix = %q, want \"...SS_HAPPY\"", rs[1].AccessTokenSuffix)
	}
	if rs[1].Service != "advertising" {
		t.Fatalf("event[1].Service = %q, want advertising", rs[1].Service)
	}
	if rs[1].ExpiresAt.Before(time.Now().Add(55 * time.Minute)) {
		t.Fatalf("event[1].ExpiresAt should be ~1h ahead, got %v", rs[1].ExpiresAt)
	}

	if got := getRetries(); len(got) != 0 {
		t.Fatalf("want 0 retry events, got %d", len(got))
	}

	if got := rt.LwACount(); got != 1 {
		t.Fatalf("LwA should be called once, got %d", got)
	}
}

// TestRoundTrip_401_TriggersRetryAndNewRefresh
// 场景:首轮业务请求拿 T1 被判 401 "Invalid token" -> SDK 清缓存 -> 再次 LwA 刷新得 T2
// -> 业务请求重试返回 200。
// 期望:
//   - refresh 事件成对出现两次(第 1 次 reason=cache-miss,第 2 次 reason=401-retry)
//   - retry 事件 1 次,FirstStatus=401,FailedAccessTokenSuffix=T1 的后 8 位,RetryStatus=200
func TestRoundTrip_401_TriggersRetryAndNewRefresh(t *testing.T) {
	getRefreshes, getRetries := captureHooks(t)

	// 按顺序返回: LwA1(T1) -> data(401) -> LwA2(T2) -> data(200)
	rt := &mockRoundTripper{
		handler: func(n int, req *http.Request) *http.Response {
			switch {
			case strings.Contains(req.URL.Path, "/auth/o2/token"):
				if req.URL.Host == "" {
					// 第一次还是第二次刷新,按 LwA 调用计数区分(header 不变,只能按次序)
				}
				// 发两次 LwA 就返回两个不同的 token
				if rt := req.Context(); rt != nil {
					_ = rt
				}
				// 通过调用次数判断是哪次
				return stubResp(200, `{"access_token":"ACCESS_T1","token_type":"bearer","expires_in":3600,"refresh_token":"ignored"}`)
			default:
				return stubResp(401, `{"message":"Unauthorized exception while handling 3P Request: Unauthorized exception while authenticating LwA token: Invalid token"}`)
			}
		},
	}
	// 用更精细的 handler:按调用次数决定响应
	rt.handler = func(n int, req *http.Request) *http.Response {
		isLwA := strings.Contains(req.URL.Path, "/auth/o2/token")
		// 按总调用次数区分:
		//   1 LwA -> T1
		//   2 data -> 401
		//   3 LwA -> T2
		//   4 data -> 200
		switch n {
		case 1:
			if !isLwA {
				t.Fatalf("call %d expected LwA, got %s", n, req.URL.Path)
			}
			return stubResp(200, `{"access_token":"ACCESS_T1","token_type":"bearer","expires_in":3600,"refresh_token":"ignored"}`)
		case 2:
			if isLwA {
				t.Fatalf("call %d expected data, got LwA", n)
			}
			return stubResp(401, `{"message":"Unauthorized exception while handling 3P Request: Unauthorized exception while authenticating LwA token: Invalid token"}`)
		case 3:
			if !isLwA {
				t.Fatalf("call %d expected LwA, got %s", n, req.URL.Path)
			}
			return stubResp(200, `{"access_token":"ACCESS_T2","token_type":"bearer","expires_in":3600,"refresh_token":"ignored"}`)
		case 4:
			if isLwA {
				t.Fatalf("call %d expected data, got LwA", n)
			}
			return stubResp(200, `{"ok":true}`)
		default:
			t.Fatalf("unexpected call %d path=%s", n, req.URL.Path)
			return nil
		}
	}
	installMockTransport(t, rt)

	auth := newTestAuth("rt_401_retry_unique_long_suffix_for_test_002")
	h := &headerInjector{auth: auth, rt: rt}

	req, _ := http.NewRequest("GET", "https://advertising-api.amazon.com/sp/campaigns/list", nil)
	resp, err := h.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip err: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("final status = %d, want 200", resp.StatusCode)
	}

	rs := getRefreshes()
	if len(rs) != 4 {
		t.Fatalf("want 4 refresh events (2 start+end pairs), got %d:\n%+v", len(rs), rs)
	}
	// 第一组(cache-miss,得 T1)
	if rs[0].Phase != pkg.EventPhaseStart || rs[0].Reason != "cache-miss" {
		t.Fatalf("event[0] bad: %+v", rs[0])
	}
	if rs[1].Phase != pkg.EventPhaseEnd || !rs[1].Success || rs[1].AccessTokenSuffix != "...CCESS_T1" {
		t.Fatalf("event[1] bad: %+v", rs[1])
	}
	// 第二组(401-retry,得 T2)
	if rs[2].Phase != pkg.EventPhaseStart || rs[2].Reason != "401-retry" {
		t.Fatalf("event[2] bad: %+v", rs[2])
	}
	if rs[3].Phase != pkg.EventPhaseEnd || !rs[3].Success || rs[3].AccessTokenSuffix != "...CCESS_T2" {
		t.Fatalf("event[3] bad: %+v", rs[3])
	}

	retries := getRetries()
	if len(retries) != 1 {
		t.Fatalf("want 1 auth retry event, got %d", len(retries))
	}
	ev := retries[0]
	if ev.Service != "advertising" {
		t.Fatalf("retry.Service = %q", ev.Service)
	}
	if ev.FirstStatus != 401 {
		t.Fatalf("retry.FirstStatus = %d, want 401", ev.FirstStatus)
	}
	if !strings.Contains(ev.FirstBodyHead, "Invalid token") {
		t.Fatalf("retry.FirstBodyHead = %q, want contains \"Invalid token\"", ev.FirstBodyHead)
	}
	if ev.FailedAccessTokenSuffix != "...CCESS_T1" {
		t.Fatalf("retry.FailedAccessTokenSuffix = %q, want \"...CCESS_T1\"", ev.FailedAccessTokenSuffix)
	}
	if ev.RetryStatus != 200 {
		t.Fatalf("retry.RetryStatus = %d, want 200", ev.RetryStatus)
	}
	if ev.RetrySkipped != "" || ev.RetryErr != nil {
		t.Fatalf("retry should not be skipped and no err, got skipped=%q err=%v", ev.RetrySkipped, ev.RetryErr)
	}

	if got := rt.LwACount(); got != 2 {
		t.Fatalf("LwA should be called twice (initial + 401 retry), got %d", got)
	}
}
