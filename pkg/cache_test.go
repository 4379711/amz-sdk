package pkg

import (
	"sync"
	"testing"
	"time"
)

func TestAccessTokenCache_GetPut(t *testing.T) {
	var c AccessTokenCache

	// 未写入 -> Get 返回空
	if tk := c.Get("rt"); tk != "" {
		t.Fatalf("empty cache Get want \"\", got %q", tk)
	}

	// 正常写入 -> Get 可取回(必须超过 safety window)
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})
	if tk := c.Get("rt"); tk != "T1" {
		t.Fatalf("fresh token Get want T1, got %q", tk)
	}

	// 还剩不到 safety window -> Get 视为过期,返回空
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(accessTokenSafetyWindow - time.Second),
	})
	if tk := c.Get("rt"); tk != "" {
		t.Fatalf("within-safety-window Get want \"\", got %q", tk)
	}

	// 已过期 -> Get 返回空
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(-time.Minute),
	})
	if tk := c.Get("rt"); tk != "" {
		t.Fatalf("expired Get want \"\", got %q", tk)
	}
}

func TestAccessTokenCache_Invalidate(t *testing.T) {
	var c AccessTokenCache
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})

	c.Invalidate("rt")
	if tk := c.Get("rt"); tk != "" {
		t.Fatalf("after Invalidate Get want \"\", got %q", tk)
	}

	// Invalidate 不存在的 key 不报错
	c.Invalidate("not-exist")
}

// TestAccessTokenCache_InvalidateIfMatch_HitsOldToken
// 场景:cache 里仍是失败时用的那个 token,应该被删除(和 Invalidate 等价)。
func TestAccessTokenCache_InvalidateIfMatch_HitsOldToken(t *testing.T) {
	var c AccessTokenCache
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})

	c.InvalidateIfMatch("rt", "T1")
	if tk := c.Get("rt"); tk != "" {
		t.Fatalf("matching-token Invalidate should delete, but Get got %q", tk)
	}
}

// TestAccessTokenCache_InvalidateIfMatch_SkipsNewerToken
// 缓存已被别的 goroutine 刷到新 token T2,持旧 T1 的请求不能误删 T2。
func TestAccessTokenCache_InvalidateIfMatch_SkipsNewerToken(t *testing.T) {
	var c AccessTokenCache
	c.Put("rt", CacheItem{
		AccessToken:            "T2",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})

	c.InvalidateIfMatch("rt", "T1")
	if tk := c.Get("rt"); tk != "T2" {
		t.Fatalf("newer token T2 was wrongly deleted, Get got %q", tk)
	}
}

// TestAccessTokenCache_InvalidateIfMatch_MissingKey
// key 不存在时静默返回,不 panic,不影响后续写入。
func TestAccessTokenCache_InvalidateIfMatch_MissingKey(t *testing.T) {
	var c AccessTokenCache
	c.InvalidateIfMatch("rt", "T1")

	c.Put("rt", CacheItem{
		AccessToken:            "T2",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})
	if tk := c.Get("rt"); tk != "T2" {
		t.Fatalf("after missing-key Invalidate + Put, want T2, got %q", tk)
	}
}

// TestAccessTokenCache_Concurrent_InvalidateIfMatch 模拟并发重试场景:
//   - 多个 goroutine 同时以过期 token T1 作为 failedToken 调用 InvalidateIfMatch
//   - 其间另一个 goroutine 刚用 T2 写入 cache
//   - 期望:最终 cache 仍是 T2,没有被任何一个持 T1 的 goroutine 误删
//
// 这里用 -race 运行可以同时检验 cache 并发读写的线程安全。
func TestAccessTokenCache_Concurrent_InvalidateIfMatch(t *testing.T) {
	const attempts = 1000
	const callers = 16
	expiresAt := time.Now().Add(30 * time.Minute)
	for attempt := 0; attempt < attempts; attempt++ {
		var c AccessTokenCache
		c.Put("rt", CacheItem{AccessToken: "T1", AccessTokenExpiredTime: expiresAt})
		start := make(chan struct{})
		var wg sync.WaitGroup
		for i := 0; i < callers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				<-start
				c.InvalidateIfMatch("rt", "T1")
			}()
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			c.InvalidateIfMatch("rt", "T1")
			c.Put("rt", CacheItem{AccessToken: "T2", AccessTokenExpiredTime: expiresAt})
		}()
		close(start)
		wg.Wait()
		if got := c.Get("rt"); got != "T2" {
			t.Fatalf("attempt %d: fresh token must survive old-token invalidation, got %q", attempt, got)
		}
	}
}
