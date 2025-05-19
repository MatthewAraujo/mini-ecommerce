package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/go-playground/validator/v10"
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

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) Order(order *types.CreateOrderPayload) (int, error) {
	if err := utils.Validate.Struct(order); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := utils.TranslateValidationErrors(validationErrors)

			response, _ := json.Marshal(errorMessages)
			return http.StatusBadRequest, fmt.Errorf("validation error: %s", response)
		}

		return http.StatusInternalServerError, fmt.Errorf("internal server error: %s", err)
	}
	ctx := context.Background()
	txQueries, tx, err := s.BeginTransaction(ctx)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	userExists, err := s.db.FindCustomerByID(ctx, order.CustomerID)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	if userExists.Name == "" {
		return http.StatusConflict, fmt.Errorf("customer doest not exists")
	}

	newOrder, err := txQueries.CreateOrder(ctx, order.CustomerID)
	if err != nil {
		logger.LogError("Could not create order", err)
		return http.StatusInternalServerError, fmt.Errorf("could not create order: %w", err)
	}

	for _, item := range order.Items {

		product, err := txQueries.FindProductByID(ctx, item.ProductID)
		if err != nil {
			if err != sql.ErrNoRows {
				logger.Warn(err.Error())
				return http.StatusInternalServerError, fmt.Errorf("Internal error")
			}
		}

		if product.Name == "" {
			return http.StatusConflict, fmt.Errorf("product does not exists")

		}

		stock, err := txQueries.GetStockByProductID(ctx, item.ProductID)
		if err != nil {
			logger.LogError("Failed to get stock for product", err)
			return http.StatusInternalServerError, fmt.Errorf("failed to fetch stock for product ID %d: %w", item.ProductID, err)
		}

		if stock.AvailableQuantity < item.Quantity {
			err = fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
			log.Printf("Insufficient stock for product ID %d: %v", item.ProductID, err)
			return http.StatusConflict, err
		}

		_, err = txQueries.AddOrderItem(ctx, repository.AddOrderItemParams{
			OrderID:   newOrder.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
		if err != nil {
			logger.LogError("Failed to add item to order", err)
			return 500, fmt.Errorf("failed to add item to order: %w", err)
		}

		_, err = txQueries.DecreaseStock(ctx, repository.DecreaseStockParams{
			ProductID:         item.ProductID,
			AvailableQuantity: item.Quantity,
		})

		if err != nil {
			logger.LogError("Failed to decrease stock for product", err)
			return 500, fmt.Errorf("failed to update stock for product ID %d: %w", item.ProductID, err)
		}
	}

	return http.StatusCreated, nil
}
