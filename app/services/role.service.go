package services

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/dto/response"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/models"
	"errors"

	"gorm.io/gorm"
)

type RoleService interface {
	GetAllRoles() ([]response.RoleResponse, error)
	CreateRole(req request.CreateRoleRequest) (*response.RoleResponse, error)
	UpdateRole(id string, req request.UpdateRoleRequest) (*response.RoleResponse, error)
	DeleteRole(id string) (*response.RoleResponse, error)
}

type roleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{db: db}
}

func (s *roleService) GetAllRoles() ([]response.RoleResponse, error) {
	var roles []models.Role
	// Find all role
	if err := s.db.Preload("Users").Find(&roles).Error; err != nil {
		return nil, errors.New("role not found")
	}

	rolesResponse := helpers.ToRolesResponse(roles)
	return rolesResponse, nil
}

func (s *roleService) CreateRole(req request.CreateRoleRequest) (*response.RoleResponse, error) {
	// Check if role exist
	var existingRole models.Role
	if err := s.db.Where("role_name = ?", req.Name).First(&existingRole).Error; err == nil {
		return nil, errors.New("role with this name is already exist")
	}

	// Create role
	role := models.Role{
		RoleName: req.Name,
	}
	if err := s.db.Create(&role).Error; err != nil {
		return nil, errors.New("failed to create role")
	}

	roleResponse := helpers.ToRoleResponse(&role)
	return &roleResponse, nil
}

func (s *roleService) UpdateRole(id string, req request.UpdateRoleRequest) (*response.RoleResponse, error) {
	var existing models.Role

	// Check if exist
	if err := s.db.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	// Gunakan MAP untuk update selective fields
	updateData := map[string]interface{}{}

	// Name: update kalau dikirim dan berbeda
	if req.Name != "" && req.Name != existing.RoleName {
		var temp models.Role
		err := s.db.Where("role_name = ? AND id != ?", req.Name, id).First(&temp).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if err == nil {
			return nil, errors.New("role with this name is already exists")
		}
		updateData["name"] = req.Name
	}

	// kalau user tidak kirim apapun
	if len(updateData) == 0 {
		return nil, errors.New("no field to update")
	}

	// update
	if err := s.db.Model(&existing).Updates(updateData).Error; err != nil {
		return nil, errors.New("failed to update role")
	}

	// reload data baru
	s.db.First(&existing, id)

	res := helpers.ToRoleResponse(&existing)
	return &res, nil
}

func (s *roleService) DeleteRole(id string) (*response.RoleResponse, error) {
	// Find role by id
	var role models.Role
	if err := s.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}

	// Delete role
	if err := s.db.Delete(&role).Error; err != nil {
		return nil, errors.New("failed to delete role")
	}

	roleResponse := helpers.ToRoleResponse(&role)
	return &roleResponse, nil
}
