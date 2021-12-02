package controller

import "github.com/gin-gonic/gin"

func (ctr Controller) GetPesananHistory(ctx *gin.Context) {
	history, err := ctr.Repo.Queries.GetPesananHistory(ctx)
	if err != nil {
		if err != nil {
			ctx.JSON(500, gin.H{
				"status": "internal server error",
			})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   history,
	})
}
