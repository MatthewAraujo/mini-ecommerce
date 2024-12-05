-- name: FindProductByID :one
SELECT * FROM products
WHERE id = $1;

-- name: FindProductByName :one
SELECT * FROM products
WHERE name = $1
LIMIT 1;

-- name: InsertProduct :one
INSERT INTO products (name, description, price ) 
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetAllProductsPagination :many 
SELECT * FROM products LIMIT $1 OFFSET $2;

-- name: CountProducts :one
SELECT COUNT(*) AS TOTAL
FROM products;