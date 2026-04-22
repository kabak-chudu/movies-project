package services

import (
	"errors"
	"fmt"
	"movies/internal/models"
	"movies/internal/repository"
)

var ErrMovieNotFound error = errors.New("такого фильма по айди не существует")
var ErrGenreNotFound error = errors.New("такого жанра по айди не существует")

type MovieService interface {
	CreateMovie(*models.CreateMovieRequest) (*models.Movie, error)
	GetAllByFilter(filter repository.MovieFilter) ([]models.Movie, error)
	GetMovieByID(id uint) (*models.Movie, error)
	DeleteMovie(id uint) error
	UpdatePATCHMovie(id uint, req *models.UpdateMovieRequest) (*models.Movie, error)
}

type movieService struct {
	movie repository.MovieRepository
	// genre repository.GenreRepository
}

func NewMovieService(
	movie repository.MovieRepository,
	// genre repository.GenreRepository,
) MovieService {
	return &movieService{movie: movie}
}

func (s *movieService) CreateMovie(req *models.CreateMovieRequest) (*models.Movie, error) {

	if err := s.validCreate(req); err != nil {
		return nil, fmt.Errorf("не удалось создать фильм ошибка: %w", err)
	}

	movie := &models.Movie{
		Title:   *req.Title,
		Country: *req.Country,
		Year:    *req.Year,
		GenreID: *req.GenreID,
	}

	if err := s.movie.Create(movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieService) GetAllByFilter(filter repository.MovieFilter) ([]models.Movie, error) {
	movies, err := s.movie.GetAll(filter)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *movieService) GetMovieByID(id uint) (*models.Movie, error) {
	movie, err := s.movie.GetByID(id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieService) DeleteMovie(id uint) error {
	return s.movie.Delete(id)
}

func (s *movieService) UpdatePATCHMovie(id uint, req *models.UpdateMovieRequest) (*models.Movie, error) {
	exists, err := s.movie.Exists(id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrMovieNotFound
	}

	movie, err := s.movie.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Country != nil {
		if *req.Country != "" {
			movie.Country = *req.Country
		}
	}

	if req.Title != nil {
		if *req.Title != "" {
			movie.Title = *req.Title
		}
	}

	if req.Year != nil {
		if *req.Year < 1900 {
			return nil, errors.New("если хочешь поменять год фильма то введи его не меньше 1900")
		}
		movie.Year = *req.Year
	}

	// if req.GenreID != nil {
	// 	exists, err := s.genre.Exists(*req.GenreID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if !exists {
	// 		return nil, ErrGenreNotFound
	// 	}
	// 	movie.GenreID = *req.GenreID
	// }

	if err := s.movie.UpdatePATCH(movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieService) validCreate(req *models.CreateMovieRequest) error {
	// if req.GenreID == nil {
	// 	return errors.New("для создания фильма надо обязательно указать айди к жанру (genre_id)")
	// }
	// exists, err := s.genre.Exists(*req.GenreID)
	// if err != nil {
	// 	return err
	// }
	// if !exists {
	// 	return ErrGenreNotFound
	// }

	if req.Country == nil {
		return errors.New("для создания фильма надо обязательно указать страну производства (country)")
	}
	if *req.Country == "" {
		return errors.New("страна производства не может быть пустой")
	}

	if req.Title == nil {
		return errors.New("для создания фильма надо обязательно указать название (title)")
	}
	if *req.Title == "" {
		return errors.New("название не может быть пустым")
	}

	if req.Year == nil {
		return errors.New("для создания фильма надо обязательно указать год (year)")
	}
	if *req.Year < 1900 {
		return errors.New("год должен быть не меньше 1900")
	}

	return nil
}
