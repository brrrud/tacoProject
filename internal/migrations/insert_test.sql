INSERT INTO pfc_info (proteins, fats, carbohydrates)
VALUES (1.3, 25.2, 16.1);


INSERT INTO pfc_info (proteins, fats, carbohydrates)
VALUES (1.7, 2.4, 6.0);


INSERT INTO product (name_product, pfc_info_id_fk)
VALUES ('meat', 1);


INSERT INTO product (name_product, pfc_info_id_fk)
VALUES ('tomato', 2);

INSERT INTO taco (name_taco) VALUES ('meat taco');
INSERT INTO taco (name_taco) VALUES ('super chease taco');

INSERT INTO taco_product VALUES (1, 1), (2, 1), (1, 2);