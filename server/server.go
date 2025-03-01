package main

import (
	"log"
	"net"

	"go-server/server/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	// listen on port 8080 with tcp protocol
	if err != nil {
		log.Println("an error occurred:", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("an error occurred when starting the grpc servre:", err)
	}
}
