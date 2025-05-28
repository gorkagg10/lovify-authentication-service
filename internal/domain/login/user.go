package login

type User struct {
	id             string
	username       string
	hashedPassword string
	sessionToken   *Token
	csrfToken      *Token
}

func NewUser(username string, hashedPassword string) *User {
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

func (u *User) SessionToken() *Token {
	return u.sessionToken
}

func (u *User) CSRFToken() *Token {
	return u.csrfToken
}

func (u *User) setSessionToken(sessionToken *Token) {
	u.sessionToken = sessionToken
}

func (u *User) setCSRFToken(csrfToken *Token) {
	u.csrfToken = csrfToken
}
