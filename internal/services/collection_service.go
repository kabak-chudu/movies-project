package services

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"

	"gorm.io/gorm"
)

var ErrGroupNotFound = errors.New("Коллекция не найдена")

type CollectionService interface {
	CreateCollection(req models.CollectionCreateRequest) (*models.Collection, error) 
	GetAllCollections() ([]models.Collection, error)
	GetCollectionByID(id uint) (*models.Collection, error)
	AddMovieToCollection(collID, movieID uint, req models.CollectionAddRequest) (*models.Collection, error)
	RemoveMovieFromCollection(collID, movieID uint) error

}

type collectionService struct {
	collection	repository.CollectionRepository
	movie		repository.MovieRepository
}

func NewCollectionService(
	collection repository.CollectionRepository,
	movie repository.MovieRepository,
	) CollectionService {
	return &collectionService{
		collection: collection,
		movie: movie,
	}
}

func (s *collectionService) CreateCollection(req models.CollectionCreateRequest) (*models.Collection, error) {
	if err := s.validateCollectionCreate(req); err != nil {
		return nil, err
	}

	newCollection := &models.Collection{
		Name: req.Name,
		UserID: req.UserID,
	}

	if err := s.collection.Create(newCollection); err != nil {
		return nil, err
	}

	return newCollection, nil
}

func (s *collectionService) GetAllCollections() ([]models.Collection, error) {
	return s.collection.GetAll()
}

func (s *collectionService) GetCollectionByID(id uint) (*models.Collection, error) {
	collection, err := s.collection.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGroupNotFound
		}

		return nil, err
	}

	return collection, err
}

func (s *collectionService) AddMovieToCollection(collID, movieID uint, req models.CollectionAddRequest) (*models.Collection, error) {
	collection, err := s.collection.GetByID(collID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	movie, err := s.movie.GetByID(movieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	if err := s.collection.AddMovie(collection, movie); err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *collectionService) RemoveMovieFromCollection(collID, movieID uint) error {
	collection, err := s.collection.GetByID(collID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGroupNotFound
		}

		return err
	}

	movie, err := s.movie.GetByID(movieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		return err
	}

	return s.collection.RemoveMovie(collection, movie)
}

func (s *collectionService) validateCollectionCreate(req models.CollectionCreateRequest) error {
	if req.Name == "" {
		return errors.New("Поле name не должно быть пустым")
	}

	if req.UserID <= 0 {
		return errors.New("Поле UserID не должно быть равно или меньше 0")
	}

	return nil
}