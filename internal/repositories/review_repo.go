package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"fmt"
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
	var count int64
	err := r.db.WithContext(c).
		Model(&models.Review{}).
		Where("user_id = ? AND split_system_id = ?", review.UserID, review.SplitSystemID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("отзыв уже существует")
	}
	return r.db.WithContext(c).Create(review).Error
}

func (r reviewRepository) GetSplitSystemReviews(c context.Context, splitSystemID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM reviews WHERE split_system_id = ?",
		splitSystemID).Error
}
