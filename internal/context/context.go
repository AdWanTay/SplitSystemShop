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
	BrandService       *services.BrandService
	TypeService        *services.TypeService
	ModeService        *services.ModeService
	EnergyClassService *services.EnergyClassService
}

func InitServices(db *gorm.DB) *AppContext {
	userRepo := repositories.NewUserRepository(db)
	splitSystemRepo := repositories.NewSplitSystemRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	favoritesRepo := repositories.NewFavoritesRepository(db)
	brandRepo := repositories.NewBrandRepository(db)
	typeRepo := repositories.NewTypeRepository(db)
	modeRepo := repositories.NewModeRepository(db)
	energyClassRepo := repositories.NewEnergyClassRepository(db)

	return &AppContext{
		UserService:        services.NewUserService(userRepo),
		SplitSystemService: services.NewSlitSystemService(splitSystemRepo),
		CartService:        services.NewCartService(cartRepo),
		FavoritesService:   services.NewFavoritesService(favoritesRepo),
		BrandService:       services.NewBrandService(brandRepo),
		TypeService:        services.NewTypeService(typeRepo),
		ModeService:        services.NewModeService(modeRepo),
		EnergyClassService: services.NewEnergyClassService(energyClassRepo),
	}
}
