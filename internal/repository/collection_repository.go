package repository

import (
	"gorm.io/gorm"
	"movies/internal/models"
)

type CollectionRepository interface {
	Create(c *models.Collection) error
	GetAll() ([]models.Collection, error)
	GetByID(id uint) (*models.Collection, error)
	AddMovie(c *models.Collection, m *models.Movie) error
	RemoveMovie(c *models.Collection, m *models.Movie) error
}

type collectionRepository struct {
	db *gorm.DB
}

func NewCollectionRepository(db *gorm.DB) CollectionRepository {
	return &collectionRepository{db: db}

}

func (r *collectionRepository) Create(c *models.Collection) error {
	if err := r.db.Create(c).Error; err != nil {
		return err
	}

	return nil
}

func (r *collectionRepository) GetAll() ([]models.Collection, error) {
	var collections []models.Collection

	if err := r.db.Find(&collections).Error; err != nil {
		return nil, err
	}

	return collections, nil
}

func (r *collectionRepository) GetByID(id uint) (*models.Collection, error) {
	var collection models.Collection

	if err := r.db.Preload("Movies").First(&collection, id).Error; err != nil {
		return nil, err
	}

	return &collection, nil
}

func (r *collectionRepository) AddMovie(c *models.Collection, m *models.Movie) error {
	if err := r.db.Model(c).Association("Movies").Append(m); err != nil {
		return err
	}

	return nil
}

func (r *collectionRepository) RemoveMovie(c *models.Collection, m *models.Movie) error {
	if err := r.db.Model(c).Association("Movies").Delete(m); err != nil {
		return err
	}

	return nil
}
