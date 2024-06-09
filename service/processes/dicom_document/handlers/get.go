package handlers

import (
	"net/http"

	"dicom-parser-app/client"

	"github.com/labstack/echo/v4"
)

// GetByID godoc
// @Summary Get a Dicom Document by ID
// @Description Retrieves a Dicom Document by its ID
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Success 200 {object} client.GetDicomDocumentResponse
// @Failure 400 {object} string "No file name provided"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /get/{id} [get]
func (h *Handler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No file name provided")
		}
		doc, err := h.service.GetDicomDocumentByID(c, docID)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}

		response := client.GetDicomDocumentResponse{
			ID:       doc.ID,
			Name:     doc.Name,
			Location: doc.Location,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// GetDataByID godoc
// @Summary Get Dicom Document data by ID
// @Description Retrieves the data of a Dicom Document by its ID
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Success 200 {object} client.GetDocumentDataResponse
// @Failure 400 {object} string "No file name provided"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /get/{id}/data [get]
func (h *Handler) GetDataByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No file name provided")
		}
		doc, err := h.service.GetDocumentDataByID(c, docID)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}
		response := client.GetDocumentDataResponse{
			ID:   doc.ID,
			Data: doc.Data.String(),
		}
		return c.JSON(http.StatusOK, response)
	}
}

// GetDicomDocumentDataByIDandTag godoc
// @Summary Get Dicom Document data by ID and Tag
// @Description Retrieves the data of a Dicom Document by its ID and a DICOM tag
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Param group body string true "DICOM Group"
// @Param element body string true "DICOM Element"
// @Success 200 {object} client.GetDicomDocumentDataByIDandTagResponse
// @Failure 400 {object} string "No file name provided"
// @Failure 400 {object} string "Could not parse group and element from request"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /get/{id}/tag [get]
func (h *Handler) GetDicomDocumentDataByIDandTag() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No file name provided")
		}
		var request client.GetDicomDocumentDataByIDandTagRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Could not parse group and element from request")
		}

		r, err := h.service.GetDicomDocumentDataByIDandTag(c, docID, request)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}
		response := client.GetDicomDocumentDataByIDandTagResponse{
			ID:    docID,
			Value: r,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// GetAll godoc
// @Summary Get all Dicom Documents
// @Description Retrieves all Dicom Documents
// @Tags get
// @Accept  json
// @Produce  json
// @Success 200 {array} client.GetDicomDocumentResponse
// @Failure 500 {object} string "Error getting dicom documents"
// @Router /get/all [get]
func (h *Handler) GetAll() echo.HandlerFunc {
	// add pagination
	return func(c echo.Context) error {
		docs, err := h.service.ListDocuments(c)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom documents")
		}
		var response []client.GetDicomDocumentResponse
		for _, doc := range docs {
			response = append(response, client.GetDicomDocumentResponse{
				ID:       doc.ID,
				Name:     doc.Name,
				Location: doc.Location,
			})
		}
		return c.JSON(http.StatusOK, response)
	}
}
