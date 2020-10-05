package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UpgradeHTTPS() gin.HandlerFunc {
	return func(c *gin.Context) {
		scheme := c.Request.URL.Scheme
		if scheme == "http" || (!viper.GetBool("development") && scheme == "") {
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
