package controllers

import (
	"alterrathree/lib/inmemdb"
	"alterrathree/models"
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
