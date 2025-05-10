package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type BrandService struct {
	repo repositories.BrandRepository
}

func NewBrandService(repo repositories.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) GetAll(c context.Context) ([]models.Brand, error) {
	all, err := s.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return all, nil
}
