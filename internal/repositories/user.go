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
	Delete(c context.Context, userId uint) error
	GetCart(c context.Context, userId uint) (*[]models.SplitSystem, error)
	GetFavorites(c context.Context, userId uint) (*[]models.SplitSystem, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetUserById(c context.Context, id uint) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(c).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUserByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(c).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) CreateUser(c context.Context, user *models.User) error {
	return u.db.WithContext(c).Create(user).Error
}
func (u userRepository) Update(c context.Context, user *models.User) error {
	return u.db.WithContext(c).Save(user).Error
}

func (u userRepository) Delete(c context.Context, userId uint) error {
	return u.db.WithContext(c).Delete(models.User{}, userId).Error
}

func (u userRepository) GetCart(c context.Context, userId uint) (*[]models.SplitSystem, error) {
	var user models.User

	err := u.db.WithContext(c).Preload("Cart").First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user.Cart, nil
}

func (u userRepository) GetFavorites(c context.Context, userId uint) (*[]models.SplitSystem, error) {
	var user models.User

	err := u.db.WithContext(c).Preload("Favorites").First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user.Favorites, nil
}
