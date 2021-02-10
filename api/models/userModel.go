package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id    int
	Name  string
	Token string
}

func (user *User) CreateUser(db *gorm.DB) error {
	result := db.Create(&user)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (user *User) UpdateUser(db *gorm.DB, new_name string) error {
	result := db.Model(&user).Update("name", new_name)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func GetUser(db *gorm.DB, token string) (*User, error) {
	var user User
	result := db.Where("token = ?", token).First(&user)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUser(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)

	if err := result.Error; err != nil {
		return nil, err
	}

	return users, nil
}
