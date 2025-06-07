package login

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	SessionToken = "session"
	CSRFToken    = "csrf"
)

type Authorization struct {
	userRepository     UserRepository
	securityRepository SecurityRepository
	tokenRepository    TokenRepository
}

func NewAuthorization(
	userRepository UserRepository,
	securityRepository SecurityRepository,
	tokenRepository TokenRepository) *Authorization {
	return &Authorization{
		userRepository:     userRepository,
		securityRepository: securityRepository,
		tokenRepository:    tokenRepository,
	}
}

func (a *Authorization) Register(ctx context.Context, username string, password string) error {
	exists, err := a.userRepository.UsernameExists(ctx, username)
	if err != nil {
		return fmt.Errorf("checking if user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	hashedPassword, err := a.securityRepository.HashPassword(password)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}
	user := NewUser(username, hashedPassword)
	if err = a.userRepository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("storing user: %w", err)
	}
	return nil
}

func (a *Authorization) Login(ctx context.Context, username string, password string) (*User, error) {
	user, err := a.userRepository.GetUser(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}

	validPassword, err := a.securityRepository.CheckPassword(user.HashedPassword(), password)
	if !validPassword {
		return nil, fmt.Errorf("invalid password")
	}
	if err != nil {
		return nil, fmt.Errorf("checking password: %w", err)
	}

	sessionToken, err := a.securityRepository.GenerateToken(SessionToken)
	if err != nil {
		return nil, fmt.Errorf("generating session token: %w", err)
	}
	err = a.tokenRepository.StoreToken(ctx, sessionToken, username)
	if err != nil {
		return nil, fmt.Errorf("storing session token: %w", err)
	}
	user.setSessionToken(sessionToken)

	csrfToken, err := a.securityRepository.GenerateToken(CSRFToken)
	if err != nil {
		return nil, fmt.Errorf("generating session token: %w", err)
	}
	err = a.tokenRepository.StoreToken(ctx, csrfToken, username)
	if err != nil {
		return nil, fmt.Errorf("storing csrf token: %w", err)
	}
	user.setCSRFToken(csrfToken)

	return user, nil
}

func isValidToken(userToken *Token) bool {
	if userToken.ExpirationDate().After(time.Now()) {
		return false
	}
	return true
}

func (a *Authorization) AuthorizeUser(ctx context.Context, username string, sessionToken string, csrfToken string) error {
	dbSessionToken, err := a.tokenRepository.GetToken(ctx, sessionToken, SessionToken, username)
	if err != nil {
		return fmt.Errorf("getting session token: %w", err)
	}
	if !isValidToken(dbSessionToken) {
		return errors.New("invalid session token")
	}
	dbCSRFToken, err := a.tokenRepository.GetToken(ctx, csrfToken, CSRFToken, username)
	if !isValidToken(dbCSRFToken) {
		return errors.New("invalid csrf token")
	}
	return nil
}
