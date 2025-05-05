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
	return s.repo.GetSplitSystem(c, id)
}

func (s *SplitSystemService) GetAllSplitSystems(c context.Context) (*[]models.SplitSystem, error) {
	return s.repo.GetAllSplitSystems(c)
}

func (s *SplitSystemService) DeleteSplitSystem(c context.Context, splitSystemID uint) error {
	return s.repo.DeleteSplitSystem(c, splitSystemID)
}
