package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type TypeRepository interface {
	GetAll(c context.Context) ([]models.Type, error)
}

func NewTypeRepository(db *gorm.DB) TypeRepository {
	return &typeRepository{db: db}
}

type typeRepository struct {
	db *gorm.DB
}

func (r typeRepository) GetAll(c context.Context) ([]models.Type, error) {
	var types []models.Type
	if err := r.db.WithContext(c).
		Find(&types).Error; err != nil {
		return nil, err
	}
	return types, nil
}
