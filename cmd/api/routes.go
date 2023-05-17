package main

import (
	ar "backend/controllers/article"
	"backend/controllers/authorization"
	"backend/controllers/category"
	"backend/controllers/fileUploader"
	controllers "backend/controllers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLFiles("./view/index.html", "./view/success.html")

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

	articles := router.Group("/api/v1/articles")
	articles.GET("/all", ar.GetArticles)
	articles.POST("/create", authorization.RequireRole, ar.CreateArticle)
	articles.GET("/find/:id", ar.GetArticleWithId)
	articles.DELETE("/delete/:id", ar.DeleteArticle)

	ctgry := router.Group("/api/v1/category")
	ctgry.GET("/all", category.GetCategories)

	file := router.Group("/api/v1/file")
	file.POST("/upload", fileUploader.FileUploader)
	file.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return router
}
