package controllers

import (
	HashPassword "backend/helpers/hash"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	log.Println(err)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "validation error",
			"message": "should enter (email , password ، first_name , last_name)",
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
