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
	GetFamilyMembers(FamilyID string) ([]models.FamilyMember, error)
	AddFamilyMembers(req request.AddFamilyMemberRequest) ([]models.FamilyMember, error)
}

type memberService struct {
	db *gorm.DB
}

func NewMemberService(db *gorm.DB) MemberService {
	return &memberService{db: db}
}

func (s *memberService) AddFamilyMembers(req request.AddFamilyMemberRequest) ([]models.FamilyMember, error) {
	var members []models.FamilyMember
	for _, member := range req.Members {
		members = append(members, models.FamilyMember{UserID: member.UserID, FamilyID: req.FamilyID, Role: constants.MemberRole(member.Role)})
	}

	if err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&members).Error; err != nil {
		return nil, errors.New("failed to create memmbers:" + err.Error())
	}

	return members, nil
}

func (s *memberService) GetFamilyMembers(FamilyID string) ([]models.FamilyMember, error) {
	var existingFamily models.Family
	if err := s.db.First(&existingFamily, "id = ?", FamilyID).Error; err != nil {
		return nil, errors.New("failed to get family:" + err.Error())
	}
	
	var members []models.FamilyMember
	if err := s.db.Where("family_id = ?", FamilyID).Find(&members).Error; err != nil {
		return nil, errors.New("failed to get members:" + err.Error())
	}
	
	return members, nil
}
