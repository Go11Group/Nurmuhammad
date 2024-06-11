package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllSolved(w *gin.Context) {
	users, err := h.Solved.GetAll()
	if err != nil {
		w.JSON(http.StatusBadGateway, gin.H{"message": "error on Get users"})
		return
	}

	for _, user := range users {
		w.JSON(http.StatusOK, user)
	}

}
