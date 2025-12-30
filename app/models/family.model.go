package models

import "time"

type Family struct {
	ID   string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string  `gorm:"type:varchar(25);uniqueIndex;not null" json:"name"`
	Desc *string `gorm:"type:text;"`

	Member []FamilyMember `gorm:"-"`

	CreatedAt time.Time
	UpdateAt  time.Time
}

func (Family) TableName() string {
	return "families"
}
