-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flat
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    house_id BIGSERIAL NOT NULL,
    price INT NOT NULL,
    rooms INT NOT NULL,
    status VARCHAR NOT NULL,
    CONSTRAINT fk_house
        FOREIGN KEY (house_id) REFERENCES house(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS flat;
-- +goose StatementEnd