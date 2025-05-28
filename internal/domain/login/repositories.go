package login

type UserRepository interface {
	UsernameExists(username string) (bool, error)
	GetUser(username) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
}

type SecurityRepository interface {
	HashPassword(password string) (string, error)
	GenerateToken() (*Token, error)
}
