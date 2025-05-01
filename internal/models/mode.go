package models

type Mode struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
