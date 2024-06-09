package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/labstack/echo/v4"

	dynamoClient "dicom-parser-app/common/dynamodb"
	"dicom-parser-app/service/processes/dicom_document/models"
)

type DicomDocument struct {
	ID       string `dynamodbav:"id" valid:"required" json:"id"`
	Name     string `dynamodbav:"name" valid:"required" json:"name"`
	Location string `dynamodbav:"location" valid:"required" json:"location"`
}

type IRepository interface {
	CreateDicomDocument(ctx echo.Context, doc models.DicomFile) (models.DicomFile, error)
	GetDicomDocumentByID(ctx echo.Context, id string) (models.DicomFile, error)
	ListDocuments(ctx echo.Context) ([]models.DicomFile, error)
	UpdateDicomDocument() string
}

type Repository struct {
	client dynamoClient.DynamoClient
	table  dynamoClient.DynamoTableParams
}

func NewRepository(client dynamoClient.DynamoClient, table dynamoClient.DynamoTableParams) *Repository {
	return &Repository{
		client: client,
		table:  table,
	}
}

func (r *Repository) CreateDicomDocument(ctx echo.Context, doc models.DicomFile) (models.DicomFile, error) {
	dicomDoc := ParseToDynamo(doc)

	item, err := attributevalue.MarshalMap(dicomDoc)
	if err != nil {
		return models.DicomFile{}, err
	}

	_, err = r.client.PutItem(ctx.Request().Context(), &dynamodb.PutItemInput{
		TableName: aws.String(r.table.TableName),
		Item:      item,
	})
	if err != nil {
		return models.DicomFile{}, err
	}

	return dicomDoc.ParseFromDynamo(), nil
}

func (r *Repository) GetDicomDocumentByID(ctx echo.Context, id string) (models.DicomFile, error) {
	doc := DicomDocument{ID: id}
	response, err := r.client.GetItem(ctx.Request().Context(), &dynamodb.GetItemInput{
		Key: doc.GetKey(), TableName: aws.String(r.table.TableName),
	})
	if err != nil {
		ctx.Logger().Error("Couldn't get info about %v: %v\n", id, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &doc)
		if err != nil {
			ctx.Logger().Error("Couldn't unmarshal response: %v\n", err)
		}
	}

	return doc.ParseFromDynamo(), nil
}

func (r *Repository) ListDocuments(ctx echo.Context) ([]models.DicomFile, error) {

	response, err := r.client.Scan(ctx.Request().Context(), &dynamodb.ScanInput{
		TableName: aws.String(r.table.TableName),
	})
	if err != nil {
		ctx.Logger().Error("Couldn't get list of documents: %v\n", err)
		return nil, err
	}

	var docs []DicomDocument
	err = attributevalue.UnmarshalListOfMaps(response.Items, &docs)
	if err != nil {
		ctx.Logger().Error("Couldn't unmarshal response: %v\n", err)
		return nil, err
	}

	var dicomDocs []models.DicomFile
	for _, doc := range docs {
		dicomDocs = append(dicomDocs, doc.ParseFromDynamo())
	}

	return dicomDocs, nil

}

func (r *Repository) UpdateDicomDocument() string {
	return "UpdateDicomDocument"
}

func (r *Repository) DeleteDicomDocument() string {
	return "DeleteDicomDocument"
}

func ParseToDynamo(d models.DicomFile) DicomDocument {
	return DicomDocument{
		ID:       d.ID,
		Name:     d.Name,
		Location: d.Location,
	}
}

func (d *DicomDocument) ParseFromDynamo() models.DicomFile {
	return models.DicomFile{
		ID:       d.ID,
		Name:     d.Name,
		Location: d.Location,
	}
}

func (d *DicomDocument) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(d.ID)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id}
}
