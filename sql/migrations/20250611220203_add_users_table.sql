-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id       BIGSERIAL PRIMARY KEY,
    name     text NOT NULL,
    email    text NOT NULL UNIQUE,
    username text NOT NULL UNIQUE,
    password text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
