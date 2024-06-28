package handler

import (
	"gateway/genproto/transportService"
	"gateway/genproto/weatherService"
)

type Handler struct {
	Weather   weatherService.WeatherServiceClient
	Transport transportService.TransportServiceClient
}

func NewHandler(Weather weatherService.WeatherServiceClient, Transport transportService.TransportServiceClient) *Handler {
	return &Handler{Weather: Weather, Transport: Transport}
}
