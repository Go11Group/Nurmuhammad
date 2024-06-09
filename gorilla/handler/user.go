package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/gorilla/mux"
)

func (h *Handler) user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)
	switch r.Method {
	case "GET":
		h.UserGetById(w, r)
	case "DELETE":
		fmt.Println("Handling DELETE request")
		h.UserDeleteById(w, r)
	case "PUT":
		h.UserUpdateById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 - Method Not Allowed"))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}

}

func (h *Handler) UserGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, err := h.User.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Encode, err: %s", err.Error())))
	}
}

func (h *Handler) UserDeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.User.DeleteById(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while Delete, err: %s", err.Error())))
		return
	}
	w.Write([]byte("Succes to delete users"))
}

func (h *Handler) UserUpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in id, %s", err)))
	}
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in id, %s", err)))
	}
	err = h.User.UpdateById(&user, id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error on update, %s", err)))
	} else {
		w.Write([]byte("Succes to update"))
	}
}
