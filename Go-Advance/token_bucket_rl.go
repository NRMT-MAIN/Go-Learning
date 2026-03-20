package main

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type TockenBucket struct {
	capacity   int
	tokens     int
	refillRate time.Duration
	mu         sync.Mutex
	lastRefill time.Time
}

func NewTockenBucket(capacity int, refillRate time.Duration) *TockenBucket {
	return &TockenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TockenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	
	tokensToAdd := int(elapsed / tb.refillRate)
	if tokensToAdd > 0 {
		tb.tokens += tokensToAdd
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastRefill = now
	}
}

func (tb *TockenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TockenBucket) GetTokens() int {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens
}

func (tb *TockenBucket) AllowN(n int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}
	return false
}

func tockenBucketExample() {
	// 10 tokens per second, capacity of 10 tokens
	//limiter := NewTockenBucket(10, time.Second)

	// for i := 0; i < 20; i++ {
	// 	if limiter.Allow() {
	// 		println("Request allowed at", time.Now().Format("15:04:05"))
	// 	} else {
	// 		println("Request denied at", time.Now().Format("15:04:05"))
	// 	}
	// 	time.Sleep(100 * time.Millisecond) // Simulate requests every 100ms
	// }


	// Token Bucket Package

	// 10 tokens per second, burst size of 20 tokens
	limiter := rate.NewLimiter(10 , 20) ; 

	// Not-Blocking
	if limiter.Allow() {
		println("Request allowed at", time.Now().Format("15:04:05"))
	} else {
		println("Request denied at", time.Now().Format("15:04:05"))
	}

	// Wait until a token is available
	ctx := context.Background()
	err := limiter.Wait(ctx)
	if err != nil {
		println("Error waiting for token:", err.Error())
	} else {
		println("Request allowed after waiting at", time.Now().Format("15:04:05"))
	}

	// Get Reservation
	reservation := limiter.Reserve()
	if reservation.OK() {
		println("Request allowed immediately at", time.Now().Format("15:04:05"))
	} else {
		println("Request denied at", time.Now().Format("15:04:05"))
	}

	err = limiter.WaitN(ctx, 5) // Wait for 5 tokens
	if err != nil {
		println("Error waiting for 5 tokens:", err.Error())
	} else {
		println("Request allowed after waiting for 5 tokens at", time.Now().Format("15:04:05"))
	}
}
