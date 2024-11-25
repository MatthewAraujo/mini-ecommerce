package db

import (
	"database/sql"
	"log"

	"github.com/MatthewAraujo/min-ecommerce/pkg/assert"
)

func NewMyPostgresSQLStorage(url string) (*sql.DB, error) {
	assert.NotNil(url, "Mysql Config cannot be nil")
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db, nil
}
