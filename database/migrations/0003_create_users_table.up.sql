CREATE TABLE users(
    id bigserial PRIMARY KEY,
    username varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    session_token bigserial references tokens(id),
    csrf_token bigserial references tokens(id)
);