package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
	"time"
)

type ReviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}
func (s *ReviewService) Create(c context.Context, review dto.NewReviewRequest, userID uint) (*models.Review, error) {
	return s.repo.Create(c, &models.Review{
		SplitSystemID: review.SplitSystemID,
		UserID:        userID,
		Rating:        review.Rating,
		Comment:       review.Comment,
		CreatedAt:     time.Time{},
	})
}
