package main

import (
	"net/http"
	"sync"
	"time"
)

type bucket struct {
	tokens   float64
	lastSeen time.Time
}

// rateLimiter is a per-client-IP token bucket. Stale buckets are swept
// periodically so a flood of distinct IPs can't grow the map without bound.
type rateLimiter struct {
	mu       sync.Mutex
	buckets  map[string]*bucket
	rps      float64
	burst    int
	stopChan chan struct{}
}

func newRateLimiter(rps float64, burst int) *rateLimiter {
	rl := &rateLimiter{
		buckets:  make(map[string]*bucket),
		rps:      rps,
		burst:    burst,
		stopChan: make(chan struct{}),
	}
	go rl.cleanupLoop()
	return rl
}

func (rl *rateLimiter) Stop() {
	close(rl.stopChan)
}

func (rl *rateLimiter) cleanupLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			cutoff := time.Now().Add(-10 * time.Minute)
			rl.mu.Lock()
			for key, b := range rl.buckets {
				if b.lastSeen.Before(cutoff) {
					delete(rl.buckets, key)
				}
			}
			rl.mu.Unlock()
		case <-rl.stopChan:
			return
		}
	}
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, ok := rl.buckets[key]
	if !ok {
		b = &bucket{tokens: float64(rl.burst), lastSeen: now}
		rl.buckets[key] = b
	}

	elapsed := now.Sub(b.lastSeen).Seconds()
	b.tokens += elapsed * rl.rps
	if b.tokens > float64(rl.burst) {
		b.tokens = float64(rl.burst)
	}
	b.lastSeen = now

	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.allow(clientIP(r)) {
			w.Header().Set("Retry-After", "1")
			writeJSON(w, http.StatusTooManyRequests, map[string]string{"error": "rate limit exceeded"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
