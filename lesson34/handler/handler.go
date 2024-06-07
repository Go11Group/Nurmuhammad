package handler

import (
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"net/http"
)

type Handler struct {
	Student *postgres.StudentRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/student/", handler.student)

	return &http.Server{Handler: mux}
}

type Book struct {
	Name, Author, Publisher string
}
