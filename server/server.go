package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"reverss/proto"
)

type server struct {
	proto.ReverseServer
}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterReverseServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

func (s *server) Do(c context.Context, request *proto.Request) (response *proto.Response, err error) {
	n := 0
	// Сreate an array of runes to safely reverse a string.
	rune := make([]rune, len(request.Message))

	for _, r := range request.Message {
		rune[n] = r
		n++
	}

	// Reverse using runes.
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output := string(rune)
	response = &proto.Response{
		Message: output,
	}

	return response, nil
}
