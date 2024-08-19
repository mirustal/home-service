// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	authgrpc "auth-service/pkg/pb"
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AuthClient is an autogenerated mock type for the AuthClient type
type AuthClient struct {
	mock.Mock
}

// DummyLogin provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) DummyLogin(ctx context.Context, in *authgrpc.DummyLoginRequest, opts ...grpc.CallOption) (*authgrpc.DummyLoginResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DummyLogin")
	}

	var r0 *authgrpc.DummyLoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.DummyLoginRequest, ...grpc.CallOption) (*authgrpc.DummyLoginResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.DummyLoginRequest, ...grpc.CallOption) *authgrpc.DummyLoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authgrpc.DummyLoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *authgrpc.DummyLoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Login(ctx context.Context, in *authgrpc.LoginRequest, opts ...grpc.CallOption) (*authgrpc.LoginResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *authgrpc.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.LoginRequest, ...grpc.CallOption) (*authgrpc.LoginResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.LoginRequest, ...grpc.CallOption) *authgrpc.LoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authgrpc.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *authgrpc.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSession provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) RefreshSession(ctx context.Context, in *authgrpc.RefreshRequest, opts ...grpc.CallOption) (*authgrpc.RefreshResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RefreshSession")
	}

	var r0 *authgrpc.RefreshResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.RefreshRequest, ...grpc.CallOption) (*authgrpc.RefreshResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.RefreshRequest, ...grpc.CallOption) *authgrpc.RefreshResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authgrpc.RefreshResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *authgrpc.RefreshRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) Register(ctx context.Context, in *authgrpc.RegisterRequest, opts ...grpc.CallOption) (*authgrpc.RegisterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *authgrpc.RegisterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.RegisterRequest, ...grpc.CallOption) (*authgrpc.RegisterResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.RegisterRequest, ...grpc.CallOption) *authgrpc.RegisterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authgrpc.RegisterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *authgrpc.RegisterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateSession provides a mock function with given fields: ctx, in, opts
func (_m *AuthClient) ValidateSession(ctx context.Context, in *authgrpc.ValidateRequest, opts ...grpc.CallOption) (*authgrpc.ValidateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSession")
	}

	var r0 *authgrpc.ValidateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.ValidateRequest, ...grpc.CallOption) (*authgrpc.ValidateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *authgrpc.ValidateRequest, ...grpc.CallOption) *authgrpc.ValidateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authgrpc.ValidateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *authgrpc.ValidateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthClient creates a new instance of AuthClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthClient {
	mock := &AuthClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
