package models

import (
	"github.com/suyashkumar/dicom"
)

type IDicomDatasetInterface struct {
}

type DicomDataset struct {
	ID         string        `json:"id"`
	DocumentID string        `json:"document_id"`
	Data       dicom.Dataset `json:"data"`
	Tags       []string      `json:"tags"`
}
