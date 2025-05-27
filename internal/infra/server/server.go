package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	authServiceGrpc "github.com/gorkagg10/lovify-authentication-service.git/grpc/auth-service"
	"github.com/gorkagg10/lovify-authentication-service.git/internal/domain/login"
)

type AuthServer struct {
	authServiceGrpc.UnimplementedAuthServiceServer
	authenticationService *login.Authentication
}

func NewAuthServer(authenticationService *login.Authentication) *AuthServer {
	return &AuthServer{
		authenticationService: authenticationService,
	}
}

func (s *AuthServer) RegisterUser(_ context.Context, req *authServiceGrpc.RegisterRequest) (*emptypb.Empty, error) {
	err := s.authenticationService.Register(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
