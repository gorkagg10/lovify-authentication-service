package cache

import "time"

type User struct {
	username       string
	hashedPassword string
}

func NewUser(
	username string,
	hashedPassword string) *User {
	return &User{
		username:       username,
		hashedPassword: hashedPassword,
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
	username       string
}

func NewToken(token string, expirationDate time.Time, username string) *Token {
	return &Token{
		token:          token,
		expirationDate: expirationDate,
		username:       username,
	}
}
