package models

type EnergyClass struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
