package app

import (
	config "AuthInGo/config/env"
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
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil,              // setup chi router
		ReadTimeout:  10 * time.Second, // 10 seconds
		WriteTimeout: 10 * time.Second, // 10 seconds
		IdleTimeout:  30 * time.Second, // 30 seconds
	}

	fmt.Println("Starting Server on", app.Config.Addr)
	return server.ListenAndServe()

}
