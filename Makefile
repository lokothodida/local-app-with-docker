DOCKER_COMPOSE := docker-compose -f docker/docker-compose.yaml
USER_ID := $(shell id -u)

.env:
	cp .env.dist .env

.PHONY: build
build:
	$(DOCKER_COMPOSE) build --build-arg USER=$(USER) --build-arg USER_ID=$(USER_ID)

.PHONY: start
start:
	$(DOCKER_COMPOSE) up -d

.PHONY: stop
stop:
	$(DOCKER_COMPOSE) down

.PHONY: sh
sh:
	$(DOCKER_COMPOSE) exec app sh

.PHONY: web
web:
	$(DOCKER_COMPOSE) exec app go run main.go
