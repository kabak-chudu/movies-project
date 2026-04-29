package repository

import (
	"movies/internal/models"
	"gorm.io/gorm"

)

type UserRepository interface {
	Create(c *models.User) error
}

type userRepository struct {
	db *gorm.DB

}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}

}

func (r *userRepository) Create(c *models.User) error {
	if err := r.db.Create(c).Error; err != nil {
		return err
	}
	
	return nil
}