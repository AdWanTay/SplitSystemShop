package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderByUserCart(c context.Context, order *models.Order) (*models.Order, error)
	UpdateOrderStatus(c context.Context, orderID uint, status string) (*models.Order, error)
	GetAll(c context.Context) ([]models.Order, error)
}
type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r orderRepository) CreateOrderByUserCart(c context.Context, order *models.Order) (*models.Order, error) {
	if err := r.db.WithContext(c).Create(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
func (r orderRepository) UpdateOrderStatus(c context.Context, orderID uint, status string) (*models.Order, error) {
	// Сначала обновим статус
	if err := r.db.WithContext(c).
		Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error; err != nil {
		return nil, err
	}

	// Затем получим обновлённый заказ
	var order models.Order
	if err := r.db.WithContext(c).
		Preload("User").
		Preload("SplitSystems").
		First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r orderRepository) GetAll(c context.Context) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.WithContext(c).Preload("User").Preload("SplitSystems").Find(&orders).Error
	return orders, err
}
