package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName	string	`json:"user_name"`
}

type UserCreateRequest struct {
	UserName	string	`json:"user_name" binding:"required"`
}