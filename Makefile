# ---- Config (edit if needed) -----------------------------------------------
DB_URL      ?= postgres://online:online@localhost:5432/online_test?sslmode=disable
JWT         ?= change_me_now
GRAPHQL_PORT?= 8080
GRPC_PORT   ?= 9090

# ---- Phony targets ----------------------------------------------------------
.PHONY: help db-up db-down run-api run-grpc graphql grpc build clean

help:
	@echo "Targets:"
	@echo "  db-up       - start postgres with docker-compose"
	@echo "  db-down     - stop postgres"
	@echo "  run-api     - run GraphQL HTTP server (:$(GRAPHQL_PORT))"
	@echo "  run-grpc    - run gRPC server (:$(GRPC_PORT))"
	@echo "  graphql     - run gqlgen generate"
	@echo "  grpc        - generate gRPC stubs from proto/"
	@echo "  build       - build both servers"
	@echo "  clean       - remove binaries"

db-up:
	docker compose up -d

db-down:
	docker compose down

run-api:
	@echo "Starting GraphQL API on :$(GRAPHQL_PORT)"
	DATABASE_URL="$(DB_URL)" JWT_SECRET="$(JWT)" go run ./cmd/api

run-grpc:
	@echo "Starting gRPC server on :$(GRPC_PORT)"
	DATABASE_URL="$(DB_URL)" go run ./cmd/grpc

graphql:
	gqlgen generate

grpc:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build:
	mkdir -p bin
	GOOS=linux  GOARCH=amd64 go build -o bin/api ./cmd/api
	GOOS=linux  GOARCH=amd64 go build -o bin/grpc ./cmd/grpc
	@echo "Built binaries in ./bin"

clean:
	rm -rf bin
