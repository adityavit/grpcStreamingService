// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: api/v1/streaming_service.proto

package api_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	StreamData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (StreamingService_StreamDataClient, error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) StreamData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (StreamingService_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], "/StreamingService/StreamData", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceStreamDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingService_StreamDataClient interface {
	Recv() (*DataResponse, error)
	grpc.ClientStream
}

type streamingServiceStreamDataClient struct {
	grpc.ClientStream
}

func (x *streamingServiceStreamDataClient) Recv() (*DataResponse, error) {
	m := new(DataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility
type StreamingServiceServer interface {
	StreamData(*DataRequest, StreamingService_StreamDataServer) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServiceServer struct {
}

func (UnimplementedStreamingServiceServer) StreamData(*DataRequest, StreamingService_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).StreamData(m, &streamingServiceStreamDataServer{stream})
}

type StreamingService_StreamDataServer interface {
	Send(*DataResponse) error
	grpc.ServerStream
}

type streamingServiceStreamDataServer struct {
	grpc.ServerStream
}

func (x *streamingServiceStreamDataServer) Send(m *DataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamData",
			Handler:       _StreamingService_StreamData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/streaming_service.proto",
}
