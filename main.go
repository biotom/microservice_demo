package main

import (
	"grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	type service struct{}

	listen, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterAddServiceServer(server, &service{})

	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
