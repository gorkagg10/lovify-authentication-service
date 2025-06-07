CREATE TABLE tokens(
    id bigserial PRIMARY KEY,
    token text NOT NULL,
    type token_type NOT NULL,
    expiration_date varchar(50),
    username varchar(32) references users(username)
);