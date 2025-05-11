package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type SplitSystemRepository interface {
	GetSplitSystem(c context.Context, systemID uint) (*models.SplitSystem, error)
	GetAllSplitSystems(c context.Context) ([]models.SplitSystem, error)
	DeleteSplitSystem(c context.Context, systemID uint) error
}

type splitSystemRepository struct {
	db *gorm.DB
}

func NewSplitSystemRepository(db *gorm.DB) SplitSystemRepository {
	return &splitSystemRepository{db: db}
}

func (r splitSystemRepository) GetAllSplitSystems(c context.Context) ([]models.SplitSystem, error) {
	var splitSystems []models.SplitSystem
	err := r.db.WithContext(c).
		Preload("Brand").
		Preload("Type").
		Preload("Modes").
		Preload("EnergyClassCooling").
		Preload("EnergyClassHeating").Find(&splitSystems).Error
	if err != nil {
		return nil, err
	}

	return splitSystems, nil

}

func (r splitSystemRepository) GetSplitSystem(c context.Context, systemID uint) (*models.SplitSystem, error) {
	var splitSystem models.SplitSystem
	err := r.db.WithContext(c).
		Preload("Brand").
		Preload("Type").
		Preload("Modes").
		Preload("EnergyClassCooling").
		Preload("EnergyClassHeating").
		First(&splitSystem, systemID).Error
	if err != nil {
		return nil, err
	}

	return &splitSystem, nil
}

func (r splitSystemRepository) DeleteSplitSystem(c context.Context, systemID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM split_systems WHERE id = ?", systemID).Error
}
