package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)
type ReviewRepository interface {
	Creat(*models.Review) error
	GetByID(id uint) (*models.Review, error)
	GetAll() ([]models.Review, error)
	Delete(id uint) error
	Update(review *models.Review) error
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Creat(review *models.Review) error {
	if review == nil {
		return nil
	}
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review
	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *gormReviewRepository) GetAll() ([]models.Review, error) {
	var reviews []models.Review
	if err := r.db.Find(&reviews).Error; err != nil{
		return nil, err
	}
	return reviews, nil
}

func (r *gormReviewRepository) Delete(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}

func (r *gormReviewRepository) Update(review *models.Review) error {
	if review == nil {
		return nil
	}
	return r.db.Save(&review).Error
}