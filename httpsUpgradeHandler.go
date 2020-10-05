package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UpgradeHTTPSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !viper.GetBool("development") && c.Request.Header.Get("x-forwarded-proto") != "https" {
			targetURL := url.URL{
				Scheme: "https", Host: c.Request.Host, Path: c.Request.URL.Path, RawQuery: c.Request.URL.RawQuery,
			}
			c.Redirect(http.StatusTemporaryRedirect, targetURL.String())
			c.Abort()
			return
		}

		c.Next()
	}
}
