// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeTraceCollectorServer is an autogenerated mock type for the UnsafeTraceCollectorServer type
type UnsafeTraceCollectorServer struct {
	mock.Mock
}

// mustEmbedUnimplementedTraceCollectorServer provides a mock function with given fields:
func (_m *UnsafeTraceCollectorServer) mustEmbedUnimplementedTraceCollectorServer() {
	_m.Called()
}

// NewUnsafeTraceCollectorServer creates a new instance of UnsafeTraceCollectorServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeTraceCollectorServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeTraceCollectorServer {
	mock := &UnsafeTraceCollectorServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
