
ALTER TABLE orders 
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE order_items 
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;


ALTER TABLE order_items 
    ADD COLUMN cell_price NUMERIC;

