package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/viper"
)

type Article struct {
	Title string
	Index int
	Date  string
	Body  template.HTML
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

		articleBody, err := ioutil.ReadFile(filepath.Join(viper.GetString("rambler_articles"), article.Name()))
		if err != nil {
			return nil, err
		}

		mdBytes := blackfriday.Run(articleBody)
		articleTmpl := template.HTML(mdBytes)
		articles = append(articles, Article{
			Title: nameIndexDate[0],
			Index: index,
			Date:  date.Format(dateFormat),
			Body:  articleTmpl,
		})
	}

	sort.Slice(articles, func(i, j int) bool { return articles[i].Index < articles[j].Index })

	return &ArticleHandler{
		Articles: articles,
	}, nil
}

func (ah *ArticleHandler) ListHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "articles.template.html", gin.H{
		"articles": ah.Articles,
	})
}

func (ah *ArticleHandler) Handler(c *gin.Context) {
	date := c.Param("article_date")
	for _, article := range ah.Articles {
		if article.Date == date {
			c.HTML(http.StatusOK, "article.template.html", article)
			return
		}
	}

	c.Redirect(http.StatusSeeOther, "/rambler")
}
