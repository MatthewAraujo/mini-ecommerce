package types

import "github.com/MatthewAraujo/min-ecommerce/service/repository"

type CostumersStore interface {
	GetAllCustomers() ([]repository.Customer, error)
}
