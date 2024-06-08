// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	client "dicom-parser-app/client"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"

	models "dicom-parser-app/service/processes/dicom_document/models"
)

// IDicomDocumentService is an autogenerated mock type for the IDicomDocumentService type
type IDicomDocumentService struct {
	mock.Mock
}

// CreateDicomDocument provides a mock function with given fields: ctx, params
func (_m *IDicomDocumentService) CreateDicomDocument(ctx echo.Context, params client.CreateDicomDocumentRequest) (*models.DicomFile, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for CreateDicomDocument")
	}

	var r0 *models.DicomFile
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, client.CreateDicomDocumentRequest) (*models.DicomFile, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, client.CreateDicomDocumentRequest) *models.DicomFile); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DicomFile)
		}
	}

	if rf, ok := ret.Get(1).(func(echo.Context, client.CreateDicomDocumentRequest) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDicomDocument provides a mock function with given fields:
func (_m *IDicomDocumentService) GetDicomDocument() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDicomDocument")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewIDicomDocumentService creates a new instance of IDicomDocumentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDicomDocumentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDicomDocumentService {
	mock := &IDicomDocumentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}