// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"
	pb "home-service/pkg/pb"

	mock "github.com/stretchr/testify/mock"
)

// AuthServer is an autogenerated mock type for the AuthServer type
type AuthServer struct {
	mock.Mock
}

// DummyLogin provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) DummyLogin(_a0 context.Context, _a1 *pb.DummyLoginRequest) (*pb.DummyLoginResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DummyLogin")
	}

	var r0 *pb.DummyLoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DummyLoginRequest) (*pb.DummyLoginResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DummyLoginRequest) *pb.DummyLoginResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DummyLoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DummyLoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) Login(_a0 context.Context, _a1 *pb.LoginRequest) (*pb.LoginResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *pb.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.LoginRequest) *pb.LoginResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.LoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSession provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) RefreshSession(_a0 context.Context, _a1 *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RefreshSession")
	}

	var r0 *pb.RefreshResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RefreshRequest) (*pb.RefreshResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RefreshRequest) *pb.RefreshResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.RefreshResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.RefreshRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) Register(_a0 context.Context, _a1 *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *pb.RegisterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RegisterRequest) *pb.RegisterResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.RegisterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.RegisterRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateSession provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) ValidateSession(_a0 context.Context, _a1 *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSession")
	}

	var r0 *pb.ValidateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ValidateRequest) (*pb.ValidateResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ValidateRequest) *pb.ValidateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ValidateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ValidateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedAuthServer provides a mock function with given fields:
func (_m *AuthServer) mustEmbedUnimplementedAuthServer() {
	_m.Called()
}

// NewAuthServer creates a new instance of AuthServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthServer {
	mock := &AuthServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
