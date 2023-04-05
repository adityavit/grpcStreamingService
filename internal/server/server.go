package streamingServer

import (
	"fmt"
	pb "github.com/adityavit/grpcStreamingService/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

type server struct {
	pb.UnimplementedStreamingServiceServer
}

func (s *server) StreamData(req *pb.DataRequest, stream pb.StreamingService_StreamDataServer) error {
	message := req.GetMessage()
	log.Printf("Received message: %v", message)

	for i := 0; i < 10; i++ {
		response := &pb.DataResponse{Message: fmt.Sprintf("%s %d", message, i)}
		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func NewServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, &server{})
	reflection.Register(s)
	return s
}
