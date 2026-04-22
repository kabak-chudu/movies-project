package services

import "movies/internal/models"

type MovieService interface {
	CreateMovie() (*models.Movie, error)
}