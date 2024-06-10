package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	dicomDocModels "dicom-parser-app/service/processes/dicom_document/models"
	"dicom-parser-app/service/processes/dicom_document/service/mocks"
	"dicom-parser-app/tests"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/suyashkumar/dicom"
)

const TestFilePath = "../../../../tests/assets/test_dicom_files/test1.dcm"

func TestDicomDatasetService_GetDicomElementByTagAndDocumentID(t *testing.T) {
	sampleDataset := tests.GetParsedDicomDocument(TestFilePath)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("test")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echo.New().NewContext(req, httptest.NewRecorder())

	tests := []struct {
		name    string
		mocks   func() *mocks.IDicomDocumentService
		want    *dicom.Element
		wantErr bool
	}{
		{
			name: "Successfully get tag",
			mocks: func() *mocks.IDicomDocumentService {
				svc := new(mocks.IDicomDocumentService)
				svc.On("GetDocumentDataByID", mock.Anything, mock.Anything).Return(dicomDocModels.DicomFile{
					Data: sampleDataset,
				}, nil)
				return svc
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mocks()
			s := NewDicomDatasetService(nil, m)
			tags, err := s.GetDicomElementByTagAndDocumentID(c, "test", "(0002,0000)")
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, tags)
		})
	}
}
