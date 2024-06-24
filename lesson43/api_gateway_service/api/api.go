package api

import (
	"github.com/Go11Group/at_lesson/lesson43/api_gateway_service/api/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	mux.POST("/station/create", h.StationCreate)
	mux.GET("/station/:id", h.StationGetId)
	mux.GET("/station", h.StationGetAll)
	mux.DELETE("/station/:id", h.StationDelete)
	mux.PUT("/station/:id", h.StationUpdate)

	return &http.Server{Handler: mux, Addr: ":8081"}
}
