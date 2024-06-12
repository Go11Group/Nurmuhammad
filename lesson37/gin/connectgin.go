package gin

import (
	"github.com/Go11Group/at_lesson/lesson37/psql"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *psql.UserRepo
}

func ConnectGin(handler Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/users", handler.GetAllUser)
	return r
}
