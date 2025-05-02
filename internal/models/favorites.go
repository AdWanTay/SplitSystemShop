package models

type Favorite struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint `gorm:"not null"`
	SplitSystemID uint `gorm:"not null"`

	User        User        `gorm:"foreignKey:UserID"`
	SplitSystem SplitSystem `gorm:"foreignKey:SplitSystemID"`
}
