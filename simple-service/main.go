package main

import (
	"context"
	"dmha/hello-grpc/person"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":30051"
)

// server is used to implement UnimplementedPersonServiceServer
type server struct {
	person.UnimplementedPersonServiceServer
}

func (s *server) Create(ctx context.Context, request *person.CreatePersonRequest) (*person.CreatePersonResponse, error) {
	log.Printf("Received: %v", request.Person.Name)

	return &person.CreatePersonResponse{
		Id: request.Person.Id,
	}, nil
}

func main() {
	fmt.Printf("Running gRPC simple service tutorial in port %s\n", port)

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	person.RegisterPersonServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
