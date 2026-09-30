package auth

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/4379711/amz-sdk/pkg"
)

var _ pkg.IAuth = (*AdAuth)(nil)

type contextTestTransport func(*http.Request) (*http.Response, error)

func (f contextTestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func contextTestAuth(t *testing.T) *AdAuth {
	t.Helper()
	a := &AdAuth{
		App:    &App{ClientID: "test-client", ClientSecret: "test-secret"},
		Seller: &Seller{CountryCode: "US"},
		Token:  &Token{RefreshToken: t.Name()},
	}
	t.Cleanup(func() { cache.Invalidate(a.RefreshToken) })
	return a
}

func installContextTokenTransport(t *testing.T, rt http.RoundTripper) {
	t.Helper()
	previous := pkg.DefaultClient
	pkg.DefaultClient = &http.Client{Transport: rt}
	t.Cleanup(func() { pkg.DefaultClient = previous })
}

func receiveContextResult[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case result := <-ch:
		return result
	case <-time.After(3 * time.Second):
		t.Fatal("等待认证操作完成超时")
		var zero T
		return zero
	}
}

type observingContext struct {
	context.Context
	waiting chan struct{}
	once    sync.Once
}

func (c *observingContext) Done() <-chan struct{} {
	c.once.Do(func() { close(c.waiting) })
	return c.Context.Done()
}

func TestAcquireAccessToken_CancelDoesNotCancelSharedRefresh(t *testing.T) {
	a := contextTestAuth(t)
	started := make(chan context.Context, 1)
	release := make(chan struct{})
	var releaseOnce sync.Once
	var tokenCalls atomic.Int32
	installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
		defer req.Body.Close()
		tokenCalls.Add(1)
		started <- req.Context()
		select {
		case <-release:
			return newResp(200, `{"access_token":"shared-token","expires_in":3600}`), nil
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}))
	t.Cleanup(func() {
		releaseOnce.Do(func() { close(release) })
		refreshFlight.Do(a.RefreshToken, func() (any, error) { return "", nil })
	})
	h := &headerInjector{auth: a}
	leaderCtx, cancelLeader := context.WithCancel(context.Background())
	defer cancelLeader()
	leader := make(chan error, 1)
	go func() {
		_, err := h.acquireAccessToken(leaderCtx, "cache-miss")
		leader <- err
	}()
	refreshCtx := receiveContextResult(t, started)
	deadline, ok := refreshCtx.Deadline()
	if !ok || time.Until(deadline) < 175*time.Second || time.Until(deadline) > 180*time.Second {
		t.Fatalf("共享刷新应有独立的 180 秒期限: %v, %v", deadline, ok)
	}
	followerCtx := &observingContext{Context: context.Background(), waiting: make(chan struct{})}
	follower := make(chan error, 1)
	go func() {
		token, err := h.acquireAccessToken(followerCtx, "cache-miss")
		if err == nil && token != "shared-token" {
			err = errors.New("等待者未收到共享刷新结果")
		}
		follower <- err
	}()
	receiveContextResult(t, followerCtx.waiting)
	cancelLeader()
	if err := receiveContextResult(t, leader); !errors.Is(err, context.Canceled) {
		t.Fatalf("首个等待者取消错误 = %v", err)
	}
	if err := refreshCtx.Err(); err != nil {
		t.Fatalf("首个等待者取消影响了共享刷新: %v", err)
	}
	releaseOnce.Do(func() { close(release) })
	if err := receiveContextResult(t, follower); err != nil {
		t.Fatal(err)
	}
	if token, err := h.acquireAccessToken(context.Background(), "cache-miss"); err != nil || token != "shared-token" {
		t.Fatalf("共享结果未缓存: token=%q err=%v", token, err)
	}
	if got := tokenCalls.Load(); got != 1 {
		t.Fatalf("刷新次数 = %d, want 1", got)
	}
}

func TestAcquireAccessToken_CanceledContextDoesNotRefresh(t *testing.T) {
	for _, cached := range []bool{false, true} {
		name := "cache-miss"
		if cached {
			name = "cache-hit"
		}
		t.Run(name, func(t *testing.T) {
			a := contextTestAuth(t)
			if cached {
				cache.Put(a.RefreshToken, pkg.CacheItem{AccessToken: "cached-token", AccessTokenExpiredTime: time.Now().Add(time.Hour)})
			}
			var calls atomic.Int32
			installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
				calls.Add(1)
				return nil, errors.New("不应发起刷新")
			}))
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := (&headerInjector{auth: a}).acquireAccessToken(ctx, "cache-miss")
			if !errors.Is(err, context.Canceled) || calls.Load() != 0 {
				t.Fatalf("已取消请求仍获取令牌: err=%v calls=%d", err, calls.Load())
			}
		})
	}
}

type trackedAuthBody struct {
	io.Reader
	closed atomic.Int32
}

func (b *trackedAuthBody) Close() error {
	b.closed.Add(1)
	return nil
}

func TestRoundTrip_TokenFailureClosesUnsentBody(t *testing.T) {
	for _, retry := range []bool{false, true} {
		phase := "initial"
		if retry {
			phase = "retry"
		}
		for _, failure := range []string{"refresh-error", "canceled", "deadline"} {
			t.Run(phase+"/"+failure, func(t *testing.T) {
				a := contextTestAuth(t)
				if retry {
					cache.Put(a.RefreshToken, pkg.CacheItem{AccessToken: "expired-token", AccessTokenExpiredTime: time.Now().Add(time.Hour)})
				}
				started := make(chan context.Context, 1)
				release := make(chan struct{})
				var releaseOnce sync.Once
				refreshErr := errors.New("模拟刷新失败")
				installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
					defer req.Body.Close()
					started <- req.Context()
					if failure == "refresh-error" {
						return nil, refreshErr
					}
					select {
					case <-release:
						return newResp(200, `{"access_token":"after-cancel","expires_in":3600}`), nil
					case <-req.Context().Done():
						return nil, req.Context().Err()
					}
				}))
				t.Cleanup(func() {
					releaseOnce.Do(func() { close(release) })
					refreshFlight.Do(a.RefreshToken, func() (any, error) { return "", nil })
				})
				ctx, cancel := context.WithCancel(context.Background())
				if failure == "deadline" {
					cancel()
					ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
				}
				defer cancel()
				body := &trackedAuthBody{Reader: strings.NewReader("payload")}
				replay := &trackedAuthBody{Reader: strings.NewReader("payload")}
				req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.example.test/resource", body)
				if err != nil {
					t.Fatal(err)
				}
				req.GetBody = func() (io.ReadCloser, error) { return replay, nil }
				var apiCalls atomic.Int32
				h := &headerInjector{auth: a, rt: contextTestTransport(func(req *http.Request) (*http.Response, error) {
					apiCalls.Add(1)
					req.Body.Close()
					return newResp(401, `{"message":"Invalid token"}`), nil
				})}
				finished := make(chan error, 1)
				go func() { _, err := h.RoundTrip(req); finished <- err }()
				refreshCtx := receiveContextResult(t, started)
				wantErr := refreshErr
				if failure == "canceled" {
					cancel()
					wantErr = context.Canceled
				}
				if failure == "deadline" {
					wantErr = context.DeadlineExceeded
				}
				if err := receiveContextResult(t, finished); !errors.Is(err, wantErr) {
					t.Fatalf("请求错误 = %v, want %v", err, wantErr)
				}
				if body.closed.Load() != 1 {
					t.Fatalf("原请求体关闭次数 = %d", body.closed.Load())
				}
				wantReplay, wantCalls := int32(0), int32(0)
				if retry {
					wantReplay, wantCalls = 1, 1
				}
				if replay.closed.Load() != wantReplay || apiCalls.Load() != wantCalls {
					t.Fatalf("重放体关闭次数=%d, API调用次数=%d", replay.closed.Load(), apiCalls.Load())
				}
				if failure != "refresh-error" && refreshCtx.Err() != nil {
					t.Fatalf("业务取消影响了共享刷新: %v", refreshCtx.Err())
				}
			})
		}
	}
}

func TestFetchAccessToken_UsesSuppliedContext(t *testing.T) {
	a := contextTestAuth(t)
	installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
		defer req.Body.Close()
		<-req.Context().Done()
		return nil, req.Context().Err()
	}))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := a.fetchAccessToken(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("刷新未保留context错误: %v", err)
	}
}

func TestBuildClient_TimeoutLeavesRefreshAvailableForLaterRequests(t *testing.T) {
	a := contextTestAuth(t)
	started := make(chan context.Context, 1)
	release := make(chan struct{})
	var releaseOnce sync.Once
	var tokenCalls atomic.Int32
	installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
		defer req.Body.Close()
		tokenCalls.Add(1)
		started <- req.Context()
		select {
		case <-release:
			return newResp(200, `{"access_token":"later-token","expires_in":3600}`), nil
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}))
	t.Cleanup(func() {
		releaseOnce.Do(func() { close(release) })
		refreshFlight.Do(a.RefreshToken, func() (any, error) { return "", nil })
	})
	previousTransport := pkg.SharedTransport
	pkg.SharedTransport = &http.Transport{}
	var apiCalls atomic.Int32
	pkg.SharedTransport.RegisterProtocol("https", contextTestTransport(func(req *http.Request) (*http.Response, error) {
		apiCalls.Add(1)
		return newResp(200, "ok"), nil
	}))
	t.Cleanup(func() { pkg.SharedTransport = previousTransport })
	client := a.BuildClient()
	client.Timeout = 100 * time.Millisecond
	finished := make(chan error, 1)
	go func() {
		response, err := client.Get("https://api.example.test/resource")
		if response != nil {
			response.Body.Close()
		}
		finished <- err
	}()
	refreshCtx := receiveContextResult(t, started)
	if err := receiveContextResult(t, finished); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Client.Timeout未及时结束等待: %v", err)
	}
	if refreshCtx.Err() != nil || apiCalls.Load() != 0 {
		t.Fatalf("超时后刷新context=%v, API调用次数=%d", refreshCtx.Err(), apiCalls.Load())
	}
	releaseOnce.Do(func() { close(release) })
	later := a.BuildClient()
	response, err := later.Get("https://api.example.test/resource")
	if err != nil {
		t.Fatal(err)
	}
	response.Body.Close()
	if response.StatusCode != 200 || tokenCalls.Load() != 1 || apiCalls.Load() != 1 {
		t.Fatalf("后续请求未复用刷新结果: status=%d refreshes=%d calls=%d", response.StatusCode, tokenCalls.Load(), apiCalls.Load())
	}
}
