package main

import (
	"context"
	"net"
	pb "new/genproto/generator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneratorServer
}

func (s *server) FindSurname(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	for i, v := range in.All {
		if in.Names == i {
			new := &pb.Response{Result: map[string]string{i: v}}
			return new, nil
		}
	}
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := server{}
	grpc := grpc.NewServer()
	pb.RegisterGeneratorServer(grpc, &s)
	err = grpc.Serve(listener)
	if err != nil {
		panic(err)
	}
}
