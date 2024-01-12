USE goexpert;
CREATE TABLE IF NOT EXISTS products
(
    id    VARCHAR(255) PRIMARY KEY,
    name VARCHAR(80),
    price DECIMAL(10, 2)
);
