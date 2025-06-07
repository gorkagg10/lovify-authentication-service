CREATE TABLE users(
    id bigserial PRIMARY KEY,
    username varchar(32) UNIQUE NOT NULL,
    password varchar(64) NOT NULL
);