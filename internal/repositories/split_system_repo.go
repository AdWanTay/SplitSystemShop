package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type SplitSystemRepository interface {
	GetSplitSystem(c context.Context, systemID uint) (*models.SplitSystem, error)
	GetAllSplitSystems(c context.Context, filters map[string]interface{}) ([]models.SplitSystem, error)
	DeleteSplitSystem(c context.Context, systemID uint) error
}

type splitSystemRepository struct {
	db *gorm.DB
}

func NewSplitSystemRepository(db *gorm.DB) SplitSystemRepository {
	return &splitSystemRepository{db: db}
}
func (r splitSystemRepository) GetAllSplitSystems(c context.Context, filters map[string]interface{}) ([]models.SplitSystem, error) {
	var splitSystems []models.SplitSystem

	query := r.db.WithContext(c).
		Preload("Brand").
		Preload("Type").
		Preload("Modes").
		Preload("EnergyClassCooling").
		Preload("EnergyClassHeating").
		Model(&models.SplitSystem{})

	// ====== Точные фильтры (IN или =)
	if brands, ok := filters["brand"]; ok {
		query = query.Where("brand_id IN ?", brands)
	}
	if types, ok := filters["type"]; ok {
		query = query.Where("type_id IN ?", types)
	}
	if modes, ok := filters["mode"]; ok {
		// фильтрация по множественным режимам (через join таблицу many2many)
		query = query.Joins("JOIN split_system_modes ssm ON ssm.split_system_id = split_systems.id").
			Where("ssm.mode_id IN ?", modes)
	}
	if hasInverter, ok := filters["has_inverter"]; ok {
		query = query.Where("has_inverter = ?", hasInverter)
	}
	if energyCooling, ok := filters["energy_class_cooling"]; ok {
		query = query.Where("energy_class_cooling_id = ?", energyCooling)
	}
	if energyHeating, ok := filters["energy_class_heating"]; ok {
		query = query.Where("energy_class_heating_id = ?", energyHeating)
	}

	// ====== Диапазоны
	rangeFields := []string{
		"recommended_area",
		"cooling_power",
		"price",
		"min_noise_level",
		"max_noise_level",
		"external_weight", "external_width", "external_height", "external_depth",
		"internal_weight", "internal_width", "internal_height", "internal_depth",
	}

	for _, field := range rangeFields {
		if min_, ok := filters[field+"_min"]; ok {
			query = query.Where(field+" >= ?", min_)
		}
		if max_, ok := filters[field+"_max"]; ok {
			query = query.Where(field+" <= ?", max_)
		}
	}

	// Выполнение запроса
	err := query.Distinct().Find(&splitSystems).Error
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
