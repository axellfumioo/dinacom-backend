package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers(page, pageSize int) ([]response.UserResponse, int64, error)
	GetUserSession(userId string) (*response.UserResponse, error)
	GetUserByID(userId string) (*response.UserResponse, error)
	CreateUser(req request.CreateUserRequest) (*response.UserResponse, error)
	UpdateUser(userId string, req request.UpdateUserRequest) (*response.UserResponse, error)
	DeleteUser(userId string) error
	ChangePassword(userId string, req request.ChangePasswordRequest) error
	GetUserByRoleName(roleName string) ([]response.UserResponse, error)
	GetTotalUsers() (int64, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetAllUsers(page, pageSize int) ([]response.UserResponse, int64, error) {
	var users []models.User
	var totalRows int64

	// Count total rows
	if err := s.db.Model(&models.User{}).Count(&totalRows).Error; err != nil {
		return nil, 0, errors.New("failed to count users")
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get users with pagination
	if err := s.db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, errors.New("failed to fetch users")
	}

	usersResponse := helpers.ToUsersResponse(users)
	return usersResponse, totalRows, nil
}

func (s *userService) GetUserSession(userId string) (*response.UserResponse, error) {
	var existingUser models.User
	if err := s.db.Preload("Role").First(&existingUser, "id = ?", userId).Error; err != nil {
		return nil, errors.New("user not registered")
	}

	userResponse := helpers.ToUserResponse(&existingUser)
	return &userResponse, nil
}

func (s *userService) GetUserByID(userId string) (*response.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("database error")
	}

	userResponse := helpers.ToUserResponse(&user)
	return &userResponse, nil
}

func (s *userService) CreateUser(req request.CreateUserRequest) (*response.UserResponse, error) {
	// Check if email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := models.User{
		FullName:    req.Name,
		Email:       req.Email,
		Password:    &hashedPassword,
		Role:        req.Role,
		PhoneNumber: &req.PhoneNumber,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	userResponse := helpers.ToUserResponse(&user)
	return &userResponse, nil
}

func (s *userService) UpdateUser(userId string, req request.UpdateUserRequest) (*response.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("database error")
	}

	// Check if email already used by another user
	if req.Email != "" && req.Email != user.Email {
		var existingUser models.User
		if err := s.db.Where("email = ? AND id != ?", req.Email, userId).First(&existingUser).Error; err == nil {
			return nil, errors.New("email already used by another user")
		}
	}

	// Update fields
	if req.Name != "" {
		user.FullName = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = &req.PhoneNumber
	}
	if req.Role != "" {
		user.Role = req.Role
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	userResponse := helpers.ToUserResponse(&user)
	return &userResponse, nil
}

func (s *userService) DeleteUser(userId string) error {
	var user models.User
	if err := s.db.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return errors.New("database error")
	}

	// Soft delete
	if err := s.db.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (s *userService) ChangePassword(userId string, req request.ChangePasswordRequest) error {
	var user models.User
	if err := s.db.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return errors.New("database error")
	}

	// Verify old password
	if err := helpers.ComparePassword(*user.Password, req.OldPassword); err != nil {
		return errors.New("old password is incorrect")
	}

	// Hash new password
	hashedPassword, err := helpers.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = &hashedPassword
	if err := s.db.Save(&user).Error; err != nil {
		return errors.New("failed to update password")
	}

	return nil
}

func (s *userService) GetUserByRoleName(roleName string) ([]response.UserResponse, error) {
	var users []models.User
	if err := s.db.Joins("JOIN roles ON roles.id = users.role_id").Where("roles.name = ?", roleName).Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("database error")
	}
	// convert to []response.UserResponse
	userResponses := helpers.ToUsersResponse(users)
	return userResponses, nil
}

func (s *userService) GetTotalUsers() (int64, error) {
	var total int64
	if err := s.db.Model(&models.User{}).Count(&total).Error; err != nil {
		return 0, errors.New("failed to count users")
	}
	return total, nil
}
