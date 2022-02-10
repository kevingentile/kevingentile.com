package main

import (
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/kevingentile/kevingentile.com/handlers"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		Fatal(err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = viper.GetString("port")
	}
	Debug("Start PORT:", port)
	engine := gin.Default()
	engine.Use(UpgradeHTTPSMiddleware())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/home")
	})
	obsLimited := engine.Group("/obs").Use(handlers.NewRateLimitHandler(time.Second * viper.GetDuration("limiter_duration")))
	{
		obsLimited.GET("/:platform/:username", handleFortniteData)
	}

	// TODO move to angular
	engine.StaticFile("/fortnite", "pages/fortnite.html")

	// Serve this file for any /fornite route
	engine.GET("/fortnite/:platform/:username", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.File("assets/obs/fortnite.html")
	})

	engine.Static("/fortnite/assets", "assets/obs")

	articleHandler, err := NewArticleHandler()
	if err != nil {
		Fatal(err)
	}

	apiGroup := engine.Group("/api").Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	{
		apiGroup.GET("/articles", articleHandler.ApiListHandler)
		apiGroup.GET("/articles/:article_date", articleHandler.ApiArticleHandler)
	}

	engine.Static("/home", "angular/kevingentile-com/dist/kevingentile-com")

	engine.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/")
	})

	Fatal(engine.Run(":" + port))
}

func initConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		Fatal(err)
	}

	// validate website home directory exists, if not create it
	webDir := path.Join(home, ".kevingentile.com")
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		Print("creating", webDir)
		if err := os.Mkdir(webDir, os.ModePerm); err != nil {
			return err
		}
	}

	// validate config file exists, if not create it
	configPath := path.Join(webDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		Print("creating", configPath)
		if _, err := os.Create(configPath); err != nil {
			return err
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetEnvPrefix("KG")
	viper.AutomaticEnv()
	viper.AddConfigPath(webDir)

	viper.SetDefault("port", 8080)
	viper.SetDefault("limiter_duration", 3)
	viper.SetDefault("rambler_articles", "articles/")

	if err := viper.ReadInConfig(); err != nil {
		Error(err)
	}

	return nil
}
