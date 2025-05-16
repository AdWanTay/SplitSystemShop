package models

type Article struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"` // Заголовок статьи
	Description string `json:"description"`
	Content     string `gorm:"type:text" json:"content"` // HTML контент
	ImageURL    string `json:"image_url"`                // путь к картинке
}
