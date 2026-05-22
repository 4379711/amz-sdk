package pkg

import "testing"

func TestTokenSuffix(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"exactly 8", "12345678", "...12345678"},
		{"longer than 8", "abcdefghijklmn", "...ghijklmn"},
		{"empty", "", "<short>"},
		{"7 chars still short", "1234567", "<short>"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := TokenSuffix(c.in)
			if got != c.want {
				t.Fatalf("TokenSuffix(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestHeadString(t *testing.T) {
	cases := []struct {
		name string
		b    []byte
		n    int
		want string
	}{
		{"shorter than n", []byte("hello"), 10, "hello"},
		{"equal to n", []byte("hello"), 5, "hello"},
		{"longer than n", []byte("hello world"), 5, "hello"},
		{"empty", nil, 10, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := HeadString(c.b, c.n)
			if got != c.want {
				t.Fatalf("HeadString(%q, %d) = %q, want %q", c.b, c.n, got, c.want)
			}
		})
	}
}

// TestHooksDefaultNoop 验证默认 hook 不会 panic 且可被重复调用。
// 业务侧未替换实现时,SDK 触发事件应保持零副作用、零 goroutine 阻塞。
func TestHooksDefaultNoop(t *testing.T) {
	OnTokenRefresh(TokenRefreshEvent{})
	OnTokenRefresh(TokenRefreshEvent{Service: "advertising", Phase: EventPhaseEnd})
	OnAuthRetry(AuthRetryEvent{})
	OnAuthRetry(AuthRetryEvent{Service: "selling_partner", FirstStatus: 401})
}
