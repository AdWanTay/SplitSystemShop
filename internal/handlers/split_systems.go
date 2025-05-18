package handlers

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/services"
	"SplitSystemShop/internal/utils"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
	"strconv"
	"time"
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
		if service.Delete(c.Context(), uint(splitSystemID)) != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка удаления товара"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Товар успешно удален",
		})
	}
}

func CreateSplitSystem(splitSystemService *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Файл обязателен"})
		}

		// Генерация пути
		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(file.Filename)
		savePath := "./web/static/uploads/" + filename

		if err = c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при сохранении файла"})
		}

		// Парсинг полей
		price, _ := strconv.Atoi(c.FormValue("price"))
		brandID, _ := strconv.ParseInt(c.FormValue("brand_id"), 10, 64)
		typeID, _ := strconv.ParseInt(c.FormValue("type_id"), 10, 64)
		recommendedArea, _ := strconv.ParseFloat(c.FormValue("recommended_area"), 64)
		coolingPower, _ := strconv.ParseFloat(c.FormValue("cooling_power"), 64)
		energyCoolID, _ := strconv.ParseInt(c.FormValue("energy_class_cooling_id"), 10, 64)
		energyHeatID, _ := strconv.ParseInt(c.FormValue("energy_class_heating_id"), 10, 64)
		minNoise, _ := strconv.ParseFloat(c.FormValue("min_noise_level"), 64)
		maxNoise, _ := strconv.ParseFloat(c.FormValue("max_noise_level"), 64)
		extWeight, _ := strconv.ParseFloat(c.FormValue("external_weight"), 64)
		extWidth, _ := strconv.Atoi(c.FormValue("external_width"))
		extHeight, _ := strconv.Atoi(c.FormValue("external_height"))
		extDepth, _ := strconv.Atoi(c.FormValue("external_depth"))
		intWeight, _ := strconv.ParseFloat(c.FormValue("internal_weight"), 64)
		intWidth, _ := strconv.Atoi(c.FormValue("internal_width"))
		intHeight, _ := strconv.Atoi(c.FormValue("internal_height"))
		intDepth, _ := strconv.Atoi(c.FormValue("internal_depth"))
		hasInverter := c.FormValue("has_inverter") == "true"

		//modeIDs := c.FormValue("modes")
		//var modes []models.Mode
		//for _, mid := range modeIDs {
		//	id, err := strconv.Atoi(mid)
		//	if err != nil {
		//		continue
		//	}
		//	modes = append(modes, models.Mode{ID: uint(id)})
		//}

		split := &models.SplitSystem{
			Title:                c.FormValue("title"),
			ShortDescription:     c.FormValue("short_description"),
			LongDescription:      c.FormValue("long_description"),
			BrandID:              uint(brandID),
			TypeID:               uint(typeID),
			Price:                price,
			HasInverter:          hasInverter,
			RecommendedArea:      recommendedArea,
			CoolingPower:         coolingPower,
			EnergyClassCoolingID: uint(energyCoolID),
			EnergyClassHeatingID: uint(energyHeatID),
			MinNoiseLevel:        minNoise,
			MaxNoiseLevel:        maxNoise,
			ExternalWeight:       extWeight,
			ExternalWidth:        extWidth,
			ExternalHeight:       extHeight,
			ExternalDepth:        extDepth,
			InternalWeight:       intWeight,
			InternalWidth:        intWidth,
			InternalHeight:       intHeight,
			InternalDepth:        intDepth,
			ImageURL:             filename,
		}

		if _, err = splitSystemService.Create(c.Context(), *split); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать товар"})
		}

		return c.Status(fiber.StatusCreated).JSON(split)
	}
}

func UpdateSplitSystem(service *services.SplitSystemService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректный ID"})
		}

		// Попытка получить файл
		file, err := c.FormFile("image")
		var filename string
		if err == nil && file != nil {
			filename = strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(file.Filename)
			savePath := "./web/static/uploads/" + filename
			if err := c.SaveFile(file, savePath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при сохранении файла"})
			}
		}

		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка при обработке формы"})
		}

		modeIDs := form.Value["modes"]
		var modes []models.Mode
		for _, mid := range modeIDs {
			id, err := strconv.Atoi(mid)
			if err != nil {
				continue
			}
			modes = append(modes, models.Mode{ID: uint(id)})
		}

		input := dto.UpdateSplitSystemRequest{
			Title:                c.FormValue("title"),
			ShortDescription:     c.FormValue("short_description"),
			LongDescription:      c.FormValue("long_description"),
			Price:                utils.ParseInt(c.FormValue("price")),
			BrandID:              utils.ParseUint(c.FormValue("brand_id")),
			TypeID:               utils.ParseUint(c.FormValue("type_id")),
			RecommendedArea:      utils.ParseFloat(c.FormValue("recommended_area")),
			CoolingPower:         utils.ParseFloat(c.FormValue("cooling_power")),
			Modes:                modes,
			HasInverter:          c.FormValue("has_inverter") == "true",
			EnergyClassCoolingID: utils.ParseUint(c.FormValue("energy_class_cooling_id")),
			EnergyClassHeatingID: utils.ParseUint(c.FormValue("energy_class_heating_id")),
			MinNoiseLevel:        utils.ParseFloat(c.FormValue("min_noise_level")),
			MaxNoiseLevel:        utils.ParseFloat(c.FormValue("max_noise_level")),
			InternalWeight:       utils.ParseFloat(c.FormValue("internal_weight")),
			InternalWidth:        utils.ParseInt(c.FormValue("internal_width")),
			InternalHeight:       utils.ParseInt(c.FormValue("internal_height")),
			InternalDepth:        utils.ParseInt(c.FormValue("internal_depth")),
			ExternalWeight:       utils.ParseFloat(c.FormValue("external_weight")),
			ExternalWidth:        utils.ParseInt(c.FormValue("external_width")),
			ExternalHeight:       utils.ParseInt(c.FormValue("external_height")),
			ExternalDepth:        utils.ParseInt(c.FormValue("external_depth")),
		}

		if filename != "" {
			input.ImageURL = &filename
		}

		if err = service.UpdateSplitSystem(c.Context(), uint(id), input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при обновлении"})
		}
		return c.JSON(fiber.Map{"message": "Обновление успешно"})
	}
}
