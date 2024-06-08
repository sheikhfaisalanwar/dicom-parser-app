package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"dicom-parser-app/service/processes/dicom_document/models"
	"dicom-parser-app/service/processes/dicom_document/service/mocks"
)

const TestFileDirectory = "../../../../tests/assets/test_dicom_files/test1.dcm"

func TestCreate(t *testing.T) {

	tests := []struct {
		name       string
		input      func() *http.Request
		setupMock  func() *mocks.IDicomDocumentService
		wantErr    bool
		wantStatus int
	}{
		{
			name: "Successfully request to create dicom document",
			input: func() *http.Request {
				body, ct := createMultipartFormData(t, "file", TestFileDirectory)

				req := httptest.NewRequest(http.MethodPost, "/", &body)
				req.Header.Add("Content-Type", ct.FormDataContentType())
				return req
			},
			setupMock: func() *mocks.IDicomDocumentService {
				mockService := mocks.NewIDicomDocumentService(t)
				mockService.On("CreateDicomDocument", mock.Anything, mock.Anything).Return(
					&models.DicomFile{
						Name:     "test1.dcm",
						Location: "location",
					}, nil)
				return mockService
			},
			wantErr:    false,
			wantStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockService := test.setupMock()
			h := NewHandler(mockService)
			e := echo.New()
			req := test.input()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.Create()(c)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.wantStatus, rec.Code)
		})

	}
}

func createMultipartFormData(t *testing.T, fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	w.Close()
	return b, w
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}
