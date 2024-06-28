package handler

import (
	pb "gateway/genproto/weatherService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetWeather(ctx *gin.Context) {
	req := &pb.Place{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Weather.GetCurrentWeather(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetNextday(ctx *gin.Context) {
	req := &pb.Place{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Weather.GetWeatherForecast(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportWeatherCondition(ctx *gin.Context) {
	req := &pb.ReportWeatherConditionRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Weather.ReportWeatherCondition(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
