package products

import (
	"context"
	"database/sql"
	"fmt"
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

	defer tx.Rollback()

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateProduct(product *types.CreateProductPayload) (int, error) {
	logger.Info("Validating product")
	if err := utils.Validate.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		return http.StatusBadRequest, fmt.Errorf("validation error: %s", errors)
	}

	ctx := context.Background()

	productAlreadyExists, err := s.db.FindProductByName(ctx, product.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	if productAlreadyExists.Name == "" {
		return http.StatusConflict, fmt.Errorf("product already exists")
	}

	logger.Info("inserting product")

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

	productBase, err := txQueries.InsertProduct(ctx,
		repository.InsertProductParams{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = txQueries.InsertStockProduct(ctx, repository.InsertStockProductParams{
		ProductID:         productBase.ID,
		AvailableQuantity: int32(product.Quantity),
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
