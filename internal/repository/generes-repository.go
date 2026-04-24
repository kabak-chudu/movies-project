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

func (c *gormGenereRepository) Create(gener *models.Genre) error {
	if gener == nil {
		return nil
	}
	return c.db.Create(gener).Error
}

func (f *gormGenereRepository) FindByID(id uint) (*models.Genre, error) {
	var gener models.Genre
	if err := f.db.First(&gener, id).Error; err != nil {
		return nil, err
	}
	return &gener, nil
}

func (f *gormGenereRepository) FindAll() ([]models.Genre, error) {
	var generes []models.Genre
	if err := f.db.Find(&generes).Error; err != nil {
		return nil, err
	}
	return generes, nil
}

func (d *gormGenereRepository) Delete(id uint) error {
	return d.db.Delete(&models.Genre{}, id).Error
}

func (u *gormGenereRepository) Update(gener *models.Genre) error {
	if gener == nil {
		return nil
	}
	return u.db.Save(&gener).Error
}
