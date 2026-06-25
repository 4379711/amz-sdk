package auth

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/4379711/amz-sdk/pkg"

	"github.com/bytedance/sonic"
	"github.com/oklog/ulid/v2"
	"golang.org/x/sync/singleflight"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type App struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Beta         bool
	AppID        string
}

type Seller struct {
	SellerName string
	SellerId   string
	// 国家编码：US
	CountryCode string
	// 店铺:SC、VC
	SellerType string
}

type SpAuth struct {
	*App
	*Seller
	*Token
}

func (c *App) String() string {
	return fmt.Sprintf("AppID: %s\nClientID: %s\nClientSecret: %s\nRedirectURL: %s", c.AppID, c.ClientID, c.ClientSecret, c.RedirectURL)
}

func (t *Token) String() string {
	return fmt.Sprintf("AccessToken: %s\nRefreshToken: %s\nTokenType: %s\nExpiresIn: %d\nExpiresAt: %s\n", t.AccessToken, t.RefreshToken, t.TokenType, t.ExpiresIn, t.ExpiresAt.Format("2006-01-02 15:04:05"))
}

func (a *SpAuth) GetLwaTokenEndpoint() string {
	switch a.CountryCode {
	case "CA":
		return "https://api.amazon.com/auth/o2/token"
	case "US":
		return "https://api.amazon.com/auth/o2/token"
	case "MX":
		return "https://api.amazon.com/auth/o2/token"
	case "BR":
		return "https://api.amazon.com/auth/o2/token"
	case "IE":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "ES":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "UK":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "FR":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "BE":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "NL":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "DE":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "IT":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "SE":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "ZA":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "PL":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "EG":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "TR":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "SA":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "AE":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "IN":
		return "https://api.amazon.co.uk/auth/o2/token"
	case "SG":
		return "https://api.amazon.co.jp/auth/o2/token"
	case "AU":
		return "https://api.amazon.co.jp/auth/o2/token"
	case "JP":
		return "https://api.amazon.co.jp/auth/o2/token"
	default:
		panic("invalid region")
	}
}

func (a *SpAuth) GetDataEndpoint() string {
	switch a.CountryCode {
	case "CA":
		return "https://sellingpartnerapi-na.amazon.com"
	case "US":
		return "https://sellingpartnerapi-na.amazon.com"
	case "MX":
		return "https://sellingpartnerapi-na.amazon.com"
	case "BR":
		return "https://sellingpartnerapi-na.amazon.com"
	case "IE":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "ES":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "UK":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "FR":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "BE":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "NL":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "DE":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "IT":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "SE":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "ZA":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "PL":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "EG":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "TR":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "SA":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "AE":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "IN":
		return "https://sellingpartnerapi-eu.amazon.com"
	case "SG":
		return "https://sellingpartnerapi-fe.amazon.com"
	case "AU":
		return "https://sellingpartnerapi-fe.amazon.com"
	case "JP":
		return "https://sellingpartnerapi-fe.amazon.com"
	default:
		panic("invalid region")
	}
}

// fetchAccessToken 向 LwA 端点换一枚新的 access_token 并**原样返回**,
// 不修改 a.Token 任何字段。
//
// 专为并发路径(acquireAccessToken 的 singleflight 回调)设计:同一 refresh_token
// 被多个 goroutine 共享时,若在回调里直接写 a.Token,只有首个 goroutine 的 a.Token
// 会被更新,其它 goroutine 持有的 *SpAuth 对象字段永远陈旧,容易埋坑。
// 因此并发路径统一用这个纯函数拿结果,写共享状态的责任交给 cache.Put。
func (a *SpAuth) fetchAccessToken() (*Token, error) {
	endpoint := a.GetLwaTokenEndpoint()
	reqBody, _ := sonic.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": a.RefreshToken,
		"client_id":     a.ClientID,
		"client_secret": a.ClientSecret,
	})
	resp, err := pkg.DefaultClient.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer(reqBody))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	return decodeTokenResponse(resp)
}

// GetAccessTokenFromEndpoint 向 LwA 端点换一枚新的 access_token 并写入 a.Token。
//
// 保留该方法是为了兼容 IAuth 接口和少量直接调用方(集成测试/首次授权调试工具)。
// 业务侧的高并发路径应该走 BuildClient 构造出来的 http.Client,让 SDK 内部的
// singleflight + cache 统一管理刷新节奏,**不要**在多 goroutine 共享的 *SpAuth
// 上直接调本方法,否则会出现"只有部分 goroutine 的 a.Token 被刷新"的假象。
func (a *SpAuth) GetAccessTokenFromEndpoint() error {
	t, err := a.fetchAccessToken()
	if err != nil {
		return err
	}
	if a.Token == nil {
		a.Token = &Token{}
	}
	a.Token.AccessToken = t.AccessToken
	a.Token.TokenType = t.TokenType
	a.Token.ExpiresIn = t.ExpiresIn
	a.Token.ExpiresAt = t.ExpiresAt
	return nil
}

// ParseToken 解析 refresh_token grant 的响应,只更新 access_token 相关字段。
//
// 不能用解析到的 Token 整体覆盖 a.Token:LwA 对 refresh_token grant 的响应
// 允许省略 refresh_token 字段(OAuth2 规范允许),若全量覆盖会把 a.RefreshToken 清空,
// 下次刷新用空 refresh_token 请求 LwA 会被拒为 invalid_token。
func (a *SpAuth) ParseToken(resp *http.Response) error {
	t, err := decodeTokenResponse(resp)
	if err != nil {
		return err
	}
	if a.Token == nil {
		a.Token = &Token{}
	}
	a.Token.AccessToken = t.AccessToken
	a.Token.TokenType = t.TokenType
	a.Token.ExpiresIn = t.ExpiresIn
	a.Token.ExpiresAt = t.ExpiresAt
	return nil
}

// decodeTokenResponse 解析 LwA token 响应体,不修改 SpAuth 状态。
func decodeTokenResponse(resp *http.Response) (*Token, error) {
	if resp == nil {
		return nil, errors.New("nil token response")
	}
	if resp.StatusCode >= 400 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return nil, errors.New(fmt.Sprintf("Fail to generate token. %s", buf.String()))
	}
	var t Token
	if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, err
	}
	t.ExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)
	return &t, nil
}

// GetRefreshToken 首次授权,用 authorization_code 换完整 Token(含 refresh_token)。
// 这是拿到 refresh_token 的唯一时机,因此全量写入 a.Token。
func (a *SpAuth) GetRefreshToken(code string) error {
	baseURL := a.GetLwaTokenEndpoint()
	resp, err := pkg.DefaultClient.PostForm(baseURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {a.RedirectURL},
		"client_id":     {a.ClientID},
		"client_secret": {a.ClientSecret}})
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	t, err := decodeTokenResponse(resp)
	if err != nil {
		return err
	}
	a.Token = t
	return nil
}
func (a *SpAuth) GetLwaCodeEndpoint() string {
	if a.SellerType == "VC" {
		return GetVCLwaCodeEndpoint(a.CountryCode)
	} else if a.SellerType == "SC" {
		return GetSCLwaCodeEndpoint(a.CountryCode)
	} else {
		panic("invalid SellerType")
	}
}
func (a *SpAuth) GetLwaURL() (url string) {
	baseURL := a.GetLwaCodeEndpoint()
	url = fmt.Sprintf("%s?application_id=%s&redirect_uri=%s&state=%s", baseURL, a.AppID, a.RedirectURL, ulid.Make().String())
	if a.Beta {
		url = fmt.Sprintf("%s&version=beta", url)
	}
	return url
}

func GetSCLwaCodeEndpoint(countryCode string) string {
	switch countryCode {
	case "CA":
		return "https://sellercentral.amazon.ca/apps/authorize/consent"
	case "US":
		return "https://sellercentral.amazon.com/apps/authorize/consent"
	case "MX":
		return "https://sellercentral.amazon.com.mx/apps/authorize/consent"
	case "BR":
		return "https://sellercentral.amazon.com.br/apps/authorize/consent"
	case "IE":
		return "https://sellercentral.amazon.ie/apps/authorize/consent"
	case "ES":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	case "UK":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	case "FR":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	case "BE":
		return "https://sellercentral.amazon.com.be/apps/authorize/consent"
	case "NL":
		return "https://sellercentral.amazon.nl/apps/authorize/consent"
	case "DE":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	case "IT":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	case "SE":
		return "https://sellercentral.amazon.se/apps/authorize/consent"
	case "ZA":
		return "https://sellercentral.amazon.co.za/apps/authorize/consent"
	case "PL":
		return "https://sellercentral.amazon.pl/apps/authorize/consent"
	case "EG":
		return "https://sellercentral.amazon.eg/apps/authorize/consent"
	case "TR":
		return "https://sellercentral.amazon.com.tr/apps/authorize/consent"
	case "SA":
		return "https://sellercentral.amazon.sa/apps/authorize/consent"
	case "AE":
		return "https://sellercentral.amazon.ae/apps/authorize/consent"
	case "IN":
		return "https://sellercentral.amazon.in/apps/authorize/consent"
	case "SG":
		return "https://sellercentral.amazon.sg/apps/authorize/consent"
	case "AU":
		return "https://sellercentral.amazon.com.au/apps/authorize/consent"
	case "JP":
		return "https://sellercentral.amazon.co.jp/apps/authorize/consent"
	default:
		panic("invalid region")
	}
}

func GetVCLwaCodeEndpoint(countryCode string) string {
	switch countryCode {
	case "CA":
		return "https://vendorcentral.amazon.ca/apps/authorize/consent"
	case "US":
		return "https://vendorcentral.amazon.com/apps/authorize/consent"
	case "MX":
		return "https://vendorcentral.amazon.com.mx/apps/authorize/consent"
	case "BR":
		return "https://vendorcentral.amazon.com.br/apps/authorize/consent"
	case "IE":
		return ""
	case "ES":
		return "https://vendorcentral.amazon.es/apps/authorize/consent"
	case "UK":
		return "https://vendorcentral.amazon.co.uk/apps/authorize/consent"
	case "FR":
		return "https://vendorcentral.amazon.fr/apps/authorize/consent"
	case "BE":
		return "https://vendorcentral.amazon.com.be/apps/authorize/consent"
	case "NL":
		return "https://vendorcentral.amazon.nl/apps/authorize/consent"
	case "DE":
		return "https://vendorcentral.amazon.de/apps/authorize/consent"
	case "IT":
		return "https://vendorcentral.amazon.it/apps/authorize/consent"
	case "SE":
		return "https://vendorcentral.amazon.se/apps/authorize/consent"
	case "ZA":
		return "https://vendorcentral.amazon.co.za/apps/authorize/consent"
	case "PL":
		return "https://vendorcentral.amazon.pl/apps/authorize/consent"
	case "EG":
		return "https://vendorcentral.amazon.me/apps/authorize/consent"
	case "TR":
		return "https://vendorcentral.amazon.me/apps/authorize/consent"
	case "SA":
		return "https://vendorcentral.amazon.me/apps/authorize/consent"
	case "AE":
		return "https://vendorcentral.amazon.com.tr/apps/authorize/consent"
	case "IN":
		return "https://www.vendorcentral.in/apps/authorize/consent"
	case "SG":
		return "https://vendorcentral.amazon.com.sg/apps/authorize/consent"
	case "AU":
		return "https://vendorcentral.amazon.com.au/apps/authorize/consent"
	case "JP":
		return "https://vendorcentral.amazon.co.jp/apps/authorize/consent"
	default:
		panic("invalid region")
	}
}

var (
	cache         = new(pkg.AccessTokenCache)
	refreshFlight singleflight.Group
)

type headerInjector struct {
	rt   http.RoundTripper
	auth *SpAuth
}

// acquireAccessToken 并发安全地获取/刷新 access_token:
// 同一 refresh_token 在同一时刻只允许一个 in-flight LwA 刷新请求,
// 其余 goroutine 等待并复用结果,避免并发启动多个 runner 时对 LwA 造成风暴。
//
// 回调里刻意不使用 GetAccessTokenFromEndpoint,因为那会把结果写进
// h.auth.Token;而 h.auth 是首个进入 singleflight 的 goroutine 所持有的
// *SpAuth,其他 goroutine 持有的是不同的 *SpAuth 对象,字段不会被同步。
// 真正的跨 goroutine 共享状态是 cache,这里只把结果写到 cache 就够了。
//
// reason 描述触发本次调用的上游原因("cache-miss" / "401-retry"),仅用于
// TokenRefreshEvent,不影响实际行为。
func (h *headerInjector) acquireAccessToken(reason string) (string, error) {
	if tk := cache.Get(h.auth.RefreshToken); tk != "" {
		return tk, nil
	}
	v, err, _ := refreshFlight.Do(h.auth.RefreshToken, func() (any, error) {
		// double-check:等锁期间可能已被其它 goroutine 刷好
		if cached := cache.Get(h.auth.RefreshToken); cached != "" {
			return cached, nil
		}
		rtSuffix := pkg.TokenSuffix(h.auth.RefreshToken)
		pkg.OnTokenRefresh(pkg.TokenRefreshEvent{
			Service:            "selling_partner",
			RefreshTokenSuffix: rtSuffix,
			Reason:             reason,
			Phase:              pkg.EventPhaseStart,
		})
		begin := time.Now()
		t, err := h.auth.fetchAccessToken()
		end := pkg.TokenRefreshEvent{
			Service:            "selling_partner",
			RefreshTokenSuffix: rtSuffix,
			Reason:             reason,
			Phase:              pkg.EventPhaseEnd,
			Duration:           time.Since(begin),
			Err:                err,
		}
		if err == nil {
			end.Success = true
			end.AccessTokenSuffix = pkg.TokenSuffix(t.AccessToken)
			end.ExpiresAt = t.ExpiresAt
		}
		pkg.OnTokenRefresh(end)
		if err != nil {
			return "", err
		}
		return cache.Put(h.auth.RefreshToken, pkg.CacheItem{
			AccessToken:            t.AccessToken,
			AccessTokenExpiredTime: t.ExpiresAt,
		}), nil
	})
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

func (h *headerInjector) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	// 调用方显式传入的 access_token 优先(例如跨服务复用 RDT)。
	// 外部注入的 token 我们无权失效/刷新,直接透传即可,不参与重试逻辑。
	if v := ctx.Value("x-amz-access-token"); v != nil {
		if s, ok := v.(string); ok && s != "" {
			req.Header.Set("x-amz-access-token", s)
			pkg.DefaultLogger.Debug("sp api using injected access_token (RDT)",
				"method", req.Method,
				"url", req.URL.String(),
				"injected_at_suffix", pkg.TokenSuffix(s))
			return h.rt.RoundTrip(req)
		}
	}

	tk, err := h.acquireAccessToken("cache-miss")
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-amz-access-token", tk)
	resp, err := h.rt.RoundTrip(req)
	if err != nil || resp == nil {
		if err != nil {
			// 首次请求的网络错误 (timeout / connection reset / DNS 等) 业务层
			// 也能看到,但 SDK 拥有更结构化的上下文 (method/url/refresh_token 尾),
			// 用 WARN 级别帮助跨请求聚合分析。
			pkg.DefaultLogger.Warn("sp api request network error",
				"method", req.Method,
				"url", req.URL.String(),
				"rt_suffix", pkg.TokenSuffix(h.auth.RefreshToken),
				"error", err.Error())
		}
		return resp, err
	}

	// 兜底重试:缓存的 ExpiresAt 与 Amazon 服务端判定偶尔会不一致
	// (时钟偏差/服务端节点缓存等),此时 SP API 会返回 403 "access token has expired"。
	// 这种瞬时错误应由 SDK 透明处理,不应暴露给上层,否则业务层会将其与
	// refresh_token 吊销混淆。
	matched, firstBody := shouldRetryAuthExpired(resp)
	if !matched {
		return resp, nil
	}
	event := pkg.AuthRetryEvent{
		Service:                 "selling_partner",
		RefreshTokenSuffix:      pkg.TokenSuffix(h.auth.RefreshToken),
		Method:                  req.Method,
		URL:                     req.URL.String(),
		FirstStatus:             resp.StatusCode,
		FirstBodyHead:           pkg.HeadString(firstBody, 256),
		FailedAccessTokenSuffix: pkg.TokenSuffix(tk),
	}
	defer func() { pkg.OnAuthRetry(event) }()

	retryReq, ok := cloneRequestForRetry(req)
	if !ok {
		// body 已消费且无法重放,不做重试,把原响应还给上层
		event.RetrySkipped = "body-not-cloneable"
		return resp, nil
	}
	// 丢弃首次响应 body,避免连接池里的连接被挂起。
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// 只在 cache 里仍是"刚刚失败的那个 token"时才失效,避免并发重试场景
	// 把别的 goroutine 刚刷出来的新 token 误删,触发冗余 LwA 刷新。
	cache.InvalidateIfMatch(h.auth.RefreshToken, tk)
	newTk, err := h.acquireAccessToken("401-retry")
	if err != nil {
		event.RetryErr = err
		return nil, err
	}
	retryReq.Header.Set("x-amz-access-token", newTk)
	retryResp, retryErr := h.rt.RoundTrip(retryReq)
	if retryErr != nil {
		event.RetryErr = retryErr
	} else if retryResp != nil {
		event.RetryStatus = retryResp.StatusCode
	}
	return retryResp, retryErr
}

// shouldRetryAuthExpired 仅在 SP API 明确告知 access_token 失效时才返回 true。
// 其它 401/403(比如真正的权限缺失、签名错误)不重试,直接抛给上层处理。
//
// 该函数会"消费"并重建 resp.Body,这样上层拿到的 resp 仍可被正常解码。
//
// SP API 的 report 下载接口实际上走的是广告 API 域名
// (advertising-api.amazon.com),body 会沿用广告 API 的 "Invalid token"
// 措辞,因此此处与 advertising/auth 匹配同一份关键字,避免跨域名调用时
// 的不一致。
//
// 返回的 body 是已完整读出的字节切片,调用方可直接拿去做诊断日志,不需要再次
// 从 resp.Body 读(resp.Body 已被重建为 NopCloser,重新读会失败)。
func shouldRetryAuthExpired(resp *http.Response) (matched bool, body []byte) {
	if resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusUnauthorized {
		return false, nil
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return false, body
	}
	lo := strings.ToLower(string(body))
	switch {
	// SP API 经典响应:"The access token you provided has expired."
	case strings.Contains(lo, "access token") && strings.Contains(lo, "expired"):
		return true, body
	// 跨域名调用(SP API 的报表下载)可能出现广告 API 的措辞
	case strings.Contains(lo, "invalid token"), strings.Contains(lo, "invalid_token"):
		return true, body
	case strings.Contains(lo, "authenticating lwa token"):
		return true, body
	case strings.TrimSpace(lo) == `{"message":"unauthorized"}`:
		return true, body
	}
	return false, body
}

// cloneRequestForRetry 克隆一个可重放的 Request。
// 对于 GET/DELETE 等无 body 的请求,直接 Clone 即可;
// 对于 POST/PUT/PATCH,要求原 Request 设置了 GetBody(Go 标准库在
// http.NewRequest(..., body) 中对 bytes.Buffer/Reader/strings.Reader 会自动设置),
// 否则无法重放,此时放弃重试。
func cloneRequestForRetry(req *http.Request) (*http.Request, bool) {
	clone := req.Clone(req.Context())
	if req.Body == nil || req.Body == http.NoBody {
		return clone, true
	}
	if req.GetBody == nil {
		return nil, false
	}
	body, err := req.GetBody()
	if err != nil {
		return nil, false
	}
	clone.Body = body
	return clone, true
}

func (a *SpAuth) BuildClient() *http.Client {
	client := *pkg.DefaultClient
	client.Transport = &headerInjector{
		auth: a,
		rt:   pkg.SharedTransport,
	}
	return &client
}
