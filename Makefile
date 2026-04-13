.PHONY: dev build test lint tidy docker-up docker-down migrate clean help

dev: ## Run with hot-reload (requires air: go install github.com/air-verse/air@latest)
	air

build: ## Compile production binary
	go build -ldflags="-w -s" -o bin/ts-background-jobs-1 ./cmd/server

test: ## Run all tests
	go test ./... -v -timeout 60s

test-cover: ## Run tests with coverage
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

lint: ## Run linters (requires golangci-lint)
	golangci-lint run ./...

tidy: ## Tidy go modules
	go mod tidy

docker-up: ## Start all services
	docker compose up -d

docker-down: ## Stop all services
	docker compose down

migrate: ## Apply database migrations (runs automatically on startup)
	@echo "Migrations run automatically on server start"

clean: ## Remove build artifacts
	rm -rf bin/ coverage.out

help: ## Show available targets
	@grep -E "^[a-zA-Z_-]+:.*##" $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
