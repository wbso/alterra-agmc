package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

// get all data
func (m *Model) GetAllUser() ([]User, error) {
	var users []User
	result := m.DB.Find(&users)
	return users, result.Error
}
