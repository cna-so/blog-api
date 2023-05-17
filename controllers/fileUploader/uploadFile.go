package fileUploader

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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
	err = ctx.SaveUploadedFile(file, "public/img/"+time.Now().Format("2006-01-02")+file.Filename)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "FILE CREATION",
			"message": err.Error(),
		})
	}
	ctx.HTML(201, "success.html", nil)
	return
}
