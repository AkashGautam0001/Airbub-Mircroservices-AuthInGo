package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
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

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password

	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		fmt.Println("Error fetching user by email", err)
		return "", err
	}

	if user == nil {
		fmt.Println("User not found")
		return "", fmt.Errorf("no user found with email: %s", email)

	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		fmt.Println("Invalid password")
		return "", fmt.Errorf("invalid password")
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "secret")))

	if err != nil {
		fmt.Println("Error generating token", err)
		return "", err
	}
	fmt.Println("User logged in successfully")
	return tokenString, nil
}
