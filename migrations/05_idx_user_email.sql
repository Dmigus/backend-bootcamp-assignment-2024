-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX idx_user_email ON "user" (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_user_email;
-- +goose StatementEnd