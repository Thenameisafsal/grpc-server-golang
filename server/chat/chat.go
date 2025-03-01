package chat

import (
	"context"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Println("message received:", message.Content)
	return &Message{
		Content: "hello world",
	}, nil

}

// func (s *Server) mustEmbedUnimplementedChatServiceServer()
