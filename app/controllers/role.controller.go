package controllers

import (
	"backend-dinakom/app/dto/request"
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/services"

	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	GetAllRoles(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
}

type roleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &roleController{roleService: roleService}
}

// GetAllRoles godoc
// @Summary Get All Roles
// @Description get all roles data
// @Tags Roles (Admin)
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /roles [get]
// @Security BearerAuth
func (ctrl *roleController) GetAllRoles(c *fiber.Ctx) error {

	roles, err := ctrl.roleService.GetAllRoles()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Roles retrieved successfully", roles)
}

// CreateRole godoc
// @Summary Create new role
// @Description Create a new role
// @Tags Roles (Admin)
// @Accept json
// @Produce json
// @Param request body request.CreateRoleRequest true "Role create data"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /roles/create [post]
// @Security BearerAuth
func (ctrl *roleController) CreateRole(c *fiber.Ctx) error {
	var req request.CreateRoleRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	role, err := ctrl.roleService.CreateRole(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.CreatedResponse(c, "Role created successfully", role)
}

// Update Role godoc
// @Summary Update Role
// @Description Update a Role data
// @Tags Roles (Admin)
// @Produce  json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /role/{id}/update [patch]
// @Security BearerAuth
// @Param request body request.UpdateRoleRequest true "Role update data"
func (ctrl *roleController) UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var req request.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	role, err := ctrl.roleService.UpdateRole(id, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Role updated successfully", role)
}

// Delete role by roleID godoc
// @Summary Delete role by roleID
// @Description Delete role by roleID
// @Tags Roles (Admin)
// @Produce  json
// @Param id path int true "role ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /roles/{id} [delete]
// @Security BearerAuth
func (ctrl *roleController) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	role, err := ctrl.roleService.DeleteRole(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, "Role deleted successfully", role)
}
