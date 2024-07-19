package api

import (
	"github.com/gin-gonic/gin"
	"new/api/handler"
)

func ConnectGin(handler *handler.Handler) *gin.Engine {

	c := gin.Default()

	c.GET("/user/:id")
	c.DELETE("/user/:id")

	return c
}
