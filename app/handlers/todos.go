package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/datastore"
)

func TodoShow(c *gin.Context) {
	ds := getDS(c)
	id := bson.ObjectIdHex(c.Param("id"))
	todo, e := ds.FindTodo(id)
	if e != nil {
		renderError(c, e)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func TodoCreate(c *gin.Context) {
	ds := getDS(c)
	todo := datastore.Todo{}
	e := c.BindJSON(&todo)
	if e != nil {
		renderError(c, e)
		return
	}

	e = ds.CreateTodo(&todo)
	if e != nil {
		renderError(c, e)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func TodoList(c *gin.Context) {
	ds := getDS(c)
	todos, e := ds.ListTodos()
	if e != nil {
		renderError(c, e)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func TodoDelete(c *gin.Context) {
	ds := getDS(c)
	id := bson.ObjectIdHex(c.Param("id"))
	e := ds.DeleteTodo(id)
	if e != nil {
		renderError(c, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func TodoUpdate(c *gin.Context) {
	ds := getDS(c)
	id := bson.ObjectIdHex(c.Param("id"))
	todo := datastore.Todo{}
	e := c.BindJSON(&todo)
	if e != nil {
		renderError(c, e)
		return
	}

	todo.Id = id
	e = ds.UpdateTodo(&todo)
	if e != nil {
		renderError(c, e)
		return
	}

	c.JSON(200, todo)
}
