CREATE TABLE IF NOT EXISTS cart_items
(
    id         SERIAL PRIMARY KEY,
    product_id INT       NOT NULL REFERENCES products (id),
    user_id    INT       NOT NULL REFERENCES users (id),
    quantity   INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_cart_product ON cart_items (user_id, product_id);
