package handlers

import (
	"dicom-parser-app/common"
	"dicom-parser-app/service/processes/dicom_document/service"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	Name() string
	RegisterRoutes(router *echo.Echo)
}

type Handler struct {
	apiName string
	service service.IDicomDocumentService
	logger  common.MyLogger
}

func NewHandler(s service.IDicomDocumentService) *Handler {
	return &Handler{
		apiName: "dicom_document",
		service: s,
		logger:  common.NewLogger(),
	}
}
