package services

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"
)

type UserService interface {
	CreateUser(req models.UserCreateRequest) (*models.User, error) 

}

type userService struct {
	user	repository.UserRepository
}

func NewUserService(
	user repository.UserRepository,
	) UserService {
	return &userService{
		user: user,
	}
}

func (s *userService) CreateUser(req models.UserCreateRequest) (*models.User, error) {
	if err := s.validateUserCreate(req); err != nil {
		return nil, err
	}

	newUser := &models.User{
		UserName: req.UserName,
	}

	if err := s.user.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}


func (s *userService) validateUserCreate(req models.UserCreateRequest) error {
	if req.UserName == "" {
		return errors.New("поле name не должно быть пустым")
	}

	return nil
}