// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "home-service/internal/models"
)

// FlatGetter is an autogenerated mock type for the FlatGetter type
type FlatGetter struct {
	mock.Mock
}

// GetFlatByID provides a mock function with given fields: ctx, flatID
func (_m *FlatGetter) GetFlatByID(ctx context.Context, flatID int) (models.Flat, error) {
	ret := _m.Called(ctx, flatID)

	if len(ret) == 0 {
		panic("no return value specified for GetFlatByID")
	}

	var r0 models.Flat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (models.Flat, error)); ok {
		return rf(ctx, flatID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) models.Flat); ok {
		r0 = rf(ctx, flatID)
	} else {
		r0 = ret.Get(0).(models.Flat)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, flatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFlatGetter creates a new instance of FlatGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlatGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlatGetter {
	mock := &FlatGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
