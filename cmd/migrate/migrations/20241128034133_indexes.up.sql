CREATE INDEX idx_order_customer_id ON "orders"(customer_id);
CREATE INDEX idx_order_order_date ON "orders"(order_date);
CREATE INDEX idx_order_status ON "orders"(status);
CREATE INDEX idx_order_item_order_product ON "order_items"(order_id, product_id);
CREATE UNIQUE INDEX idx_customer_email ON "customers"(email);
CREATE INDEX idx_product_name ON "products"(name);
CREATE INDEX idx_stock_product_id ON "stocks"(product_id);