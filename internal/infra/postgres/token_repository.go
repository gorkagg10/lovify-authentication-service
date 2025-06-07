package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gorkagg10/lovify-authentication-service/internal/domain/login"
)

type TokenRepository struct {
	pgClient *sql.DB
}

func NewTokenRepository(pgClient *sql.DB) *TokenRepository {
	return &TokenRepository{
		pgClient: pgClient,
	}
}

func (t *TokenRepository) StoreToken(ctx context.Context, token *login.Token, username string) error {
	var tokenID int64
	if err := t.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO tokens (token, type, expiration_date, username) 
				VALUES ($1, $2, $3, $4)
				RETURNING id;`, token.Token(), token.TokenType(), token.ExpirationDate().String(), username).Scan(&tokenID); err != nil {
		return fmt.Errorf("inserting new user in database: %w", err)
	}
	return nil
}

func (t *TokenRepository) GetToken(ctx context.Context, token string, tokenType string, username string) (*login.Token, error) {
	var dbToken Token
	if err := t.pgClient.QueryRowContext(
		ctx,
		`SELECT token, token_type, expiration_date, username from tokens WHERE token = $1, token_type = $2, username = $3;`,
		token, tokenType, username,
	).Scan(&dbToken); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("token not found")
		}
		return nil, fmt.Errorf("fetching token: %w", err)
	}
	expirationDate, err := time.Parse(time.RFC3339, dbToken.ExpirationDate)
	if err != nil {
		return nil, fmt.Errorf("parsing token expiration date: %w", err)
	}
	return login.NewToken(dbToken.Token, dbToken.TokenType, expirationDate), nil
}
