CREATE TYPE product_category AS ENUM ('Makanan/Minuman', 'Pakaian', 'Teknologi');

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255)     NOT NULL,
    price       DOUBLE PRECISION NOT NULL,
    description TEXT             NOT NULL,
    category    product_category NOT NULL,
    created_at  TIMESTAMP        NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP        NOT NULL DEFAULT NOW()
);