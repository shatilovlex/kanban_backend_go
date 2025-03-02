include .env

LOCAL_BIN:=$(CURDIR)/bin

ENV_DIR = .env

name = kanban_scheme

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


env-prepare: ## copy .env from example
	cp -n .env.example .env

audit: ## Run linter checks
	golangci-lint run ./...

tidy: ## go  mod tidy run and reformat code
	go mod tidy
	go fmt ./...
	golangci-lint run --fix ./...

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-add:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create $(name) sql

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

sqlc:
	$(LOCAL_BIN)/sqlc generate -f sqlc.yaml

