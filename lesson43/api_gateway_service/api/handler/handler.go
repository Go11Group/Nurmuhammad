package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) ClientCreate(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	body := ctx.Request.Body

	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080"+url, body)
	fmt.Println("http://localhost:8080" + url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": string(resp)})
}
