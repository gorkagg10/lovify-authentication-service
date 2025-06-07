package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	autherrors "github.com/gorkagg10/lovify-authentication-service/errors"
	"github.com/gorkagg10/lovify-authentication-service/internal/domain/login"
)

type UserRepository struct {
	pgClient *sql.DB
}

func NewUserRepository(pgClient *sql.DB) *UserRepository {
	return &UserRepository{pgClient}
}

func (u *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	var userId int64
	if err := u.pgClient.QueryRowContext(
		ctx,
		`SELECT id from users where username = $1;`, username).Scan(&userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, autherrors.ErrUserAlreadyExists
	}
	return true, nil
}

func (u *UserRepository) GetUser(ctx context.Context, username string) (*login.User, error) {
	var user User
	if err := u.pgClient.QueryRowContext(
		ctx,
		`SELECT username, password from users where username = $1;`, username).Scan(&user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("getting user from database: %w", err)
	}
	return login.NewUser(user.Username, user.Password), nil
}

func (u *UserRepository) CreateUser(ctx context.Context, user *login.User) error {
	var userID int64
	if err := u.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO users (username, password)
				VALUES($1, $2)
				RETURNING id;`, user.Username(), user.HashedPassword()).Scan(&userID); err != nil {
		return fmt.Errorf("inserting new user in database: %w", err)
	}
	return nil
}
