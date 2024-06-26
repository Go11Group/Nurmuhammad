package main

import (
	"context"
	"fmt"
	"log"
	pb "new/genproto/generator"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	Menu(ctx, gen)

}

func Menu(ctx context.Context, gen pb.GeneratorClient) {
	var request int
	fmt.Println(`
	1 - Add Book
	2 - Search Book
	3 - Borrow Book
	4 - Exit`)
	fmt.Print("Enter number which request you choose>>> ")
	fmt.Scan(&request)
	switch request {
	case 1:
		AddBook(ctx, gen)
	case 2:
		SearchBook(ctx, gen)
	case 3:
		BorrowBook(ctx, gen)

	}
}

func AddBook(ctx context.Context, gen pb.GeneratorClient) {
	book := pb.AddBookRequest{}
	fmt.Print("Enter book name>>> ")
	fmt.Scan(&book.Title)
	fmt.Print("Enter book author name>>> ")
	fmt.Scan(&book.Author)
	fmt.Print("Enter book published year>>> ")
	fmt.Scan(&book.YearPublished)

	book2, err := gen.AddBook(ctx, &book)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("book created succesfully:", book2)
	}
	Menu(ctx, gen)
}

func BorrowBook(ctx context.Context, gen pb.GeneratorClient) {
	book := pb.BorrowBookRequest{}
	fmt.Print("Enter book id>>> ")
	fmt.Scan(&book.BookId)
	fmt.Print("Enter user id>>> ")
	fmt.Scan(&book.UserId)
	req, _ := gen.BorrowBook(ctx, &book)
	if req.Succes {
		fmt.Println("Succes to borrow")
	} else {
		fmt.Println("Unsucces to borrow")
	}
	Menu(ctx, gen)

}

func SearchBook(ctx context.Context, gen pb.GeneratorClient) {
	query := pb.SearchBookRequest{}
	query.Query = "select id,name,author,year from books"
	book, err := gen.SearchBook(ctx, &query)
	if err != nil {
		log.Println(err)
	}
	for _, v := range book.Books {
		fmt.Println(v)
	}
	Menu(ctx, gen)

}
