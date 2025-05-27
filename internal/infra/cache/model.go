package cache

type User struct {
	username       string
	hashedPassword string
	sessionToken   string
	csrfToken      string
}

func NewUser(
	username string,
	hashedPassword string,
	sessionToken string,
	csrfToken string) User {
	return User{
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
