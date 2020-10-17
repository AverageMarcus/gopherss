.DEFAULT_GOAL := default

IMAGE ?= docker.cluster.fun/averagemarcus/gopherss:latest
PLATFORMS ?= linux/amd64,linux/arm/v7,linux/arm64

export DOCKER_CLI_EXPERIMENTAL=enabled

.PHONY: test # Run all tests, linting and format checks
test: lint check-format run-tests

.PHONY: lint # Perform lint checks against code
lint:
	@go vet && golint -set_exit_status ./...

.PHONY: check-format # Checks code formatting and returns a non-zero exit code if formatting errors found
check-format:
	@gofmt -e -l .

.PHONY: format # Performs automatic format fixes on all code
format:
	@gofmt -s -w .

.PHONY: run-tests # Runs all tests
run-tests:
	@go test

.PHONY: fetch-deps # Fetch all project dependencies
fetch-deps:
	@go mod tidy

.PHONY: build # Build the project
build: lint check-format fetch-deps
	@go build -o gopherss .

.PHONY: docker-build # Build the docker image
docker-build:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--output "type=docker,push=false" \
		--tag $(IMAGE) \
		.

.PHONY: docker-publish # Push the docker image to the remote registry
docker-publish:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform $(PLATFORMS) \
		--output "type=image,push=true" \
		--tag $(IMAGE) \
		.

.PHONY: run # Run the application
run:
	@go run .

.PHONY: ci # Perform CI specific tasks to perform on a pull request
ci:
	@echo "⚠️ 'ci' unimplemented"

.PHONY: release # Release the latest version of the application
release:
	@echo "⚠️ 'released' unimplemented"

.PHONY: help # Show this list of commands
help:
	@echo "gopherss"
	@echo "Usage: make [target]"
	@echo ""
	@echo "target	description" | expand -t20
	@echo "-----------------------------------"
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1	\2/' | expand -t20

default: test build
