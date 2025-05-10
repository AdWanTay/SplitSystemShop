package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type TypeService struct {
	repo repositories.TypeRepository
}

func NewTypeService(repo repositories.TypeRepository) *TypeService {
	return &TypeService{repo: repo}
}

func (s *TypeService) GetAll(c context.Context) ([]models.Type, error) {
	all, err := s.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return all, nil
}
