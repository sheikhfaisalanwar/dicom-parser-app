package service

import (
	"fmt"
	"strings"

	"dicom-parser-app/client"
	"dicom-parser-app/service/processes/dicom_dataset/models"
	"dicom-parser-app/service/processes/dicom_dataset/repository"
	"dicom-parser-app/service/processes/dicom_document/service"

	"github.com/labstack/echo/v4"
	"github.com/suyashkumar/dicom"
)

type IDicomDatasetService interface {
	// CreateDicomDataset Creates a record in the database
	CreateDicomDataset(ctx echo.Context, doc models.DicomDataset) (models.DicomDataset, error)

	// GetDicomTagsByDocumentID returns all dicom dataset tags for a document
	GetDicomTagsByDocumentID(ctx echo.Context, id string) (models.DicomDataset, error)

	// GetDicomElementByTagAndDocumentID find and returns a dicom element in a document by tag
	GetDicomElementByTagAndDocumentID(ctx echo.Context, id string, tag string) (*dicom.Element, error)

	// GetDicomElementByTagNameAndDocumentID find and returns a dicom element in a document by tag
	GetDicomElementByTagNameAndDocumentID(ctx echo.Context, id string, tagName string) (*dicom.Element, error)

	// GetDicomImageByDocumentID returns a dicom image from a document
	GetDicomImageByDocumentID(ctx echo.Context, id string) (models.DicomDataset, error)
}

type DicomDatasetService struct {
	repository      repository.IRepository
	documentService service.IDicomDocumentService
}

func NewDicomDatasetService(repository repository.IRepository, documentService service.IDicomDocumentService) IDicomDatasetService {
	return &DicomDatasetService{
		repository:      repository,
		documentService: documentService,
	}
}

// CreateDicomDataset Creates a record in the database
// This is unused for now since DICOM data is large and we can't store it in dynamoDB
func (d *DicomDatasetService) CreateDicomDataset(ctx echo.Context, doc models.DicomDataset) (models.DicomDataset, error) {
	return d.repository.CreateDicomDataset(ctx, doc)
}

// GetDicomTagsByDocumentID returns a dicom dataset tags as string representations
func (d *DicomDatasetService) GetDicomTagsByDocumentID(ctx echo.Context, id string) (models.DicomDataset, error) {
	dataset := models.DicomDataset{DocumentID: id}
	doc, err := d.documentService.GetDocumentDataByID(ctx, id)
	if err != nil {
		ctx.Logger().Error("Could not get document data by id")
		return dataset, err
	}

	// TODO: implement blob storage for DICOM data intead of parsing it everytime
	var tags []string
	for iter := doc.Data.FlatStatefulIterator(); iter.HasNext(); {
		tags = append(tags, iter.Next().Tag.String())
	}
	dataset.Tags = tags
	return dataset, nil
}

// GetDicomElementByTagAndDocumentID find and returns a dicom element in a document by tag
func (d *DicomDatasetService) GetDicomElementByTagAndDocumentID(ctx echo.Context, id string, tag string) (*dicom.Element, error) {
	doc, err := d.documentService.GetDocumentDataByID(ctx, id)
	if err != nil {
		ctx.Logger().Error("Could not get document data by id")
		return &dicom.Element{}, err
	}

	for iter := doc.Data.FlatStatefulIterator(); iter.HasNext(); {
		element := iter.Next()
		if element.Tag.String() == tag {
			ctx.Logger().Infof("Element found with tag: %s", tag)
			return element, nil
		}
	}
	return &dicom.Element{}, client.ErrorFailedToFindDataByTag
}

// GetDicomElementByTagNameAndDocumentID find and returns a dicom element in a document by tag
func (d *DicomDatasetService) GetDicomElementByTagNameAndDocumentID(ctx echo.Context, id string, tagName string) (*dicom.Element, error) {
	doc, err := d.documentService.GetDocumentDataByID(ctx, id)
	if err != nil {
		ctx.Logger().Error("Could not get document data by id")
		return &dicom.Element{}, err
	}

	for iter := doc.Data.FlatStatefulIterator(); iter.HasNext(); {
		element := iter.Next()
		if strings.Contains(element.String(), tagName) {
			fmt.Println(element)
			ctx.Logger().Infof("Element found with tag: %s", tagName)
			return element, nil
		}
		fmt.Println("Element not found")
		ctx.Logger().Error("Element not found")
	}
	return &dicom.Element{}, nil
}

// GetDicomImageByDocumentID returns a dicom image from a document
func (d *DicomDatasetService) GetDicomImageByDocumentID(ctx echo.Context, id string) (models.DicomDataset, error) {
	dataset := models.DicomDataset{DocumentID: id}
	_, err := d.documentService.GetDocumentDataByID(ctx, id)
	if err != nil {
		ctx.Logger().Error("Could not get document data by id")
		return dataset, err
	}
	return dataset, nil
}
