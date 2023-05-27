-- name: CreatePayment :exec
INSERT INTO payments (id, user_id, va_number, amount)
VALUES ($1, $2, $3, $4);

-- name: GetPayment :one
SELECT id, user_id, va_number, amount, status, created_at, updated_at
FROM payments
WHERE id = $1
LIMIT 1;

-- name: UpdatePaymentStatus :exec
UPDATE payments
SET status = $2
WHERE id = $1;