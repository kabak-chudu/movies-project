package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name	string	`json:"name"`
	User	User 	`json:"-"`
	UserID	uint 	`json:"user_id" gorm:"not null;index"`
	Movies	[]Movie	`json:"movies" gorm:"many2many:collection_movies;"`
}

type User struct {
	gorm.Model
	Username	string			`json:"username"`
	Collections	[]Collection	`json:"collections"`
}

type CollectionCreateRequest struct {
	Name	string	`json:"name" binding:"required"`
	UserID	uint	`json:"user_id" binding:"required"`
}

type CollectionAddRequest struct {
	MovieID	uint	`json:"movie_id" binding:"required"`
}

