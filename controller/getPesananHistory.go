package controller

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetPesananHistory(ctx *gin.Context) {
	getTime := time.Now().UTC()
	queryString := ctx.Query("date")
	if queryString != "" {
		tmp, err := time.Parse("2006-1-2", strings.Trim(queryString, " "))
		if err != nil {
			ctx.JSON(403, gin.H{
				"status":  "forbiden",
				"message": "Pastikan anda memasukkan tangal dengan format YYY-MM-DD",
			})
			return
		}

		getTime = tmp
	}
	history, err := ctr.Repo.Queries.GetPesananHistory(ctx, getTime)
	if err != nil {
		if err != nil {
			ctx.JSON(500, gin.H{
				"status":  "internal server error",
				"message": err,
			})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   history,
	})
}
