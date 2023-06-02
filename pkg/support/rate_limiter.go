package support

import (
	"sync"
	"time"
)

type RateLimiter struct {
	capacity  int
	tokens    int
	refill    time.Duration
	lastCheck time.Time
	mu        sync.Mutex
}

func NewRateLimiter(capacity int, refill time.Duration) *RateLimiter {
	return &RateLimiter{
		capacity:  capacity,
		tokens:    capacity,
		refill:    refill,
		lastCheck: time.Now(),
	}
}

func (rl *RateLimiter) Acquire() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck)
	rl.lastCheck = now

	rl.tokens += int(elapsed / rl.refill)
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}
