package dicomdocument

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dynamocommon "dicom-parser-app/common/dynamodb"
	"dicom-parser-app/service/processes/dicom_document/handlers"
	"dicom-parser-app/service/processes/dicom_document/repository"
	"dicom-parser-app/service/processes/dicom_document/service"
)

type Config struct {
	DicomDocumentDynamoTable string `default:"dicom-documents-LOCAL" envconfig:"DICOM_DOCUMENT_DYNAMO_TABLE"`
}

func NewConfig() *Config {
	return &Config{
		DicomDocumentDynamoTable: "dicom-documents-LOCAL",
	}
}

func InitServer() {
	config := NewConfig()
	// initialize echo
	e := echo.New()
	e.Debug = true

	// Naive logger
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

	e.GET("/", handlers.Setup())

	// initialize storage
	storage := repository.NewDicomFileStorage()
	dynamoClient := dynamocommon.NewDynamoDBClient(e.Logger)
	repo := repository.NewRepository(dynamoClient, dynamocommon.DynamoTableParams{
		TableName: config.DicomDocumentDynamoTable,
	})

	// initialize service
	s := service.NewDicomDocumentService(storage, repo)

	// initialize handler
	h := handlers.NewHandler(s)

	h.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)
}
