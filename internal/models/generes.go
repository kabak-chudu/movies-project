package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `json:"name"`
}

type CreateGenreRequest struct {
	Name *string `json:"name" binding:"required"`
}

type UpdateGenreRequest struct {
	Name *string `json:"name"`
}
