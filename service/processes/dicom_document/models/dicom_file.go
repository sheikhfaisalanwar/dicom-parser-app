package models

import (
	"github.com/suyashkumar/dicom"
)

type IDicomFile interface {
	GetName() string
	SetName(name string)
	GetData() dicom.Dataset
	SetData(data dicom.Dataset)
	ParseData() (dicom.Dataset, error)
	GetLocation() string
	SetLocation(location string)
}

type DicomFile struct {
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

func (d *DicomFile) GetLocation() string {
	return d.Location
}

func (d *DicomFile) SetLocation(location string) {
	d.Location = location
}
