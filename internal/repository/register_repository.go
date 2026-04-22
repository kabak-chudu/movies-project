package repository

import (
	"errors"
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
	if err := r.checkUsername(user); err != nil {
		return err
	}

	return r.db.Create(&user).Error
}

func (r *gormRegisterRepository) checkUsername(req *models.User) error {
	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("пользователь с таким Username уже существует")
	}

	return nil
}
