package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Go11Group/at_lesson/lesson43/api_gateway_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) StationCreate(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	body := ctx.Request.Body

	req, err := http.NewRequest("POST", "http://localhost:8080"+url, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
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
	var message string
	if err := json.Unmarshal(resp, &message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse response"})
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": message})
}

func (h *handler) StationGetId(ctx *gin.Context) {
	id := ctx.Param("id")
	getURL := fmt.Sprintf("http://localhost:8080/station/%s", id)
	resp, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var station models.Station
	fmt.Println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error user encode json",
		})
	}
	ctx.JSON(http.StatusOK, station)

}

func (h *handler) StationGetAll(ctx *gin.Context) {
	getURL := "http://localhost:8080/station"
	resp, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var station []models.Station
	fmt.Println(resp.Body)
	if err := json.NewDecoder(resp.Body).Decode(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error user encode json",
		})
	}
	ctx.JSON(http.StatusOK, station)
}

func (h *handler) StationDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	getURL := fmt.Sprintf("http://localhost:8080/station/%s", id)
	req, err := http.NewRequest("DELETE", getURL, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := h.client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}
	var message string
	if err := json.Unmarshal(res, &message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse response"})
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": message})

}

func (h *handler) StationUpdate(ctx *gin.Context) {
	body := ctx.Request.Body
	id := ctx.Param("id")
	getURL := fmt.Sprintf("http://localhost:8080/station/%s", id)

	req, err := http.NewRequest("PUT", getURL, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := h.client.Do(req)
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
	var message string
	if err := json.Unmarshal(resp, &message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse response"})
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": message})

}
