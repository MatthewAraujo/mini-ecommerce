// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package repository

import (
	"context"
)

const getProductByID = `-- name: GetProductByID :one
SELECT id, name, description, price FROM products
WHERE id = $1
`

func (q *Queries) GetProductByID(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}