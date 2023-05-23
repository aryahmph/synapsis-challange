CREATE TABLE IF NOT EXISTS carts
(
    id         SERIAL PRIMARY KEY,
    user_id    INT       NOT NULL REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    INT       NOT NULL REFERENCES carts (id),
    product_id INT       NOT NULL REFERENCES products (id),
    quantity   INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);