package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type EnergyClassService struct {
	repo repositories.EnergyClassRepository
}

func NewEnergyClassService(repo repositories.EnergyClassRepository) *EnergyClassService {
	return &EnergyClassService{repo: repo}
}

func (s *EnergyClassService) GetAll(c context.Context) ([]models.EnergyClass, error) {
	all, err := s.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return all, nil
}
