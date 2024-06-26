// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// IStorage is an autogenerated mock type for the IStorage type
type IStorage struct {
	mock.Mock
}

// StoreDicomFile provides a mock function with given fields: id, file
func (_m *IStorage) StoreDicomFile(id string, file multipart.FileHeader) (string, error) {
	ret := _m.Called(id, file)

	if len(ret) == 0 {
		panic("no return value specified for StoreDicomFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, multipart.FileHeader) (string, error)); ok {
		return rf(id, file)
	}
	if rf, ok := ret.Get(0).(func(string, multipart.FileHeader) string); ok {
		r0 = rf(id, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, multipart.FileHeader) error); ok {
		r1 = rf(id, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIStorage creates a new instance of IStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStorage {
	mock := &IStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
