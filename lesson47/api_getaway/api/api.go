package api

import (
	"gateway/api/handler"
	pb "gateway/genproto/transportService"
	p "gateway/genproto/weatherService"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewHandler(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	weather := p.NewWeatherServiceClient(conn)
	transport := pb.NewTransportServiceClient(conn)
	h := handler.NewHandler(weather, transport)
	router.GET("/weather", h.GetWeather)
	router.GET("/nextday/weather", h.GetNextday)
	router.PUT("/report/weather", h.ReportWeatherCondition)

	router.GET("/bus/schedule", h.GetBus)
	router.GET("/bus/location", h.TrackBusLocation)
	router.PUT("/report/traffic", h.ReportTraffic)

	return router
}
