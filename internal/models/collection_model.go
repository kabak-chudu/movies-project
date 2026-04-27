package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name	string	`json:"name"`
	User	User 	`json:"-"`
	UserID	uint 	`json:"user_id" gorm:"not null;index"`
	Movies	[]Movie	`json:"movies" gorm:"many2many:collection_movies;"`
}

type CollectionCreateRequest struct {
	Name	string	`json:"name" binding:"required"`
	UserID	uint	`json:"user_id" binding:"required"`
}

type CollectionUpdateRequest struct {
	MovieID	uint	`json:"movie_id" binding:"required"`
}
