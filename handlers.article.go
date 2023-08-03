package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
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
	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   fmt.Sprintf("article-%v", articleId),
			"payload": article,
		},
	)
}
