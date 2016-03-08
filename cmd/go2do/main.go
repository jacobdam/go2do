package main

import (
	"os"

	"github.com/jacobdam/go2do/app"
	"golang.org/x/crypto/bcrypt"
)

func f(pw string) string {
	out, _ := bcrypt.GenerateFromPassword([]byte(pw), 0)
	return string(out)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("APP_ENV")
	if port == "" {
		env = "development"
	}
	app := app.App{Port: port, Env: env}
	app.Run()
}
