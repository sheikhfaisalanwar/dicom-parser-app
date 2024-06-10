package handlers

import (
	"errors"
	"net/http"

	"dicom-parser-app/client"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Setup() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	}
}

func (h *Handler) GetDataByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No file name provided")
		}
		doc, err := h.service.GetDicomTagsByDocumentID(c, docID)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}
		response := client.GetDocumentTagsResponse{
			DocumentID: doc.DocumentID,
			Tags:       doc.Tags,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) GetDataByIDAndTag() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No document id provided")
		}
		var req client.GetElementByTagRequest
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		c.Logger().Infof("Getting data for document %s and tag %s", docID, req.Tag)

		element, err := h.service.GetDicomElementByTagAndDocumentID(c, docID, req.Tag)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}
		if element == nil {
			return c.JSON(http.StatusNotFound, "Data by tag not found")
		}

		return c.JSON(http.StatusOK, element)
	}
}

func (h *Handler) GetDataByIDAndTagName() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No document id provided")
		}
		var req client.GetElementByTagNameRequest
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		c.Logger().Infof("Getting data for document %s and tag name %s", docID, req.TagName)

		element, err := h.service.GetDicomElementByTagNameAndDocumentID(c, docID, req.TagName)
		if err != nil {
			if errors.Is(err, client.ErrorFailedToFindDataByTag) {
				return c.JSON(http.StatusNotFound, "Data by tag name not found")
			}
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Error getting dicom document")
		}

		return c.JSON(http.StatusOK, element)
	}
}
