package services

import (
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type DoctorChatService interface {
	GetAllDoctors() ([]models.User, error)
	GetAllUserDoctorChats(userID string) ([]models.DoctorChatRoom, error)
	GetOrCreateDoctorChatRoom(userID string, doctorID string) (*models.DoctorChatRoom, error)
	CreateDoctorChatroom(userID string, doctorID string) (*models.DoctorChatRoom, error)
}

type doctorChatService struct {
	db *gorm.DB
}

func NewDoctorChatService(db *gorm.DB) DoctorChatService {
	return &doctorChatService{db: db}
}

func (s *doctorChatService) GetAllDoctors() ([]models.User, error) {
	var doctors []models.User
	if err := s.db.Joins("role ON role.id = users.role_id").Where("role.role_name = ?", "DOCTOR").Find(&doctors).Error; err != nil {
		return nil, err
	}

	if len(doctors) == 0 {
		return nil, errors.New("no doctors found")
	}

	return doctors, nil
}

func (s *doctorChatService) GetAllUserDoctorChats(userID string) ([]models.DoctorChatRoom, error) {
	var existingUser models.User
	if err := s.db.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}
	var existingChatRoom []models.DoctorChatRoom
	if err := s.db.Preload("Doctor").Where("user_id = ?", userID).Find(&existingChatRoom).Error; err != nil {
		return nil, err
	}

	return existingChatRoom, nil
}

func (s *doctorChatService) GetOrCreateDoctorChatRoom(userID string, doctorID string) (*models.DoctorChatRoom, error) {
	var chatRoom models.DoctorChatRoom
	if err := s.db.Where("user_id = ? AND doctor_id = ?", userID, doctorID).First(&chatRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newChatRoom := models.DoctorChatRoom{
				UserID:   userID,
				DoctorID: doctorID,
			}
			if err := s.db.Create(&newChatRoom).Error; err != nil {
				return nil, err
			}
			return &newChatRoom, nil
		}

		return nil, err
	}

	return &chatRoom, nil
}

func (s *doctorChatService) CreateDoctorChatroom(userID string, doctorID string) (*models.DoctorChatRoom, error) {
	var existingChatRoom models.DoctorChatRoom
	if err := s.db.Where("user_id = ? AND doctor_id = ?", userID, doctorID).First(&existingChatRoom).Error; err == nil {
		return nil, errors.New("chat room is already exist")
	}

	newChatRoom := models.DoctorChatRoom{
		DoctorID: doctorID,
		UserID:   userID,
	}
	if err := s.db.Create(&newChatRoom).Error; err != nil {
		return nil, errors.New("failed to create doctor chat room: " + err.Error())
	}

	return &newChatRoom, nil
}
