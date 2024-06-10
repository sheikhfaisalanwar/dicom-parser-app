# dicom-parser-app

## SUMMARY
This is a simple docker application that allows you to 
- Create DICOM document records
- Read DICOM document records
  - Read all the Data Elements in a document
  - Read a specific element in a document by its tag
- Convert DICOM document records to a PNG image

## Service Stack
- Golang 1.22.4 for backend server logic
- AWS DynamoDB for data storage
- Docker for local development

## How to Run

### Prerequisites
- Docker
- Docker Compose
- Make
- Go 1.22.4

1. Clone the repository
2. Run the following command to install the dependencies
```
make init
make run-docker
```

## HOW TO USE

[Here's](https://www.postman.com/faisalanwar21/workspace/sheikh-personalprojects/request/1162575-33980d5b-6093-4dc6-a1e4-fc26ef61b925?action=share&creator=1162575&ctx=documentation&active-environment=1162575-05cf4a21-4e9d-483a-b740-e68ecd6083ce) the link to a POSTMAN collection that you can use to interact with the API

## Run the following command to run the tests
```
make unit-test
```

## STRUCTURE
The application is structured as follows:

````
├── client // A client wrapper than can be used to interact with the API
├── common
│   └── dynamodb // A wrapper around the AWS SDK for DynamoDB
├── docker
│   └── dynamodb // docker configuations for the dyanmodb
│       └── tables // table definitions
├── docs // API documentation
├── service 
│   └── processes // The main business logic of the application
│       └── dicom_document 
│           ├── handlers // The handlers for the API
│           ├── models // The models for the API
│           ├── repository // The repository for the API
│           └── service // The service for the API
├── store
│   └── dicom_files // A emulated store for the DICOM files
├── tests
│   └── assets
│       └── test_dicom_files // Test DICOM files

````


## Improvements
- Add more test coverage
- Add standardized error definitions
- Add standardized response definitions
- Add stricter request validation/authorization middlewares
- Async parsing of the DICOM files
- Replace the local store with Azure Blob Storage Emulator for local dev
- Refactor dicom parsing package implementation leakage
- Cross service interaction should be handlers or message bus



