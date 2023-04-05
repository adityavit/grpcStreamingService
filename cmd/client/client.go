package main

import (
	"context"
	pb "github.com/adityavit/grpcStreamingService/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	client := pb.NewStreamingServiceClient(conn)
	// Send a request to the server
	req := &pb.DataRequest{Message: "Hello from the client!"}
	stream, err := client.StreamData(ctx, req)
	if err != nil {
		log.Fatalf("failed to stream data: %v", err)
	}

	// Receive and print the stream of responses
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Printf("Received response: %v", resp.GetMessage())
	}
}
