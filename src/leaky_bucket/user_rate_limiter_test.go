package leaky_bucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserRateLimiter struct {
	rate      int           //每分钟允许的请求数
	interval  time.Duration //添加令牌时间间隔
	tokens    int           //当前令牌数量
	maxTokens int           //令牌桶容量
	mu        sync.Mutex
}

func NewUserRateLimiter(maxToken int, ratePerMinute int) *UserRateLimiter {
	interval := time.Minute / time.Duration(ratePerMinute)
	limiter := &UserRateLimiter{
		rate:      ratePerMinute,
		interval:  interval,
		tokens:    maxToken,
		maxTokens: maxToken,
	}
	go limiter.refillToken()
	return limiter
}

func (u *UserRateLimiter) refillToken() {
	ticker := time.NewTicker(u.interval)
	for range ticker.C {
		u.mu.Lock()
		if u.tokens < u.maxTokens {
			u.tokens++
			fmt.Print("Refill token:", u.tokens)
		}
		u.mu.Unlock()
	}
}

func (u *UserRateLimiter) AllowRequest() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.tokens > 0 {
		u.tokens--
		return true
	}
	return false
}

func TestUserRateLimiter(t *testing.T) {
	rateLimiter := NewUserRateLimiter(10, 10)
	for i := 0; i < 15; i++ {
		if rateLimiter.AllowRequest() {
			fmt.Println("allow request", i)
		} else {
			fmt.Println("limit request", i)
		}
		time.Sleep(4 * time.Second) // 每 4 秒发送一次请求
	}
}
