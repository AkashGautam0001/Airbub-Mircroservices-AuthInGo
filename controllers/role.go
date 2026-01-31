package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")

	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role"))
		return
	}

	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", err)
		return
	}

	role, err := rc.RoleService.GetRoleById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get role", err)
		return
	}
	if role == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role not found"))
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role Fetched Successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to get roles", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles Fetched Successfully", roles)
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	roleIdStr := chi.URLParam(r, "roleId")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid User ID", err)
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Role ID", err)
		return
	}

	err = rc.RoleService.AssignRoleToUser(userId, roleId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to assign role to user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role assigned to user successfully", nil)
}
