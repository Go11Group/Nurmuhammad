package main

import (
	"database/sql"
	"fmt"
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
	var choice int
	fmt.Printf(`
1-run Transport proto
2-run Weather proto`)
	fmt.Printf("\n\nChoose one of them>>> ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		RunTransport(db)
	case 2:
		RunWeather()
	case 3:
		log.Fatal("ARE YOU OK THERE IS NO THIS KIND OF CHOICE!!!")
	}
}

func RunTransport(db *sql.DB) {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewTransportService(postgres.ConnectBook(db))
	grpc := grpc.NewServer()
	p.RegisterTransportServiceServer(grpc, s)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

func RunWeather() {
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
