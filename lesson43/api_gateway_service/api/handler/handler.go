package handler

import (
	"net/http"
)

type handler struct {
	client http.Client
}

func NewHandler() *handler {
	return &handler{}
}
