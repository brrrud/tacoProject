-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product
(
    id_product     SERIAL PRIMARY KEY,
    name_product   VARCHAR NOT NULL UNIQUE,
    pfc_info_id_fk BIGINT  NOT NULL,
    weight_product         DECIMAL NOT NULL,
    CONSTRAINT pfc_info_fk FOREIGN KEY (pfc_info_id_fk) REFERENCES pfc_info (id_pfc_info)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product;
-- +goose StatementEnd
