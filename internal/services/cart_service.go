package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type CartService struct {
	repo repositories.CartRepository
}

func NewCartService(repo repositories.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) GetCart(c context.Context, userID uint) ([]models.SplitSystem, error) {
	return s.repo.GetUserCart(c, userID)
}

func (s *CartService) AddToCart(c context.Context, userID, systemID uint) error {
	return s.repo.AddToCart(c, userID, systemID)
}

func (s *CartService) RemoveFromCart(c context.Context, userID, systemID uint) error {
	return s.repo.RemoveFromCart(c, userID, systemID)
}

func (s *CartService) ClearCart(c context.Context, userID uint) error {
	return s.repo.ClearCart(c, userID)
}
