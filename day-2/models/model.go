package models

import (
	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

func New(db *gorm.DB) (*Model, error) {
	db.AutoMigrate(&User{})
	return &Model{
		DB: db,
	}, nil
}
