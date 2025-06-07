package postgres

type User struct {
	ID       int64
	Username string
	Password string
}

type Token struct {
	ID             int64
	Token          string
	TokenType      string
	ExpirationDate string
	Username       string
}
