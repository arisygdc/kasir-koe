package controller

import (
	"kasir/database"
)

type Controller struct {
	Repo database.PostgreSQL
}

func NewController(dbconn database.PostgreSQL) (controller *Controller) {
	controller = new(Controller)
	controller.Repo = dbconn
	return
}
