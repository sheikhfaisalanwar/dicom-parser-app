package handlers

import (
	"net/http"

	"dicom-parser-app/client"

	"github.com/labstack/echo/v4"
)

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
			Data: doc.Data,
		}
		return c.JSON(http.StatusOK, response)
	}
}

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
