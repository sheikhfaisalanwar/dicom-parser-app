package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(
	router *echo.Echo,
	m ...echo.MiddlewareFunc,
) {
	group := router.Group("/dicom_data")
	group.Use(m...)

	// GET
	group.GET("/:id/tags", h.GetDocumentTagsByID())
	group.GET("/:id/tag", h.GetDataByIDAndTag())
	group.GET("/:id/tag-name", h.GetDataByIDAndTagName())
	group.GET("/:id/image", h.GetImageByDocumentID())

	group.GET("/", h.Setup())

}
