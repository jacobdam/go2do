package app

import (
	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/handlers"
)

func configRoutes(e *gin.Engine) {
	// e.GET("/users/:id", handlers.UserShow)
	// e.POST("/users", handlers.UserCreate)
	e.POST("/todos", handlers.TodoCreate)
}
