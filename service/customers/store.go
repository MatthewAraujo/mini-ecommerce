package customers

import (
	"context"

	"github.com/MatthewAraujo/min-ecommerce/service/repository"
)

type Store struct {
	db *repository.Queries
}

func NewStore(db *repository.Queries) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllCustomers() ([]repository.Customer, error) {
	customers, err := s.db.FinAllCustomers(context.Background())
	if err != nil {
		return nil, err
	}

	return customers, nil
}
