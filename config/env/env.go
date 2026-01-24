package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env file found")
	}

	fmt.Println("Environment variables loaded")
}

func GetString(key string, fallback string) string {
	load()

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return value
}
