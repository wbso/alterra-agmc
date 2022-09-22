package main

import (
	"alterrasix/app/book"
	"alterrasix/app/user"
	"alterrasix/repository"
	"fmt"
	"os"

	"alterrasix/config"
	"alterrasix/handler"
	"github.com/joho/godotenv"
)

func run() error {
	err := godotenv.Load()

	// connect db to mysql or sqlite
	db, err := config.ConnectMysql()
	// db, err := config.ConnectSQlite()
	if err != nil {
		return fmt.Errorf("error while connecting to database %w", err)
	}
	// create repository instance
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository()

	// create service instance
	userService := user.New(userRepo)
	bookService := book.New(bookRepo)

	// create handler instance
	h := handler.New(userService, bookService)

	return h.ListenAndServe(envDefault("PORT", ":5000"))
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func envDefault(key, def string) string {
	env := os.Getenv(key)
	if env == "" {
		return def
	}
	return env
}
