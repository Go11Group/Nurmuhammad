package main

import (
	"context"
	"log"
	"net"
	pb "new/genproto/generator"
	"new/storage"
	"new/storage/tables"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneratorServer
	b *tables.BookRepo
}

func main() {
	db, err := storage.ConnnectDb()
	if err != nil {
		log.Fatal("Error on connect db", err)
	}
	defer db.Close()
	book := tables.ConnectBook(db)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := server{b: book}
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, &s)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	id := &pb.AddBookResponse{}
	err := s.b.CreateBook(req, id)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	books, err := s.b.SearchBook(req)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	err := s.b.AddBook(req)

	if err != nil {
		return &pb.BorrowBookResponse{Succes: false}, nil
	} else {
		return &pb.BorrowBookResponse{Succes: true}, nil
	}
}
