package models

import "time"

type User struct {
	ID               string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email            string  `gorm:"uniqueIndex;type:varchar(100);not null"`
	Password         *string `gorm:"type:varchar(100);default:null"`
	FullName         string
	PhoneNumber      *string `gorm:"type:varchar(100);default:null"`
	StravaIntegrated bool    `gorm:"default:false" json:"strava_integrated"`

	// Relations
	RoleID *string

	Role           *Role           `gorm:"foreignKey:RoleID;references:ID"`
	Profile        *UserProfile    `gorm:"foreignKey:UserID;references:ID" json:"profile"`
	StravaToken    *StravaToken    `gorm:"-"`
	FoodScans      []FoodScan      `gorm:"-" json:"food_scans"`
	UserMeals      []UserMeal      `gorm:"-" json:"user_meals"`
	AIChats        []AiChat        `gorm:"-" json:"ai_chats"`
	AIChatMessages []AIChatMessage `gorm:"-" json:"ai_chat_messages"`
	MembersOf      []FamilyMember  `gorm:"-" json:"member_of"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
