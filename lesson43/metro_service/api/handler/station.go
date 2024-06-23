package handler

import (
	"fmt"
	"github.com/Go11Group/at_lesson/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	stn := models.CreateStation{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) DeleteStation(ctx *gin.Context) {
	id := ctx.Param(":id")
	err := h.Station.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateStation(ctx *gin.Context) {
	id := ctx.Param(":id")
	station := models.CreateStation{}
	if err := ctx.BindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := h.Station.Update(id, &station)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, "OKAY")

}

func (h *handler) GetByIdStation(ctx *gin.Context) {
	id := ctx.Param(":id")

	station, err := h.Station.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, station)
}

func (h *handler) GetAllStation(ctx *gin.Context) {
	stations, err := h.Station.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, stations)
}
