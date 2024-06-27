package server

import (
	"context"
	pb "new/genproto/generator"
	"new/storage/tables"
)

type server struct {
	pb.UnimplementedGeneratorServer
	b *tables.BookRepo
}

func NewBookService(g *tables.BookRepo) *server {
	return &server{b: g}
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
