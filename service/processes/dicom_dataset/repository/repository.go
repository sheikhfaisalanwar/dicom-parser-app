package repository

import (
	"strings"

	dynamoClient "dicom-parser-app/common/dynamodb"
	"dicom-parser-app/service/processes/dicom_dataset/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/suyashkumar/dicom"
)

type DicomDataset struct {
	ID         string   `dynamodbav:"id" valid:"required" json:"id"`
	DocumentID string   `dynamodbav:"document_id" valid:"required" json:"document_id"`
	Data       string   `dynamodbav:"data" valid:"required" json:"data"`
	Tags       []string `dynamodbav:"tags" valid:"required" json:"tags"`
}

// ParseToDynamo converts a DicomDataset to a DicomDataset for DynamoDB
func ParseToDynamo(doc models.DicomDataset) DicomDataset {
	var tags []string
	for iter := doc.Data.FlatStatefulIterator(); iter.HasNext(); {
		tags = append(tags, iter.Next().Tag.String())
	}
	return DicomDataset{
		ID:         doc.ID,
		DocumentID: doc.DocumentID,
		Data:       doc.Data.String(),
		Tags:       tags,
	}
}

func (d *DicomDataset) ParseFromDynamo() models.DicomDataset {
	reader := strings.NewReader(d.Data)

	dataset, err := dicom.ParseUntilEOF(reader, nil)
	if err != nil {
		return models.DicomDataset{}
	}
	return models.DicomDataset{
		ID:         d.ID,
		DocumentID: d.DocumentID,
		Data:       dataset,
		Tags:       d.Tags,
	}
}

type IRepository interface {
	CreateDicomDataset(ctx echo.Context, doc models.DicomDataset) (models.DicomDataset, error)
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

func (r *Repository) CreateDicomDataset(ctx echo.Context, doc models.DicomDataset) (models.DicomDataset, error) {
	dicomDoc := ParseToDynamo(doc)

	item, err := attributevalue.MarshalMap(dicomDoc)
	if err != nil {
		return models.DicomDataset{}, err
	}

	_, err = r.client.PutItem(ctx.Request().Context(), &dynamodb.PutItemInput{
		TableName: aws.String(r.table.TableName),
		Item:      item,
	})
	if err != nil {
		return models.DicomDataset{}, err
	}

	return dicomDoc.ParseFromDynamo(), nil
}
