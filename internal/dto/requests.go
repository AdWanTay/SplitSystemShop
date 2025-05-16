package dto

import "SplitSystemShop/internal/models"

type RegistrationRequest struct {
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangeCredentialsRequest struct {
	NewPhoneNumber string `json:"new_phone_number"`
	NewLastName    string `json:"new_last_name"`
	NewFirstName   string `json:"new_first_name"`
	NewPatronymic  string `json:"new_patronymic"`
	NewEmail       string `json:"new_email"`
}

type FiltersQuery struct {
	Brands               []uint  `json:"brands,omitempty"`
	RecommendedArea      uint    `json:"recommended_area,omitempty"`
	CoolingPower         uint    `json:"cooling_power,omitempty"`
	Types                []uint  `json:"types,omitempty"`
	MinPrice             uint    `json:"min_price,omitempty"`
	MaxPrice             uint    `json:"max_price,omitempty"`
	HasInverter          bool    `json:"has_inverter,omitempty"`
	MinNoiseLevel        float64 `json:"min_noise_level,omitempty"`
	MaxNoiseLevel        float64 `json:"max_noise_level,omitempty"`
	Modes                []uint  `json:"modes,omitempty"`
	EnergyClassCoolingID uint    `json:"energy_class_cooling_id,omitempty"`
	EnergyClassHeatingID uint    `json:"energy_class_heating_id,omitempty"`
	ExternalWeight       float64 `json:"external_weight,omitempty"`
	ExternalWidth        int     `json:"external_width,omitempty"`
	ExternalHeight       int     `json:"external_height,omitempty"`
	ExternalDepth        int     `json:"external_depth,omitempty"`
	InternalDepth        int     `json:"internal_depth,omitempty"`
	InternalWidth        int     `json:"internal_width,omitempty"`
	InternalWeight       float64 `json:"internal_weight,omitempty"`
	InternalHeight       int     `json:"internal_height,omitempty"`
}

type NewReviewRequest struct {
	SplitSystemID uint   `json:"split_system_id"`
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
}

type NewArticleRequest struct {
	Title       string `json:"title"`       // Заголовок статьи
	Description string `json:"description"` // Краткое описание (для карточки)
	Content     string `json:"content"`     // HTML контент статьи (из Quill)
	ImageURL    string `json:"image_url"`
}

type NewSplitSystemRequest struct {
	Title                string        `json:"title"`
	ShortDescription     string        `json:"short_description"`
	LongDescription      string        `json:"long_description"`
	BrandID              uint          `json:"brand_id"`
	TypeID               uint          `json:"type_id"`
	Price                int           `json:"price"`
	HasInverter          bool          `json:"has_inverter"`
	RecommendedArea      float64       `json:"recommended_area"`
	CoolingPower         float64       `json:"cooling_power"`
	EnergyClassCoolingID uint          `json:"energy_class_cooling_id"`
	EnergyClassHeatingID uint          `json:"energy_class_heating_id"`
	Modes                []models.Mode `json:"modes"`
	MinNoiseLevel        float64       `json:"min_noise_level"`
	MaxNoiseLevel        float64       `json:"max_noise_level"`
	ExternalWeight       float64       `json:"external_weight"`
	ExternalWidth        int           `json:"external_width"`
	ExternalHeight       int           `json:"external_height"`
	ExternalDepth        int           `json:"external_depth"`
	InternalWeight       float64       `json:"internal_weight"`
	InternalWidth        int           `json:"internal_width"`
	InternalHeight       int           `json:"internal_height"`
	InternalDepth        int           `json:"internal_depth"`
	ImageURL             string        `json:"image_url"`
}
type UpdateSplitSystemRequest struct {
	Title                string        `json:"title,omitempty"`
	ShortDescription     string        `json:"short_description,omitempty"`
	LongDescription      string        `json:"long_description,omitempty"`
	Price                int           `json:"price,omitempty"`
	BrandID              uint          `json:"brand_id,omitempty"`
	TypeID               uint          `json:"type_id,omitempty"`
	RecommendedArea      float64       `json:"recommended_area,omitempty"`
	CoolingPower         float64       `json:"cooling_power,omitempty"`
	HasInverter          bool          `json:"has_inverter,omitempty"`
	EnergyClassCoolingID uint          `json:"energy_class_cooling_id,omitempty"`
	Modes                []models.Mode `json:"modes,omitempty"`
	EnergyClassHeatingID uint          `json:"energy_class_heating_id,omitempty"`
	MinNoiseLevel        float64       `json:"min_noise_level,omitempty"`
	MaxNoiseLevel        float64       `json:"max_noise_level,omitempty"`
	InternalWeight       float64       `json:"internal_weight,omitempty"`
	InternalWidth        int           `json:"internal_width,omitempty"`
	InternalHeight       int           `json:"internal_height,omitempty"`
	InternalDepth        int           `json:"internal_depth,omitempty"`
	ExternalWeight       float64       `json:"external_weight,omitempty"`
	ExternalWidth        int           `json:"external_width,omitempty"`
	ExternalHeight       int           `json:"external_height,omitempty"`
	ExternalDepth        int           `json:"external_depth,omitempty"`
	ImageURL             *string       `json:"image_url,omitempty"`
}
