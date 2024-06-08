SHELL := /bin/bash
.SHELLFLAGS = -o pipefail -e -c

export TEST_FILES_LOCATION = ./tests/assets/test_dicom_files/

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  init       - Initialize the project"
	@echo "  build      - Build the project"
	@echo "  run        - Run the project"
	@echo "  run-docker - Run the project in a docker container"

.PHONY: init
init:
	go mod tidy
	go mod vendor

.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: run
run:
	go run main.go

.PHONY: run-docker
run-docker:
	docker compose up --build --force-recreate --detach