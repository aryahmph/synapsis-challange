-- name: CreateOrder :one
INSERT INTO orders (user_id, payment_id)
VALUES ($1, $2)
RETURNING id;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3)
RETURNING id;