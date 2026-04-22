package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title   string `json:"title"`
	Country string `json:"country"`
	Year    int    `json:"int"`
	Genre   *Genre `json:"-"`
	GenreID uint   `json:"genre_id" gorm:"not null;index"`
}
