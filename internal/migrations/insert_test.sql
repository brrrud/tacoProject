INSERT INTO pfc_info(proteins, fats, carbohydrates)
VALUES (20.5, 10.4, 13.2);

INSERT INTO product (name_product, pfc_info_id_fk, weight_product)
VALUES ('Chicken', 1, 100);

INSERT INTO pfc_info(proteins, fats, carbohydrates)
VALUES (21.5, 15.4, 10.2);

INSERT INTO product (name_product, pfc_info_id_fk, weight_product)
VALUES ('Meat', 2, 149);

INSERT INTO pfc_info(proteins, fats, carbohydrates)
VALUES (1.0, 2.4, 1.2);

INSERT INTO product (name_product, pfc_info_id_fk, weight_product)
VALUES ('Tomato', 3, 149);


INSERT INTO taco (name_taco) VALUES ('chicken_taco');
INSERT INTO taco (name_taco) VALUES ('meat_taco');

INSERT INTO taco_product VALUES (1,1);
INSERT INTO taco_product VALUES (2,2);
INSERT INTO taco_product VALUES (1, 3);