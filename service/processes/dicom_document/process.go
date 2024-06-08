package dicomdocument

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"

	"dicom-parser-app/service/processes/dicom_document/handlers"
	"dicom-parser-app/service/processes/dicom_document/repository"
	"dicom-parser-app/service/processes/dicom_document/service"
)

func InitServer() {
	e := echo.New()
	e.GET("/", handlers.Setup())

	// initialize storage
	storage := repository.NewDicomFileStorage()

	// initialize service
	s := service.NewDicomDocumentService(storage)

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
