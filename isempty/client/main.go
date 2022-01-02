package main

import (
	"context"
	pb "github.com/qa-tools-family/grpc-demo/isempty/user"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	var username string
	if len(os.Args) > 1 {
		username = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var r *pb.GetUserResponse
	if username != "" && username != "empty" {
		r, err = c.GetUser(ctx, &pb.GetUserRequest{Class: "name", Username: &username})
	} else if username == "empty" {
		username = ""
		r, err = c.GetUser(ctx, &pb.GetUserRequest{Class: "name", Username: &username})
	} else {
		r, err = c.GetUser(ctx, &pb.GetUserRequest{Class: "name"})
	}

	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Response: %s", r.Class)
}
