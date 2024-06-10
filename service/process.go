package service

import (
	"encoding/json"
	"fmt"
	"os"

	"dicom-parser-app/common"
	dynamocommon "dicom-parser-app/common/dynamodb"
	datasetHandlers "dicom-parser-app/service/processes/dicom_dataset/handlers"
	dicomDatasetRepository "dicom-parser-app/service/processes/dicom_dataset/repository"
	dicomDatasetService "dicom-parser-app/service/processes/dicom_dataset/service"
	"dicom-parser-app/service/processes/dicom_document/handlers"
	"dicom-parser-app/service/processes/dicom_document/repository"
	"dicom-parser-app/service/processes/dicom_document/service"

	"github.com/labstack/echo/v4"
)

type Config struct {
	DicomDocumentDynamoTable string `default:"dicom-documents-LOCAL" envconfig:"DICOM_DOCUMENT_DYNAMO_TABLE"`
	DicomDatasetDynamoTable  string `default:"dicom-datasets-LOCAL" envconfig:"DICOM_DATASET_DYNAMO_TABLE"`
}

func NewConfig() *Config {
	return &Config{
		DicomDocumentDynamoTable: "dicom-documents-LOCAL",
		DicomDatasetDynamoTable:  "dicom-datasets-LOCAL",
	}
}

func InitServer() {
	config := NewConfig()

	fmt.Println("Starting dicom-parser service")
	// initialize echo
	e := echo.New()
	e.Debug = true
	logMiddleware := common.NewLoggerMiddleware()

	e.Use(logMiddleware)

	e.GET("/", handlers.Setup())

	// initialize storage
	storage := repository.NewDicomFileStorage()
	dynamoClient := dynamocommon.NewDynamoDBClient(e.Logger)
	documentRepo := repository.NewRepository(dynamoClient, dynamocommon.DynamoTableParams{
		TableName: config.DicomDocumentDynamoTable,
	})
	// initialize dataset repository
	datasetRepo := dicomDatasetRepository.NewRepository(dynamoClient, dynamocommon.DynamoTableParams{
		TableName: config.DicomDatasetDynamoTable,
	})

	// initialize document service
	s := service.NewDicomDocumentService(storage, documentRepo)

	//dependency
	datasetService := dicomDatasetService.NewDicomDatasetService(datasetRepo, s)

	// initialize dataset handler
	datasethandlers := datasetHandlers.NewHandler(datasetService)

	datasethandlers.RegisterRoutes(e, logMiddleware)

	// initialize document handler
	h := handlers.NewHandler(s)

	h.RegisterRoutes(e, logMiddleware)
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
