package app

import (
	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/handlers"
)

func configRoutes(e *gin.Engine) {
	// log.Println("configRoutes")
	// e.GET("/users/:id", handlers.UserShow)
	// e.POST("/users", handlers.UserCreate)
	e.POST("/todos", handlers.TodoCreate)
	e.GET("/todos/:id", handlers.TodoShow)
	e.GET("/todos", handlers.TodoList)
	e.DELETE("/todos/:id", handlers.TodoDelete)
	e.PATCH("/todos/:id", handlers.TodoUpdate)
}
