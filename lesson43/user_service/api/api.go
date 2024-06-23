package api

import (
	"database/sql"
	"net/http"
	"user/api/handler"

	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)

	mux.POST("user/create", h.CreateUser)
	mux.GET("user", h.GetAllUser)
	mux.GET("user/:id", h.GetByIdUser)
	mux.PUT("user/:id", h.UpdateUser)
	mux.DELETE("user/:id", h.DeleteUser)
	return &http.Server{Handler: mux, Addr: ":8082"}
}
