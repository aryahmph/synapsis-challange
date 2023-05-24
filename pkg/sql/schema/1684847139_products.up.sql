CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255)     NOT NULL,
    price       DOUBLE PRECISION NOT NULL,
    description TEXT             NOT NULL,
    category    VARCHAR(255)     NOT NULL,
    created_at  TIMESTAMP        NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP        NOT NULL DEFAULT NOW()
);

INSERT INTO products(name, price, description, category)
VALUES ('Macbook Air M1', 20000000, 'Lorem ipsum sit dolor amet', 'Teknologi'),
       ('Macbook Pro M1', 40000000, 'Lorem ipsum sit dolor amet', 'Teknologi'),
       ('Iphone 13', 9000000, 'Lorem ipsum sit dolor amet', 'Teknologi'),
       ('Kaos Erigo', 100000, 'Lorem ipsum sit dolor amet', 'Pakaian'),
       ('Celana Erigo', 150000, 'Lorem ipsum sit dolor amet', 'Pakaian'),
       ('Sweater Erigo', 250000, 'Lorem ipsum sit dolor amet', 'Pakaian'),
       ('Chicken Karage', 30000, 'Lorem ipsum sit dolor amet', 'Makanan'),
       ('Ayam Geprek', 12000, 'Lorem ipsum sit dolor amet', 'Makanan')
