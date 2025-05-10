package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type ModeRepository interface {
	GetAll(c context.Context) ([]models.Mode, error)
}

func NewModeRepository(db *gorm.DB) ModeRepository {
	return &modeRepository{db: db}
}

type modeRepository struct {
	db *gorm.DB
}

func (r modeRepository) GetAll(c context.Context) ([]models.Mode, error) {
	var modes []models.Mode
	if err := r.db.WithContext(c).
		Find(&modes).Error; err != nil {
		return nil, err
	}
	return modes, nil
}
