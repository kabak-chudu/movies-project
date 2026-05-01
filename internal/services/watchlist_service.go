package services

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"

	"gorm.io/gorm"
)

var ErrWatchlistNotFound = errors.New("watchlist не найдена")

type WatchlistService interface {
	CreateWatchList(req models.CreateWatchlistRequest) (*models.Watchlist, error)
	GetWatchListByUserID(models.Watchlist) (*models.Watchlist, error)
	AddMovieToWatchList(watchListID uint, req models.WatchlistAddRequest) (*models.Movie, error)
	RemoveWatchlist(watchlistID uint) error
}

type watchlistService struct {
	watchlist repository.WatchlistRepository
	movie     repository.MovieRepository
	user      repository.LoginRepository
}

func NewWatchListService(
	watchlist repository.WatchlistRepository,
	movie repository.MovieRepository,
	user repository.LoginRepository,
) WatchlistService {
	return &watchlistService{
		watchlist: watchlist,
		movie:     movie,
		user:      user,
	}
}

func (s *watchlistService) CreateWatchList(req models.CreateWatchlistRequest) (*models.Watchlist, error) {
	if err := s.validateWatchlistCreate(req); err != nil {
		return nil, err
	}

	newWatchlist := &models.Watchlist{
		Name:   *req.Name,
		UserID: *req.UserID,
	}

	if err := s.watchlist.Create(newWatchlist); err != nil {
		return nil, err
	}

	return newWatchlist, nil
}

func (s *watchlistService) GetWatchListByUserID(w models.Watchlist) (*models.Watchlist, error) {
	watchlist, err := s.watchlist.GetByUserID(w.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usera по такому айди не существует")
		}

		return nil, err
	}

	return watchlist, nil
}

func (s *watchlistService) AddMovieToWatchList(watchListID uint, req models.WatchlistAddRequest) (*models.Movie, error) {
	watchlist, err := s.watchlist.GetByID(watchListID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWatchlistNotFound
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

	for _, m := range watchlist.Movies {
		if m.ID == movie.ID {
			return nil, errors.New("фильм уже есть в этой подборке")
		}
	}

	if err := s.watchlist.AddMovie(watchlist, movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *watchlistService) RemoveWatchlist(watchlistID uint) error {
	if err := s.watchlist.RemoveWatchlistByID(watchlistID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrWatchlistNotFound
		}
		return err
	}
	return nil
}

func (s *watchlistService) validateWatchlistCreate(req models.CreateWatchlistRequest) error {
	if req.Name == nil {
		return errors.New("поле name обязателен")
	}
	if *req.Name == "" {
		return errors.New("поле name не должно быть пустым")
	}
	if req.UserID == nil {
		return errors.New("поле user_id обязателен")
	}
	if *req.UserID <= 0 {
		return errors.New("поле UserID не должно быть равно или меньше 0")
	}
	if _, err := s.user.GetByID(*req.UserID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usera по такому айди не существует")
		}
	}

	return nil
}
