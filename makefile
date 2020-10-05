.DEFAULT_GOAL := help

.PHONY: generate
generate: ## Invoke go generate
	go generate ./...

.PHONY: help
help: ## help
	@grep -E '^[0-9a-zA-Z_/()$$-]+:.*?## .*$$' $(lastword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'