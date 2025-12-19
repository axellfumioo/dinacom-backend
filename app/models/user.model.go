package models

import "time"

type User struct {
	UserID      string    `gorm:"primaryKey" json:"user_id"`
	Email       string    `gorm:"uniqueIndex type:varchar(100); not null" json:"email"`
	Password    *string   `gorm:"type:varchar(100);default:null" json:"password"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `gorm:"type:varchar(100);default:null" json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
