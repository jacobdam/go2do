package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/datastore"
)

func TodoCreate(c *gin.Context) {
	ds := getDS(c)
	todo := datastore.Todo{}
	if e := c.BindJSON(&todo); e != nil {
		renderError(c, e)
		return
	}
	ds.CreateTodo(&todo)

	c.JSON(http.StatusCreated, todo)
}
