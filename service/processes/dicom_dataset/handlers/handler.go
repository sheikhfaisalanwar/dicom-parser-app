package handlers

import (
	"dicom-parser-app/service/processes/dicom_dataset/service"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	Name() string
	RegisterRoutes(router *echo.Echo)
}

type Handler struct {
	apiName string
	service service.IDicomDatasetService
}

func NewHandler(s service.IDicomDatasetService) *Handler {
	return &Handler{
		apiName: "dicom_document",
		service: s,
	}
}
