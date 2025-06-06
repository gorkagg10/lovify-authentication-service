package login

type UserRepository interface {
	UsernameExists(username string) (bool, error)
	GetUser(username string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
}

type SecurityRepository interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword string, password string) (bool, error)
	GenerateToken() (*Token, error)
}
