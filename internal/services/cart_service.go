package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type CartService struct {
	repo        repositories.CartRepository
	userService *UserService
}

func NewCartService(repo repositories.CartRepository, userService *UserService) *CartService {
	return &CartService{repo: repo, userService: userService}
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

func (s *CartService) LoadCartModuleData(c context.Context, userID uint) (*dto.CartModuleResponse, error) {
	cart, err := s.userService.GetCart(c, userID)
	if err != nil {
		return nil, err
	}

	favorites, err := s.userService.GetFavorites(c, userID)
	if err != nil {
		return nil, err
	}

	response := dto.NewCartModuleResponse(cart, favorites)
	return &response, nil
}
