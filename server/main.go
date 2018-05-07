package main

import (
	"fmt"
	"net"

	pb "github.com/morix1500/cloudendpoint/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func echo(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(", Metadata:", md)
	msg := &pb.Msg{Message: in.Message.Message}
	return &pb.Response{Message: msg}, nil
}

type server struct{}

func (server) Echo1(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Print("Echo1 Received: ", in.Message.Message)
	return echo(ctx, in)
}

func (server) Echo2(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Print("Echo1 Received: ", in.Message.Message)
	return echo(ctx, in)
}

func (server) Echo3(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Print("Echo1 Received: ", in.Message.Message)
	return echo(ctx, in)
}

func (server) Echo4(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Print("Echo1 Received: ", in.Message.Message)
	return echo(ctx, in)
}

func main() {
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
