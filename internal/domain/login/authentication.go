package login

import (
	"fmt"
)

type Authentication struct {
	userRepository  UserRepository
	TokenRepository SecurityRepository
}

func NewAuthentication(userRepository UserRepository, tokenRepository SecurityRepository) *Authentication {
	return &Authentication{
		userRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (a *Authentication) Register(username string, password string) error {
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
	if err = a.userRepository.StoreUser(user); err != nil {
		return fmt.Errorf("storing user: %w", err)
	}
	return nil
}

func (a *Authentication) Login(username string) error {
	user, err := a.userRepository.GetUser(username)
	if err != nil {
		return fmt.Errorf("getting user: %w", err)
	}

	sessionToken, err := a.TokenRepository.GenerateToken()
	if err != nil {
		return fmt.Errorf("generating session token: %w", err)
	}
	user.setSessionToken(sessionToken)

	csrfToken, err := a.TokenRepository.GenerateToken()
	if err != nil {
		return fmt.Errorf("generating session token: %w", err)
	}
	user.setCSRFToken(csrfToken)

	if err = a.userRepository.StoreUser(user); err != nil {
		return fmt.Errorf("storing user: %w", err)
	}
	return nil
}
