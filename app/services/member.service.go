package services

import (
	"backend-dinakom/app/constants"
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberService interface {
	AddFamilyMembers(req request.AddFamilyMemberRequest) (any, error)
}

type memberService struct {
	db *gorm.DB
}

func NewMemberService(db *gorm.DB) MemberService {
	return &memberService{db: db}
}

func (s *memberService) AddFamilyMembers(req request.AddFamilyMemberRequest) (any, error) {
	var members []models.FamilyMember
	for _, member := range req.Members {
		members = append(members, models.FamilyMember{UserID: member.UserID, FamilyID: req.FamilyID, Role: constants.MemberRole(member.Role)})
	}

	if err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&members).Error; err != nil {
		return nil, errors.New("failed to create memmbers:" + err.Error())
	}

	return members, nil
}
