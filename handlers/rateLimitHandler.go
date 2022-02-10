package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevingentile/kevingentile.com/rateLimiter"
)

// RateLimit a route using a shared RateLimiter channel
func NewRateLimitHandler(tickDuration time.Duration) gin.HandlerFunc {
	rl, err := rateLimiter.NewRateLimiter(tickDuration)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		rl.Wait(func() {
			c.Next()
		})
	}
}
