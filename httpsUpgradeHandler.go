package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UpgradeHTTPSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		scheme := c.Request.URL.Scheme
		if scheme == "http" || (scheme == "" && !viper.GetBool("development")) {
			targetURL := url.URL{
				Scheme: "https", Host: c.Request.Host, Path: c.Request.URL.Path, RawQuery: c.Request.URL.RawQuery,
			}
			Debug("HTTPS upgrade source:", c.Request.URL.String(), "| target: ", targetURL.String())
			c.Redirect(http.StatusTemporaryRedirect, targetURL.String())
			c.Abort()
			return
		}

		c.Next()
	}
}
