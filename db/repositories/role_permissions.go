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

func NewRolePermissionsRepository(db *sql.DB) RolePermissionsRepository {
	return &RolePermissionsRepositoryImpl{
		db: db,
	}
}

func (r *RolePermissionsRepositoryImpl) GetRolePermissionById(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query := `
		SELECT role_id, permission_id, created_at, updated_at
		FROM role_permissions
		WHERE role_id = ? AND permission_id = ?;
	`
	row := r.db.QueryRow(query, roleId, permissionId)

	rolePermission := &models.RolePermission{}
	err := row.Scan(&rolePermission.RoleID, &rolePermission.PermissionID, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return rolePermission, nil
}

func (r *RolePermissionsRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {
	query := `
		SELECT role_id, permission_id, created_at, updated_at
		FROM role_permissions
		WHERE role_id = ?;
	`
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		if err := rows.Scan(&rolePermission.RoleID, &rolePermission.PermissionID, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	return rolePermissions, nil
}
func (r *RolePermissionsRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?);"
	_, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err
	}
	return &models.RolePermission{RoleID: roleId, PermissionID: permissionId}, nil
}

func (r *RolePermissionsRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?;"
	_, err := r.db.Exec(query, roleId, permissionId)
	return err
}
func (r *RolePermissionsRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	query := `
		SELECT role_id, permission_id, created_at, updated_at
		FROM role_permissions;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		if err := rows.Scan(&rolePermission.RoleID, &rolePermission.PermissionID, &rolePermission.CreatedAt, &rolePermission.UpdatedAt); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	return rolePermissions, nil
}
