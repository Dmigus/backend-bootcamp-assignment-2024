-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS house
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    address VARCHAR NOT NULL,
    year INT NOT NULL,
    developer VARCHAR,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS house;
-- +goose StatementEnd