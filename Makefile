

# Используем bin в текущей директории для установки плагина
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: install-swagger
install-swagger:
	go version
	GOBIN=$(LOCAL_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

.PHONY: gen-http-endpoints
gen-http-endpoints: install-swagger
	$(LOCAL_BIN)/oapi-codegen -generate "types,std-http" -o ./internal/controllers/gen.go -package controllers ./api/openapiv3/api.yaml


.PHONY: lint
lint:
	@golangci-lint run
