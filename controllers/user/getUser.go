package controllers

import (
	"backend/controllers/authorization"
	"backend/helpers/hash"
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	var body models.UserLogin
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usr, ok := body.GetUser()
	if ok != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}
	if HashPassword.ComparePassWithHash(usr.Password, body.Password) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "your password is invalid",
		})
		return
	}
	token, err := authorization.GenerateJwtToken(usr.Email, usr.Role)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"email":        usr.Email,
		"access_token": token,
	})
}
