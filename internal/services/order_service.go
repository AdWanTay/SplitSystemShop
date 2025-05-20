package services

import (
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
	"fmt"
	"time"
)

type OrderService struct {
	repo     repositories.OrderRepository
	userRepo repositories.UserRepository
}

func NewOrderService(repo repositories.OrderRepository, userRepo repositories.UserRepository) *OrderService {
	return &OrderService{repo: repo, userRepo: userRepo}
}

func (s *OrderService) CreateOrderByUserCart(c context.Context, userID uint) (*models.Order, error) {
	cart, err := s.userRepo.GetCart(c, userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении корзины")
	}
	if len(cart) == 0 {
		return nil, fmt.Errorf("корзина пуста")
	}

	totalPrice := 0
	for _, item := range cart {
		totalPrice += item.Price
	}

	order := &models.Order{
		UserID:       userID,
		SplitSystems: cart,
		CreatedAt:    time.Time{},
		TotalPrice:   totalPrice,
	}
	if err = s.userRepo.ClearCart(c, userID); err != nil {
		return nil, fmt.Errorf("ошибка при очистке корзины")
	}

	return s.repo.CreateOrderByUserCart(c, order)
}

func (s *OrderService) UpdateOrderStatus(c context.Context, orderID uint, status string) (*models.Order, error) {
	return s.repo.UpdateOrderStatus(c, orderID, status)
}

func (s *OrderService) GetAll(c context.Context) ([]models.Order, error) {
	return s.repo.GetAll(c)
}
