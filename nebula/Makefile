# Makefile

GO_BINARY_NAME := nebula-server-go
DOCKER_IMAGE_NAME := NEBULA-SERVER-GO

build:
	go build -o bin/nebula-server main.go
	docker build -t $(GO_BINARY_NAME) ./bin
	docker image save -o $(DOCKER_IMAGE_NAME) $(GO_BINARY_NAME)