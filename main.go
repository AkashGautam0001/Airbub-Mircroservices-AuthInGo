package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
)

func main() {
	config.Load()

	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	app.Run()
}
