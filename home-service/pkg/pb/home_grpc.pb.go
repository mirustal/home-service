// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: home.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HouseService_CreateHouse_FullMethodName      = "/home.HouseService/CreateHouse"
	HouseService_GetFlatsInHouse_FullMethodName  = "/home.HouseService/GetFlatsInHouse"
	HouseService_SubscribeToHouse_FullMethodName = "/home.HouseService/SubscribeToHouse"
	HouseService_CreateFlat_FullMethodName       = "/home.HouseService/CreateFlat"
	HouseService_UpdateFlat_FullMethodName       = "/home.HouseService/UpdateFlat"
)

// HouseServiceClient is the client API for HouseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HouseServiceClient interface {
	CreateHouse(ctx context.Context, in *CreateHouseRequest, opts ...grpc.CallOption) (*CreateHouseResponse, error)
	GetFlatsInHouse(ctx context.Context, in *GetFlatsInHouseRequest, opts ...grpc.CallOption) (*GetFlatsInHouseResponse, error)
	SubscribeToHouse(ctx context.Context, in *SubscribeToHouseRequest, opts ...grpc.CallOption) (*SubscribeToHouseResponse, error)
	CreateFlat(ctx context.Context, in *CreateFlatRequest, opts ...grpc.CallOption) (*CreateFlatResponse, error)
	UpdateFlat(ctx context.Context, in *UpdateFlatRequest, opts ...grpc.CallOption) (*UpdateFlatResponse, error)
}

type houseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHouseServiceClient(cc grpc.ClientConnInterface) HouseServiceClient {
	return &houseServiceClient{cc}
}

func (c *houseServiceClient) CreateHouse(ctx context.Context, in *CreateHouseRequest, opts ...grpc.CallOption) (*CreateHouseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHouseResponse)
	err := c.cc.Invoke(ctx, HouseService_CreateHouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseServiceClient) GetFlatsInHouse(ctx context.Context, in *GetFlatsInHouseRequest, opts ...grpc.CallOption) (*GetFlatsInHouseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFlatsInHouseResponse)
	err := c.cc.Invoke(ctx, HouseService_GetFlatsInHouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseServiceClient) SubscribeToHouse(ctx context.Context, in *SubscribeToHouseRequest, opts ...grpc.CallOption) (*SubscribeToHouseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubscribeToHouseResponse)
	err := c.cc.Invoke(ctx, HouseService_SubscribeToHouse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseServiceClient) CreateFlat(ctx context.Context, in *CreateFlatRequest, opts ...grpc.CallOption) (*CreateFlatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFlatResponse)
	err := c.cc.Invoke(ctx, HouseService_CreateFlat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *houseServiceClient) UpdateFlat(ctx context.Context, in *UpdateFlatRequest, opts ...grpc.CallOption) (*UpdateFlatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateFlatResponse)
	err := c.cc.Invoke(ctx, HouseService_UpdateFlat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HouseServiceServer is the server API for HouseService service.
// All implementations must embed UnimplementedHouseServiceServer
// for forward compatibility.
type HouseServiceServer interface {
	CreateHouse(context.Context, *CreateHouseRequest) (*CreateHouseResponse, error)
	GetFlatsInHouse(context.Context, *GetFlatsInHouseRequest) (*GetFlatsInHouseResponse, error)
	SubscribeToHouse(context.Context, *SubscribeToHouseRequest) (*SubscribeToHouseResponse, error)
	CreateFlat(context.Context, *CreateFlatRequest) (*CreateFlatResponse, error)
	UpdateFlat(context.Context, *UpdateFlatRequest) (*UpdateFlatResponse, error)
	mustEmbedUnimplementedHouseServiceServer()
}

// UnimplementedHouseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHouseServiceServer struct{}

func (UnimplementedHouseServiceServer) CreateHouse(context.Context, *CreateHouseRequest) (*CreateHouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHouse not implemented")
}
func (UnimplementedHouseServiceServer) GetFlatsInHouse(context.Context, *GetFlatsInHouseRequest) (*GetFlatsInHouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlatsInHouse not implemented")
}
func (UnimplementedHouseServiceServer) SubscribeToHouse(context.Context, *SubscribeToHouseRequest) (*SubscribeToHouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToHouse not implemented")
}
func (UnimplementedHouseServiceServer) CreateFlat(context.Context, *CreateFlatRequest) (*CreateFlatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlat not implemented")
}
func (UnimplementedHouseServiceServer) UpdateFlat(context.Context, *UpdateFlatRequest) (*UpdateFlatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlat not implemented")
}
func (UnimplementedHouseServiceServer) mustEmbedUnimplementedHouseServiceServer() {}
func (UnimplementedHouseServiceServer) testEmbeddedByValue()                      {}

// UnsafeHouseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HouseServiceServer will
// result in compilation errors.
type UnsafeHouseServiceServer interface {
	mustEmbedUnimplementedHouseServiceServer()
}

func RegisterHouseServiceServer(s grpc.ServiceRegistrar, srv HouseServiceServer) {
	// If the following call pancis, it indicates UnimplementedHouseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HouseService_ServiceDesc, srv)
}

func _HouseService_CreateHouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).CreateHouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HouseService_CreateHouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).CreateHouse(ctx, req.(*CreateHouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HouseService_GetFlatsInHouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlatsInHouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).GetFlatsInHouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HouseService_GetFlatsInHouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).GetFlatsInHouse(ctx, req.(*GetFlatsInHouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HouseService_SubscribeToHouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeToHouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).SubscribeToHouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HouseService_SubscribeToHouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).SubscribeToHouse(ctx, req.(*SubscribeToHouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HouseService_CreateFlat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).CreateFlat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HouseService_CreateFlat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).CreateFlat(ctx, req.(*CreateFlatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HouseService_UpdateFlat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HouseServiceServer).UpdateFlat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HouseService_UpdateFlat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HouseServiceServer).UpdateFlat(ctx, req.(*UpdateFlatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HouseService_ServiceDesc is the grpc.ServiceDesc for HouseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HouseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "home.HouseService",
	HandlerType: (*HouseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHouse",
			Handler:    _HouseService_CreateHouse_Handler,
		},
		{
			MethodName: "GetFlatsInHouse",
			Handler:    _HouseService_GetFlatsInHouse_Handler,
		},
		{
			MethodName: "SubscribeToHouse",
			Handler:    _HouseService_SubscribeToHouse_Handler,
		},
		{
			MethodName: "CreateFlat",
			Handler:    _HouseService_CreateFlat_Handler,
		},
		{
			MethodName: "UpdateFlat",
			Handler:    _HouseService_UpdateFlat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home.proto",
}
