package models

type SplitSystem struct {
	ID                   uint `gorm:"primaryKey"`
	BrandID              uint
	Brand                Brand
	TypeID               uint
	Type                 Type
	Price                int // копейки
	Inverter             bool
	RecommendedArea      float64 // м²
	CoolingPower         float64 // кВт
	Modes                []Mode  `gorm:"many2many:split_system_modes"`
	EnergyClassCoolingID uint
	EnergyClassCooling   EnergyClass `gorm:"foreignKey:EnergyClassCoolingID"`
	EnergyClassHeatingID uint
	EnergyClassHeating   EnergyClass `gorm:"foreignKey:EnergyClassHeatingID"`

	MinNoiseLevel float64 // дБ
	MaxNoiseLevel float64

	ExternalWeight float64 // кг
	ExternalWidth  int     // мм
	ExternalHeight int
	ExternalDepth  int

	InternalWeight float64 // кг
	InternalWidth  int
	InternalHeight int
	InternalDepth  int

	Images []SplitSystemImage `gorm:"foreignKey:SplitSystemID"`
}
