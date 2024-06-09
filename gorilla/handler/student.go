package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/gorilla/mux"
)

func (h *Handler) student(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)
	switch r.Method {
	case "GET":
		h.StudentGetById(w, r)
	case "DELETE":
		fmt.Println("Handling DELETE request")
		h.StudentDeleteById(w, r)
	case "PUT":
		h.StudentUpdateById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("405 - Method Not Allowed"))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}

}

func (h *Handler) StudentGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, err := h.Student.GetById(id)
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

func (h *Handler) StudentDeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Student.DeleteById(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while Delete, err: %s", err.Error())))
		return
	}
	w.Write([]byte("Succes to delete student"))
}

func (h *Handler) StudentUpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in id, %s", err)))
	}
	var student model.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in id, %s", err)))
	}
	err = h.Student.UpdateById(&student, id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error on update, %s", err)))
	} else {
		w.Write([]byte("Succes to update"))
	}
}
