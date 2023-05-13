package controllers

import (
	controllers "backend/controllers/user"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	user := router.Group("/api/v1/user")
	user.POST("/create", controllers.CreateUser)

	return router
}
