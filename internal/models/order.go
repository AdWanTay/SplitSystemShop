package models

import (
	"time"
)

type Order struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	UserID       uint          `json:"user_id"`
	User         User          `gorm:"foreignKey:UserID" json:"user"`
	SplitSystems []SplitSystem `gorm:"many2many:order_split_systems" json:"split_systems"`
	CreatedAt    time.Time     `json:"created_at"`
	TotalPrice   int           `json:"total_price"`
	Status       string        `gorm:"default:'В обработке'" json:"status"`
}
