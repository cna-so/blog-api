package main

import (
	"backend/controllers/authorization"
	controllers "backend/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	user := router.Group("/api/v1/user")
	user.POST("/create", controllers.CreateUser)
	user.POST("/login", controllers.GetUser)
	user.GET("/validate", authorization.RequireLogin, func(context *gin.Context) {
		context.JSON(http.StatusOK, "your token is valid")
	})

	return router
}
