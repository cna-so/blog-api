package fileUploader

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"os"
)

func FileUploader(ctx *gin.Context) {
	if _, err := os.Stat("public/img"); os.IsNotExist(err) {
		err = os.Mkdir("public/img", os.ModePerm)
		// TODO: handle error
	}
	file, err := ctx.FormFile("image_file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "FILE VALIDATION",
			"message": err.Error(),
		})
	}
	id, _ := uuid.DefaultGenerator.NewV4()
	err = ctx.SaveUploadedFile(file, fmt.Sprintf("public/img/%s-%s", id, file.Filename))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "FILE CREATION",
			"message": err.Error(),
		})
	}
	ctx.JSON(201,
		gin.H{
			"file": fmt.Sprintf("public/img/%s-%s", id, file.Filename),
		})
}
