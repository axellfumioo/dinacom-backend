package services

import (
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	extservices "backend-dinakom/external/services"
	"errors"
	"time"

	"gorm.io/gorm"
)

type NotificationService interface {
	DailyReminder() ([]response.UserResponse, error)
}

type notificationService struct {
	db *gorm.DB
}

func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationService{db: db}
}

func (s *notificationService) DailyReminder() ([]response.UserResponse, error) {
	var existingRole models.Role
	if err := s.db.First(&existingRole, "role_name = ?", "USER").Error; err != nil {
		return nil, errors.New("failed to find role: " + err.Error())
	}

	loc := time.Now().Location()
	startDate := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		0, 0, 0, 0,
		loc,
	)
	endDate := startDate.AddDate(0, 0, 1)
	var existingUsers []models.User
	if err := s.db.
		Table("users AS u").
		Joins(`
		LEFT JOIN user_meals AS um
		ON um.user_id = u.id
		AND um.created_at >= ?
		AND um.created_at < ?
	`, startDate, endDate).
		Where("u.role_id = ?", existingRole.ID).
		Where("um.id IS NULL").
		Find(&existingUsers).Error; err != nil {
		return nil, errors.New("failed to get users: " + err.Error())
	}

	if len(existingUsers) == 0 {
		return nil, errors.New("failed to get users: users not found")
	}

	for _, user := range existingUsers {
		err := extservices.FetchFonnteAPI(user.FullName, *user.PhoneNumber)	
		if (err != nil) {
			return nil, errors.New(err.Error())
		}
	}

	userResponse := helpers.ToUsersResponse(existingUsers)
	return userResponse, nil
}
