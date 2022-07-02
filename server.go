package main

import (
	"google.golang.org/grpc"
	"grpctutorial1/chat"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("unable to lsten on poret 9000, %w", err)
	}
	s := chat.Server{}
	grpcserver := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcserver, &s)
	if err := grpcserver.Serve(lis); err != nil {
		log.Fatalf("unable to start grpc server on port 9000, %v", err)
	}

}
