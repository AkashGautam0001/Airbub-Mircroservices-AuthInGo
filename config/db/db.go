package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DB_NAME", "auth_dev")

	fmt.Println("Connecting to Database", cfg.DBName, cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to DB", err)
		return nil, err
	}

	fmt.Println("Trying to connect to database...")

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error connecting to DB", err)
		return nil, pingErr
	}

	fmt.Println("Connected to database successfully!", cfg.DBName)

	return db, nil
}
