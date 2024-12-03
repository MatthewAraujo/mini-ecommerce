package types

import "database/sql"

type CostumersService interface {
	CreateCustomer(c *CreateCustomerPayload) (int, error)
	Login(c *LoginCustomerPayload) (string, int, error)
}

type CreateCustomerPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type LoginCustomerPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type CreateProductPayload struct {
	Name        string         `json:"name" validate:"required"`
	Description sql.NullString `json:"description" validate:"required,max=255"`
	Price       string         `json:"price" validate:"required"`
	Quantity    int            `json:"quantity" validate:"required"`
}

type ProductService interface {
	CreateProduct(p *CreateProductPayload) (int, error)
}
