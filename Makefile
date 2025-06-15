# .PHONY tells Make that these are not filenames but tasks
.PHONY: build test lint docker run

include .env

# Default target: run the application
run: docker-compose-run postgres_ready go_run

# Build the Go application
build:
	@echo "Building the Go application..."
	go build -o api ./cmd/api


# Run tests
test-integration:
	@echo "Running Integration tests..."
	INTEGRATION_TEST=true go test ./... -v


# Run Docker Compose services
docker-compose-run:
	@echo "Starting PostgreSQL and Redis containers with Docker Compose..."
	docker-compose up -d  # Starts the containers in detached mode


# Ensure PostgreSQL is ready before running the application
postgres_ready:
	@echo "Waiting for PostgreSQL to start..."
	@until docker exec pg-game pg_isready -U postgres; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 1; \
	done

# Run the Go application
go_run:
	@echo "Running the Go application..."
	go run main.go serve

# Stop and remove Docker Compose containers
stop:
	@echo "Stopping and removing Docker Compose containers..."
	docker-compose down  # Stops and removes containers, networks, and volumes defined in docker-compose.yml
