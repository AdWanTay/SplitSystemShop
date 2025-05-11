package dto

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
