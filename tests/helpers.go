package tests

import "github.com/suyashkumar/dicom"

func GetParsedDicomDocument(filepath string) dicom.Dataset {
	dataset, err := dicom.ParseFile(filepath, nil)
	if err != nil {
		panic(err)
	}
	return dataset
}
