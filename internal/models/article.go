package models

type Article struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`                    // Заголовок статьи
	Content string `gorm:"type:text" json:"content"` // HTML контент
}
