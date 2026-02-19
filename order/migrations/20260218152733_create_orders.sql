-- +goose Up
CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    order_uuid UUID NOT NULL,
    user_uuid  UUID NOT NULL,
    part_uuids UUID[], 
    total_price numeric(12,2), 
    transaction_uuid UUID,     
    payment_method   text, 
    status          text  
);

-- +goose Down
drop table order;