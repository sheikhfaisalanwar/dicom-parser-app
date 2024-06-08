package service

import (
	"github.com/labstack/echo/v4"

	"dicom-parser-app/client"
	"dicom-parser-app/service/processes/dicom_document/models"
	"dicom-parser-app/service/processes/dicom_document/repository"
)

type IDicomDocumentService interface {
	// GetDicomDocument returns a dicom document
	GetDicomDocument() string

	// UploadDicomDocument uploads a dicom document
	CreateDicomDocument(ctx echo.Context, params client.CreateDicomDocumentRequest) (*models.DicomFile, error)
}

type DicomDocumentService struct {
	store repository.IStorage
}

func NewDicomDocumentService(store repository.IStorage) IDicomDocumentService {
	return &DicomDocumentService{
		store: store,
	}
}

func (d *DicomDocumentService) GetDicomDocument() string {
	return "Dicom Document"
}

func (d *DicomDocumentService) CreateDicomDocument(ctx echo.Context, params client.CreateDicomDocumentRequest) (*models.DicomFile, error) {

	dicomDoc := &models.DicomFile{
		Name: params.File.Filename,
	}
	location, err := d.store.StoreDicomFile(params.File)
	if err != nil {
		return dicomDoc, err
	}

	dicomDoc.Location = location

	//dicomDoc, err = d.parseDicomFile(dicomDoc)

	return dicomDoc, nil
}

func (d *DicomDocumentService) parseDicomFile(doc *models.DicomFile) (*models.DicomFile, error) {
	data, err := doc.ParseData()
	if err != nil {
		return doc, err
	}
	doc.Data = data

	return doc, nil
}

func (d *DicomDocumentService) GetDicomDocumentData(doc *models.DicomFile) (string, error) {
	return doc.Data.String(), nil
}
