package main

import "golang.org/x/crypto/bcrypt"

func f(pw string) string {
	out, _ := bcrypt.GenerateFromPassword([]byte(pw), 0)
	return string(out)
}

func main() {
	app := app.App{}
}
