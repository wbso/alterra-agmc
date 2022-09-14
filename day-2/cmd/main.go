package main

import (
	"errors"
	"fmt"
	"os"

	"alterratwo/controllers"
	"alterratwo/routes"

	"github.com/joho/godotenv"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}

	c := controllers.New()
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
