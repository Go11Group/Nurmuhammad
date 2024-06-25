package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "translater/protos/translate"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslaterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TranslateRequest{
		Word:            []string{"Hello", "World"},
		SourceLanguage:  "en",
		TargetLanguages: []string{"uz"},
	}
	res, err := c.GetTranslateWord(ctx, req)
	if err != nil {
		log.Fatalf("Could not translate: %v", err)
	}
	for key, value := range res.TranslatedWord {
		fmt.Printf("%s -> %s\n", key, value)
	}

}
