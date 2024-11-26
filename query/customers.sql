-- name: FinAllCustomers :many
select * from customers;

-- name: InsertCustomers :one
INSERT INTO customers (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: FindCustomerByID :one
SELECT * FROM customers
WHERE id = $1;