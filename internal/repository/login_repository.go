package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Login(*models.User) error
	GetByID(id uint) (*models.User, error)
}

type gormLoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(
	db *gorm.DB,
) LoginRepository {
	return &gormLoginRepository{db: db}
}

func (r *gormLoginRepository) Login(user *models.User) error {
	err := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error

	return err
}

func (r *gormLoginRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
