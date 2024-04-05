-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS taco
(
    id_taco   SERIAL PRIMARY KEY,
    name_taco VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS taco;
-- +goose StatementEnd
