package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

// Config : Config holds server configuration
type Config struct {
	Addr string // PORT
}

// Application : Application struct
type Application struct {
	Config Config
	Store  db.Storage
}

// Constructor
func NewConfig() Config {
	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
		Store:  *db.NewStorage(),
	}
}

func (app *Application) Run() error {
	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error connecting to DB", err)
		return err
	}
	ur := repo.NewUsersRepository(db)
	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionsRepository(db)
	urr := repo.NewUserRolesRepository(db)

	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr, rpr, urr)

	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)

	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter, rRouter), // setup chi router
		ReadTimeout:  10 * time.Second,                     // 10 seconds
		WriteTimeout: 10 * time.Second,                     // 10 seconds
		IdleTimeout:  30 * time.Second,                     // 30 seconds
	}

	fmt.Println("Starting Server on", app.Config.Addr)
	return server.ListenAndServe()

}
