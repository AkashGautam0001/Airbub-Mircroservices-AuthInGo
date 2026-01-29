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

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Invalid input data", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}
