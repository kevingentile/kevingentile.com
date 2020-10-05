package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpgradeHTTPS() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL
		if url.Scheme == "http" {
			url.Scheme = "https"
			c.Redirect(http.StatusPermanentRedirect, url.String())
			return
		}

		c.Next()
	}
}
