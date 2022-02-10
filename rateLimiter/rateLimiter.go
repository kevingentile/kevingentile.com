package rateLimiter

import (
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
}

// NewRateLimiter creates a rate limiter that deilvers a tick every tickPeriod
func NewRateLimiter(tickPeriod time.Duration) (*RateLimiter, error) {
	return &RateLimiter{
		ticker: time.NewTicker(tickPeriod),
	}, nil
}

// Wait is a blocking operation that waits until a tick is received at the end of the tickPeriod
func (rl *RateLimiter) Wait(next func()) {
	<-rl.ticker.C
	next()
}
