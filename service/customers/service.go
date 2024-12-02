package customers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	configs "github.com/MatthewAraujo/min-ecommerce/config"
	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/service/auth"
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

	defer tx.Rollback()

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateCustomer(customer *types.CreateCustomerPayload) (int, error) {
	logger.Info("Validating customers")
	if err := utils.Validate.Struct(customer); err != nil {
		errors := err.(validator.ValidationErrors)
		return http.StatusBadRequest, fmt.Errorf("validation error: %s", errors)
	}

	ctx := context.Background()

	emailAlreadyExists, err := s.db.FindCustomerByEmail(ctx, customer.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	if emailAlreadyExists.Email != "" {
		return http.StatusConflict, fmt.Errorf("email already has been used")
	}

	logger.Info("inserting customers")
	hashedPassword, err := auth.HashPassword(customer.Password)
	if err != nil {
		logger.Warn(err.Error())
		return http.StatusInternalServerError, fmt.Errorf("Internal error")
	}
	_, err = s.db.InsertCustomers(ctx,
		repository.InsertCustomersParams{
			Name:     customer.Name,
			Email:    customer.Email,
			Password: hashedPassword,
		})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *Service) Login(customer *types.LoginCustomerPayload) (string, int, error) {
	logger.Info("Service.Login", "Searching customer by email")
	u, err := s.db.FindCustomerByEmail(context.Background(), customer.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.LogError("Service.Login", fmt.Errorf("customer not found: %s", customer.Email))
			return "", http.StatusNotFound, fmt.Errorf("customer not found")
		}
		logger.LogError("Service.Login", fmt.Errorf("error finding customer: %w", err))
		return "", http.StatusInternalServerError, fmt.Errorf("error finding customer: %w", err)
	}

	logger.Info("Service.Login", "Customer found, verifying password")
	if !auth.ComparePasswords(u.Password, []byte(customer.Password)) {
		logger.LogError("Service.Login", fmt.Errorf("invalid password for customer: %s", customer.Email))
		return "", http.StatusUnauthorized, fmt.Errorf("invalid password")
	}

	logger.Info("Service.Login", "Password verified, generating token")
	secret := []byte(configs.Envs.JWT.JWTSecret)
	token, err := auth.CreateJWT(secret, string(u.ID))
	if err != nil {
		logger.LogError("Service.Login", fmt.Errorf("error creating token: %w", err))
		return "", http.StatusInternalServerError, fmt.Errorf("error creating token: %w", err)
	}

	logger.Info("Service.Login", "Token generated successfully")
	return token, http.StatusAccepted, nil
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
