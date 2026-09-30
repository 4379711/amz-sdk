package auth

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestGetLwaURL_PreservesRedirectURI(t *testing.T) {
	for _, redirectURI := range []string{
		"https://app.example/callback",
		"https://app.example/callback?tenant=1",
		"https://app.example/callback?tenant=1&next=a+b",
		"https://app.example/a+b",
		"https://app.example/a%2Fb",
		"https://app.example/回调?next=%2F商品&tenant=示例",
	} {
		t.Run(redirectURI, func(t *testing.T) {
			a := contextTestAuth(t)
			a.RedirectURL = redirectURI
			a.SellerType = "SC"
			a.AppID = "test-app&+%"
			parsed, err := url.Parse(a.GetLwaURL())
			if err != nil {
				t.Fatal(err)
			}
			query, err := url.ParseQuery(parsed.RawQuery)
			if err != nil {
				t.Fatal(err)
			}
			if got := query.Get("redirect_uri"); got != redirectURI {
				t.Fatalf("redirect_uri=%q, want %q", got, redirectURI)
			}
			if query.Get("application_id") != a.AppID {
				t.Fatalf("application_id=%q", query.Get("application_id"))
			}
			if _, err := ulid.ParseStrict(query.Get("state")); err != nil {
				t.Fatalf("state 不是有效的ULID: %v", err)
			}
			if query.Has("next") || query.Has("tenant") {
				t.Fatalf("回调地址参数泄漏到授权参数中: %v", query)
			}
			var exchangeURI string
			installContextTokenTransport(t, contextTestTransport(func(req *http.Request) (*http.Response, error) {
				defer req.Body.Close()
				if err := req.ParseForm(); err != nil {
					return nil, err
				}
				exchangeURI = req.PostForm.Get("redirect_uri")
				return newResp(200, `{"access_token":"test-access","refresh_token":"test-refresh","expires_in":3600}`), nil
			}))
			if err := a.GetRefreshToken("test-code"); err != nil {
				t.Fatal(err)
			}
			if exchangeURI != query.Get("redirect_uri") {
				t.Fatalf("授权及换码回调地址不一致: %q != %q", query.Get("redirect_uri"), exchangeURI)
			}
		})
	}
}

func TestGetLwaURL_Beta(t *testing.T) {
	a := contextTestAuth(t)
	a.SellerType = "SC"
	for _, beta := range []bool{false, true} {
		a.Beta = beta
		parsed, err := url.Parse(a.GetLwaURL())
		if err != nil {
			t.Fatal(err)
		}
		query := parsed.Query()
		if beta && query.Get("version") != "beta" {
			t.Fatalf("beta版本缺失: %v", query)
		}
		if !beta && query.Has("version") {
			t.Fatalf("非beta授权包含version: %v", query)
		}
	}
}
