# Variáveis
MIGRATION_DIR=cmd/migrate/migrations

build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom

migration:
	@migrate create -ext sql -dir $(MIGRATION_DIR) $(NAME)

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

# Prevenir conflitos com targets
%:
	@:
