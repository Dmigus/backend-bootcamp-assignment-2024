-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id uuid PRIMARY KEY,
    email VARCHAR NOT NULL,
    salt bytea NOT NULL,
    password_hash bytea NOT NULL,
    role VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd