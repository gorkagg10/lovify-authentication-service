package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	service "github.com/gorkagg10/lovify-authentication-service/grpc/auth-service"
	"github.com/gorkagg10/lovify-authentication-service/internal/domain/login"
	"github.com/gorkagg10/lovify-authentication-service/internal/infra/base64"
	"github.com/gorkagg10/lovify-authentication-service/internal/infra/cache"
	"github.com/gorkagg10/lovify-authentication-service/internal/infra/server"
)

func main() {
	port := 8081

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error("failed to listen", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("listening", slog.String("port", fmt.Sprintf(":%d", port)))

	authServer := setupAuthServer()
	srv := setupGrpcServer(authServer)

	if err = srv.Serve(lis); err != nil {
		slog.Error("failed to serve", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func setupGrpcServer(authServer *server.AuthServer) *grpc.Server {
	grpcServer := grpc.NewServer()
	service.RegisterAuthServiceServer(grpcServer, authServer)
	return grpcServer
}

func setupAuthServer() *server.AuthServer {
	userRepository := cache.NewUserRepository(map[string]cache.User{})
	securityRepository := base64.NewSecurityRepository()

	authenticationService := login.NewAuthentication(userRepository, securityRepository)
	return server.NewAuthServer(authenticationService)
}
