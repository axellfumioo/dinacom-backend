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
	GetAllMemberStatistics(userId string, roomId string) ([]models.User, error)
	AddFamilyMembers(userID string, req request.AddFamilyMemberRequest) ([]models.FamilyMember, error)
	DeleteFamilyMember(userID string, familyID string, memberID string) (*models.FamilyMember, error)
}

type memberService struct {
	db *gorm.DB
}

func NewMemberService(db *gorm.DB) MemberService {
	return &memberService{db: db}
}

func (s *memberService) AddFamilyMembers(userID string, req request.AddFamilyMemberRequest) ([]models.FamilyMember, error) {
	var currentMember models.FamilyMember
	if err := s.db.First(&currentMember, "user_id = ? AND family_id = ?", userID, req.FamilyID).Error; err != nil {
		return nil, errors.New("failed to get current member:" + err.Error())
	}

	if currentMember.Role != "PARENT" {
		return nil, errors.New("access denied:you dont have access to this feature")
	}

	var members []models.FamilyMember
	for _, member := range req.Members {
		members = append(members, models.FamilyMember{UserID: member.UserID, FamilyID: req.FamilyID, Role: constants.MemberRole(member.Role)})
	}

	if err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&members).Error; err != nil {
		return nil, errors.New("failed to create memmbers:" + err.Error())
	}

	return members, nil
}

func (s *memberService) GetAllMemberStatistics(userId string, roomId string) ([]models.User, error) {
	var users []models.User

	if err := s.db.
		Model(&models.User{}).
		Joins("JOIN family_members fm ON fm.user_id = users.id").
		Where("fm.family_id = ?", roomId).
		Where("users.id != ?", userId).
		Preload("Profile").
		Preload("Role").
		Preload("MemberOf").
		Preload("UserMeals").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

func (s *memberService) DeleteFamilyMember(userID string, familyID string, memberID string) (*models.FamilyMember, error) {
	var currentMember models.FamilyMember
	if err := s.db.First(&currentMember, "user_id = ? AND family_id = ?", userID, familyID).Error; err != nil {
		return nil, errors.New("failed to get current member:" + err.Error())
	}

	// Check if member is not a parent
	if currentMember.Role != "PARENT" {
		return nil, errors.New("access denied:you dont have access to this feature")
	}

	// Check if exist
	var exist models.FamilyMember
	if err := s.db.First(&exist, "user_id = ? AND family_id = ?", userID, familyID).Error; err != nil {
		return nil, errors.New("failed to get member:" + err.Error())
	}

	if err := s.db.Delete(&exist).Error; err != nil {
		return nil, errors.New("failed to delete member:" + err.Error())
	}

	return &exist, nil
}
