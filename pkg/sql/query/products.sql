-- name: GetProduct :one
SELECT id, name, price, description, category, created_at, updated_at
FROM products
WHERE id = $1
LIMIT 1;

-- name: ListProducts :many
SELECT id, name, price, description, category, created_at, updated_at
FROM products;

-- name: ListProductsByCategory :many
SELECT id, name, price, description, category, created_at, updated_at
FROM products
WHERE category = $1;