package controllers

import "alterratwo/lib/inmemdb"

type Controller struct {
	BookDB inmemdb.DB
}

func New() *Controller {
	return &Controller{
		BookDB: *inmemdb.New(),
	}
}
