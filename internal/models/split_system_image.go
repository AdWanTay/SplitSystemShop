package models

type SplitSystemImage struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	SplitSystemID uint   `json:"split_system_id"`
	ImageURL      string `json:"image_url"`
}
