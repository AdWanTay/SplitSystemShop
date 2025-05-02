package repositories

import (
	"gorm.io/gorm"
)

type SplitSystemRepository interface {
	//GetUserByEmail(c context.Context, email string) (*models.User, error)
	//CreateUser(c context.Context, user *models.User) error
	//GetUserById(c context.Context, id uint) (*models.User, error)
	//Update(c context.Context, user *models.User) error
	//Delete(c context.Context, userId uint) error
}

type splitSystemRepository struct {
	db *gorm.DB
}

func NewSplitSystemRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
