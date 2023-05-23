CREATE TYPE order_status AS ENUM ('Pending', 'Processing', 'Shipped','Delivered');

CREATE TABLE IF NOT EXISTS orders
(
    id         SERIAL PRIMARY KEY,
    user_id    INT          NOT NULL REFERENCES users (id),
    payment_id INT          NOT NULL REFERENCES payments (id),
    status     order_status NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT       NOT NULL REFERENCES orders (id),
    product_id INT       NOT NULL REFERENCES products (id),
    quantity   INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

