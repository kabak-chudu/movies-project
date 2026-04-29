package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)

type GenereRepository interface {
	Create(*models.Genre) error
	FindByID(id uint) (*models.Genre, error)
	FindAll() ([]models.Genre, error)
	Delete(id uint) error
	Update(genre *models.Genre) error
}

type gormGenereRepository struct {
	db *gorm.DB
}

func NewGenereRepository(db *gorm.DB) GenereRepository {
	return &gormGenereRepository{db: db}
}

func (r *gormGenereRepository) Create(gener *models.Genre) error {
	if gener == nil {
		return nil
	}
	return r.db.Create(gener).Error
}

func (r *gormGenereRepository) FindByID(id uint) (*models.Genre, error) {
	var gener models.Genre
	if err := r.db.First(&gener, id).Error; err != nil {
		return nil, err
	}
	return &gener, nil
}

func (r *gormGenereRepository) FindAll() ([]models.Genre, error) {
	var generes []models.Genre
	if err := r.db.Find(&generes).Error; err != nil {
		return nil, err
	}
	return generes, nil
}

func (r *gormGenereRepository) Delete(id uint) error {
	return r.db.Delete(&models.Genre{}, id).Error
}

func (r *gormGenereRepository) Update(gener *models.Genre) error {
	if gener == nil {
		return nil
	}
	return r.db.Save(&gener).Error
}
