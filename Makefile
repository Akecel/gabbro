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
gabbro/build: ## Build.
	@scripts/build.sh github.com/akecel/gabbro