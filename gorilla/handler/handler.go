package handler

import (
	"net/http"

	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gorilla/mux"
)

type Handler struct {
	Student *postgres.StudentRepo
}

func NewHandler(handler Handler) *http.Server {
	m := mux.NewRouter()

	m.HandleFunc("/student/{id}", handler.student)
	return &http.Server{Handler: m}
}

type Book struct {
	Name, Author, Publisher string
}
