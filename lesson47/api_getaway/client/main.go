package main

import (
	"gateway/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient(":50052")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	r := api.NewHandler(conn)
	r.Run()
}
