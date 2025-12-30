package models

import (
	"backend-dinakom/app/constants"
	"time"
)

type FamilyMember struct {
	ID   string               `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Role constants.MemberRole `gorm:"type:varchar(20);default:'CHILD'"` // PARENT - CHILD

	FamilyID string `gorm:"type:uuid;not null"`
	UserID   string `gorm:"type:uuid;not null"`

	User   *User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Family *Family `gorm:"foreignKey:FamilyID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (FamilyMember) TableName() string {
	return "family_members"
}
