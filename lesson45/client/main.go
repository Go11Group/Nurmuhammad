package main

import (
	"context"
	"fmt"
	pb "new/genproto/generator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	gen := pb.NewGeneratorClient(conn)
	var names string
	fmt.Scan(&names)
	req := &pb.Request{
		All:   map[string]string{"Nurmuhammad": "Meliqo'ziyev", "Faxriddin": "Rahimberdiyev"},
		Names: names,
	}
	resp, err := gen.FindSurname(context.Background(), req)

	if err != nil {
		panic(err)
	}
	fmt.Println("Seat - Name")
	for k, v := range resp.Result {
		fmt.Printf("%s   - %s\n", v, k)
	}
}
