package handlers

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}

		splitSystem, err := service.GetSplitSystem(c.Context(), uint(splitSystemID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Ошибка получения товара",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"item": splitSystem,
		})
	}
}

func GetAllSplitSystems(splitSystemService *services.SplitSystemService, userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем фильтры из query-параметров

		qFilters := dto.FiltersQuery{}
		err := c.QueryParser(&qFilters)
		if err != nil {
			return err
		}

		type FiltersQuery struct {
			Brand []string
			Type  []string
			Mode  []string
		}

		sliceFilters := FiltersQuery{}
		err = c.QueryParser(&sliceFilters)
		if err != nil {
			return err
		}
		filtersList := []string{
			"recommended_area_min", "recommended_area_max",
			"cooling_power_min", "cooling_power_max",
			"price_min", "price_max",
			"has_inverter",
			"min_noise_level_min", "min_noise_level_max",
			"max_noise_level_min", "max_noise_level_max",
			"energy_class_cooling", "energy_class_heating",
			"external_weight_min", "external_weight_max",
			"external_width_min", "external_width_max",
			"external_height_min", "external_height_max",
			"external_depth_min", "external_depth_max",
			"internal_depth_min", "internal_depth_max",
			"internal_width_min", "internal_width_max",
			"internal_weight_min", "internal_weight_max",
			"internal_height_min", "internal_height_max",
			"sort",
		}
		filter := map[string]interface{}{}
		for _, name := range filtersList {
			value := c.Query(name, "")
			if value != "" {
				filter[name] = value
			}
		}

		if sliceFilters.Brand != nil {
			filter["brand"] = sliceFilters.Brand
		}
		if sliceFilters.Type != nil {
			filter["type"] = sliceFilters.Type
		}
		if sliceFilters.Mode != nil {
			filter["mode"] = sliceFilters.Mode
		}
		// Получаем список систем с фильтрами
		splitSystems, err := splitSystemService.GetAllSplitSystems(c.Context(), filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка получения товаров"})
		}

		// Проверяем, есть ли корзина у пользователя
		userID := c.Locals("userId")
		var cart []models.SplitSystem
		var favorites []models.SplitSystem
		if userID != nil {
			cart, err = userService.GetCart(c.Context(), userID.(uint))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка получения корзины"})
			}
			favorites, err = userService.GetFavorites(c.Context(), userID.(uint))
		}

		// Формируем ответ
		response := dto.CatalogResponse{}
		response.New(cart, favorites, splitSystems)

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func DeleteSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		splitSystemID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Неверный id товара"})
		}
		if service.DeleteSplitSystem(c.Context(), uint(splitSystemID)) != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка удаления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Товар успешно удален",
		})
	}
}
