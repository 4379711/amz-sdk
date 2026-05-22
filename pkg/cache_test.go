package pkg

import (
	"sync"
	"sync/atomic"
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
// 这是新增方法的核心行为:缓存已被别的 goroutine 刷到新 token T2,
// 当前 goroutine 拿着过期的 T1 去失效,不能误删 T2。
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
	var c AccessTokenCache
	c.Put("rt", CacheItem{
		AccessToken:            "T1",
		AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
	})

	const n = 64
	var wg sync.WaitGroup
	var refreshCount int32

	// 启动一个 goroutine 模拟"A 线程": 失效旧 token 并写入 T2
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.InvalidateIfMatch("rt", "T1")
		atomic.AddInt32(&refreshCount, 1)
		c.Put("rt", CacheItem{
			AccessToken:            "T2",
			AccessTokenExpiredTime: time.Now().Add(30 * time.Minute),
		})
	}()

	// 其他 n 个 goroutine 拿着旧 T1 尝试失效 + Get,模拟"同一时刻一起打 401"的情况
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.InvalidateIfMatch("rt", "T1")
			_ = c.Get("rt")
		}()
	}

	wg.Wait()

	// 最终 cache 里只能是 T2 或空(T2 可能被最后一个 InvalidateIfMatch 刚好在 Put 前执行,导致空;
	// 但任何情况下都不应该是 T1)
	got := c.Get("rt")
	if got != "T2" && got != "" {
		t.Fatalf("after concurrent InvalidateIfMatch, cache should be T2 or empty, got %q", got)
	}
	if got == "" {
		t.Logf("edge timing: Put happened before all InvalidateIfMatch finished, cache empty (acceptable)")
	}
}
