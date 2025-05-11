package models

import "time"

type Review struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	SplitSystemID uint        `gorm:"not null" json:"split_system_id"`
	SplitSystem   SplitSystem `json:"split_system"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `json:"user"`

	Rating    int       `gorm:"not null" json:"rating"` // 1–5
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
