package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)

	chiRouter.Use(middlewares.RateLimitMiddleware)
	// TODO: Add IP specific rate limiting, and jwt based also

	// Add reverse proxy
	chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.in", "/fakestoreservice"))

	chiRouter.Get("/ping", controllers.PingHandler)

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)
	return chiRouter

}
