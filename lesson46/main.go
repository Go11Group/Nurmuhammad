package main

import (
	"log"
	"net"
	pb "new/genproto/weatherService"
	"new/service"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewWeatherService()
	grpc := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpc, s)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
