package services

import (
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
	item, err := s.repo.GetSplitSystem(c, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *SplitSystemService) GetAllSplitSystems(c context.Context) (*[]models.SplitSystem, error) {
	items, err := s.repo.GetAllSplitSystems(c)
	if err != nil {
		return nil, err
	}

	return items, nil
}
