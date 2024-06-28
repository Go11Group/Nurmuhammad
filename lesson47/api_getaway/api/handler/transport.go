package handler

import (
	pb "gateway/genproto/transportService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetBus(ctx *gin.Context) {
	req := &pb.BusNumber{}

	resp, err := h.Transport.GetBusSchedule(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackBusLocation(ctx *gin.Context) {
	req := &pb.BusNumber{}

	resp, err := h.Transport.TrackBusLocation(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportTraffic(ctx *gin.Context) {
	req := &pb.ReportTrafficJamRequest{}

	resp, err := h.Transport.ReportTrafficJam(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
