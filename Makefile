#PROJECT_ROOT_DIR hold the absolute path to this project root dir.
PROJECT_ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: build
build:
	@go mod download && go mod verify
	@cd cmd/cart/api && go build

.PHONY: docker-migrations
docker-migrations:
	@docker exec -it go-base-service go run cmd/base/migration/main.go postgres $(command)

.PHONY: docker-create-migrations
docker-create-migrations:
	@docker exec -it go-base-service go run cmd/base/migration/main.go  -dir /go/src/github.com/domibustosb/go_base/db/$(type) postgres create null $(name)
