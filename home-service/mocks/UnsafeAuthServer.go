// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeAuthServer is an autogenerated mock type for the UnsafeAuthServer type
type UnsafeAuthServer struct {
	mock.Mock
}

// mustEmbedUnimplementedAuthServer provides a mock function with given fields:
func (_m *UnsafeAuthServer) mustEmbedUnimplementedAuthServer() {
	_m.Called()
}

// NewUnsafeAuthServer creates a new instance of UnsafeAuthServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeAuthServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeAuthServer {
	mock := &UnsafeAuthServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
