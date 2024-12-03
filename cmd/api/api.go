package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/service/customers"
	"github.com/MatthewAraujo/min-ecommerce/service/products"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type APIServer struct {
	addr  string
	db    *repository.Queries
	dbTx  *sql.DB
	redis *redis.Client
}

func NewAPIServer(addr string, db *repository.Queries, dbTx *sql.DB, redis *redis.Client) *APIServer {
	return &APIServer{
		addr:  addr,
		db:    db,
		dbTx:  dbTx,
		redis: redis,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// if the api changes in the future we can just change the version here, and the old version will still be available
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	costumersService := customers.NewService(s.db, s.dbTx)
	customersHandler := customers.NewHandler(costumersService)
	customersHandler.RegisterRoutes(subrouter.PathPrefix("/customers").Subrouter())

	productService := products.NewService(s.db, s.dbTx)
	productHandler := products.NewHandler(productService, *s.db)
	productHandler.RegisterRoutes(subrouter.PathPrefix("/product").Subrouter())

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
