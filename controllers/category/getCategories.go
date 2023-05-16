package category

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(ctx *gin.Context) {
	var category models.CategoryModel

	categories, err := category.GetCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "retrieving error",
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
