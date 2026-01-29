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
	uc.UserService.GetUserById()
	w.Write([]byte("User GetUserById Endpoint"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User CreateUser Endpoint"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")

	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
		w.Write([]byte("Something went wrong while login"))
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		w.Write([]byte("Invalid request payload"))
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "User login failed", err)
		return
	}

	response := map[string]any{
		"message": "User logged in successfully",
		"data":    jwtToken,
		"success": true,
		"error":   nil,
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", response)
}
