package main

import (
	"errors"
	"fmt"
	"os"

	"alterratwo/config"
	"alterratwo/controllers"
	"alterratwo/models"
	"alterratwo/routes"

	"github.com/joho/godotenv"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}

	// connect db to mysql or sqlite
	db, err := config.ConnectMysql()
	// db, err := config.ConnectSQlite()
	if err != nil {
		return fmt.Errorf("error while connecting to database %w", err)
	}
	m, err := models.New(db)
	if err != nil {
		return err
	}

	// m.GetAllUser()
	// return nil
	c := controllers.New(m)
	r := routes.New(c)
	return r.ListenAndServe(":1323")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
