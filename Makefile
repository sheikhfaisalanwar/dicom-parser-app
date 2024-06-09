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


.PHONY: unit-test
unit-test: ## Run all tests other than the integration tests in the tests folder
	go test -v ./... | grep -v '\(tests\)'


.PHONY: integration-test
integration-test: ## Run the integration tests in the tests folder
	go test -v ./tests/...

.PHONY: docker-dynamodb-local-setup
docker-dynamodb-local-setup:
	echo Checking if Dynamo is running
	until docker run --rm -it --network dicom-parser-app --env-file ./docker/.env \
       -it amazon/aws-cli dynamodb list-tables --endpoint-url http://dynamodb-local:8000; do \
	   echo "Dynamo is not running yet"; \
	   sleep 2; \
   	done

	echo Creating tables
	docker run -v ./docker/dynamodb:/home/dynamodb --workdir /home/dynamodb --env-file ./docker/.env \
   		--rm --network dicom-parser-app \
	   	-it amazon/aws-cli dynamodb create-table --cli-input-json file://./docker/dynamodb/tables/dicom-document-localhost.json \
	    --endpoint-url http://dynamodb-local:8000