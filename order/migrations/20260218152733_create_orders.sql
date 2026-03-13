-- +goose Up
CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    order_uuid UUID NOT NULL,
    user_uuid UUID NOT NULL,
    part_uuids UUID [],
    total_price NUMERIC(12, 2),
    transaction_uuid UUID,
    payment_method TEXT,
    status TEXT
);
-- +goose Down
DROP TABLE order;