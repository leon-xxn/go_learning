package leaky_bucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type LeakyBucket struct {
	capacity     int // 桶的容量
	rate         int // 每秒漏出的请求数量
	currentWater int // 当前水量
	mu           sync.Mutex
	ticker       *time.Ticker // 漏水的速率
}

func NewLeakBucket(capacity int, rate int) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		rate:     rate,
		ticker:   time.NewTicker(time.Second), // 每秒漏水一次
	}
}

func (b *LeakyBucket) startLeaking() {
	go func() {
		for range b.ticker.C {
			b.mu.Lock()
			if b.currentWater > 0 {
				b.currentWater -= b.rate
				if b.currentWater < 0 {
					b.currentWater = 0
				}
				fmt.Print("current water:", b.currentWater)
			}
			b.mu.Unlock()
		}
	}()
}

// AllowRequest 允许请求进入桶
func (b *LeakyBucket) AllowRequest() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.currentWater < b.capacity {
		b.currentWater++
		return true
	}
	return false
}

func (b *LeakyBucket) StopLeaking() {
	b.ticker.Stop()
}

func TestLeakyBucket(t *testing.T) {
	leakyBucket := NewLeakBucket(10, 1)
	leakyBucket.startLeaking()
	for i := 0; i < 15; i++ {
		if leakyBucket.AllowRequest() {
			fmt.Println("handle request")
		} else {
			fmt.Println("reject request")
		}
		time.Sleep(200 * time.Millisecond)
	}
	leakyBucket.StopLeaking()
}
