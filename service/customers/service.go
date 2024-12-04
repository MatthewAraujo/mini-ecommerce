package customers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := utils.TranslateValidationErrors(validationErrors)

			response, _ := json.Marshal(errorMessages)
			return http.StatusBadRequest, fmt.Errorf("validation error: %s", response)
		}

		return http.StatusInternalServerError, fmt.Errorf("internal server error: %s", err)
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
			Role:     repository.UserRole(customer.Role),
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
	token, err := auth.CreateJWT(secret, u.ID, string(u.Role))
	if err != nil {
		logger.LogError("Service.Login", fmt.Errorf("error creating token: %w", err))
		return "", http.StatusInternalServerError, fmt.Errorf("error creating token: %w", err)
	}

	logger.Info("Service.Login", "Token generated successfully")
	return token, http.StatusAccepted, nil
}
