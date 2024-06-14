package gin

import (
	"exam/dbcon"
	"exam/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLesson(c *gin.Context) {

	filter := dbcon.FilterLesson{Limit: 10}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	lesson, err := h.Lesson.GetAllLesson(filter)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	c.JSON(http.StatusOK, lesson)

}

func (h *Handler) GetLesson(c *gin.Context) {
	id := c.Param("id")
	lesson, err := h.Lesson.GetLesson(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "error on Get lesson"})
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) InsertToLesson(c *gin.Context) {
	lesson := models.Lesson{}
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.Lesson.InsertToLesson(lesson)
	if err != nil {
		er := fmt.Sprintf("Error to Insert lesson : %s", err)
		c.Writer.Write([]byte(er))
	} else {
		c.JSON(http.StatusOK, lesson)
	}

}

func (h *Handler) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	err := h.Lesson.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, "delete succesful")
	}

}

func (h *Handler) UpdatedLesson(c *gin.Context) {
	id := c.Param("id")
	lesson := models.Lesson{}
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newlesson, err := h.Lesson.UpdateLesson(id, lesson)
	if err != nil {
		err = fmt.Errorf("error on Update lesson: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, newlesson)
	}
}
