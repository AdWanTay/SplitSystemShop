package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type FavoritesService struct {
	repo repositories.FavoritesRepository
}

func NewFavoritesService(repo repositories.FavoritesRepository) *FavoritesService {
	return &FavoritesService{repo: repo}
}

func (s *FavoritesService) GetFavorites(c context.Context, userID uint) ([]models.SplitSystem, error) {
	return s.repo.GetUserFavorites(c, userID)
}

func (s *FavoritesService) AddToFavorites(c context.Context, userID, systemID uint) error {
	return s.repo.AddToFavorites(c, userID, systemID)
}

func (s *FavoritesService) RemoveFromFavorites(c context.Context, userID, systemID uint) error {
	return s.repo.RemoveFromFavorites(c, userID, systemID)
}

func (s *FavoritesService) ClearFavorites(c context.Context, userID uint) error {
	return s.repo.ClearFavorites(c, userID)
}
