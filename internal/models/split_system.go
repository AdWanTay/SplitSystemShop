package models

type SplitSystem struct {
	ID                   uint        `gorm:"primaryKey" json:"id"`
	Title                string      `json:"title"`
	BrandID              uint        `json:"brand_id"`
	Brand                Brand       `json:"brand"`
	TypeID               uint        `json:"type_id"`
	Type                 Type        `json:"type"`
	Price                int         `json:"price"` // копейки
	HasInverter          bool        `json:"has_inverter"`
	RecommendedArea      float64     `json:"recommended_area"` // м²
	CoolingPower         float64     `json:"cooling_power"`    // кВт
	Modes                []Mode      `gorm:"many2many:split_system_modes" json:"modes"`
	EnergyClassCoolingID uint        `json:"energy_class_cooling_id"`
	EnergyClassCooling   EnergyClass `json:"energy_class_cooling"`
	EnergyClassHeatingID uint        `json:"energy_class_heating_id"`
	EnergyClassHeating   EnergyClass `json:"energy_class_heating"`

	MinNoiseLevel float64 `json:"min_noise_level"` // дБ
	MaxNoiseLevel float64 `json:"max_noise_level"`

	ExternalWeight float64 `json:"external_weight"` // кг
	ExternalWidth  int     `json:"external_width"`  // мм
	ExternalHeight int     `json:"external_height"`
	ExternalDepth  int     `json:"external_depth"`

	InternalWeight float64 `json:"internal_weight"` // кг
	InternalWidth  int     `json:"internal_width"`
	InternalHeight int     `json:"internal_height"`
	InternalDepth  int     `json:"internal_depth"`
	ImageURL       string  `gorm:"foreignKey:SplitSystemID" json:"image_url"`

	AverageRating float64 `gorm:"average_rating; default:0" json:"average_rating"`
}
