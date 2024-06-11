package handler

import (
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) problem(w *gin.Context) {
	switch w.Request.Method {
	case "GET":
		h.ProblemGetById(w)
	case "DELETE":
		h.ProblemDeleteById(w)
	case "PUT":
		h.ProblemUpdateById(w)
	default:
		w.JSON(http.StatusMethodNotAllowed, gin.H{"message": "405 - Method Not Allowed"})
	}

}

func (h *Handler) ProblemGetById(w *gin.Context) {
	id := w.Param("id")

	book, err := h.Problem.ProblemsGetById(id)
	if err != nil {
		w.Writer.Write([]byte("Error get user id : " + id))
		return
	}

	w.JSON(http.StatusFound, book)

}

func (h *Handler) ProblemDeleteById(w *gin.Context) {
	id := w.Param("id")

	err := h.Problem.ProblemDeleteById(id)
	if err != nil {
		w.Writer.Write([]byte("Error delete user id : " + id))
		return
	}
	w.Writer.Write([]byte("succes to delete user id : " + id))
}

func (h *Handler) ProblemUpdateById(w *gin.Context) {

	id, err := strconv.Atoi(w.Param("id"))
	if err != nil {
		w.Writer.Write([]byte("Error user id"))
	}
	var problem model.Problem
	if err := w.BindJSON(&problem); err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.Problem.ProblemUpdateById(&problem, id)
	if err != nil {
		w.Writer.Write([]byte("Error to update user id : " + strconv.Itoa(id)))
	} else {
		w.Writer.Write([]byte("succes to update user id : " + strconv.Itoa(id)))
	}
}
