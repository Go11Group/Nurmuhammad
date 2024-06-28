package handler

import (
	pb "gateway/genproto/transportService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetBus(ctx *gin.Context) {
	req := &pb.BusNumber{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Transport.GetBusSchedule(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackBusLocation(ctx *gin.Context) {
	req := &pb.BusNumber{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Transport.TrackBusLocation(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportTraffic(ctx *gin.Context) {
	req := &pb.ReportTrafficJamRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp, err := h.Transport.ReportTrafficJam(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
