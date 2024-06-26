package service

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"dicom-parser-app/client"
	"dicom-parser-app/service/processes/dicom_document/models"
	"dicom-parser-app/service/processes/dicom_document/repository/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const TestFileDirectory = "../../../../tests/assets/test_dicom_files/test1.dcm"

func TestNewDicomDocumentService(t *testing.T) {
	store := mocks.NewIStorage(t)
	repo := mocks.NewIRepository(t)
	s := NewDicomDocumentService(store, repo)
	assert.NotNil(t, s)
}

func TestCreateDicomDocument(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("test")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	tests := []struct {
		input     client.CreateDicomDocumentRequest
		setupMock func() (*mocks.IStorage, *mocks.IRepository)
		want      models.DicomFile
		wantErr   bool
	}{
		{
			input: client.CreateDicomDocumentRequest{
				File: multipart.FileHeader{
					Filename: "test.dcm",
				},
			},
			setupMock: func() (*mocks.IStorage, *mocks.IRepository) {
				store := mocks.NewIStorage(t)
				repo := mocks.NewIRepository(t)
				store.On("StoreDicomFile", mock.Anything, mock.Anything).Return("location", nil)
				repo.On("CreateDicomDocument", mock.Anything, mock.Anything).Return(models.DicomFile{
					Name: "test.dcm",
				}, nil)

				return store, repo
			},
			want: models.DicomFile{
				Name: "test.dcm",
			},
			wantErr: false,
		},
		{
			input: client.CreateDicomDocumentRequest{
				File: multipart.FileHeader{
					Filename: "test.dcm",
				},
			},
			setupMock: func() (*mocks.IStorage, *mocks.IRepository) {
				store := mocks.NewIStorage(t)
				repo := mocks.NewIRepository(t)
				store.On("StoreDicomFile", mock.Anything, mock.Anything).Return("", errors.New("error"))
				return store, repo
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		store, repo := test.setupMock()
		s := NewDicomDocumentService(store, repo)

		c := echo.New().NewContext(req, httptest.NewRecorder())
		doc, err := s.CreateDicomDocument(c, test.input)
		if test.wantErr {
			assert.Error(t, err)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, test.want.Name, doc.Name)
	}
}

//func TestGetDicomDocumentDataByIDandTag(t *testing.T) {
//	tests := []struct {
//		name      string
//		id        string
//		request   client.GetDicomDocumentDataByIDandTagRequest
//		setupMock func() *mocks.IRepository
//		want      string
//		wantErr   bool
//	}{
//		{
//			name: "Valid ID and Tag",
//			id:   "testID",
//			request: client.GetDicomDocumentDataByIDandTagRequest{
//				Group:   0002,
//				Element: 0000,
//			},
//			setupMock: func() *mocks.IRepository {
//				repo := new(mocks.IRepository)
//				repo.On("GetDicomDocumentByID", mock.Anything, "testID").Return(models.DicomFile{
//					Name:     "test1.dcm",
//					Location: TestFileDirectory,
//				}, nil)
//				return repo
//			},
//			want:    "testValue",
//			wantErr: false,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			repo := test.setupMock()
//			s := NewDicomDocumentService(nil, repo, nil)
//
//			got, err := s.GetDicomDocumentDataByIDandTag(nil, test.id, test.request)
//			if test.wantErr {
//				assert.Error(t, err)
//				return
//			}
//			assert.NoError(t, err)
//			assert.NotNil(t, got)
//		})
//	}
//}

//func TestCreateDicomDocument_Error(t *testing.T) {
//	store := new(MockStorage)
//	s := NewDicomDocumentService(store)
//	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("test")))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	c := echo.New().NewContext(req, httptest.NewRecorder())
//	params := client.CreateDicomDocumentRequest{
//		File: &client.File{
//			Filename: "test.dcm",
//			Data:     []byte("test"),
//		},
//	}
//	store.On("StoreDicomFile", params.File).Return("", errors.New("error"))
//	_, err := s.CreateDicomDocument(c, params)
//	assert.Error(t, err)
//}
//
//func TestParseDicomFile(t *testing.T) {
//	store := new(MockStorage)
//	s := NewDicomDocumentService(store).(*DicomDocumentService)
//	doc := &models.DicomFile{
//		Name:     "test.dcm",
//		Location: "location",
//		Data:     []byte("test"),
//	}
//	parsedDoc, err := s.ParseDicomFile(doc)
//	assert.NoError(t, err)
//	assert.Equal(t, []byte("test"), parsedDoc.Data)
//}
//
//func TestGetDicomDocumentData(t *testing.T) {
//	store := new(MockStorage)
//	s := NewDicomDocumentService(store).(*DicomDocumentService)
//	doc := &models.DicomFile{
//		Name:     "test.dcm",
//		Location: "location",
//		Data:     []byte("test"),
//	}
//	data, err := s.GetDicomDocumentData(doc)
//	assert.NoError(t, err)
//	assert.Equal(t, "test", data)
//}
