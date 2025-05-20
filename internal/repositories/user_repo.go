package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(c context.Context, email string) (*models.User, error)
	CreateUser(c context.Context, user *models.User) error
	GetUserById(c context.Context, id uint) (*models.User, error)
	Update(c context.Context, user *models.User) error
	Delete(c context.Context, userID uint) error
	GetFavorites(c context.Context, userID uint) ([]models.SplitSystem, error)
	GetCart(c context.Context, userID uint) ([]models.SplitSystem, error)
	GetOrders(c context.Context, userID uint) ([]models.Order, error)
	ClearCart(c context.Context, userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) GetUserById(c context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(c).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetUserByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) CreateUser(c context.Context, user *models.User) error {
	return r.db.WithContext(c).Create(user).Error
}
func (r userRepository) Update(c context.Context, user *models.User) error {
	return r.db.WithContext(c).Save(user).Error
}

func (r userRepository) Delete(c context.Context, userID uint) error {
	err := r.db.WithContext(c).Model(&models.User{ID: userID}).Association("Cart").Clear()
	if err != nil {
		return err
	}
	err = r.db.WithContext(c).Model(&models.User{ID: userID}).Association("Favorites").Clear()
	if err != nil {
		return err
	}
	return r.db.WithContext(c).Delete(models.User{}, userID).Error
}
func (r userRepository) ClearCart(c context.Context, userID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM user_cart WHERE user_id = ?",
		userID).Error
}

func (r userRepository) GetFavorites(c context.Context, userID uint) ([]models.SplitSystem, error) {
	var user models.User
	err := r.db.WithContext(c).Preload("Favorites").Preload("Favorites.Type").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return user.Favorites, nil
}

func (r userRepository) GetCart(c context.Context, userID uint) ([]models.SplitSystem, error) {
	var user models.User
	err := r.db.WithContext(c).Preload("Cart").Preload("Cart.Type").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return user.Cart, nil
}

func (r userRepository) GetOrders(c context.Context, userID uint) ([]models.Order, error) {
	var user models.User
	err := r.db.WithContext(c).Preload("Orders").Preload("Orders.SplitSystems").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return user.Orders, nil
}
