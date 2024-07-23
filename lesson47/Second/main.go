package main

import (
	"log"
	"net"
	p "new/genproto/transportService"
	pb "new/genproto/weatherService"
	"new/service"
	"new/storage"
	"new/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.ConnnectDb()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewTransportService(postgres.ConnectBook(db))
	s1 := service.NewWeatherService()
	grpc := grpc.NewServer()
	p.RegisterTransportServiceServer(grpc, s)
	pb.RegisterWeatherServiceServer(grpc, s1)
	log.Println("server is running on :50052 ...")
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
