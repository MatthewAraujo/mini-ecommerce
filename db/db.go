package db

import (
	"database/sql"
	"log"

	"github.com/MatthewAraujo/min-ecommerce/pkg/assert"
	_ "github.com/jackc/pgx/v5/stdlib" // Importa o driver para registrar automaticamente
)

func NewMyPostgresSQLStorage(url string) (*sql.DB, error) {
	assert.NotNil(url, "Config cannot be nil")
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return db, nil
}
