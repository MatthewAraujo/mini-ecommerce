// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"database/sql"
)

type Customer struct {
	ID      int32          `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
}

type Order struct {
	ID         int32         `json:"id"`
	CustomerID int32         `json:"customer_id"`
	OrderDate  sql.NullTime  `json:"order_date"`
	Status     string        `json:"status"`
}

type OrderItem struct {
	ID        int32 `json:"id"`
	OrderID   int32 `json:"order_id"`
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type Product struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       string         `json:"price"`
}

type Stock struct {
	ID                int32 `json:"id"`
	ProductID         int32 `json:"product_id"`
	AvailableQuantity int32 `json:"available_quantity"`
}
