package types

import "github.com/MatthewAraujo/min-ecommerce/repository"

type CostumersService interface {
	CreateCustomer(c *CreateCustomerPayload) (repository.Customer, int, error)
}

type CreateCustomerPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}
