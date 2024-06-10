package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(
	router *echo.Echo,
	m ...echo.MiddlewareFunc,
) {
	group := router.Group("/dicom_document")
	group.Use(m...)
	// GET
	group.GET("/:id", h.GetByID())
	group.GET("/:id/data", h.GetDataByID())
	group.GET("/", h.GetAll())
	group.GET("/:id/tag", h.GetDicomDocumentDataByIDandTag())

	// POST
	group.POST("/create", h.Create())
}
