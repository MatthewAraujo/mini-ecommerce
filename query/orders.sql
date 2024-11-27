-- name: GetOrderByID :one
SELECT * FROM orders
WHERE id = $1;

-- name: GetOrderItemsByOrderID :many
SELECT * FROM order_items
WHERE order_id = $1;

-- name: CompleteOrder :execrows
UPDATE orders
SET status = 'completed'
WHERE id = $1;

-- name: CancelOrder :execrows
UPDATE orders
SET status = 'canceled'
WHERE id = $1;

-- name: AddOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, order_date, status)
VALUES ($1, NOW(), 'pending')
RETURNING *;

