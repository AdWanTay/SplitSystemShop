package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	LastName    string `gorm:"not null"`
	FirstName   string `gorm:"not null"`
	Patronymic  string `gorm:"not null"`
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
	Password    string `gorm:"not null"`
	Reviews     []Review
	Role        string        `gorm:"default:'user'"`
	Cart        []SplitSystem `gorm:"many2many:user_cart"`
	Favorites   []SplitSystem `gorm:"many2many:user_favorites"`
}
