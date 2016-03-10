package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/datastore"
	"github.com/jacobdam/go2do/config"
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

func NewApp(env string, port string) (app *App, err error) {
	app = &App{Port: port, Env: env}
	e := app.createDatastore()
	if e != nil {
		err = AppError{e}
		return
	}

	e = app.createRouter()
	if e != nil {
		err = AppError{e}
		return
	}
	return
}

func (app *App) Run() {
	log.Println("Start listening on :" + app.Port)
	e := http.ListenAndServe(":"+app.Port, app.Handler)

	if e != nil {
		log.Fatal("Cannot start app: ", e.Error())
	}
}

func (app *App) createDatastore() error {
	dsConf := config.DB[app.Env]
	ds, err := datastore.NewDataStore(&dsConf)
	if err != nil {
		return AppError{err}
	}

	app.DataStore = ds
	return nil
}

func (app *App) createRouter() error {
	router := gin.New()
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	router.Use(func(c *gin.Context) {
		dsCopy := app.DataStore.Copy()
		defer dsCopy.Close()
		c.Set("ds", dsCopy)
		c.Next()
	})

	configRoutes(router)

	app.Handler = router
	return nil
}
