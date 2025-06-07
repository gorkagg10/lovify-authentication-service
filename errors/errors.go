package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrUserAlreadyExistsMsg              = "USER_ALREADY_EXISTS"
	ErrDatabaseQueryFailedMsg            = "DATABASE_QUERY_FAILED"
	ErrHashedPasswordGenerationFailedMsg = "HASHED_PASSWORD_GENERATION_FAILED"
)

var (
	StatusUserAlreadyExists              = status.New(codes.InvalidArgument, ErrUserAlreadyExistsMsg)
	StatusDatabaseQueryFailed            = status.New(codes.Internal, ErrDatabaseQueryFailedMsg)
	StatusHashedPasswordGenerationFailed = status.New(codes.Internal, ErrHashedPasswordGenerationFailedMsg)

	ErrUserAlreadyExists              = StatusUserAlreadyExists.Err()
	ErrDatabaseQueryFailed            = StatusDatabaseQueryFailed.Err()
	ErrHashedPasswordGenerationFailed = StatusHashedPasswordGenerationFailed.Err()
)
