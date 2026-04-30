package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Login(*models.Login) error
}

type gormLoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(
	db *gorm.DB,
) LoginRepository {
	return &gormLoginRepository{db: db}
}

func (r *gormLoginRepository) Login(user *models.Login) error {
	if user == nil {
		return nil
	}
	if err := r.check(user); err != nil {
		return err
	}

	return nil
}

func (r *gormLoginRepository) check(req *models.Login) error {
	var user models.User

	err := r.db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error
	if err != nil {
		return err
	}

	return nil
}
