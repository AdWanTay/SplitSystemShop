package context

import (
	"SplitSystemShop/internal/repositories"
	"SplitSystemShop/internal/services"
	"gorm.io/gorm"
)

type AppContext struct {
	UserService        *services.UserService
	SplitSystemService *services.SplitSystemService
}

func InitServices(db *gorm.DB) *AppContext {
	userRepo := repositories.NewUserRepository(db)
	splitSystemRepo := repositories.NewSplitSystemRepository(db)

	return &AppContext{
		UserService:        services.NewUserService(userRepo),
		SplitSystemService: services.NewSlitSystemService(splitSystemRepo),
	}
}
