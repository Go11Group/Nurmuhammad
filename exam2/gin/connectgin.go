package gin

import (
	"exam/dbcon"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User       *dbcon.UserRepo
	Course     *dbcon.CourseRepo
	Lesson     *dbcon.LessonRepo
	Enrollment *dbcon.EnrollmentRepo
}

func ConnectGin(handler *Handler) *gin.Engine {
	c := gin.Default()
	c.GET("/user", handler.GetAllUser)
	return c
}
