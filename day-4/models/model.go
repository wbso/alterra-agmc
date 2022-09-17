package models

import (
	"gorm.io/gorm"
)

type Model struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Model, error) {
	db.AutoMigrate(&User{})
	return &Model{
		db: db,
	}, nil
}
