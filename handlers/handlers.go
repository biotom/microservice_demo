package handlers

import (
	"context"
	pb "micro_example/pb"
)

//Server represents a gRPC server
type Server struct{}

//Add adds two numbers
func (s *Server) Add(ctx context.Context, r *pb.AddRequest) (*pb.AddResponse, error) {
	response := &pb.AddResponse{}
	response.Result = r.A + r.B

	return response, nil
}

//Anagram finds if two strings are anagrams and returns a boolean describing the result
func (s *Server) Anagram(ctx context.Context, r *pb.IsAnagramRequest) (*pb.IsAnagramResponse, error) {
	response := &pb.IsAnagramResponse{}
	response.AnagramBool = false

	return response, nil
}
