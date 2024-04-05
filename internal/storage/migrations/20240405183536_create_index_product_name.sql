-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_product_name ON taco(name_taco);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_product_name;
-- +goose StatementEnd
