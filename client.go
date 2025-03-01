package main

import (
	"context"
	"go-server/server/chat"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// a connection without transport security
	if err != nil {
		log.Println("err=", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Content: "hello from client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Println("err=", err)
		return
	}
	log.Println("response=", response.Content)

}
