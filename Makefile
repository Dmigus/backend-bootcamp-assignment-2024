

# Используем bin в текущей директории для установки плагина генерации http ручек из openapi
LOCAL_BIN:=$(CURDIR)/bin

bin/oapi-codegen:
	GOBIN=$(LOCAL_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

.PHONY: gen-http-endpoints
gen-http-endpoints: bin/oapi-codegen
	#$(LOCAL_BIN)/oapi-codegen -generate "types,std-http" -o ./internal/controllers/gen.go -package controllers ./api/openapiv3/api.yaml
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
