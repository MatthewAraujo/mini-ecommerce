# Variáveis
MIGRATION_DIR=cmd/migrate/migrations

build:
	@go build -o bin/ecom cmd/main.go

test:
	@echo "Testing..."
	@go test -v ./...
	
run: build
	@echo "Building..."
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

watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi
