package models

import "time"

type User struct {
	ID          string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email       string  `gorm:"uniqueIndex;type:varchar(100);not null"`
	Password    *string `gorm:"type:varchar(100);default:null"`
	FullName    string
	PhoneNumber *string `gorm:"type:varchar(100);default:null"`

	RoleID *string
	Role   *Role `gorm:"foreignKey:RoleID;references:ID"`

	Profile   *UserProfile `gorm:"foreignKey:UserID;references:ID" json:"profile"`
	FoodScans []FoodScan   `gorm:"-" json:"food_scans"`
	UserMeals []UserMeal   `gorm:"-" json:"user_meals"`
	AIChats   []AiChat     `gorm:"-" json:"ai_chats"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
