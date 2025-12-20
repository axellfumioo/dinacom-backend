package models

import "time"

type User struct {
	UserID      string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email       string  `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	Password    *string `gorm:"type:varchar(100);default:null" json:"password"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `gorm:"type:varchar(100);default:null" json:"phone_number"`
	Role        string  `gorm:"type:varchar(20);default:'USER'" json:"role"`

	Profile UserProfile `gorm:"foreignKey:UserID;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
