package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/jacobdam/go2do/app/datastore"
)

func getDS(c *gin.Context) *datastore.DataStore {
	ds, _ := c.Get("ds")

	return ds.(*datastore.DataStore)
}
