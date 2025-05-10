package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type BrandRepository interface {
	GetAll(c context.Context) ([]models.Brand, error)
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}

type brandRepository struct {
	db *gorm.DB
}

func (r brandRepository) GetAll(c context.Context) ([]models.Brand, error) {
	var brands []models.Brand
	if err := r.db.WithContext(c).
		Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}
