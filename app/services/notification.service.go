package services

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type NotificationService interface {
	DailyReminder() (any, error)
}

type notificationService struct {
	db *gorm.DB
}

func NewNotificationService(db *gorm.DB) NotificationService {
	return &notificationService{db: db}
}

func (s *notificationService) DailyReminder() (any, error) {
	var existingRole models.Role
	if err := s.db.First(&existingRole, "role_name = ?", "USER").Error; err != nil {
		return nil, errors.New("failed to find role: " + err.Error())
	}

	var existingUsers []models.User
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	if err := s.db.
		Table("users u").
		Joins(`
        LEFT JOIN user_meals usm
        ON usm.user_id = u.id
        AND usm.created_at >= ?
        AND usm.created_at < ?
    `, today, tomorrow).
		Where("u.role_id = ?", existingRole.ID).
		Where("usm.id IS NULL").
		Find(&existingUsers).Error;
	err != nil {
		return nil, errors.New("failed to get users")
	}

	userResponse := helpers.ToUsersResponse(existingUsers)
	return userResponse, nil
}
