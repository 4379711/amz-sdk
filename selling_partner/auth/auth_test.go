package auth

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

func newResp(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// TestShouldRetryAuthExpired_SPAPI 覆盖 SP API 侧 access_token 失效的多种 body。
//
// SP API 的报表下载等接口实际走的是 advertising-api.amazon.com,body 会和
// 广告 API 一致为 "Invalid token",因此这里必须能覆盖广告 API 的措辞。
func TestShouldRetryAuthExpired_SPAPI(t *testing.T) {
	cases := []struct {
		name   string
		status int
		body   string
		want   bool
	}{
		{
			name:   "sp api 经典的 expired 响应",
			status: 403,
			body:   `{"errors":[{"code":"Unauthorized","message":"The access token you provided has expired."}]}`,
			want:   true,
		},
		{
			name:   "sp api 报表下载跨走到 ad api 的 LwA 失效响应",
			status: 401,
			body:   `{"message":"Unauthorized exception while handling 3P Request: Unauthorized exception while authenticating LwA token: Invalid token"}`,
			want:   true,
		},
		{
			name:   "简化形式 Unauthorized",
			status: 401,
			body:   `{"message":"Unauthorized"}`,
			want:   true,
		},
		{
			name:   "2xx 响应不重试",
			status: 200,
			body:   `{"ok":true}`,
			want:   false,
		},
		{
			name:   "真实的 403 权限不足不重试",
			status: 403,
			body:   `{"errors":[{"code":"Unauthorized","message":"Access to requested resource is denied."}]}`,
			want:   false,
		},
		{
			name:   "5xx 服务端错误不重试",
			status: 500,
			body:   `{"errors":[{"code":"InternalFailure"}]}`,
			want:   false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resp := newResp(c.status, c.body)
			got, body := shouldRetryAuthExpired(resp)
			if got != c.want {
				t.Fatalf("status=%d body=%s: got %v, want %v", c.status, c.body, got, c.want)
			}
			if c.status == 401 || c.status == 403 {
				if string(body) != c.body {
					t.Fatalf("returned body mismatch: got %q, want %q", string(body), c.body)
				}
			}
			read, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("resp.Body 已被消费,无法再读: %v", err)
			}
			if !bytes.Equal(read, []byte(c.body)) {
				t.Fatalf("resp.Body 未正确重建: got %q, want %q", string(read), c.body)
			}
		})
	}
}
