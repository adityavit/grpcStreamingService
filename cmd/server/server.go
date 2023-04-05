package main

import (
	"fmt"
	"log"
	"net"

	server "github.com/adityavit/grpcStreamingService/internal/server"
)

func main() {
	fmt.Println("Starting new grpc server")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := server.NewServer()
	log.Println("Listening on :9000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
