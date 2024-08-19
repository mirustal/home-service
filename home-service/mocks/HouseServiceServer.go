// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"
	pb "home-service/pkg/pb"

	mock "github.com/stretchr/testify/mock"
)

// HouseServiceServer is an autogenerated mock type for the HouseServiceServer type
type HouseServiceServer struct {
	mock.Mock
}

// CreateFlat provides a mock function with given fields: _a0, _a1
func (_m *HouseServiceServer) CreateFlat(_a0 context.Context, _a1 *pb.CreateFlatRequest) (*pb.CreateFlatResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateFlat")
	}

	var r0 *pb.CreateFlatResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateFlatRequest) (*pb.CreateFlatResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateFlatRequest) *pb.CreateFlatResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateFlatResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateFlatRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseServiceServer) CreateHouse(_a0 context.Context, _a1 *pb.CreateHouseRequest) (*pb.CreateHouseResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateHouse")
	}

	var r0 *pb.CreateHouseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateHouseRequest) (*pb.CreateHouseResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateHouseRequest) *pb.CreateHouseResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateHouseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateHouseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlatsInHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseServiceServer) GetFlatsInHouse(_a0 context.Context, _a1 *pb.GetFlatsInHouseRequest) (*pb.GetFlatsInHouseResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetFlatsInHouse")
	}

	var r0 *pb.GetFlatsInHouseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetFlatsInHouseRequest) (*pb.GetFlatsInHouseResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetFlatsInHouseRequest) *pb.GetFlatsInHouseResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetFlatsInHouseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetFlatsInHouseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToHouse provides a mock function with given fields: _a0, _a1
func (_m *HouseServiceServer) SubscribeToHouse(_a0 context.Context, _a1 *pb.SubscribeToHouseRequest) (*pb.SubscribeToHouseResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeToHouse")
	}

	var r0 *pb.SubscribeToHouseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.SubscribeToHouseRequest) (*pb.SubscribeToHouseResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.SubscribeToHouseRequest) *pb.SubscribeToHouseResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.SubscribeToHouseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.SubscribeToHouseRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFlat provides a mock function with given fields: _a0, _a1
func (_m *HouseServiceServer) UpdateFlat(_a0 context.Context, _a1 *pb.UpdateFlatRequest) (*pb.UpdateFlatResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFlat")
	}

	var r0 *pb.UpdateFlatResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateFlatRequest) (*pb.UpdateFlatResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateFlatRequest) *pb.UpdateFlatResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.UpdateFlatResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateFlatRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedHouseServiceServer provides a mock function with given fields:
func (_m *HouseServiceServer) mustEmbedUnimplementedHouseServiceServer() {
	_m.Called()
}

// NewHouseServiceServer creates a new instance of HouseServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHouseServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *HouseServiceServer {
	mock := &HouseServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
