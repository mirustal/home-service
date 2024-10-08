// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	models "auth-service/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserGetter is an autogenerated mock type for the UserGetter type
type UserGetter struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, userID
func (_m *UserGetter) GetUser(ctx context.Context, userID string) (models.User, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserGetter creates a new instance of UserGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserGetter {
	mock := &UserGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
