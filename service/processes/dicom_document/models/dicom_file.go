package models

import (
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type IDicomFile interface {
	GetName() string
	SetName(name string)
	GetData() dicom.Dataset
	SetData(data dicom.Dataset)
	ParseData() (dicom.Dataset, error)
	FindByTag(tag tag.Tag) (string, error)
	GetLocation() string
	SetLocation(location string)
}

type DicomFile struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Data     dicom.Dataset `json:"data"`
	Location string        `json:"location"`
}

func (d *DicomFile) GetName() string {
	return d.Name
}

func (d *DicomFile) SetName(name string) {
	d.Name = name
}

func (d *DicomFile) GetData() dicom.Dataset {
	return d.Data
}

func (d *DicomFile) SetData(data dicom.Dataset) {
	d.Data = data
}

func (d *DicomFile) ParseData() (dicom.Dataset, error) {
	data, err := dicom.ParseFile(d.Location, nil)

	if err != nil {
		return d.Data, err
	}
	return data, nil
}

func (d *DicomFile) FindByTag(tag tag.Tag) (string, error) {
	value, err := d.Data.FindElementByTagNested(tag)
	if err != nil {
		return "", err
	}
	return value.String(), nil
}

func (d *DicomFile) GetLocation() string {
	return d.Location
}

func (d *DicomFile) SetLocation(location string) {
	d.Location = location
}
