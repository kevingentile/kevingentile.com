package main

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Article struct {
	Title string
	Index int
	Date  string
}

type ArticleHandler struct {
	Articles []Article
}

const dateFormat string = "Jan-02-2006"

func NewArticleHandler() (*ArticleHandler, error) {
	var articles []Article
	articlesDir, err := ioutil.ReadDir(viper.GetString("rambler_articles"))
	if err != nil {
		return nil, err
	}

	for _, article := range articlesDir {
		nameExtension := strings.Split(article.Name(), ".")
		nameIndexDate := strings.Split(nameExtension[0], "_")
		index, err := strconv.Atoi(nameIndexDate[1])
		if err != nil {
			return nil, err
		}

		date, err := time.Parse(dateFormat, nameIndexDate[2])
		if err != nil {
			return nil, err
		}
		articles = append(articles, Article{Title: nameIndexDate[0], Index: index, Date: date.Format(dateFormat)})
	}

	sort.Slice(articles, func(i, j int) bool { return articles[i].Index < articles[j].Index })

	return &ArticleHandler{
		Articles: articles,
	}, nil
}

func (ah *ArticleHandler) Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "articles.template.html", gin.H{
		"articles": ah.Articles,
	})
}
