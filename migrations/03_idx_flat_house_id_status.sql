-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_flat_house_id_status ON flat (house_id, status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_flat_house_id_status;
-- +goose StatementEnd