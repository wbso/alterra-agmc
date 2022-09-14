package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

// get all data
func (m *Model) GetAllUser() ([]User, error) {
	var users []User
	result := m.db.Find(&users)
	return users, result.Error
}

func (m *Model) GetUserByID(id int) (User, error) {
	var user User
	res := m.db.First(&user, id)
	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func (m *Model) UpdateUser(user *User) (*User, error) {
	res := m.db.Save(user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (m *Model) DeleteUser(user *User) error {
	res := m.db.Delete(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *Model) CreateUser(user User) (User, error) {
	result := m.db.Create(&user)
	return user, result.Error
}
