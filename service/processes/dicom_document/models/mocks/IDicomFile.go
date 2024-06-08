// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dicom "github.com/suyashkumar/dicom"
)

// IDicomFile is an autogenerated mock type for the IDicomFile type
type IDicomFile struct {
	mock.Mock
}

// GetData provides a mock function with given fields:
func (_m *IDicomFile) GetData() dicom.Dataset {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 dicom.Dataset
	if rf, ok := ret.Get(0).(func() dicom.Dataset); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dicom.Dataset)
	}

	return r0
}

// GetLocation provides a mock function with given fields:
func (_m *IDicomFile) GetLocation() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLocation")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *IDicomFile) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ParseData provides a mock function with given fields:
func (_m *IDicomFile) ParseData() (dicom.Dataset, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ParseData")
	}

	var r0 dicom.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func() (dicom.Dataset, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() dicom.Dataset); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dicom.Dataset)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetData provides a mock function with given fields: data
func (_m *IDicomFile) SetData(data dicom.Dataset) {
	_m.Called(data)
}

// SetLocation provides a mock function with given fields: location
func (_m *IDicomFile) SetLocation(location string) {
	_m.Called(location)
}

// SetName provides a mock function with given fields: name
func (_m *IDicomFile) SetName(name string) {
	_m.Called(name)
}

// NewIDicomFile creates a new instance of IDicomFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDicomFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDicomFile {
	mock := &IDicomFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
