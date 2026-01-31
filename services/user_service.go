package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) error
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

func (u *UserServiceImpl) GetUserById(id int) (*models.User, error) {
	fmt.Println("Getting User by ID in UserService")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) error {
	fmt.Println("Creating User in UserService")
	password := payload.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
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

	fmt.Println(user.Id, user.Email)

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	fmt.Println("jwtPayload", jwtPayload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "secret")))

	if err != nil {
		fmt.Println("Error generating token", err)
		return "", err
	}
	fmt.Println("User logged in successfully")
	return tokenString, nil
}
