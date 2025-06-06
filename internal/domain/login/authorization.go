package login

import (
	"fmt"
	"time"
)

type Authorization struct {
	userRepository  UserRepository
	TokenRepository SecurityRepository
}

func NewAuthorization(userRepository UserRepository, tokenRepository SecurityRepository) *Authorization {
	return &Authorization{
		userRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (a *Authorization) Register(username string, password string) error {
	exists, err := a.userRepository.UsernameExists(username)
	if err != nil {
		return fmt.Errorf("checking if user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	hashedPassword, err := a.TokenRepository.HashPassword(password)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}
	user := NewUser(username, hashedPassword)
	if err = a.userRepository.CreateUser(user); err != nil {
		return fmt.Errorf("storing user: %w", err)
	}
	return nil
}

func (a *Authorization) Login(username string, password string) (*User, error) {
	user, err := a.userRepository.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}

	validPassword, err := a.TokenRepository.CheckPassword(user.HashedPassword(), password)
	if !validPassword {
		return nil, fmt.Errorf("invalid password")
	}
	if err != nil {
		return nil, fmt.Errorf("checking password: %w", err)
	}

	sessionToken, err := a.TokenRepository.GenerateToken()
	if err != nil {
		return nil, fmt.Errorf("generating session token: %w", err)
	}
	user.setSessionToken(sessionToken)

	csrfToken, err := a.TokenRepository.GenerateToken()
	if err != nil {
		return nil, fmt.Errorf("generating session token: %w", err)
	}
	user.setCSRFToken(csrfToken)

	if err = a.userRepository.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("storing user: %w", err)
	}
	return user, nil
}

func isValidToken(token string, userToken *Token) bool {
	if token != userToken.Token() || userToken.ExpirationDate().After(time.Now()) {
		return false
	}
	return true
}

func (a *Authorization) Authorize(username string, sessionToken string, csrfToken string) error {
	user, err := a.userRepository.GetUser(username)
	if err != nil {
		return fmt.Errorf("getting user: %w", err)
	}
	if !isValidToken(sessionToken, user.SessionToken()) {
		return fmt.Errorf("invalid session token")
	}
	if !isValidToken(csrfToken, user.CSRFToken()) {
		return fmt.Errorf("invalid csrf token")
	}
	return nil
}
