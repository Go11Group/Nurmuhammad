package gin

import (
	"exam/dbcon"

	"github.com/gin-gonic/gin"
)

// create handler struct to use methods for each table
type Handler struct {
	User       *dbcon.UserRepo
	Course     *dbcon.CourseRepo
	Lesson     *dbcon.LessonRepo
	Enrollment *dbcon.EnrollmentRepo
}

func ConnectGin(handler *Handler) *gin.Engine {
	// create gin
	c := gin.Default()

	// User API
	c.GET("/user", handler.GetAllUser)
	c.GET("/user/:id", handler.GetUser)
	c.POST("/user", handler.InsertToUser)
	c.DELETE("/user/:id", handler.DeleteUser)
	c.PUT("/user/:id", handler.UpdatedUser)

	// course API
	c.GET("/course", handler.GetAllCourse)
	c.GET("/course/:id", handler.GetCourse)
	c.POST("/course", handler.InsertToCourse)
	c.DELETE("/course/:id", handler.DeleteCourse)
	c.PUT("/course/:id", handler.UpdatedCourse)

	// lesson API
	c.GET("/lesson", handler.GetAllLesson)
	c.GET("/lesson/:id", handler.GetLesson)
	c.POST("/lesson", handler.InsertToLesson)
	c.DELETE("/lesson/:id", handler.DeleteLesson)
	c.PUT("/lesson/:id", handler.UpdatedLesson)

	// enrollment API
	c.GET("/enrollment", handler.GetAllEnrollment)
	c.GET("/enrollment/:id", handler.GetEnrollment)
	c.POST("/enrollment", handler.InsertToEnrollment)
	c.DELETE("/enrollment/:id", handler.DeleteEnrollment)
	c.PUT("/enrollment/:id", handler.UpdatedEnrollment)

	// Extra API
	c.GET("/users/:user_id/courses", handler.GetUserCourses)
	c.GET("/courses/:course_id/lessons", handler.CourseLessons)
	c.GET("/courses/:course_id/enrollments", handler.CourseUsers)
	c.GET("/users/search", handler.SearchUsers)
	c.GET("/courses/popular", handler.GetPopularCourses)

	// return gin to listen using run()
	return c
}
