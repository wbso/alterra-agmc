package controllers

import (
	"alterratwo/lib/inmemdb"
	"alterratwo/models"
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
