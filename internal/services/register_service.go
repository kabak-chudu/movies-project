package services

import (
	"errors"
	"fmt"
	"movies/internal/models"
	"movies/internal/repository"
)

type RegisterService interface {
	Register(*models.UserCreateRequest) (*models.User, error)
}

type registerService struct {
	register repository.RegisterRepository
}

func NewRegisterService(
	register repository.RegisterRepository,
) RegisterService {
	return &registerService{register: register}
}

func (s *registerService) Register(req *models.UserCreateRequest) (*models.User, error) {
	if err := s.validateCreate(req); err != nil {
		return nil, fmt.Errorf("не удалось зарегаться ошибка: %w", err)
	}

	user := &models.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	if err := s.register.Register(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *registerService) validateCreate(req *models.UserCreateRequest) error {
	if req.Username == nil {
		return errors.New("Username обязательно должен быть")
	}
	if *req.Username == "" {
		return errors.New("Username не должен быть пустым")
	}

	if req.Password == nil {
		return errors.New("Password обязательно должен быть")
	}
	if *req.Password == "" {
		return errors.New("Password не должен быть пустым")
	}
	return nil
}
