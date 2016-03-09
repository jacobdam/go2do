package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/datastore"
)

func UserShow(c *gin.Context) {

}

func UserCreate(c *gin.Context) {
	ds := getDS(c)
	reg := datastore.UserRegistration{}
	if e := c.BindJSON(&reg); e != nil {
		renderError(c, e)
		return
	}

	ds.CreateUser(&reg)

	c.JSON(200, gin.H{})
}
