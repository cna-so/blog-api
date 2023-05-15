package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteArticle(ctx *gin.Context) {
	var article models.Article
	var err error
	article.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	deleteArticle, err := article.DeleteArticle()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("article with id %s successfully deleted", deleteArticle),
	})
}
