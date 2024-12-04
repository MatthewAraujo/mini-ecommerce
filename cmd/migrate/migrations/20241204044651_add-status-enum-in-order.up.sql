ALTER TABLE orders
ADD CONSTRAINT check_order_status
CHECK (status IN ('pending', 'send', 'arrived'));
