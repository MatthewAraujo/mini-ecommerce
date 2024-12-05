package types

import "github.com/MatthewAraujo/min-ecommerce/repository"

type ValidationErrorResponse struct {
	Field      string `json:"field"`
	Validation string `json:"validation"`
	Value      string `json:"value,omitempty"`
	Message    string `json:"message"`
}

type CostumersService interface {
	CreateCustomer(c *CreateCustomerPayload) (int, error)
	Login(c *LoginCustomerPayload) (string, int, error)
}

type CreateCustomerPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
	Role     string `json:"role" validate:"required,oneof=user admin"`
}

type LoginCustomerPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type ProductService interface {
	CreateProduct(p *CreateProductPayload) (int, error)
	GetAllProducts(p *GetAllProductsPayload) (GetAllProductsResponse, int, error)
}

type CreateProductPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,max=255"`
	Price       string `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

type GetAllProductsPayload struct {
	Page int `json:"page" validate:"required,min=1"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	TotalPages  int `json:"totalPages"`
	TotalItems  int `json:"totalItems"`
	PageSize    int `json:"pageSize"`
}

type GetAllProductsResponse struct {
	Products   []repository.Product
	Pagination Pagination
}

type OrderService interface {
	Order(o *CreateOrderPayload) (int, error)
}

type CreateOrderPayload struct {
	CustomerID int32           `json:"customer_id" validate:"required,min=1"`
	Items      []OrderItemData `json:"items" validate:"required,dive"`
}

type OrderItemData struct {
	ProductID int32 `json:"product_id" validate:"required"`
	Quantity  int32 `json:"quantity" validate:"required,min=1"`
}
