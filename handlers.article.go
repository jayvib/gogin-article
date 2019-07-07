package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	db DB
}

func (h *handlers) showIndexPage(c *gin.Context) {
	articles, _ := h.db.getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	article, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	render(c, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "applicaton/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		// html
		c.HTML(http.StatusOK, templateName, data)
	}
}
