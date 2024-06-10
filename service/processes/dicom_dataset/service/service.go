package service

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"dicom-parser-app/client"
	"dicom-parser-app/service/processes/dicom_dataset/models"
	"dicom-parser-app/service/processes/dicom_dataset/repository"
	"dicom-parser-app/service/processes/dicom_document/service"

	"github.com/labstack/echo/v4"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
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
	GetDicomImageByDocumentID(ctx echo.Context, id string) (*os.File, error)
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
			ctx.Logger().Infof("Element found with tag: %s", tagName)
			return element, nil
		}
	}
	return &dicom.Element{}, nil
}

// GetDicomImageByDocumentID returns a dicom image from a document
func (d *DicomDatasetService) GetDicomImageByDocumentID(ctx echo.Context, id string) (*os.File, error) {
	//data := models.DicomDataset{DocumentID: id}
	var pngFile *os.File
	dataset, err := d.documentService.GetDocumentDataByID(ctx, id)
	if err != nil {
		ctx.Logger().Error("Could not get document data by id")
		return pngFile, err
	}

	// Find pixel data elements
	var allPixelData []*dicom.Element
	for iter := dataset.Data.FlatStatefulIterator(); iter.HasNext(); {
		e := iter.Next()
		if e.Tag == tag.PixelData {
			allPixelData = append(allPixelData, e)
			ctx.Logger().Infof("Element found with tag: %s", tag.PixelData)
		}
	}
	fmt.Println(allPixelData)

	pixelDataElement, err := dataset.Data.FindElementByTagNested(tag.PixelData)
	if err != nil {
		ctx.Logger().Error("Error finding pixel data in document")
		return pngFile, err
	}
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)
	fmt.Println(pixelDataInfo.Frames)
	fmt.Println(len(pixelDataInfo.Frames))
	for i, fr := range pixelDataInfo.Frames {
		img, _ := fr.GetImage()
		pngFile, err = os.Create(fmt.Sprintf("image_%d.png", i))
		if err != nil {
			fmt.Println(err)
			ctx.Logger().Error("Error creating image file")
			return pngFile, err
		}
		err := png.Encode(pngFile, img)
		if err != nil {
			fmt.Println(err)
			ctx.Logger().Error("Error encoding image to png")
			return pngFile, err
		}
		err = pngFile.Close()
		if err != nil {
			fmt.Println(err)
			ctx.Logger().Error("Error closing image file")
			return pngFile, err
		}
	}

	return pngFile, nil
}
