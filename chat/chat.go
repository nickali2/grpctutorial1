package chat

import (
	"context"
	"log"
)

type Server struct {
	//pdfpb.UnimplementedGreetServiceServer
	//pdfpb. mustEmbedUnimplementedGreeterServer()

}

func (*Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("received a message from client: %s", message.Body)
	return &Message{Body: "hello from server"}, nil
}
