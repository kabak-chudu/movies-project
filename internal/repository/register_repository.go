package repository

import (
	"movies/internal/models"

	"gorm.io/gorm"
)

type RegisterRepository interface {
	Register(*models.User) error
}

type gormRegisterRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(
	db *gorm.DB,
) RegisterRepository {
	return &gormRegisterRepository{db: db}
}

func (r *gormRegisterRepository) Register(user *models.User) error {
	if user == nil {
		return nil
	}

	return r.db.Create(&user).Error
}
