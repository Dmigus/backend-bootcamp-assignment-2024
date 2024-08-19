include .env

.PHONY: run-storage
run-storage: bin/goose
	docker-compose up -d --wait postgres
	@$(LOCAL_BIN)/goose -dir ./migrations postgres "host=localhost port=$(STORAGE_PORT) user=$(STORAGE_USER) password=$(STORAGE_PASSWORD) dbname=$(STORAGE_DATABASE) sslmode=disable" up

.PHONY: run-all
run-all: run-storage
	docker-compose build -q
	docker-compose up -d --wait renting

stop-all:
	docker-compose down


.PHONY: coverage
coverage:
	@go test -tags=unit,integration -coverprofile=coverage.out -parallel=10 ./...
	@go tool cover -html coverage.out -o coverage.html
	@rm coverage.out

# Используем bin в текущей директории для установки плагина генерации и миграции
LOCAL_BIN:=$(CURDIR)/bin

bin/goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.19.2

bin/oapi-codegen:
	GOBIN=$(LOCAL_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

.PHONY: gen-http-endpoints
gen-http-endpoints: bin/oapi-codegen
	$(LOCAL_BIN)/oapi-codegen -generate "types,std-http" -include-tags "noAuth" -o ./internal/controllers/auth/gen.go -package auth ./api/openapiv3/api.yaml
	$(LOCAL_BIN)/oapi-codegen -generate "types,std-http" -include-tags "authOnly,moderationsOnly" -o ./internal/controllers/renting/gen.go -package renting ./api/openapiv3/api.yaml

bin/sqlc:
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0

.PHONY: sqlc
sqlc: bin/sqlc
	$(LOCAL_BIN)/sqlc generate --file=./sqlc/sqlc.yaml

.PHONY: lint
lint:
	@golangci-lint run
