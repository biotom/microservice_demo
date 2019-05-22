package main

import "context"

func (s *server) Add(ctx context.Context, r *pb.AddReqest) (*pb.AddResponse, err) {
	response := &pb.AddResponse{}
	response.Result = r.A + r.B

	return response, nil
}
