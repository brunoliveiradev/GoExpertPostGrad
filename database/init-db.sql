USE goexpert;
CREATE TABLE IF NOT EXISTS products
(
    id    VARCHAR(255) PRIMARY KEY,
    name VARCHAR(80),
    price DECIMAL(10, 2)
);

INSERT INTO products (id, name, price)
VALUES ('1', 'CANETA BIC', 10.24);
INSERT INTO products (id, name, price)
VALUES ('2', 'CANETA VERMELHA', 20.48);
INSERT INTO products (id, name, price)
VALUES ('3', 'CANETA PRETA', 30.72);
INSERT INTO products (id, name, price)
VALUES ('4', 'CANETA CANETA', 40.96);
INSERT INTO products (id, name, price)
VALUES ('5', 'CANETA AZUL', 51.20);
INSERT INTO products (id, name, price)
VALUES ('6', 'AZUL CANETA', 61.20);