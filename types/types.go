package types

import (
	"context"

	"github.com/MatthewAraujo/min-ecommerce/repository"
)

type CostumersService interface {
	GetAllCustomers() ([]repository.Customer, int, error)
	Order(context.Context, int32, []repository.OrderItem) (int, error)
}
