package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var body models.Article
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Validation Error",
			"message": "(title , description , category_id , creator) are required and (category_id , creator) must be integer",
		})
		return
	}
	articleId, err := body.InsertArticle()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Creation Error",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("article with id %s succssfully created", articleId),
	})
}
