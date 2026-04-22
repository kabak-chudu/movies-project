package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title   string `json:"title"`
	Country string `json:"country"`
	Year    int    `json:"year"`
	Genre   *Genre `json:"-"`
	GenreID uint `json:"genre_id" gorm:"not null;index"`
}

type CreateMovieRequest struct {
	Title   *string `json:"title"`
	Country *string `json:"country"`
	Year    *int    `json:"year"`
	Genre   *Genre  `json:"-"`
	GenreID *uint `json:"genre_id"`
}

type UpdateMovieRequest struct {
	Title   *string `json:"title"`
	Country *string `json:"country"`
	Year    *int    `json:"year"`
	Genre   *Genre  `json:"-"`
	GenreID *uint `json:"genre_id"`
}
