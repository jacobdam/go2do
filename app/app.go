package app

import (
	"log"
	"net/http"

	"github.com/jacobdam/go2do/app/datastore"
)

type AppError struct {
	error
}

type App struct {
	DataStore *datastore.DataStore
	Handler   http.Handler
	Port      string
	Env       string
}

func NewApp() (app *App, err error) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	app = &App{Handler: h}
	return
}

func (app *App) Run() {
	log.Println("Start listening on :" + app.Port)
	http.ListenAndServe(":"+app.Port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}))
}
