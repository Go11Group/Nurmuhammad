package gin

import (
	"exam/dbcon"
	"exam/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all courses and you can use filter and limits
func (h *Handler) GetAllCourse(c *gin.Context) {

	// there i am creating variable to filter and i am giving limit default 10 if user didn't
	// insert limit it gets 10 as default value
	filter := dbcon.FilterCourse{Limit: 10}

	// there i am reading from query params and write it into filter
	// and i am checking is it giving error nil or any error.
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// there i am using my GetAllCourse function
	// and it returns me slice of courses and pages count
	courses, err := h.Course.GetAllCourse(filter)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	// there i am writing it to api as json
	c.JSON(http.StatusOK, *courses)

}

// this is the function wich gets only one course by id
func (h *Handler) GetCourse(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Course.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "error on Get course"})
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) InsertToCourse(c *gin.Context) {
	course := models.Course{}
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.Course.InsertToCourse(course)
	if err != nil {
		er := fmt.Sprintf("Error to Insert course : %s", err)
		c.Writer.Write([]byte(er))
	} else {
		c.JSON(http.StatusOK, course)
	}

}

func (h *Handler) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	err := h.Course.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, "delete succesful")
	}

}

func (h *Handler) UpdatedCourse(c *gin.Context) {
	id := c.Param("id")
	course := models.Course{}
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newcourse, err := h.Course.UpdateCourse(id, course)
	if err != nil {
		err = fmt.Errorf("error on Update course: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, *newcourse)
	}
}

func (h *Handler) CourseLessons(c *gin.Context) {
	id := c.Param("course_id")
	course, err := h.Course.GetCourseLessons(id)
	if err != nil {
		err = fmt.Errorf("error on Get course's lessons: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, *course)
	}
}

func (h *Handler) CourseUsers(c *gin.Context) {
	id := c.Param("course_id")
	course, err := h.Course.GetCourseUsers(id)
	if err != nil {
		err = fmt.Errorf("error on Get course's users: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, *course)
	}
}

func (h *Handler) GetPopularCourses(c *gin.Context) {
	var timePeriod models.TimePeriod
	if err := c.ShouldBindQuery(&timePeriod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.Course.GetPopularCourses(timePeriod)
	if err != nil {
		err = fmt.Errorf("error on Get popular courses: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, *result)
	}
}
