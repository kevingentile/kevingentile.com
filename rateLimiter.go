package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	limiter <-chan time.Time
}

// NewRateLimiter creates a new rate limiter with the default duration
func NewRateLimiter(duration time.Duration) (*RateLimiter, error) {
	return &RateLimiter{
		limiter: time.Tick(duration),
	}, nil
}

// RateLimit a route using a shared RateLimiter channel
func (rl *RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		<-rl.limiter
		c.Next()
	}
}
