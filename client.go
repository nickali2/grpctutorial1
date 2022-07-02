package main

import (
	"context"
	"google.golang.org/grpc"
	"grpctutorial1/chat"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect , %w", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{Body: "hello from client"}
	mes, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("couldnt connect to say hello, %w", err)
	}
	log.Printf("response from server: %s", mes.Body)
}
