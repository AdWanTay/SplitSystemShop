package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type SplitSystemService struct {
	repo repositories.SplitSystemRepository
}

func NewSlitSystemService(repo repositories.SplitSystemRepository) *SplitSystemService {
	return &SplitSystemService{repo: repo}
}

func (s *SplitSystemService) GetSplitSystem(c context.Context, id uint) (*models.SplitSystem, error) {
	return s.repo.GetSplitSystem(c, id)
}

func (s *SplitSystemService) GetAllSplitSystems(c context.Context, filters map[string]interface{}) ([]models.SplitSystem, error) {
	return s.repo.GetAllSplitSystems(c, filters)
}

func (s *SplitSystemService) Delete(c context.Context, splitSystemID uint) error {
	return s.repo.Delete(c, splitSystemID)
}

func (s *SplitSystemService) Create(c context.Context, input models.SplitSystem) (*models.SplitSystem, error) {
	split := models.SplitSystem{
		Title:                input.Title,
		ShortDescription:     input.ShortDescription,
		LongDescription:      input.LongDescription,
		BrandID:              input.BrandID,
		TypeID:               input.TypeID,
		Price:                input.Price,
		HasInverter:          input.HasInverter,
		RecommendedArea:      input.RecommendedArea,
		CoolingPower:         input.CoolingPower,
		EnergyClassCoolingID: input.EnergyClassCoolingID,
		EnergyClassHeatingID: input.EnergyClassHeatingID,
		MinNoiseLevel:        input.MinNoiseLevel,
		MaxNoiseLevel:        input.MaxNoiseLevel,
		ExternalWeight:       input.ExternalWeight,
		ExternalWidth:        input.ExternalWidth,
		ExternalHeight:       input.ExternalHeight,
		ExternalDepth:        input.ExternalDepth,
		InternalWeight:       input.InternalWeight,
		InternalWidth:        input.InternalWidth,
		InternalHeight:       input.InternalHeight,
		InternalDepth:        input.InternalDepth,
		ImageURL:             input.ImageURL,
	}

	err := s.repo.Create(c, &split)
	if err != nil {
		return nil, err
	}
	return &split, nil
}

func (s *SplitSystemService) UpdateSplitSystem(c context.Context, id uint, input dto.UpdateSplitSystemRequest) error {
	system, err := s.repo.GetSplitSystem(c, id)
	if err != nil {
		return err
	}

	system = &models.SplitSystem{
		Title:                input.Title,
		ShortDescription:     input.ShortDescription,
		LongDescription:      input.LongDescription,
		BrandID:              input.BrandID,
		TypeID:               input.TypeID,
		Price:                input.Price,
		HasInverter:          input.HasInverter,
		RecommendedArea:      input.RecommendedArea,
		CoolingPower:         input.CoolingPower,
		Modes:                input.Modes,
		EnergyClassCoolingID: input.EnergyClassCoolingID,
		EnergyClassHeatingID: input.EnergyClassHeatingID,
		MinNoiseLevel:        input.MinNoiseLevel,
		MaxNoiseLevel:        input.MaxNoiseLevel,
		ExternalWeight:       input.ExternalWeight,
		ExternalWidth:        input.ExternalWidth,
		ExternalHeight:       input.ExternalHeight,
		ExternalDepth:        input.ExternalDepth,
		InternalWeight:       input.InternalWeight,
		InternalWidth:        input.InternalWidth,
		InternalHeight:       input.InternalHeight,
		InternalDepth:        input.InternalDepth,
		ImageURL:             "",
		AverageRating:        input.InternalWeight,
	}
	if input.ImageURL != nil {
		system.ImageURL = *input.ImageURL
	}
	return s.repo.Update(c, system)
}
