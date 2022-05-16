package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
