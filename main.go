package main

import (
	"fmt"

	"dicom-parser-app/service/processes/dicom_document"
)

func main() {
	fmt.Println("Starting dicom-parser service")
	dicomdocument.InitServer()
}
