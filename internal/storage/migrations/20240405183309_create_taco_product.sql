-- +goose Up
-- +goose StatementBegin
CREATE TABLE taco_product
(
    taco_id    INT REFERENCES taco (id_taco),
    product_id INT REFERENCES product (id_product),
    CONSTRAINT taco_product_pk PRIMARY KEY (taco_id, product_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS taco_product;
-- +goose StatementEnd
