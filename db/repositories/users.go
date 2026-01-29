package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UsersRepository interface {
	GetByID() (*models.User, error)
	Create(username string, email string, password string) error
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteByID() error
}

type UsersRepositoryImpl struct {
	db *sql.DB
}

func NewUsersRepository(_db *sql.DB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		db: _db,
	}
}

func (u *UsersRepositoryImpl) Create(username string, email string, password string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?);"

	result, err := u.db.Exec(query, username, email, password)

	if err != nil {
		fmt.Println("Error creating user in User Repository", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error creating user in User Repository", rowErr)
		return rowErr
	}

	fmt.Println("User created in User Repository", rowsAffected)
	return nil
}

func (u *UsersRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Creating user in User Repository")

	// Step 1 : Prepare Query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?;"

	// Step 2 : Execute Query
	row := u.db.QueryRow(query, 1)

	// Step 3 : Process the result
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error querying User by ID in User Repository", err)
			return nil, err
		}
	}
	fmt.Println("Querying User by ID in User Repository", user)

	return user, nil
}

func (u *UsersRepositoryImpl) GetAll() ([]*models.User, error) {
	fmt.Println("Getting all users in User Repository")
	return nil, nil
}

func (u *UsersRepositoryImpl) DeleteByID() error {
	fmt.Println("Deleting user in User Repository")
	return nil
}

func (u *UsersRepositoryImpl) GetByEmail(email string) (*models.User, error) {

	query := "SELECT id, email, password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error querying User by email in User Repository", err)
			return nil, err
		}
	}
	fmt.Println("Getting user by email in User Repository")
	return user, nil
}
