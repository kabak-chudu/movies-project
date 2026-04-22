package services

import (
	"movies/internal/models"
	"movies/internal/repository"
)

type ReviewService interface {
	CreateReview(*models.CreateReviewRequest) (*models.Review, error)
	GetReviewByID(id uint) (*models.Review, error)
	GetReviewAll() ([]models.Review, error)
	DeleteReview(id uint) error
	UpdateReview(id uint, req *models.UpdateReviewRequest) (*models.Review, error)
}

type reviewService struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewService(
	reviewRepo repository.ReviewRepository,
) ReviewService {
	return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) CreateReview(req *models.CreateReviewRequest) (*models.Review, error) {
	review := &models.Review{
		Rating:  *req.Rating,
		Comment: *req.Comment,
		MovieID: *req.MovieID,
	}

	if err := s.reviewRepo.Creat(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (s *reviewService) GetReviewByID(id uint) (*models.Review, error) {
	review, err := s.reviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *reviewService) GetReviewAll() ([]models.Review, error) {
	review, err := s.reviewRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *reviewService) DeleteReview(id uint) error {
	return s.reviewRepo.Delete(id)
}

func (s *reviewService) UpdateReview(id uint, req *models.UpdateReviewRequest) (*models.Review, error) {
	review, err := s.reviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Rating != nil {
		review.Rating = *req.Rating
	}
	if req.Comment != nil {
		review.Comment = *req.Comment
	}
	if req.MovieID != nil {
		review.MovieID = *req.MovieID
	}
	if err := s.reviewRepo.Update(review); err != nil {
		return nil, err
	}
	return review, nil
}
