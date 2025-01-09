package main

import (
	"auth_service/pkg/api/auth_v1"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	auth_v1.UnimplementedAuthServer
}

func (s *server) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &auth_v1.RegisterResponse{
		Token: base64.URLEncoding.EncodeToString(b),
	}, nil
}

func (s *server) GetMe(ctx context.Context, req *auth_v1.GetMeRequest) (*auth_v1.GetMeResponse, error) {
	if req.Id < 1 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &auth_v1.GetMeResponse{
		User: &auth_v1.User{
			Id:    req.Id,
			Name:  "John Doe",
			Email: "johndoe@go.dev",
			Role:  auth_v1.Role_ROLE_ADMIN,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	//
	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterAuthServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
