package main

import (
	"context"
	"log"
	"micro_example/pb"

	"google.golang.org/grpc"
)

func main() {
	var connection *gprc.ClientConn

	connection, err := grpc.Dial(":4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewAddServiceClient(connection)
	addResponse, err := client.Add(context.Background(), &pb.AddRequest{A: 1, B: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sum of provided numbers: %v", addResponse.Result)

	anagramResponse, err := client.Anagram(context.Background(), &pb.IsAnagramRequest{Str1: "cat", Str2: "cat"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("2 provided strings are anagrams: %v", anagramResponse)
}
