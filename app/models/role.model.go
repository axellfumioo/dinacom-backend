package models

import "time"

type Role struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	RoleName  string    `gorm:"uniqueIndex;not null" json:"role_name"`
	Users     []User    `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}
