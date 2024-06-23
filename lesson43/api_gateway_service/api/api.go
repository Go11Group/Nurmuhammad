package api

import (
	"github.com/Go11Group/at_lesson/lesson43/api_gateway_service/api/handler"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	mux.POST("/station/create", h.ClientCreate)

	return &http.Server{Handler: mux, Addr: ":8081"}
}
