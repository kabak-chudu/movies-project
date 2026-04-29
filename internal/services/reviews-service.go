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
)ReviewService{
	return &reviewService{reviewRepo: reviewRepo}
}

func (c *reviewService) CreateReview(req *models.CreateReviewRequest) (*models.Review, error){
	review := &models.Review{
		Rating: *req.Rating,
		Comment: *req.Comment,
		MovieID: *req.MovieID,
	}

	if err := c.reviewRepo.Creat(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (c *reviewService) GetReviewByID(id uint) (*models.Review, error){
	review, err := c.reviewRepo.GetByID(id)
	 if err != nil {
		return nil, err
	 }
	 return review, nil
}

func (c *reviewService) GetReviewAll() ([]models.Review, error){
	review, err := c.reviewRepo.GetAll()
	if err != nil{
		return nil, err
	}
	return review, nil
}

func (d *reviewService) DeleteReview(id uint) error{
	return d.reviewRepo.Delete(id)
}

func (u *reviewService) UpdateReview(id uint, req *models.UpdateReviewRequest) (*models.Review, error){
	review, err := u.reviewRepo.GetByID(id)
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
	if err := u.reviewRepo.Update(review); err != nil {
		return nil, err
	}
	return review, nil
}

