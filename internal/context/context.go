package context

import (
	"SplitSystemShop/internal/repositories"
	"SplitSystemShop/internal/services"
	"gorm.io/gorm"
)

type AppContext struct {
	UserService        *services.UserService
	SplitSystemService *services.SplitSystemService
	CartService        *services.CartService
	FavoritesService   *services.FavoritesService
}

func InitServices(db *gorm.DB) *AppContext {
	userRepo := repositories.NewUserRepository(db)
	splitSystemRepo := repositories.NewSplitSystemRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	favoritesRepo := repositories.NewFavoritesRepository(db)

	return &AppContext{
		UserService:        services.NewUserService(userRepo),
		SplitSystemService: services.NewSlitSystemService(splitSystemRepo),
		CartService:        services.NewCartService(cartRepo),
		FavoritesService:   services.NewFavoritesService(favoritesRepo),
	}
}
