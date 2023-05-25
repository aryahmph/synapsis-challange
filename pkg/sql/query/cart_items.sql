-- name: CreateCartItem :one
INSERT INTO cart_items (product_id, user_id, quantity)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetCartItem :one
SELECT id, product_id, user_id, quantity, created_at, updated_at
FROM cart_items
WHERE id = $1
LIMIT 1;

-- name: VerifyAvailableCartItem :one
SELECT id
FROM cart_items
WHERE product_id = $1
  AND user_id = $2
LIMIT 1;

-- name: UpdateCartItemQuantity :exec
UPDATE cart_items
SET quantity = quantity + $2
WHERE id = $1;

-- name: ListCartItemsByCartID :many
SELECT id, product_id, user_id, quantity, created_at, updated_at
FROM cart_items
WHERE user_id = $1;

-- name: DeleteCartItem :exec
DELETE
FROM cart_items
WHERE id = $1;