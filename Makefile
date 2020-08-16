DOCKER_COMPOSE := docker-compose -f docker/docker-compose.yaml
USER_ID := $(shell id -u)

.env:
	cp .env.dist .env

.PHONY: build
build: .env
	$(DOCKER_COMPOSE) build --build-arg USER=$(USER) --build-arg USER_ID=$(USER_ID)

.PHONY: start
start: .env
	$(DOCKER_COMPOSE) up -d

.PHONY: stop
stop: .env
	$(DOCKER_COMPOSE) down

.PHONY: sh
sh: .env
	$(DOCKER_COMPOSE) exec app sh

.PHONY: web
web: .env
	$(DOCKER_COMPOSE) exec app go run main.go
