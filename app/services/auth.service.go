package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(req request.RegisterRequest) (any, error)
	Login(req request.LoginRequest) (string, error)
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db gorm.DB) AuthService {
	return &authService{db: &db}
}

func (s *authService) Register(req request.RegisterRequest) (any, error) {
	var countUsers int64
	s.db.Model(&models.User{}).Count(&countUsers)

	// use value type, not pointer
	var existingRole models.Role
	roleName := "USER"
	if countUsers < 3 {
		roleName = "ADMIN"
	}
	if err := s.db.First(&existingRole, "role_name = ?", roleName).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	var existingUser models.User
	if err := s.db.First(&existingUser, "email = ?", req.Email).Error; err == nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FullName: req.Name,
		Email:    req.Email,
		Password: &hashedPassword,
		RoleID:   &existingRole.ID,
		Profile: &models.UserProfile{
			DateOfBirth: req.DateOfBirth,
			Gender:      req.Gender,
		},
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	userResponse := helpers.ToUserResponse(user)
	return &userResponse, nil
}

func (s *authService) Login(req request.LoginRequest) (string, error) {
	var existingUser *models.User

	if err := s.db.First(&existingUser, "email = ?", req.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("user is not registered")
		}
		return "", err
	}

	if err := helpers.ComparePassword(req.Password, *existingUser.Password); err != nil {
		return "", errors.New("incorrect password")
	}

	access_token, err := helpers.GenerateToken(existingUser.ID, existingUser.Email, "USER")
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return access_token, nil
}
