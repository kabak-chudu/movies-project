package models

import "gorm.io/gorm"

type Watchlist struct {
	gorm.Model
	Name   string  `json:"name"`
	User   *User   `json:"-"`
	UserID uint    `json:"user_id" gorm:"not null;index"`
	Movies []Movie `json:"movies" gorm:"many2many:collection_movies;"`
}

type CreateWatchlistRequest struct {
	Name   *string `json:"name"`
	User   *User   `json:"-"`
	UserID *uint   `json:"user_id" gorm:"not null;index"`
}

type WatchlistAddRequest struct {
	Movie   *Movie `json:"-"`
	MovieID *uint   `json:"movie_id" gorm:"not null;index"`
}
