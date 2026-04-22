package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)

type MovieFilter struct {
	GenreID *uint
	Year    *int
}

type MovieRepository interface {
	Create(*models.Movie) error
	GetByID(id uint) (*models.Movie, error)
	GetAll(MovieFilter) ([]models.Movie, error)
	Delete(id uint) error
	UpdatePATCH(movie *models.Movie) error
	Exists(id uint) (bool, error)
}

type gormMovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(
	db *gorm.DB,
) MovieRepository {
	return &gormMovieRepository{db: db}
}

func (r *gormMovieRepository) Create(movie *models.Movie) error {
	if movie == nil {
		return nil
	}

	return r.db.Create(movie).Error
}

func (r *gormMovieRepository) GetByID(id uint) (*models.Movie, error) {
	var movie models.Movie

	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *gormMovieRepository) GetAll(filter MovieFilter) ([]models.Movie, error) {
	var movies []models.Movie

	query := r.db.Model(&models.Movie{})

	if filter.GenreID != nil {
		query = query.First(&models.Movie{}).Where("genre_id = ?", filter.GenreID)
	}
	if filter.Year != nil {
		query = query.First(&models.Movie{}).Where("year = ?", filter.Year)
	}
	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *gormMovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}

func (r *gormMovieRepository) UpdatePATCH(movie *models.Movie) error {
	if movie == nil {
		return nil
	}

	return r.db.Save(movie).Error
}

func (r *gormMovieRepository) Exists(id uint) (bool, error) {
	var count int64

	if err := r.db.Model(&models.Movie{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
