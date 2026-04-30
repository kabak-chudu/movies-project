package services

import (
	"errors"
	"fmt"
	"movies/internal/models"
	"movies/internal/repository"
)

type LoginService interface {
	Login(*models.Login) (*models.User, error)
}

type loginService struct {
	login repository.LoginRepository
}

func NewLoginService(
	login repository.LoginRepository,
) LoginService {
	return &loginService{login: login}
}

func (s *loginService) Login(req *models.Login) (*models.User, error) {
	if err := s.validValues(req); err != nil {
		return nil, fmt.Errorf("не удалось залогиниться ошибка: %w", err)
	}

	user := &models.User{
		Username: *req.Username,
		Password: *req.Password,
	}

	if err := s.login.Login(user); err != nil {
		return nil, errors.New("username или password неверные")
	}

	return user, nil
}

func (s *loginService) validValues(req *models.Login) error {
	if req.Username == nil {
		return errors.New("username обязательно должен быть")
	}
	if *req.Username == "" {
		return errors.New("username не должен быть пустым")
	}

	if req.Password == nil {
		return errors.New("password обязательно должен быть")
	}
	if *req.Password == "" {
		return errors.New("password не должен быть пустым")
	}
	return nil
}
