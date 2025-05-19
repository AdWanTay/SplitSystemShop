package models

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	LastName    string        `gorm:"not null" json:"last_name"`
	FirstName   string        `gorm:"not null" json:"first_name"`
	Patronymic  string        `gorm:"not null" json:"patronymic"`
	Email       string        `gorm:"unique" json:"email"`
	PhoneNumber string        `gorm:"unique" json:"phone_number"`
	Password    string        `gorm:"not null" json:"-"`
	Reviews     []Review      `json:"reviews,omitempty"`
	Role        string        `gorm:"default:'user'" json:"role"`
	Cart        []SplitSystem `gorm:"many2many:user_cart" json:"cart"`
	Favorites   []SplitSystem `gorm:"many2many:user_favorites" json:"favorites"`
	Orders      []Order       `gorm:"foreignKey:UserID" json:"orders"`
}
