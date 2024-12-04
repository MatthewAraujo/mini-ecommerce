-- name: FinAllCustomers :many
select * from customers;

-- name: InsertCustomers :one
INSERT INTO customers (name, email, password, role)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: FindCustomerByID :one
SELECT * FROM customers
WHERE id = $1;

-- name: FindCustomerByEmail :one
SELECT * 
FROM customers
WHERE email = $1
LIMIT 1;