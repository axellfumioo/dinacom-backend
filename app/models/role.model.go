package models

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RoleName  string    `gorm:"uniqueIndex;not null" json:"role_name"`
	Users     *[]User    `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}
