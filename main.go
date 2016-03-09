package main

import (
	"log"
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
	if env == "" {
		env = "development"
	}
	app, e := app.NewApp(env, port)
	if e != nil {
		log.Println(e)
		return
	}

	// ur := datastore.UserRegistration{Username: "phuc.dam", Password: "secret"}
	// u, e := app.DataStore.CreateUser(&ur)
	// if e != nil {
	// 	log.Println(e)
	// } else {
	// 	log.Println(u)
	// }

	app.Run()
}
