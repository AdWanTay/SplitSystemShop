package repositories

import (
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type FavoritesRepository interface {
	GetUserFavorites(c context.Context, userID uint) ([]models.SplitSystem, error)
	AddToFavorites(c context.Context, userID, systemID uint) error
	RemoveFromFavorites(c context.Context, userID, systemID uint) error
	ClearFavorites(c context.Context, userID uint) error
}

func NewFavoritesRepository(db *gorm.DB) FavoritesRepository {
	return &favoritesRepository{db: db}
}

type favoritesRepository struct {
	db *gorm.DB
}

func (r favoritesRepository) GetUserFavorites(c context.Context, userID uint) ([]models.SplitSystem, error) {
	var user models.User
	if err := r.db.WithContext(c).
		Preload("Favorites.Brand").
		Preload("Favorites.Type").
		Preload("Favorites.Modes").
		Preload("Favorites.EnergyClassCooling").
		Preload("Favorites.EnergyClassHeating").
		First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Favorites, nil

}

func (r favoritesRepository) AddToFavorites(c context.Context, userID, systemID uint) error {
	return r.db.WithContext(c).Exec(
		"INSERT INTO user_favorites (user_id, split_system_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
		userID, systemID).Error
}

func (r favoritesRepository) RemoveFromFavorites(c context.Context, userID, systemID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM user_favorites WHERE user_id = ? AND split_system_id = ?",
		userID, systemID).Error
}

func (r favoritesRepository) ClearFavorites(c context.Context, userID uint) error {
	return r.db.WithContext(c).Exec(
		"DELETE FROM user_favorites WHERE user_id = ?",
		userID).Error
}
