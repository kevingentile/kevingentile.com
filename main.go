package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var limiter <-chan time.Time

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = viper.GetString("port")
	}
	log.Println("Start PORT:", port)
	engine := gin.Default()

	engine.Use(UpgradeHTTPS())
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.template.html", nil)
	})

	articleHandler, err := NewArticleHandler()
	if err != nil {
		log.Fatal(err)
	}
	engine.GET("/rambler", articleHandler.ListHandler)
	engine.GET("/rambler/:article_date", articleHandler.Handler)

	rl, err := NewRateLimiter(time.Second * viper.GetDuration("limiter_duration"))
	if err != nil {
		log.Fatal(err)
	}

	obsLimited := engine.Group("/obs").Use(rl.RateLimit())
	{
		obsLimited.GET("/:platform/:username", handleFortniteData)
	}

	engine.StaticFile("/fortnite", "pages/fortnite.html")

	// Serve this file for any /fornite route
	engine.GET("/fortnite/:platform/:username", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.File("assets/obs/fortnite.html")
	})

	engine.Static("/assets", "assets")
	engine.Static("/images", "images")

	engine.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/")
	})

	log.Fatal(engine.Run(":" + port))

}

func initConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// validate website home directory exists, if not create it
	webDir := path.Join(home, ".kevingentile.com")
	if _, err := os.Stat(webDir); os.IsNotExist(err) {
		fmt.Println("creating", webDir)
		if err := os.Mkdir(webDir, os.ModePerm); err != nil {
			return err
		}
	}

	// validate config file exists, if not create it
	configPath := path.Join(webDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("creating", configPath)
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
		log.Println(err)
	}

	return nil
}
