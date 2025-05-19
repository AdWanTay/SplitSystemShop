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
	ReviewService      *services.ReviewService
	ArticleService     *services.ArticleService
	OrderService       *services.OrderService
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
	reviewRepo := repositories.NewReviewRepository(db)
	articleRepo := repositories.NewArticleRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	userService := services.NewUserService(userRepo)
	reviewService := services.NewReviewService(reviewRepo)
	articleService := services.NewArticleService(articleRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)

	return &AppContext{
		UserService:        userService,
		SplitSystemService: services.NewSlitSystemService(splitSystemRepo),
		CartService:        services.NewCartService(cartRepo, userService),
		FavoritesService:   services.NewFavoritesService(favoritesRepo),
		BrandService:       services.NewBrandService(brandRepo),
		TypeService:        services.NewTypeService(typeRepo),
		ModeService:        services.NewModeService(modeRepo),
		EnergyClassService: services.NewEnergyClassService(energyClassRepo),
		ReviewService:      reviewService,
		ArticleService:     articleService,
		OrderService:       orderService,
	}
}
