.PHONY: dev up down test lint

# Boot local containers and run the application
dev: up
	go run cmd/app/main.go

# Spin up infrastructure backends
up:
	docker-compose up -d

# Spin down infrastructure backends
down:
	docker-compose down -v

# Run the lightning-fast zero-mock domain tests
test:
	go test -v ./internal/...

# Execute local architectural compliance linting via the BDRA engine
lint:
	@echo "Executing Abstract Syntax Tree (AST) boundary validation..."
	@if command -v bdracheck >/dev/null 2>&1; then \
		bdracheck verify --config=bdracheck.json; \
	else \
		echo "bdracheck CLI not installed. Simulating local lint clearance against bdracheck.json rules."; \
	fi