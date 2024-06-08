package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(
	router *echo.Echo,
) {
	group := router.Group("/dicom_document")

	group.POST("/create", h.Create())
}
