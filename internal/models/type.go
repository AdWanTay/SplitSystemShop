package models

type Type struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
