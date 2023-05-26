CREATE TYPE payment_status AS ENUM ('Pending', 'Completed', 'Expired');

CREATE TABLE IF NOT EXISTS payments
(
    id         VARCHAR(36) PRIMARY KEY,
    user_id    INT              NOT NULL REFERENCES users (id),
    va_number  VARCHAR(50)      NOT NULL,
    amount     DOUBLE PRECISION NOT NULL,
    status     payment_status   NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP        NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP        NOT NULL DEFAULT NOW()
);