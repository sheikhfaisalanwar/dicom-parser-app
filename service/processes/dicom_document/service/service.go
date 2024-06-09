package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/suyashkumar/dicom/pkg/tag"

	"dicom-parser-app/client"
	"dicom-parser-app/service/processes/dicom_document/models"
	"dicom-parser-app/service/processes/dicom_document/repository"
)

type IDicomDocumentService interface {
	// GetDicomDocumentByID returns a dicom document
	GetDicomDocumentByID(ctx echo.Context, id string) (models.DicomFile, error)

	// GetDocumentDataByID returns a dicom documents tags by the id
	GetDocumentDataByID(ctx echo.Context, id string) (models.DicomFile, error)

	// GetDicomDocumentDataByIDandTag returns a dicom documents data filtered by the tag
	GetDicomDocumentDataByIDandTag(ctx echo.Context, id string, request client.GetDicomDocumentDataByIDandTagRequest) (string, error)

	// ListDocuments returns a list of dicom documents
	ListDocuments(ctx echo.Context) ([]models.DicomFile, error)

	// CreateDicomDocument stores a dicom document and creates a record in the database
	CreateDicomDocument(ctx echo.Context, params client.CreateDicomDocumentRequest) (*models.DicomFile, error)
}

type DicomDocumentService struct {
	store      repository.IStorage
	repository repository.IRepository
}

func NewDicomDocumentService(store repository.IStorage, repository repository.IRepository) IDicomDocumentService {
	return &DicomDocumentService{
		store:      store,
		repository: repository,
	}
}

func (d *DicomDocumentService) GetDicomDocumentByID(ctx echo.Context, id string) (models.DicomFile, error) {
	return d.repository.GetDicomDocumentByID(ctx, id)
}

func (d *DicomDocumentService) GetDocumentDataByID(ctx echo.Context, id string) (models.DicomFile, error) {
	doc, err := d.repository.GetDicomDocumentByID(ctx, id)
	if err != nil {
		return models.DicomFile{}, err
	}
	if doc.Location == "" {
		return models.DicomFile{}, errors.New("document location is not set")
	}
	parsedDoc, err := d.parseDicomFile(doc)
	if err != nil {
		return models.DicomFile{}, err
	}
	return parsedDoc, nil
}

func (d *DicomDocumentService) GetDicomDocumentDataByIDandTag(ctx echo.Context, id string, request client.GetDicomDocumentDataByIDandTagRequest) (string, error) {
	doc, err := d.repository.GetDicomDocumentByID(ctx, id)
	if err != nil {
		return "", err
	}
	if doc.Location == "" {
		return "", errors.New("document location is not set")
	}
	parsedDoc, err := d.parseDicomFile(doc)
	if err != nil {
		return "", err
	}

	tagToFind := tag.Tag{
		Group:   request.Group,
		Element: request.Element,
	}

	// TODO: Should prevent parser package implementation details from leaking
	tagValue, err := parsedDoc.FindByTag(tagToFind)
	if err != nil {
		return "", err
	}
	return tagValue, nil

}

func (d *DicomDocumentService) ListDocuments(ctx echo.Context) ([]models.DicomFile, error) {
	return d.repository.ListDocuments(ctx)
}

func (d *DicomDocumentService) CreateDicomDocument(ctx echo.Context, params client.CreateDicomDocumentRequest) (*models.DicomFile, error) {

	dicomDoc := &models.DicomFile{
		ID:   uuid.Must(uuid.NewRandom()).String(),
		Name: params.File.Filename,
	}

	// TODO: make async with rabbitmq
	location, err := d.store.StoreDicomFile(dicomDoc.ID, params.File)
	if err != nil {
		ctx.Logger().Error(err)
		return dicomDoc, err
	}

	dicomDoc.Location = location

	storedDoc, err := d.repository.CreateDicomDocument(ctx, *dicomDoc)
	if err != nil {
		ctx.Logger().Error(err)
		return dicomDoc, err
	}

	return &storedDoc, nil
}

func (d *DicomDocumentService) parseDicomFile(doc models.DicomFile) (models.DicomFile, error) {
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
