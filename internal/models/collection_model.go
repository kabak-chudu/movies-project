package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name	string	`json:"name"`
	UserID	uint	`json:"user_id"`
	Movies	[]Movie	`json:"movies" gorm:"many2many:collection_movies;"`
}

type User struct {
	gorm.Model
	Username	string			`json:"username"`
	Collections	[]Collection	`json:"collections"`
}

type CollectionCreate struct {
	Name	string	`json:"name" binding:"required"`
	UserID	uint	`json:"user_id" binding:"required"`
}

type AddMovie struct {
	MovieID	uint	`json:"movie_id" binding:"required"`
}