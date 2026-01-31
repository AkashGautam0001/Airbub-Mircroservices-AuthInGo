package db

import (
	"AuthInGo/models"
	"database/sql"
)

type RolePermissionsRepository interface {
	GetRolePermissionById(roleId int64, permissionId int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionsRepositoryImpl struct {
	db *sql.DB
}

// func NewRolePermissionsRepository(db *sql.DB) RolePermissionsRepository {
// 	return &RolePermissionsRepositoryImpl{
// 		db: db,
// 	}
// }

// TODO: implement RolePermissionsRepository methods
