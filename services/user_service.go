package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
}

type UserServiceImpl struct {
	userRepository db.UsersRepository
}

func NewUserService(_userRepository db.UsersRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Creating User in UserService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating User in UserService")
	password := "hashedPassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create("username1", "user1@exmple.com", hashedPassword)
	return nil
}

func (u *UserServiceImpl) LoginUser() error {
	fmt.Println("Login User in UserService")
	response := utils.CheckPasswordHash("hashedPassword", "$2a$10$PrlQ9muydfYibkA5l2H9jOqoeW2nxroxFTL2XRKOAq8V35351ReBu")
	fmt.Println("Login Successful", response)
	return nil
}
