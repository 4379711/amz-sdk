package pkg

import "time"

// TokenRefreshEvent 描述一次"真正发生的"LwA access_token 换取过程。
//
// 发送时机:singleflight 回调进入 HTTP 请求阶段时就算一次事件;cache 命中、
// singleflight 双检直接复用别人结果的场景不发(避免淹没日志)。
//
// 一次刷新按时序会收到两个事件:Phase=EventPhaseStart / EventPhaseEnd。
// 末尾事件通过 Success 字段区分成功失败,便于外部 logger 做 start+duration 对齐。
type TokenRefreshEvent struct {
	// Service 触发方,固定取值:"advertising" / "selling_partner"。
	Service string

	// RefreshTokenSuffix 已脱敏的 refresh_token 最后 8 位,用于关联同一 seller
	// 的多次事件,安全地写入日志。
	RefreshTokenSuffix string

	// Reason 触发本次刷新的上游原因:
	//   "cache-miss"  首次或 token 自然临近过期
	//   "401-retry"   前一次业务请求被 Amazon 判为无效 token,强制刷新重试
	Reason string

	// Phase 当前处于刷新的哪个阶段:EventPhaseStart / EventPhaseEnd。
	Phase string

	// === 以下字段仅 Phase=EventPhaseEnd 时填充 ===

	// Success 是否成功;为 false 时 Err 必定非 nil。
	Success bool
	// Duration 从发起 HTTP 到收到响应的耗时。
	Duration time.Duration
	// AccessTokenSuffix 新 access_token 最后 8 位(仅成功时)。
	AccessTokenSuffix string
	// ExpiresAt 新 token 的预期过期时刻(仅成功时)。
	ExpiresAt time.Time
	// Err 失败原因(仅失败时)。
	Err error
}

// AuthRetryEvent 描述一次 SDK 层的 401/403 兜底重试。
//
// 触发条件:shouldRetryAuthExpired 命中(SDK 判定是 access_token 被 Amazon
// 拒绝的瞬时错误),SDK 将清缓存、重刷 token、重发一次原请求。
type AuthRetryEvent struct {
	Service            string
	RefreshTokenSuffix string
	Method             string
	URL                string

	// FirstStatus / FirstBodyHead 首次响应的状态码与 body 前 256 字节,
	// 便于诊断 shouldRetryAuthExpired 匹配了哪条规则。
	FirstStatus   int
	FirstBodyHead string

	// FailedAccessTokenSuffix 首次请求用的 token 后 8 位。
	FailedAccessTokenSuffix string

	// 以下字段表示"重试本身的结果":
	//   - RetrySkipped 非空:未执行重试,原因如 "body-not-cloneable";
	//     此时 RetryStatus/RetryErr 均为零值。
	//   - RetrySkipped 空且 RetryErr 非 nil:网络错误,重试请求未拿到响应。
	//   - RetrySkipped 空且 RetryStatus != 0:重试已经拿到 HTTP 响应。
	RetrySkipped string
	RetryStatus  int
	RetryErr     error
}

// EventPhase* 表示一次刷新事件的时序阶段,用于配对 start/end 两条日志。
const (
	EventPhaseStart = "start"
	EventPhaseEnd   = "end"
)

// OnTokenRefresh / OnAuthRetry 是可选的事件 hook。业务侧可在启动时(例如 main
// 里)替换为自定义实现以接入项目日志系统。默认实现为空,SDK 不产生任何日志副作用。
//
// 两个变量只应在程序启动时设置一次,不要在运行中并发修改(SDK 内部没有锁
// 保护)。如果业务侧确实需要运行中切换,应自行实现同步。
var (
	OnTokenRefresh = func(TokenRefreshEvent) {}
	OnAuthRetry    = func(AuthRetryEvent) {}
)

// TokenSuffix 返回字符串的最后 8 个字符,前面加 "..."。用于安全地把
// refresh_token / access_token 写入日志,避免泄漏完整凭据。
// 对不足 8 位的字符串(异常或空值)返回固定占位符。
func TokenSuffix(s string) string {
	if len(s) < 8 {
		return "<short>"
	}
	return "..." + s[len(s)-8:]
}

// HeadString 返回 b 的前 n 字节作为字符串。b 短于 n 时直接全部返回。
// 专用于把 HTTP 响应 body 前若干字节写入诊断事件,既保留诊断价值又避免
// 日志膨胀 / 误把长 body 的敏感部分打出去。
func HeadString(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n])
}
