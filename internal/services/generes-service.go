package services

import (
	"movies/internal/models"
	"movies/internal/repository"
)

type GenereService interface {
	CreateGenere(req *models.CreateGenreRequest) (*models.Genre, error)
	GetGenerByID(id uint) (*models.Genre, error)
	GetAllGeneres() ([]models.Genre, error)
	DeleteGener(id uint) error
	UpdatePATCHGener(id uint, req *models.UpdateGenreRequest) (*models.Genre, error)
}

type genereService struct {
	genreRepo repository.GenereRepository
}

func NewGenereteService(
	genreRepo repository.GenereRepository,
) GenereService {
	return &genereService{genreRepo: genreRepo}
}
func (s *genereService) CreateGenere(req *models.CreateGenreRequest) (*models.Genre, error) {
	gener := &models.Genre{
		Name: *req.Name,
	}
	if err := s.genreRepo.Create(gener); err != nil {
		return nil, err
	}
	return gener, nil
}
func (s *genereService) GetGenerByID(id uint) (*models.Genre, error) {
	gener, err := s.genreRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return gener, nil
}
func (s *genereService) GetAllGeneres() ([]models.Genre, error) {
	generes, err := s.genreRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return generes, nil
}
func (s *genereService) DeleteGener(id uint) error {
	return s.genreRepo.Delete(id)
}
func (s *genereService) UpdatePATCHGener(id uint, req *models.UpdateGenreRequest) (*models.Genre, error) {
	gener, err := s.genreRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		gener.Name = *req.Name
	}
	if err := s.genreRepo.Update(gener); err != nil {
		return nil, err
	}
	return gener, nil
}
