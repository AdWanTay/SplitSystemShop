package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type EnergyClassRepository interface {
	GetAll(c context.Context) ([]models.EnergyClass, error)
}

func NewEnergyClassRepository(db *gorm.DB) EnergyClassRepository {
	return &energyClassRepository{db: db}
}

type energyClassRepository struct {
	db *gorm.DB
}

func (r energyClassRepository) GetAll(c context.Context) ([]models.EnergyClass, error) {
	var classes []models.EnergyClass
	if err := r.db.WithContext(c).
		Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
