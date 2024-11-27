-- name: DecreaseStock :execrows
UPDATE stock
SET available_quantity = available_quantity - $2
WHERE product_id = $1
AND available_quantity >= $2;

-- name: GetStockByProductID :one
SELECT * FROM stock
WHERE product_id = $1;


