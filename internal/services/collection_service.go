package services

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"
)

type CollectionService interface {
	CreateCollection(name string, userID uint) (models.Collection, error)
	GetAllCollections() ([]models.Collection, error)
	GetCollectionByID(id uint) (*models.Collection, error)
	AddMovieToCollection(colID, movieID uint) error
	RemoveMovieFromCollection(colID, movieID uint) error

}

type collectionService struct {
	collection	repository.CollectionRepository
	// movieRepo	*repository.MovieRepository
}

func NewCollectionService(r repository.CollectionRepository) CollectionService {
	return &collectionService{collection: r}
}

func (s *collectionService) CreateCollection(name string, userID uint) (models.Collection, error) {
	if name == "" {
		return models.Collection{}, errors.New("Название подборки не может быть пустым")
	}

	newCollection := models.Collection{Name: name, UserID: userID}
	err := s.collection.Create(&newCollection)
	return newCollection, err
}

func (s *collectionService) GetAllCollections() ([]models.Collection, error) {
	return s.collection.GetAll()
}

func (s *collectionService) GetCollectionByID(id uint) (*models.Collection, error) {
	collection, err := s.collection.GetByID(id)
	if err != nil {
		return &models.Collection{}, errors.New("Подборка не найдена")
	}
	return collection, err
}

func (s *collectionService) AddMovieToCollection(colID, movieID uint) error {
	collection, err := s.collection.GetByID(colID)
	if err != nil {
		return errors.New("Подборка не найдена")
	}

	movie := &models.Movie{}
	movie.ID = movieID

	return s.collection.AddMovie(collection, movie)
}

func (s *collectionService) RemoveMovieFromCollection(colID, movieID uint) error {
	collection, err := s.collection.GetByID(colID)
	if err != nil {
		return errors.New("Подборка не найдена")
	}

	movie := &models.Movie{}
	movie.ID = movieID

	return s.collection.RemoveMovie(collection, movie)
}