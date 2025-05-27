package login

type User struct {
	id             string
	username       string
	hashedPassword string
	sessionToken   string
	csrfToken      string
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

func (u *User) SessionToken() string {
	return u.sessionToken
}

func (u *User) CSRFToken() string {
	return u.csrfToken
}

func (u *User) setSessionToken(sessionToken string) {
	u.sessionToken = sessionToken
}

func (u *User) setCSRFToken(csrfToken string) {
	u.csrfToken = csrfToken
}
