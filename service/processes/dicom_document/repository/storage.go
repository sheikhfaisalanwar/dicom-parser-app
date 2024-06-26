package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Replace with Azure Blob Storage SDK client
var DicomFileStorageUri = os.Getenv("DICOM_FILE_STORAGE_URI")

// go:generate mockgen -destination=mocks/mock_storage.go -package=mocks dicom-parser-app/service/processes/dicom_document/repository IStorage
type IStorage interface {
	// StoreDicomFile stores a dicom file
	StoreDicomFile(id string, file multipart.FileHeader) (string, error)
}

type DicomFileStorage struct {
	Location string
}

func NewDicomFileStorage() IStorage {
	return &DicomFileStorage{
		Location: DicomFileStorageUri,
	}
}

func (d *DicomFileStorage) StoreDicomFile(id string, file multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	dir := d.Location + "/" + id

	// Destination created with id as part of the path
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(dir, filepath.Base(file.Filename))
	dst, err := os.Create(filePath)
	if err != nil {
		return filePath, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return filePath, err
	}
	return filePath, nil
}
