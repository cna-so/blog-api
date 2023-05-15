package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticleWithId(ctx *gin.Context) {
	var article models.Article
	var err error
	article.ID, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	article, err = article.GetArticleWithId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}
