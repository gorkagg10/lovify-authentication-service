package login

import "context"

type UserRepository interface {
	UsernameExists(context context.Context, username string) (bool, error)
	GetUser(context context.Context, username string) (*User, error)
	CreateUser(context context.Context, user *User) error
}

type SecurityRepository interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword string, password string) (bool, error)
	GenerateToken(tokenType string) (*Token, error)
}

type TokenRepository interface {
	StoreToken(ctx context.Context, token *Token, username string) error
	GetToken(ctx context.Context, token string, tokenType string, username string) (*Token, error)
}
