package products

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/types"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	db         *repository.Queries
	dbTx       *sql.DB
	redisStore *redis.Client
}

func NewService(db *repository.Queries, dbTx *sql.DB, redisStore *redis.Client) *Service {
	return &Service{
		db:         db,
		dbTx:       dbTx,
		redisStore: redisStore,
	}
}

func (s *Service) BeginTransaction(ctx context.Context) (*repository.Queries, *sql.Tx, error) {
	tx, err := s.dbTx.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateProduct(product *types.CreateProductPayload) (int, error) {
	logger.Info("Validating product")

	if err := utils.Validate.Struct(product); err != nil {
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
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	productAlreadyExists, err := txQueries.FindProductByName(ctx, product.Name)
	if err != nil && err != sql.ErrNoRows {
		logger.Warn(err.Error())
		return http.StatusInternalServerError, fmt.Errorf("Internal error")
	}

	if productAlreadyExists.Name != "" {
		return http.StatusConflict, fmt.Errorf("product already exists")
	}

	logger.Info("inserting product")
	productBase, err := txQueries.InsertProduct(ctx,
		repository.InsertProductParams{
			Name:        product.Name,
			Description: utils.ToNullString(product.Description),
			Price:       product.Price,
		})

	if err != nil {
		logger.LogError("ERRO", err)
		return http.StatusInternalServerError, fmt.Errorf("internal error")
	}

	logger.Info("inserting stock")
	_, err = txQueries.InsertStockProduct(ctx, repository.InsertStockProductParams{
		ProductID:         productBase.ID,
		AvailableQuantity: int32(product.Quantity),
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *Service) GetAllProducts(p *types.GetAllProductsPayload) (types.GetAllProductsResponse, int, error) {
	logger.Info("Get All Products")
	if err := utils.Validate.Struct(p); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := utils.TranslateValidationErrors(validationErrors)

			response, _ := json.Marshal(errorMessages)
			return types.GetAllProductsResponse{}, http.StatusBadRequest, fmt.Errorf("validation error: %s", response)
		}

		return types.GetAllProductsResponse{}, http.StatusInternalServerError, fmt.Errorf("internal server error: %s", err)
	}

	ctx := context.Background()
	offset := (p.Page - 1) * utils.PAGINATION_LIMIT

	products, err := s.db.GetAllProductsPagination(ctx, repository.GetAllProductsPaginationParams{
		Limit:  int32(utils.PAGINATION_LIMIT),
		Offset: int32(offset),
	})
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return types.GetAllProductsResponse{}, http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	totalProducts, err := s.db.CountProducts(ctx)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return types.GetAllProductsResponse{}, http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	totalPages := math.Ceil(float64(totalProducts) / float64(utils.PAGINATION_LIMIT))
	PAGINATION_LIMIT := int64(utils.PAGINATION_LIMIT)
	productsRemaining := totalProducts - (int64(p.Page)-1)*PAGINATION_LIMIT

	pagination := types.Pagination{
		CurrentPage: p.Page,
		TotalPages:  int(totalPages),
		TotalItems:  int(totalProducts),
		PageSize:    int(math.Min(float64(utils.PAGINATION_LIMIT), float64(productsRemaining))),
	}

	response := types.GetAllProductsResponse{
		Products:   products,
		Pagination: pagination,
	}

	return response, http.StatusOK, nil
}

func (s *Service) getTopTenFromRedis() ([]repository.Product, error) {
	ctx := context.Background()
	products := "products:top-ten"

	productsJSON, err := s.redisStore.Get(ctx, products).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("cache expired or key not found")
		}
		return nil, fmt.Errorf("failed to get posts from Redis: %w", err)
	}

	var pd []repository.Product
	if err := json.Unmarshal([]byte(productsJSON), &pd); err != nil {
		return nil, fmt.Errorf("failed to decode posts JSON: %w", err)
	}

	return pd, nil
}

func (s *Service) setTopTenToRedis(products []*repository.Product) error {
	ctx := context.Background()
	productsQuery := "products:top-ten"

	productsJSON, err := json.Marshal(products)
	if err != nil {
		return fmt.Errorf("failed to serialize posts: %w", err)
	}

	err = s.redisStore.Set(ctx, productsQuery, productsJSON, 24*time.Second).Err()
	if err != nil {
		return fmt.Errorf("failed to set posts in Redis: %w", err)
	}

	return nil
}

func (s *Service) GetMostSelledProducts() ([]repository.Product, int, error) {
	products, err := s.getTopTenFromRedis()

	if err != nil && len(products) < 10 {
		products, err := s.db.GetTop10MostSoldProducts(context.Background())
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		return products, http.StatusOK, nil
	}
	return products, http.StatusOK, nil
}
