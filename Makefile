BIN := "./.bin/app"
DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"
DOCKER_COMPOSE_TEST_FILE := "./deployments/docker-compose.tests.yml"
APP_NAME := "tg-notebot"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" cmd/app/main.go

run: build 
	 $(BIN)

test: 
	go test --short -race ./internal/...

.PHONY: build test

up:
	docker-compose --env-file ./app.env -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up --build -d

down:
	docker-compose --env-file ./app.env -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} down --volumes

mocks:
	mockgen -source=./internal/repository/repository.go -destination ./internal/repository/mocks/mock.go
	mockgen -source=./internal/services/service.go -destination ./internal/services/mocks/mock.go
