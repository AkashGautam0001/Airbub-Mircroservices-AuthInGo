package db

import (
	"fmt"
)

type UsersRepository interface {
	Create() error
}

type UsersRepositoryImpl struct {
	// db *sql.DB
}

func NewUsersRepository() *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		// db: db
	}
}

func (u *UsersRepositoryImpl) Create() error {
	fmt.Println("Creating user in User Repository")
	return nil
}
