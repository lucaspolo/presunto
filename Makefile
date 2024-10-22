IMAGE_NAME := presunto
TAG := latest

GOLANGCI_VERSION := v1.60.1
export GOBIN := $(PWD)/tools

GOLANGCI_VERSION := v1.60.1
export GOLANGCI_TIMEOUT := 8m
golangci-lint := $(GOBIN)/golangci-lint-$(GOLANGCI_VERSION)
export GOLANGCI_LINT_BIN := $(golangci-lint)
$(golangci-lint):
	@echo "Installing golangci-lint $(GOLANGCI_VERSION)..." && \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b tools $(GOLANGCI_VERSION) && \
	mv $(GOBIN)/golangci-lint $(GOBIN)/golangci-lint-$(GOLANGCI_VERSION)

dev-run:
	@go run cmd/server.go

# Comando para buildar a imagem Docker
build:
	docker build -t $(IMAGE_NAME):$(TAG) .

# Comando para rodar o container
run: build
	docker run --rm -p 8080:8080 $(IMAGE_NAME):$(TAG)

# Comando para limpar a imagem Docker
clean:
	docker rmi $(IMAGE_NAME):$(TAG)

test:
	@go test ./...

lint: $(golangci-lint)
	$(GOBIN)/golangci-lint-$(GOLANGCI_VERSION) run

build-swagger:
	@cd cmd; go run github.com/swaggo/swag/cmd/swag@latest init -g server.go -o ../docs
