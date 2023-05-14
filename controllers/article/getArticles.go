package controllers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(ctx *gin.Context) {
	var article models.Article
	articles, err := article.GetArticles()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})

}
