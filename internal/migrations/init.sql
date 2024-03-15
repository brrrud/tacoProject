CREATE TABLE IF NOT EXISTS taco_factory
(
    id_taco_factory      SERIAL PRIMARY KEY,
    address_taco_factory VARCHAR NOT NULL

);

CREATE TABLE IF NOT EXISTS taco
(
    id_taco   SERIAL PRIMARY KEY,
    name_taco VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS pfc_info
(
    id_pfc_info   SERIAL PRIMARY KEY,
    proteins      DECIMAL NOT NULL,
    fats          DECIMAL NOT NULL,
    carbohydrates DECIMAL NOT NULL
);

CREATE TABLE IF NOT EXISTS product
(
    id_product     SERIAL PRIMARY KEY,
    name_product   VARCHAR NOT NULL UNIQUE,
    pfc_info_id_fk BIGINT  NOT NULL,
    weight         DECIMAL NOT NULL,
    CONSTRAINT pfc_info_fk FOREIGN KEY (pfc_info_id_fk) REFERENCES pfc_info (id_pfc_info)
);

CREATE TABLE taco_product
(
    taco_id    INT REFERENCES taco (id_taco),
    product_id INT REFERENCES product (id_product),
    CONSTRAINT taco_product_pk PRIMARY KEY (taco_id, product_id)
);

CREATE INDEX idx_product_name ON product(name_product);
