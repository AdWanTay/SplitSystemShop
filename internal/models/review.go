package models

import "time"

type Review struct {
	ID            uint `gorm:"primaryKey"`
	SplitSystemID uint `gorm:"not null"`
	SplitSystem   SplitSystem

	UserID uint `gorm:"not null"`
	User   User

	Rating    int       `gorm:"not null"` // 1–5
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
