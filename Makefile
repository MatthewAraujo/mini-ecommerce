# Variáveis
MIGRATION_DIR=cmd/migrate/migrations

build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom

# Criação de uma nova migração
migration:
	@migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)

# Executar migrações "up"
migrate-up:
	@go run cmd/migrate/main.go up

# Reverter migrações "down"
migrate-down:
	@go run cmd/migrate/main.go down

# Prevenir conflitos com targets
%:
	@:
