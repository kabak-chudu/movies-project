package repository

import (
	"errors"
	"log/slog"
	"movies/internal/models"

	"gorm.io/gorm"
)

var ErrMovieNotFound error = errors.New("такого фильма по айди не существует")

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
	db     *gorm.DB
	logger *slog.Logger
}

func NewMovieRepository(
	db *gorm.DB,
	logger *slog.Logger,
) MovieRepository {
	return &gormMovieRepository{db: db, logger: logger}
}

func (r *gormMovieRepository) Create(movie *models.Movie) error {
	r.logger.Debug("repo.movie.create")
	if movie == nil {
		return nil
	}
	if err := r.db.Create(movie).Error; err != nil {
		r.logger.Error("repo.movie.create", "error", err.Error())
		return err
	}
	return nil
}

func (r *gormMovieRepository) GetByID(id uint) (*models.Movie, error) {
	r.logger.Debug("repo.movie.GetByID")

	var movie models.Movie

	if err := r.db.Preload("Genre").First(&movie, id).Error; err != nil {
		r.logger.Error("repo.movie.GetByID", "error", err.Error(), "id", id)
		return nil, err
	}

	return &movie, nil
}

func (r *gormMovieRepository) GetAll(filter MovieFilter) ([]models.Movie, error) {
	r.logger.Debug("repo.movie.GetAll")

	var movies []models.Movie

	query := r.db.Model(&models.Movie{}).Preload("Genre")

	if filter.GenreID != nil {
		query = query.Where("genre_id = ?", filter.GenreID)
	}
	if filter.Year != nil {
		query = query.Where("year = ?", filter.Year)
	}
	if err := query.Find(&movies).Error; err != nil {
		r.logger.Error("repo.movie.GetAll", "error", err.Error())
		return nil, err
	}

	return movies, nil
}

func (r *gormMovieRepository) Delete(id uint) error {
	r.logger.Debug("repo.movie.Delete")
	result := r.db.Delete(&models.Movie{}, id)

	if result.Error != nil {
		r.logger.Error("repo.movie.Delete", "error", result.Error, "id", id)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrMovieNotFound
	}

	return nil
}

func (r *gormMovieRepository) UpdatePATCH(movie *models.Movie) error {
	r.logger.Debug("repo.movie.Update")
	if movie == nil {
		return nil
	}
	if err := r.db.Save(movie).Error; err != nil {
		r.logger.Error("repo.movie.Update", "error", err.Error())
		return err
	}
	return nil
}

func (r *gormMovieRepository) Exists(id uint) (bool, error) {
	r.logger.Debug("repo.movie.Exists")
	var count int64

	if err := r.db.Model(&models.Movie{}).Where("id = ?", id).Count(&count).Error; err != nil {
		r.logger.Error("repo.movie.Exists", "error", err.Error(), "id", id)
		return false, err
	}

	return count > 0, nil
}
