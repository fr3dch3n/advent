SERVICE = advent
.DEFAULT_GOAL:= help

.PHONY: install
.SILENT: install
install: ## install dependencies
ifeq ($(CI),true)
	$(info "Running in CI mode")
else
	pre-commit install
	go mod download
	go mod tidy
	go mod verify
endif

.PHONY: pre-commit-install
.SILENT: pre-commit-install
pre-commit-install: ## install pre-commit
	pre-commit install

.PHONY: pre-commit
.SILENT: pre-commit
pre-commit: pre-commit-install ## run pre-commit
	pre-commit run --all-files

# GOLANG
.PHONY: test
.SILENT: test
test: install ## run tests
	gotestsum

.PHONY: build
.SILENT: build
build: install ## build binary
	go build -o $(SERVICE) .

.PHONY: run
.SILENT: run
run: install ## run binary
	go run main.go

# HELP
.PHONY: help
.SILENT: help
help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
