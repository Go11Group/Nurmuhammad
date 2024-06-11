package handler

import (
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) user(w *gin.Context) {
	switch w.Request.Method {
	case "GET":
		h.UserGetById(w)
	case "DELETE":
		h.UserDeleteById(w)
	case "PUT":
		h.UserUpdateById(w)
	default:
		w.JSON(http.StatusMethodNotAllowed, gin.H{"message": "405 - Method Not Allowed"})
	}

}

func (h *Handler) UserGetById(w *gin.Context) {

	id := w.Param("id")

	book, err := h.User.UserGetById(id)
	if err != nil {
		w.JSON(http.StatusBadGateway, gin.H{"message": "error on Get user"})
		return
	}
	w.JSON(http.StatusOK, book)
}

func (h *Handler) UserDeleteById(w *gin.Context) {
	id := w.Param("id")
	err := h.User.UserDeleteById(id)
	if err != nil {
		w.JSON(http.StatusBadGateway, gin.H{"message": "error on delete user"})
		return
	}
	w.Writer.Write([]byte("Success to delete user with ID: " + id))
}

func (h *Handler) UserUpdateById(w *gin.Context) {

	id, err := strconv.Atoi(w.Param("id"))
	if err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": "error on id user"})
	}
	var user model.User

	if err := w.BindJSON(&user); err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	w.JSON(http.StatusOK, user)

	err = h.User.UserUpdateById(&user, id)
	if err != nil {
		w.Writer.Write([]byte("Error to update user : "))
	} else {
		w.Writer.Write([]byte("Success to update user : "))
	}
}
