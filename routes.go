package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
