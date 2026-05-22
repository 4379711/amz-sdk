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

// TestShouldRetryAuthExpired_AdAPI 覆盖 Ad API 实际观测到的几类 body。
//
// 历史 bug:老实现要求 body 同时包含 "access token" 和 "expired",
// 但 Ad API 的典型响应是 "...Invalid token" / 纯 "Unauthorized",都不带
// "expired",导致 SDK 的兜底重试对 Ad API 完全失效,业务层大量 401。
func TestShouldRetryAuthExpired_AdAPI(t *testing.T) {
	cases := []struct {
		name   string
		status int
		body   string
		want   bool
	}{
		{
			name:   "ad api 最常见的 LwA 失效响应",
			status: 401,
			body:   `{"message":"Unauthorized exception while handling 3P Request: Unauthorized exception while authenticating LwA token: Invalid token"}`,
			want:   true,
		},
		{
			name:   "ad api 简化形式的 Unauthorized",
			status: 401,
			body:   `{"message":"Unauthorized"}`,
			want:   true,
		},
		{
			name:   "sp api 沿用的 expired 措辞",
			status: 403,
			body:   `{"message":"The access token you provided has expired."}`,
			want:   true,
		},
		{
			name:   "OAuth2 标准错误码 invalid_token",
			status: 401,
			body:   `{"error":"invalid_token","error_description":"The access token is invalid"}`,
			want:   true,
		},
		{
			name:   "2xx 响应不重试",
			status: 200,
			body:   `{"ok":true}`,
			want:   false,
		},
		{
			name:   "4xx 但不是认证相关(比如 profile 权限不足)",
			status: 403,
			body:   `{"message":"Access to this profile is not allowed"}`,
			want:   false,
		},
		{
			name:   "5xx 服务端错误不重试",
			status: 500,
			body:   `{"message":"Internal server error"}`,
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
			// body 参数始终是首次读完的字节切片(非 4xx 认证相关路径为 nil)
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
