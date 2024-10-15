IMAGE_NAME := presunto
TAG := latest

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
