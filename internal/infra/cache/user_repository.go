package cache

import (
	"fmt"
	"github.com/gorkagg10/lovify-authentication-service/internal/domain/login"
)

type UserRepository struct {
	database map[string]User
}

func NewUserRepository(database map[string]User) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (u *UserRepository) UsernameExists(username string) (bool, error) {
	_, ok := u.database[username]
	return ok, nil
}

func (u *UserRepository) GetUser(username string) (*login.User, error) {
	user, ok := u.database[username]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return login.NewUser(user.Username(), user.HashedPassword()), nil
}

func (u *UserRepository) CreateUser(user *login.User) error {
	u.database[user.Username()] = *NewUser(user.Username(), user.HashedPassword(), nil, nil)
	return nil
}

func (u *UserRepository) UpdateUser(user *login.User) error {
	sessionToken := NewToken(user.SessionToken().Token(), user.SessionToken().ExpirationDate())
	csrfToken := NewToken(user.CSRFToken().Token(), user.CSRFToken().ExpirationDate())
	u.database[user.Username()] = *NewUser(user.Username(), user.HashedPassword(), sessionToken, csrfToken)
	return nil
}
