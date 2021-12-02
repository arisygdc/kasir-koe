package controller

import "github.com/gin-gonic/gin"

func (ctr Controller) GetMenu(ctx *gin.Context) {
	menu, err := ctr.Repo.Queries.GetMenuAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   menu,
	})
}
