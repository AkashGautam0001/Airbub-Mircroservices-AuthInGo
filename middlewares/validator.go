package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO

		if err := utils.ReadJsonRequest(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", err)
			return
		}

		//Understand it deeply
		req_context := r.Context() //
		ctx := context.WithValue(req_context, "payload", payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO

		if err := utils.ReadJsonRequest(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", err)
			return
		}

		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
