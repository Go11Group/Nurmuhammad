package service

import (
	pb "new/genproto/weatherService"
)

type server struct {
	pb.UnimplementedWeatherServiceServer
}

func NewWeatherService() *server {
	return &server{}
}
