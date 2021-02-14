package main

import (
	"log"
	"net"

	"github.com/Sumeet-Ranjan-Parida/Microservices-with-Go/gRPC/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	gs := grpc.NewServer()

	chat.RegisterChatServiceServer(gs, &s)

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000: %v", err)
	}
}
