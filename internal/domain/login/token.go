package login

import "time"

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

func (t *Token) Token() string {
	return t.token
}

func (t *Token) ExpirationDate() time.Time {
	return t.expirationDate
}
