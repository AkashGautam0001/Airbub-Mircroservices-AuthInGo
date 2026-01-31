package db

import (
	"AuthInGo/models"
	"database/sql"
)

type PermissionsRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionsRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionsRepository(db *sql.DB) PermissionsRepository {
	return &PermissionsRepositoryImpl{
		db: db,
	}
}

func (r *PermissionsRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?;"
	row := r.db.QueryRow(query, id)

	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
func (r *PermissionsRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?;"
	row := r.db.QueryRow(query, name)
	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
func (r *PermissionsRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	permissions := []*models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r *PermissionsRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO permissions (name, description, resource, action) VALUES (?, ?, ?, ?);"
	result, err := r.db.Exec(query, name, description, resource, action)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetPermissionById(id)
}

func (r *PermissionsRepositoryImpl) DeletePermissionById(id int64) error {
	query := "DELETE FROM permissions WHERE id = ?;"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PermissionsRepositoryImpl) UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ? WHERE id = ?;"
	_, err := r.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, err
	}
	return r.GetPermissionById(id)
}
