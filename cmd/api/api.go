package api

import (
	"log"
	"net/http"

	"github.com/MatthewAraujo/min-ecommerce/service/customers"
	"github.com/MatthewAraujo/min-ecommerce/service/repository"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type APIServer struct {
	addr  string
	db    *repository.Queries
	redis *redis.Client
}

func NewAPIServer(addr string, db *repository.Queries, redis *redis.Client) *APIServer {
	return &APIServer{
		addr:  addr,
		db:    db,
		redis: redis,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// if the api changes in the future we can just change the version here, and the old version will still be available
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	CostumersStore := customers.NewStore(s.db)
	customersHandler := customers.NewHandler(CostumersStore)
	customersHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
