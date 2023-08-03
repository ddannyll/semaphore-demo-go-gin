package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func render (c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json": 
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, 
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

func getArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	article, err := getArticleById(articleId)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// valid
	render(c, 
		gin.H{
			"title":   fmt.Sprintf("article-%v", articleId),
			"payload": article,
		},
		"article.html",
	)
}
