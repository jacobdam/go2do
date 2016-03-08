package app

import (
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
  datastore.
}
