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

// GetDocumentTagsByID godoc
// @Summary Get All Dicom Document tags by ID
// @Description Retrieves the tags of a Dicom Document by its ID
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Success 200 {object} client.GetDocumentTagsResponse
// @Failure 400 {object} string "No document ID provided"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /dicom_data/{id}/tags [get]
func (h *Handler) GetDocumentTagsByID() echo.HandlerFunc {
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

// GetDataByIDAndTag godoc
// @Summary Get Dicom Document data by ID and Tag
// @Description Retrieves the data of a Dicom Header Attribute by its ID and a DICOM tag
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Param tag body string true "DICOM Tag"
// @Success 200 {object} string
// @Failure 400 {object} string "No file name provided"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /dicom_data/{id}/tag [get]
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

// GetDataByIDAndTagName godoc
// @Summary Get Dicom Document data by ID and Tag Name
// @Description Retrieves the data of a Dicom Header Attribute by its ID and a DICOM tag name
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Param tagName body string true "DICOM Tag Name"
// @Success 200 {object} string
// @Failure 400 {object} string "No document ID provided"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Error getting dicom document"
// @Router /dicom_data/{id}/tag-name [get]
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

// GetImageByDocumentID godoc
// @Summary Get Dicom Document image by ID
// @Description Retrieves the image of a Dicom Document by its ID
// @Tags get
// @Accept  json
// @Produce  json
// @Param id path string true "Dicom Document ID"
// @Success 200 {object} string
// @Failure 400 {object} string "No document ID provided
// @Failure 500 {object} string "Failed to generate image from document"
// @Router /dicom_data/{id}/image [get]
func (h *Handler) GetImageByDocumentID() echo.HandlerFunc {
	return func(c echo.Context) error {
		docID := c.Param("id")
		if docID == "" {
			return c.JSON(http.StatusBadRequest, "No document id provided")
		}

		c.Logger().Infof("Getting image for document %s", docID)

		pngFile, err := h.service.GetDicomImageByDocumentID(c, docID)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, "Failed to generate image from document")
		}
		if pngFile == nil {
			return c.JSON(http.StatusInternalServerError, "Failed to generate image from document")
		}

		return c.File(pngFile.Name())
	}
}
