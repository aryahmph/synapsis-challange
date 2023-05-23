CREATE TYPE payment_status AS ENUM ('Pending', 'Completed');

CREATE TABLE IF NOT EXISTS payments
(
    id         SERIAL PRIMARY KEY,
    user_id    INT            NOT NULL REFERENCES users (id),
    amount     INT            NOT NULL,
    status     payment_status NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP      NOT NULL DEFAULT NOW()
);