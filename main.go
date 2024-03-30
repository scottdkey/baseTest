package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/scottdkey/baseTest"
)

type server struct{}

func (s *server) YourRPCMethod(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	return &pb.ResponseMessage{YourResponseField: "Hello, " + req.YourFieldName}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})

	log.Printf("Server listening on localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
