package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	MovieID uint   `json:"movie_id" gorm:"not null;index"`
}
type CreateReviewRequest struct {
	Rating  *int    `json:"rating"`
	Comment *string `json:"comment"`
	MovieID *uint   `json:"movie_id"`
}
type UpdateReviewRequest struct {
	Rating  *int    `json:"rating"`
	Comment *string `json:"comment"`
	MovieID *uint   `json:"movie_id"`
}
