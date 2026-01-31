package db

import (
	"AuthInGo/models"
	"AuthInGo/utils"
	"database/sql"
	"fmt"
	"strings"
)

type UserRolesRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, permissionName string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRolesRepositoryImpl struct {
	db *sql.DB
}

func NewUserRolesRepository(db *sql.DB) UserRolesRepository {
	return &UserRolesRepositoryImpl{
		db: db,
	}
}

func (r *UserRolesRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?
	`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *UserRolesRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?);"
	_, err := r.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRolesRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?;"
	_, err := r.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRolesRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `
		SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM user_roles ur
		JOIN role_permissions rp ON ur.role_id = rp.role_id
		JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ?
	`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r *UserRolesRepositoryImpl) HasPermission(userId int64, permissionName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		JOIN role_permissions rp ON ur.role_id = rp.role_id
		JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND p.name = ?
	`
	var count int64
	err := r.db.QueryRow(query, userId, permissionName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRolesRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name = ?
	`
	var count int64
	err := r.db.QueryRow(query, userId, roleName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRolesRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil
	}

	query := `
		SELECT COUNT(*) = ?
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (?)
		GROUP BY ur.user_id
	`

	roleNamesStr := strings.Join(roleNames, ",")

	row := r.db.QueryRow(query, len(roleNames), userId, roleNamesStr)

	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return hasAllRoles, nil
}

func (r *UserRolesRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return false, nil
	}

	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = strings.TrimRight(placeholders, ",")

	query := fmt.Sprintf(`
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (%s)
	`, placeholders)

	args := utils.StringSliceToInterface(roleNames)
	args = append([]interface{}{userId}, args...)

	row := r.db.QueryRow(query, args...)

	var hasAnyRole bool
	if err := row.Scan(&hasAnyRole); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return hasAnyRole, nil
}
