package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(c context.Context, review *models.Review) error
	GetSplitSystemReviews(c context.Context, splitSystemID uint) error
}
type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r reviewRepository) Create(c context.Context, review *models.Review) error {
	return r.db.WithContext(c).Create(review).Error
}

func (r reviewRepository) GetSplitSystemReviews(c context.Context, splitSystemID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM reviews WHERE split_system_id = ?",
		splitSystemID).Error
}
