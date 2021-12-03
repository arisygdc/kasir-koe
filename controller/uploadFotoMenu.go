package controller

import (
	"kasir/config"
	"os"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) UploadFotoMenu(ctx *gin.Context) {
	file, err := ctx.FormFile("menu")
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
			"err":    err,
		})
		return
	}

	file.Filename = "dummy.jpg"
	filePath := config.GetPathImageMenu() + file.Filename
	checkFile, fErr := os.Open(filePath)
	if fErr != nil {
		os.Remove(filePath)
	}
	defer checkFile.Close()

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
			"err":    err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
