package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetUserCart(c context.Context, userID uint) ([]models.SplitSystem, error)
	AddToCart(c context.Context, userID, systemID uint) error
	RemoveFromCart(c context.Context, userID, systemID uint) error
	ClearCart(c context.Context, userID uint) error
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

type cartRepository struct {
	db *gorm.DB
}

func (r cartRepository) GetUserCart(c context.Context, userID uint) ([]models.SplitSystem, error) {
	var user models.User
	if err := r.db.WithContext(c).
		Preload("Cart.Brand").
		Preload("Cart.Type").
		Preload("Cart.Modes").
		Preload("Cart.EnergyClassCooling").
		Preload("Cart.EnergyClassHeating").
		First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Cart, nil

}

func (r cartRepository) AddToCart(c context.Context, userID, systemID uint) error {
	return r.db.WithContext(c).Exec(
		"INSERT INTO user_cart (user_id, split_system_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
		userID, systemID).Error
}

func (r cartRepository) RemoveFromCart(c context.Context, userID, systemID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM user_cart WHERE user_id = ? AND split_system_id = ?",
		userID, systemID).Error
}

func (r cartRepository) ClearCart(c context.Context, userID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM user_cart WHERE user_id = ?",
		userID).Error
}
