package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{UserService: _userService}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")
	id := r.Context().Value("userID")
	fmt.Println("User ID from context:", id)

	data, err := uc.UserService.GetUserById(id.(int))
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching user", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", data)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	err := uc.UserService.CreateUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User created successfully", nil)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Invalid input data", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
