-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- seeder data
-- INSERT INTO permissions (name, description, resource, action) VALUES
-- ('user:read', 'Permission to read user data', 'user', 'read'),
-- ('user:write', 'Permission to write user data', 'user', 'write'),
-- ('user:delete', 'Permission to delete user data', 'user', 'delete'),
-- ('user:manage', 'Permission to manage users', 'user', 'manage'),
-- ('role:read', 'Permission to read role data', 'role', 'read'),
-- ('role:write', 'Permission to write role data', 'role', 'write'),
-- ('role:delete', 'Permission to delete role data', 'role', 'delete'),
-- ('role:manage', 'Permission to manage roles', 'role', 'manage'),
-- ('permission:read', 'Permission to read permission data', 'permission', 'read'),
-- ('permission:write', 'Permission to write permission data', 'permission', 'write'),
-- ('permission:delete', 'Permission to delete permission data', 'permission', 'delete'),
-- ('permission:manage', 'Permission to manage permissions', 'permission', 'manage');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
