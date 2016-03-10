package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ResponseError `json:"error"`
}

func renderError(c *gin.Context, err error) {
	res := ErrorResponse{ResponseError{Message: err.Error()}}
	log.Println(err)
	c.JSON(500, res)
}
