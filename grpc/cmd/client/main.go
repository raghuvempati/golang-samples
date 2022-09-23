package main

import (
	"context"
	"fmt"
	"log"

	api "go-samples/grpc/api/v1"

	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress = "localhost:8080"

func main() {
	conn, err := gRPC.Dial(serverAddress, gRPC.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewBookServiceClient(conn)

	// add book
	bookDTO := &api.Book{
		Title:       "Go Programming",
		Author:      "John Doe",
		Description: "Go is a programming language",
		Language:    "English",
		FinishTime:  timestamppb.Now(),
	}
	resCreate, err := client.CreateBook(context.Background(), &api.CreateBookRequest{Book: bookDTO})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
	}
	log.Printf("Book created with bid: %d\n", resCreate.BookId)

	// retrieve a book
	var bookId int64 = 1
	resRetrieve, err := client.RetrieveBook(context.Background(), &api.RetrieveBookRequest{BookId: bookId})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
	} else {
		log.Printf("Book retrieved: %v\n", resRetrieve.Book.String())
	}

	// update a book
	var bookIdUpdate int64 = 3
	bookUpdate := &api.Book{
		Id:          bookIdUpdate,
		Title:       "Go Programming-updated",
		Author:      "John Doe",
		Description: "Go is a programming language",
		Language:    "English",
		FinishTime:  timestamppb.Now(),
	}
	_, err = client.UpdateBook(context.Background(), &api.UpdateBookRequest{Book: bookUpdate})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book updated: %v\n", bookUpdate.String())
	}

	// delete book
	var bookIdDelete int64 = 2
	_, err = client.DeleteBook(context.Background(), &api.DeleteBookRequest{BookId: bookIdDelete})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book deleted bid: %v\n", bookIdDelete)
	}

	// list book
	resList, err := client.ListBooks(context.Background(), &api.ListBooksRequest{})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book list: %v\n", resList.Books)
	}

}
