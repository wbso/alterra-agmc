package controllers

import (
	"alterrafour/lib/inmemdb"
	"alterrafour/models"
)

type Controller struct {
	BookDB inmemdb.DB
	Model  *models.Model
}

func New(m *models.Model) *Controller {
	return &Controller{
		BookDB: *inmemdb.New(),
		Model:  m,
	}
}
