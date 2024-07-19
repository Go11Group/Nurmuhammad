package handler

import (
	"net/http"
	"new/api/auth"
	"new/storage/redis"
	"new/structs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Login(c *gin.Context) {
	req := structs.UserInfo{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	}
	res, err := h.User.GetUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
	}
	tk, err := auth.GeneratedAccessJWTToken(res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
	}
	c.JSON(200, gin.H{"acces": tk})
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "user id is incorrect"})
		return
	}
	res, err := redis.GetUserById(id, h.User)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, &res)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "user id is incorrect"})
		return
	}
	err = h.User.DeleteUserById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "succesfully deleted"})
}
