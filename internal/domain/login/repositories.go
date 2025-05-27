package login

type UserRepository interface {
	UsernameExists(username string) (bool, error)
	GetUser(username string) (*User, error)
	StoreUser(user *User) error
}

type SecurityRepository interface {
	HashPassword(password string) (string, error)
	GenerateToken() (string, error)
}
