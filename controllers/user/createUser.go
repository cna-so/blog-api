package controllers

import (
	"backend/helpers"
	HashPassword "backend/helpers/hash"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	user.Password, err = HashPassword.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create hash password",
		})
	}
	statusCode, err := user.CreateUser()

	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
	}
}
