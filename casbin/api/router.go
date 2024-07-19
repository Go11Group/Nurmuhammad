package api

import (
	"new/api/handler"
	"new/api/middleware"

	"github.com/gin-gonic/gin"
)

func ConnectGin(handler *handler.Handler) *gin.Engine {

	c := gin.Default()

	c.POST("/login", handler.Login)

	user := c.Group("/user")
	user.Use(middleware.Check)
	user.Use(middleware.CheckPermissionMiddleware(handler.Enforcer))

	user.GET("/:id", handler.GetUserProfile)
	user.DELETE("/:id", handler.DeleteUser)

	return c
}
