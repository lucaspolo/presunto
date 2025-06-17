IMAGE_NAME := presunto
TAG := latest

dev-run-http:
	@go run cmd/http/server.go

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

lint:
	go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint run

build-swagger:
	@cd cmd; go run github.com/swaggo/swag/cmd/swag@latest init -g server.go -o ../docs
