package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type ModeService struct {
	repo repositories.ModeRepository
}

func NewModeService(repo repositories.ModeRepository) *ModeService {
	return &ModeService{repo: repo}
}

func (s *ModeService) GetAll(c context.Context) ([]models.Mode, error) {
	all, err := s.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return all, nil
}
