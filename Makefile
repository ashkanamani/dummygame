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
test:
	@echo "Running Go tests..."
	go test ./...

# Run Docker Compose services
docker-compose-run:
	@echo "Starting PostgreSQL and Redis containers with Docker Compose..."
	docker-compose up -d  # Starts the containers in detached mode


# Ensure PostgreSQL is ready before running the application
postgres_ready:
	@echo "Waiting for PostgreSQL to start..."
	@until docker exec postgresdb pg_isready -U postgres; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 1; \
	done

# Run the Go application
go_run:
	@echo "Running the Go application..."
	go run ./cmd/api

# Stop and remove Docker Compose containers
stop:
	@echo "Stopping and removing Docker Compose containers..."
	docker-compose down  # Stops and removes containers, networks, and volumes defined in docker-compose.yml
