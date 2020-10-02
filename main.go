package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var limiter <-chan time.Time

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	log.Println("Start PORT:", viper.GetInt("port"))
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.template.html", nil)
	})

	engine.GET("/rambler", func(c *gin.Context) {
		articlesDir, err := ioutil.ReadDir(viper.GetString("rambler_articles"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		//TODO load articles at process start
		var articles []Article
		for _, article := range articlesDir {
			nameExtension := strings.Split(article.Name(), ".")
			nameIndex := strings.Split(nameExtension[0], "_")
			index, err := strconv.Atoi(nameIndex[1])
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			articles = append(articles, Article{Title: nameIndex[0], Index: index})
		}

		sort.Slice(articles, func(i, j int) bool { return articles[i].Index < articles[j].Index })
		c.HTML(http.StatusOK, "articles.template.html", gin.H{
			"articles": articles,
		})
	})
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

	log.Fatal(engine.Run(":" + viper.GetString("port")))

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
	viper.SetDefault("rambler_articles", "markdown/")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	return nil
}
