package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}

func getArticle(c *gin.Context) {

	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	article, err := getArticleByID(articleID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
