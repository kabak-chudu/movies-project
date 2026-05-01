package repository

import (
	"errors"
	"movies/internal/models"

	"gorm.io/gorm"
)

type WatchlistRepository interface {
	Create(c *models.Watchlist) error
	GetByUserID(id uint) (*models.Watchlist, error)
	RemoveWatchlistByID(id uint) error
	AddMovie(c *models.Watchlist, m *models.Movie) error
}

type gormWatchlistRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) WatchlistRepository {
	return &gormWatchlistRepository{db: db}

}

func (r *gormWatchlistRepository) Create(watchlist *models.Watchlist) error {
	if err := r.db.Create(watchlist).Error; err != nil {
		return err
	}

	return nil
}

func (r *gormWatchlistRepository) GetByUserID(id uint) (*models.Watchlist, error) {
	var watchlist models.Watchlist

	if err := r.db.Preload("Movies").First(&watchlist, id).Error; err != nil {
		return nil, err
	}

	return &watchlist, nil
}

func (r *gormWatchlistRepository) AddMovie(c *models.Watchlist, m *models.Movie) error {
	if err := r.db.Model(c).Association("Movies").Append(m); err != nil {
		return err
	}

	return nil
}

func (r *gormWatchlistRepository) RemoveWatchlistByID(id uint) error {
	res := r.db.Delete(&models.Watchlist{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("watchlista по такому айди не существует")
	}

	return nil
}
