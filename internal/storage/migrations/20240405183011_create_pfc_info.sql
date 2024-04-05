-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pfc_info
(
    id_pfc_info   SERIAL PRIMARY KEY,
    proteins      DECIMAL NOT NULL,
    fats          DECIMAL NOT NULL,
    carbohydrates DECIMAL NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pfc_info;
-- +goose StatementEnd
