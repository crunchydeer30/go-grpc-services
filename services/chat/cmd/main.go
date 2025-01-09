package main

import (
	"chat/pkg/api/chat_v1"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	chat_v1.UnimplementedChatServer
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	//
	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
