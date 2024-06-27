package main

import (
	"log"
	"net"
	pb "new/genproto/generator"
	server "new/service"
	"new/storage"
	"new/storage/tables"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.ConnnectDb()
	if err != nil {
		log.Fatal("Error on connect db", err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewBookService(tables.ConnectBook(db))
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, s)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
