package main

import (
	"context"
	"fmt"
	pb "github.com/qa-tools-family/grpc-demo/isempty/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement user.UserServer.
type server struct {
	pb.UnimplementedUserServer
}

// GetUser implements user.UserServer
func (s *server) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if r.Username != nil {
		fmt.Println("receive Username:", *r.Username)
	} else {
		fmt.Println("not received Username")
	}
	return &pb.GetUserResponse{Class: r.Class}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
