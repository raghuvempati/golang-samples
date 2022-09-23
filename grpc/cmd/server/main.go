package main

import (
	"context"
	"log"
	"net"
	"time"

	internal "go-samples/grpc/internal"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoUri := "mongodb://sndsea.southeastasia.cloudapp.azure.com:27017"

	log.Println("Starting listening on port 8080")
	port := ":8080"

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Mongo Repository
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("testing")

	// create a new mongo repository
	var repository internal.BookRepository = internal.NewMongoBookRepository(db)
	srv := internal.NewRPCServer(repository)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
