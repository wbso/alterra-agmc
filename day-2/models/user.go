package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() (*Model, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Book{})
	return &Model{
		DB: db,
	}, nil
}
