package gin

import (
	"exam/dbcon"
	"exam/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// get all enrollments and you can use filter and limits
func (h *Handler) GetAllEnrollment(c *gin.Context) {

	// there i am creating variable to filter and i am giving limit default 10 if user didn't
	// insert limit it gets 10 as default value
	filter := dbcon.FilterEnrollment{Limit: 10}

	// there i am reading from query params and write it into filter
	// and i am checking is it giving error nil or any error.
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// there i am using my GetAllCourse function
	// and it returns me slice of courses and pages count
	enrollments, err := h.Enrollment.GetAllEnrollment(filter)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	// there i am writing it to api as json
	c.JSON(http.StatusOK, enrollments)

}

// this is the function wich gets only one enrollment by id
func (h *Handler) GetEnrollment(c *gin.Context) {
	id := c.Param("id")
	enrollment, err := h.Enrollment.GetEnrollment(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "error on Get enrollment"})
	}
	c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) InsertToEnrollment(c *gin.Context) {
	enrollment := models.Enrollment{EnrollmentDate: time.Now()}
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.Enrollment.InsertToEnrollment(enrollment)
	if err != nil {
		er := fmt.Sprintf("Error to Insert enrollment : %s", err)
		c.Writer.Write([]byte(er))
	} else {
		c.JSON(http.StatusOK, enrollment)
	}

}

func (h *Handler) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	err := h.Enrollment.DeleteEnrollment(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, "delete succesful")
	}

}

func (h *Handler) UpdatedEnrollment(c *gin.Context) {
	id := c.Param("id")
	enrollment := models.Enrollment{}
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newenrollment, err := h.Enrollment.UpdateEnrollment(id, enrollment)
	if err != nil {
		err = fmt.Errorf("error on Update enrollment: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, newenrollment)
	}
}
