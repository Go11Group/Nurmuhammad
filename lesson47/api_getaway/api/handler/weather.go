package handler

import (
	pb "gateway/genproto/weatherService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetWeather(ctx *gin.Context) {
	req := &pb.Place{}

	resp, err := h.Weather.GetCurrentWeather(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetNextday(ctx *gin.Context) {
	req := &pb.Place{}

	resp, err := h.Weather.GetWeatherForecast(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportWeatherCondition(ctx *gin.Context) {
	req := &pb.ReportWeatherConditionRequest{}

	resp, err := h.Weather.ReportWeatherCondition(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
