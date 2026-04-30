package services

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"
	"gorm.io/gorm"
)

var ErrCollectionNotFound = errors.New("коллекция не найдена")

type CollectionService interface {
	CreateCollection(req models.CollectionCreateRequest) (*models.Collection, error) 
	GetAllCollections() ([]models.Collection, error)
	GetCollectionByID(id uint) (*models.Collection, error)
	AddMovieToCollection(collID uint, req models.CollectionAddRequest) (*models.Collection, error)
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
	//	UserID: req.UserID,
	}

	if err := s.collection.Create(newCollection); err != nil {
		return nil, err
	}

	return newCollection, nil
}

func (s *collectionService) GetAllCollections() ([]models.Collection, error) {
	collection, err := s.collection.GetAll(); if err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *collectionService) GetCollectionByID(id uint) (*models.Collection, error) {
	collection, err := s.collection.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCollectionNotFound
		}

		return nil, err
	}

	return collection, nil
}

func (s *collectionService) AddMovieToCollection(collID uint, req models.CollectionAddRequest) (*models.Collection, error) {
	collection, err := s.collection.GetByID(collID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCollectionNotFound
		}
		return nil, err
	}

	movie, err := s.movie.GetByID(req.MovieID)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrMovieNotFound
        }
        return nil, err
    }

	for _, m := range collection.Movies {
        if m.ID == movie.ID {
            return nil, errors.New("фильм уже есть в этой подборке")
        }
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
			return ErrCollectionNotFound
		}
		return err
	}
	
	var found bool
	for _, m := range collection.Movies {
		if m.ID == movieID {
			found = true
			break
		}
	}

	if !found {
		return ErrMovieNotFound
	}

	movie := &models.Movie{}
	movie.ID = movieID

	return s.collection.RemoveMovie(collection, movie)
}

func (s *collectionService) validateCollectionCreate(req models.CollectionCreateRequest) error {
	if req.Name == "" {
		return errors.New("поле name не должно быть пустым")
	}

	//	if req.UserID <= 0 {
	//		return errors.New("поле UserID не должно быть равно или меньше 0")
	//	}

	return nil
}