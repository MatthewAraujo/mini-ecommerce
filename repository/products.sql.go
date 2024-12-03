// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package repository

import (
	"context"
	"database/sql"
)

const findProductByID = `-- name: FindProductByID :one
SELECT id, name, description, price FROM products
WHERE id = $1
`

func (q *Queries) FindProductByID(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, findProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const findProductByName = `-- name: FindProductByName :one
SELECT id, name, description, price FROM products
WHERE name = $1
`

func (q *Queries) FindProductByName(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRowContext(ctx, findProductByName, name)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const insertProduct = `-- name: InsertProduct :one
INSERT INTO products (name, description, price ) 
VALUES ($1, $2, $3)
RETURNING id, name, description, price
`

type InsertProductParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
}

func (q *Queries) InsertProduct(ctx context.Context, arg InsertProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, insertProduct, arg.Name, arg.Description, arg.Price)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}
