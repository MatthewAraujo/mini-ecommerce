package customers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/go-playground/validator"
)

type Service struct {
	db *repository.Queries

	dbTx *sql.DB
}

func NewService(db *repository.Queries, dbTx *sql.DB) *Service {
	return &Service{
		db:   db,
		dbTx: dbTx,
	}
}

func (s *Service) BeginTransaction(ctx context.Context) (*repository.Queries, *sql.Tx, error) {
	tx, err := s.dbTx.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	defer tx.Rollback()

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateCustomer(customer *types.CreateCustomerPayload) (repository.Customer, int, error) {
	if err := utils.Validate.Struct(customer); err != nil {
		errors := err.(validator.ValidationErrors)
		return repository.Customer{}, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors)
	}

	ct, err := s.db.InsertCustomers(context.Background(),
		repository.InsertCustomersParams{
			Name:     customer.Name,
			Email:    customer.Email,
			Password: customer.Password,
		})

	if err != nil {
		return repository.Customer{}, http.StatusInternalServerError, err
	}

	return ct, http.StatusCreated, nil
}

func (s *Service) Order(ctx context.Context, customerID int32, orderItems []repository.OrderItem) (int, error) {
	txQueries, tx, err := s.BeginTransaction(ctx)
	if err != nil {
		return 500, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	order, err := txQueries.CreateOrder(ctx, customerID)
	if err != nil {
		log.Printf("Failed to create order for customer ID %d: %v", customerID, err)
		return http.StatusInternalServerError, fmt.Errorf("could not create order: %w", err)
	}

	for i, item := range orderItems {
		if i == 1 {
			err = fmt.Errorf("simulated error: failed to process item %d", i+1)
			log.Printf("Error: %v", err)
			return 400, err
		}

		stock, err := txQueries.GetStockByProductID(ctx, item.ProductID)
		if err != nil {
			log.Printf("Failed to get stock for product ID %d: %v", item.ProductID, err)
			return http.StatusInternalServerError, fmt.Errorf("failed to fetch stock for product ID %d: %w", item.ProductID, err)
		}

		if stock.AvailableQuantity < item.Quantity {
			err = fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
			log.Printf("Insufficient stock for product ID %d: %v", item.ProductID, err)
			return http.StatusConflict, err
		}

		_, err = txQueries.AddOrderItem(ctx, repository.AddOrderItemParams{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
		if err != nil {
			log.Printf("Failed to add item to order ID %d: %v", order.ID, err)
			return 500, fmt.Errorf("failed to add item to order: %w", err)
		}

		_, err = txQueries.DecreaseStock(ctx, repository.DecreaseStockParams{
			ProductID:         item.ProductID,
			AvailableQuantity: item.Quantity,
		})
		if err != nil {
			log.Printf("Failed to decrease stock for product ID %d: %v", item.ProductID, err)
			return 500, fmt.Errorf("failed to update stock for product ID %d: %w", item.ProductID, err)
		}
	}

	log.Printf("Transaction completed successfully for customer ID %d, order ID %d", customerID, order.ID)
	return http.StatusCreated, nil
}
