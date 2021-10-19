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
	Title string        `json:"title"`
	Index int           `json:"index"`
	Date  string        `json:"date"`
	Body  template.HTML `json:"body"`
}

type ArticleSummary struct {
	Title string `json:"title"`
	Index int    `json:"index"`
	Date  string `json:"date"`
}

type ArticleHandler struct {
	Articles         []Article
	ArticleSummaries []ArticleSummary
}

const dateFormat string = "Jan-02-2006"

func NewArticleHandler() (*ArticleHandler, error) {
	var articles []Article
	var articleSummaries []ArticleSummary

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

		articleSummaries = append(articleSummaries, ArticleSummary{
			Title: nameIndexDate[0],
			Index: index,
			Date:  date.Format(dateFormat),
		})
	}

	sort.Slice(articles, func(i, j int) bool { return articles[i].Index < articles[j].Index })
	sort.Slice(articleSummaries, func(i, j int) bool { return articles[i].Index < articles[j].Index })

	return &ArticleHandler{
		Articles:         articles,
		ArticleSummaries: articleSummaries,
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

func (ah *ArticleHandler) ApiListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, ah.ArticleSummaries)
}

func (ah *ArticleHandler) ApiArticleHandler(c *gin.Context) {
	date := c.Param("article_date")
	for _, article := range ah.Articles {
		if article.Date == date {
			c.JSON(http.StatusOK, article)
			return
		}
	}

	c.Status(http.StatusNotFound)
}
