package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dicom-parser-app/client"
)

func Setup() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	}
}

// Create godoc
// @Summary Create a new Dicom Document
// @Description Uploads a new Dicom Document to the server and creates a record in the database
// @Tags create
// @Accept  mpfd
// @Produce  json
// @Param file formData file true "Dicom file"
// @Success 200 {object} client.CreateDicomDocumentResponse
// @Failure 400 {object} string "Could not get multipart form"
// @Failure 400 {object} string "No file found in form"
// @Failure 500 {object} string "Error uploading dicom document"
// @Router /create [post]
func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Could not get multipart form")
		}
		file := form.File["file"]

		if len(file) == 0 {
			return c.JSON(http.StatusBadRequest, "No file found in form")
		}
		uploadedDoc := file[0]

		if uploadedDoc == nil {
			return c.JSON(http.StatusBadRequest, "No file found in form")
		}

		request := client.CreateDicomDocumentRequest{
			File: *uploadedDoc,
		}

		doc, err := h.service.CreateDicomDocument(c, request)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error uploading dicom document")
		}
		c.Logger().Info(doc)

		createResponse := client.CreateDicomDocumentResponse{
			ID:       doc.ID,
			Name:     doc.Name,
			Location: doc.Location,
		}

		return c.JSON(http.StatusOK, createResponse)
	}
}
