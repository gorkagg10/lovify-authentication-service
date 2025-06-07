package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrUserAlreadyExistsMsg = "USER_ALREADY_EXISTS"
)

var (
	StatusUserAlreadyExists = status.New(codes.InvalidArgument, ErrUserAlreadyExistsMsg)

	ErrUserAlreadyExists = StatusUserAlreadyExists.Err()
)
