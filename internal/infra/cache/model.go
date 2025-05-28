package cache

import "time"

type User struct {
	username       string
	hashedPassword string
	sessionToken   *Token
	csrfToken      *Token
}

func NewUser(
	username string,
	hashedPassword string,
	sessionToken *Token,
	csrfToken *Token) *User {
	return &User{
		username:       username,
		hashedPassword: hashedPassword,
		sessionToken:   sessionToken,
		csrfToken:      csrfToken,
	}
}

func (u *User) Username() string {
	return u.username
}

func (u *User) HashedPassword() string {
	return u.hashedPassword
}

type Token struct {
	token          string
	expirationDate time.Time
}

func NewToken(token string, expirationDate time.Time) *Token {
	return &Token{
		token:          token,
		expirationDate: expirationDate,
	}
}
