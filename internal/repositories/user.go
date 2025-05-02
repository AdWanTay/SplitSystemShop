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
