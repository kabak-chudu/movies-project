package services

import (
	"errors"
	"fmt"
	"log/slog"
	"movies/internal/models"
	"movies/internal/repository"

	"gorm.io/gorm"
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
	movie  repository.MovieRepository
	logger *slog.Logger
	genre  repository.GenereRepository
}

func NewMovieService(
	movie repository.MovieRepository,
	genre repository.GenereRepository,
	logger *slog.Logger,
) MovieService {
	return &movieService{movie: movie, logger: logger, genre: genre}
}

func (s *movieService) CreateMovie(req *models.CreateMovieRequest) (*models.Movie, error) {
	if err := s.validCreate(req); err != nil {
		s.logger.Warn("failed create validator",
			"layer", "service",
			"reason", err.Error(),
		)
		return nil, fmt.Errorf("не удалось создать фильм ошибка: %w", err)
	}

	movie := &models.Movie{
		Title:   *req.Title,
		Country: *req.Country,
		Year:    *req.Year,
		GenreID: *req.GenreID,
	}

	if err := s.movie.Create(movie); err != nil {
		s.logger.Error("failed create movie",
			"error", err.Error(),
			"layer", "service",
		)
		return nil, err
	}

	s.logger.Info("succsecful created",
		"layer", "service",
		"movie_id", movie.ID,
		"movie_title", movie.Title,
		"country", movie.Country,
		"year", movie.Year,
	)
	return movie, nil
}

func (s *movieService) GetAllByFilter(filter repository.MovieFilter) ([]models.Movie, error) {
	movies, err := s.movie.GetAll(filter)
	if err != nil {
		s.logger.Error("failed get movies",
			"layer", "service",
			"error", err.Error(),
		)
		return nil, err
	}
	s.logger.Info("succesful finded movies", "layer", "service")
	return movies, nil
}

func (s *movieService) GetMovieByID(id uint) (*models.Movie, error) {
	movie, err := s.movie.GetByID(id)
	if err != nil {
		s.logger.Error("failed get movie",
			"layer", "service",
			"error", err.Error(),
			"movie_id", id,
		)
		return nil, err
	}

	s.logger.Info("finded movie by id",
		"layer", "service",
		"movie_id", id,
	)
	return movie, nil
}

func (s *movieService) DeleteMovie(id uint) error {
	if err := s.movie.Delete(id); err != nil {
		s.logger.Error("failed delete movie",
			"layer", "service",
			"error", err.Error(),
			"id", id,
		)
		return err
	}
	s.logger.Info("succesful movie deleted",
		"layer", "service",
		"id", id,
	)
	return nil
}

func (s *movieService) UpdatePATCHMovie(id uint, req *models.UpdateMovieRequest) (*models.Movie, error) {
	exists, err := s.movie.Exists(id)
	if err != nil {
		s.logger.Warn("fail find movie by id",
			"layer", "service",
			"reason", err.Error(),
			"movie_id", id,
		)
		return nil, err
	}
	if !exists {
		s.logger.Warn("not found movie",
			"layer", "service",
			"movie_id", id,
		)
		return nil, ErrMovieNotFound
	}

	movie, err := s.movie.GetByID(id)
	if err != nil {
		s.logger.Error("fail get movie by id",
			"layer", "service",
			"error", err.Error(),
			"movie_id", id,
		)
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
			s.logger.Warn("validate failed", "layer", "service", "error", "если хочешь поменять год фильма то введи его не меньше 1900", "year", *req.Year)
			return nil, errors.New("если хочешь поменять год фильма то введи его не меньше 1900")
		}
		movie.Year = *req.Year
	}

	if req.GenreID != nil {
		_, err := s.genre.FindByID(*req.GenreID)
		if err != nil {
			if errors.Is(err, ErrGenreNotFound) {
				return nil, err
			}
			return nil, err
		}
		movie.GenreID = *req.GenreID
	}

	if err := s.movie.UpdatePATCH(movie); err != nil {
		s.logger.Error("fail update movie by id", "layer", "service", "error", err.Error(), "movie_id", id)
		return nil, err
	}

	s.logger.Info("succesful update", "layer", "service", "movie_id", id)
	return movie, nil
}

func (s *movieService) validCreate(req *models.CreateMovieRequest) error {
	if req.GenreID == nil {
		return errors.New("для создания фильма надо обязательно указать айди к жанру (genre_id)")
	}
	_, err := s.genre.FindByID(*req.GenreID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("такого жанра по айди не существует")
		}
		return err
	}

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
