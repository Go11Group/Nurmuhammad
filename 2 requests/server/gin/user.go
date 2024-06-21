package gin

import (
	"exam/dbcon"
	"exam/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get all users and you can use filter and limits
func (h *Handler) GetAllUser(c *gin.Context) {

	// there i am creating variable to filter and i am giving limit default 10 if user didn't
	// insert limit it gets 10 as default value
	filter := dbcon.Filter{Limit: 10}

	// there i am reading from query params and write it into filter
	// and i am checking is it giving error nil or any error.
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// there i am using my GetAllCourse function
	// and it returns me slice of courses and pages count
	users, err := h.User.GetAll(filter)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	// there i am writing it to api as json
	c.JSON(http.StatusOK, *users)

}

// this is the function wich gets only one user by id
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "error on Get user"})
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) InsertToUser(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.User.InsertToUser(user)
	if err != nil {
		er := fmt.Sprintf("Error to Insert user : %s", err)
		c.Writer.Write([]byte(er))
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := h.User.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, "delete succesful")
	}
}

func (h *Handler) UpdatedUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newuser, err := h.User.UpdateUser(id, user)
	if err != nil {
		err = fmt.Errorf("error on Update user: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusAccepted, *newuser)
	}
}

func (h *Handler) GetUserCourses(c *gin.Context) {
	id := c.Param("user_id")
	user, err := h.User.GetUserCourses(id)
	if err != nil {
		err = fmt.Errorf("error on Get user courses: %s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusAccepted, *user)

}

func (h *Handler) SearchUsers(c *gin.Context) {
	filter := dbcon.SearchFilter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	users, err := h.User.SearchUsers(filter)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	c.JSON(http.StatusOK, *users)
}
