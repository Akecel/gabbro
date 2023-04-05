# Color
RED		:= $(shell printf "\033[31m")
GREEN	:= $(shell printf "\033[32m")
YELLOW	:= $(shell printf "\033[33m")
BLUE	:= $(shell printf "\033[34m")
BOLD	:= $(shell printf "\033[1m")
RESET	:= $(shell printf "\033[m")

DOCKER_COMPOSE ?= docker-compose
SAIL ?= ./vendor/bin/sail
CS ?= ./vendor/bin/phpcs
LINTER-FIX ?= ./vendor/bin/pint
LINTER-TEST ?= ./vendor/bin/pint --test

.PHONY: help
help: ## Provides help information on available commands.
	@printf "Usage: make <command>\n\n"
	@printf "Commands:\n"
	@awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
	  printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)

.PHONY: gabbro/format
gabbro/format: ## Run gofmt.
	@gofmt -s -w .

.PHONY: gabbro/build
gabbro/build: ## Run gofmt.
	@scripts/build.sh github.com/akecel/gabbro