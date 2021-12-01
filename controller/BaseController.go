package controller

import "kasir/database/postgres"

type Controller struct {
	Repo *postgres.Queries
}

func NewController(dbconn *postgres.Queries) (controller *Controller) {
	controller = new(Controller)
	controller.Repo = dbconn
	return
}
