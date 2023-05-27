-- name: CreateOrder :one
INSERT INTO orders (user_id, payment_id)
VALUES ($1, $2)
RETURNING id;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetOrder :one
SELECT id, user_id, payment_id, status, created_at, updated_at
FROM orders
WHERE id = $1
LIMIT 1;

-- name: GetOrderByPaymentID :one
SELECT id, user_id, payment_id, status, created_at, updated_at
FROM orders
WHERE payment_id = $1
LIMIT 1;

-- name: ListOrderItemsByOrderID :many
SELECT id, order_id, product_id, quantity, created_at, updated_at
FROM order_items
WHERE order_id = $1;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $2
WHERE id = $1;