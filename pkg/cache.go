package pkg

import (
	"sync"
	"time"
)

type TokenCache interface {
	Get(key any) string
	Put(key any, value CacheItem) string
}

type CacheItem struct {
	AccessToken            string
	AccessTokenExpiredTime time.Time
}

type AccessTokenCache struct {
	sync.Map
}

// accessTokenSafetyWindow 是 access_token 过期前的提前失效窗口。
//
// 必须 **严格大于** DefaultClient.Timeout(当前 180s),否则会触发下面的边界场景:
//  1. Get 判断"token 还剩 3:01" 返回缓存 token
//  2. HTTP 请求进入处理(DNS/TCP/TLS/等待服务端),最长允许 180s
//  3. 请求到达 Amazon 时距离 Get 已过去接近 3min,token 真过期了
//  4. Amazon 返回 403 "access token has expired"
//
// 留 2min 余量抵抗网络抖动和 Amazon 侧排队。access_token 有效期 60min,
// 即使提前 5min 失效,每 55min 刷一次,不会对 LwA 造成压力。
const accessTokenSafetyWindow = 5 * time.Minute

func (s *AccessTokenCache) Get(key any) string {
	value, ok := s.Load(key)
	if ok {
		item := value.(CacheItem)
		if item.AccessTokenExpiredTime.After(time.Now().Add(accessTokenSafetyWindow)) {
			return item.AccessToken
		}
	}
	return ""
}

func (s *AccessTokenCache) Put(key any, value CacheItem) string {
	s.Store(key, value)
	return value.AccessToken
}

// Invalidate 移除给定 refresh_token 对应的缓存条目。
// 当下游接口明确告知"access token 已过期"时(缓存的 ExpiresAt 与 Amazon
// 服务端判定不一致的偶发场景——时钟偏差/节点缓存),调用方应显式 Invalidate,
// 让下一次 Get 返回空,触发 singleflight 重新刷新。
//
// 注意:名字不叫 Delete 是为了避免和嵌入的 sync.Map.Delete 方法产生歧义/误用。
func (s *AccessTokenCache) Invalidate(key any) {
	s.Map.Delete(key)
}

// InvalidateIfMatch 仅当缓存中当前的 AccessToken 仍等于 failedToken 时才删除,
// 专用于多 goroutine 并发重试的场景:
//
//	t0  A 用 T1 请求 -> 401
//	t1  A Invalidate,singleflight 刷出 T2,写入 cache
//	t2  B 用 T1 请求 -> 401 (B 是在 A 刷新前发出去的,返回慢了)
//	t3  B 若无脑 Invalidate,会把 A 刚写的 T2 删掉,触发多余的一次 LwA 刷新
//
// 改用 InvalidateIfMatch(key, T1) 后:t3 时 cache 里已是 T2,failedToken=T1
// 不匹配,不会误删;B 紧跟着的 acquireAccessToken 从 cache 直接命中 T2 复用。
//
// 对 Amazon 配额压力不大,但能让"token 自然过期那一刻"的刷新次数收敛到 1 次。
//
// DEBUG 日志区分三种结果,便于排查"并发竞态是否真的发生":
//   - hit:成功删除了失败 token
//   - skipped-newer-token:缓存里已经是别的 goroutine 刷出来的新 token,
//     本方法保护了新 token 不被误删(就是避免冗余刷新的关键)
//   - miss:key 不存在或值类型异常
func (s *AccessTokenCache) InvalidateIfMatch(key any, failedToken string) {
	value, ok := s.Load(key)
	if !ok {
		DefaultLogger.Debug("access token cache invalidate-if-match: miss",
			"reason", "key-not-found",
			"failed_at_suffix", TokenSuffix(failedToken))
		return
	}
	item, ok := value.(CacheItem)
	if !ok {
		DefaultLogger.Debug("access token cache invalidate-if-match: miss",
			"reason", "type-assertion-failed",
			"failed_at_suffix", TokenSuffix(failedToken))
		return
	}
	if item.AccessToken == failedToken {
		s.Map.Delete(key)
		DefaultLogger.Debug("access token cache invalidate-if-match: hit",
			"failed_at_suffix", TokenSuffix(failedToken))
		return
	}
	DefaultLogger.Debug("access token cache invalidate-if-match: skipped-newer-token",
		"failed_at_suffix", TokenSuffix(failedToken),
		"current_at_suffix", TokenSuffix(item.AccessToken))
}
