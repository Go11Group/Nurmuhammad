package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAllSolved(w http.ResponseWriter, r *http.Request) {
	users, err := h.Solved.GetAll()
	if err != nil {
		w.Write([]byte("error on read"))
		return
	}

	for _, v := range users {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error while Encode, err: %s", err.Error())))
		}
	}
}
