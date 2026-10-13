.PHONY: all build test lint clean up down migrate run-api run-worker run-cli

all: build test

build:
	go build -v -o bin/restoreforge-api.exe ./cmd/api
	go build -v -o bin/restoreforge-worker.exe ./cmd/worker
	go build -v -o bin/restoreforge.exe ./cmd/restoreforge

test:
	go test -v -race ./...

test-integration:
	go test -v -tags=integration ./tests/integration/...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/ tmp/ restore-output/

up:
	docker compose -f deployments/docker/docker-compose.yml up -d

down:
	docker compose -f deployments/docker/docker-compose.yml down

migrate:
	@echo "Applying database migrations..."

run-api:
	go run ./cmd/api/main.go

run-worker:
	go run ./cmd/worker/main.go

run-cli:
	go run ./cmd/restoreforge/main.go

