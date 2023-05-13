package controllers

import (
	"backend/helpers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"error": err,
		})
		return
	}
	user, err = helpers.ValidateUser(user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	statusCode, err := user.CreateUser()
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	}
}
