.PHONY: cli help toml2cli

default: cli

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

toml2cli: ## Rebuild the CLI interface
	toml2cli -in-file=cli.toml -out-file=main.go

cli: toml2cli ## Make the CLI binary for local testing.
	CGO_ENABLED=0 go build -o ope ./
