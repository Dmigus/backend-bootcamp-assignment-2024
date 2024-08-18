-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id uuid PRIMARY KEY,
    email VARCHAR NOT NULL,
    salt char(16) NOT NULL,
    password_hash VARCHAR NOT NULL,
    role VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd