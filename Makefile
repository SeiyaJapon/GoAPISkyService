#!/bin/bash

DOCKER_API = go-api
DB_CONTAINER = postgres
OS := $(shell uname)

ifeq ($(OS),Darwin)
	UID = $(shell id -u)
else ifeq ($(OS),Linux)
	UID = $(shell id -u)
else
	UID = 1000
endif

## —— 📦 The amazing API Go Contenerizada Makefile 📦 ——————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— 🐋 Docker 🐋 ——————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————

build: create-network up ## First run / installation will call this target

rebuild: destroy build ## Rebuild all containers

building: ## Rebuild all containers without cache
	U_ID=${UID} docker-compose build --no-cache

up: ## Start the docker environment
	U_ID=${UID} docker-compose up -d --remove-orphans

restart: rebuild up ## Rebuild and start the containers

run: ## Start the containers
	docker network create app-network || true
	U_ID=${UID} docker-compose up -d

stop: ## Stop the containers
	U_ID=${UID} docker-compose stop

restart: ## Restart the containers
	$(MAKE) stop && $(MAKE) run

rebuild-all: ## Rebuilds all the containers
	U_ID=${UID} docker-compose build

create-network: ## Create Docker network
	docker network create app-network || true

down: ## Stop and remove containers
	docker-compose down --remove-orphans

destroy: ## Remove all containers, images, volumes
	docker-compose down --rmi all --volumes --remove-orphans

## —— 🦫 Go API 🦫 ———————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————————

logs: ## Tail the API logs
	U_ID=${UID} docker logs -f ${DOCKER_API}

logs-postgres: ## Tail the PostgreSQL logs
	U_ID=${UID} docker logs -f ${DB_CONTAINER}

enter: ## SSH into the API container
	U_ID=${UID} docker exec -it ${DOCKER_API} /bin/sh

enter-postgres: ## SSH into the PostgreSQL container
	U_ID=${UID} docker exec -it ${DB_CONTAINER} /bin/bash

reload-api: ## Restart only the API container
	docker-compose restart api

test: ## Run tests in the API container
	U_ID=${UID} docker exec -it ${DOCKER_API} go test ./...

show-endpoints: ## Show available API endpoints
	@echo "API: http://localhost:8080"
	@echo "Base de datos: PostgreSQL: localhost:5432"
