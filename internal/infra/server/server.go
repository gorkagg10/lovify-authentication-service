package server

import (
	"context"
	"github.com/gorkagg10/lovify-authentication-service/util"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/protobuf/types/known/emptypb"

	authServiceGrpc "github.com/gorkagg10/lovify-authentication-service/grpc/auth-service"
	"github.com/gorkagg10/lovify-authentication-service/internal/domain/login"
)

type AuthServer struct {
	authServiceGrpc.UnimplementedAuthServiceServer
	authenticationService *login.Authorization
}

func NewAuthServer(authenticationService *login.Authorization) *AuthServer {
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

func (s *AuthServer) Login(_ context.Context, req *authServiceGrpc.LoginRequest) (*authServiceGrpc.LoginResponse, error) {
	user, err := s.authenticationService.Login(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &authServiceGrpc.LoginResponse{
		SessionToken: &authServiceGrpc.Token{
			Token:          util.ValueToPointer(user.SessionToken().Token()),
			ExpirationDate: timestamppb.New(user.SessionToken().ExpirationDate()),
		},
		CsrfToken: &authServiceGrpc.Token{
			Token:          util.ValueToPointer(user.CSRFToken().Token()),
			ExpirationDate: timestamppb.New(user.CSRFToken().ExpirationDate()),
		},
	}, nil
}

func (s *AuthServer) Authorize(_ context.Context, req *authServiceGrpc.AuthorizationRequest) (*emptypb.Empty, error) {
	err := s.authenticationService.Authorize(req.GetUsername(), req.GetSessionToken(), req.GetCsrfToken())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
