package models

type SplitSystemImage struct {
	ID            uint `gorm:"primaryKey"`
	SplitSystemID uint
	ImageURL      string
}
