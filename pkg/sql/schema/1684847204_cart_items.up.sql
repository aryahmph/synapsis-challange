CREATE TABLE IF NOT EXISTS cart_items
(
    id         INT PRIMARY KEY REFERENCES products (id),
    user_id    INT       NOT NULL REFERENCES users (id),
    quantity   INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_cart_product ON cart_items (id, user_id);