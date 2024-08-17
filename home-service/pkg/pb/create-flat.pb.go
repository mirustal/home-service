// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: create-flat.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFlatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseId int32 `protobuf:"varint,1,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	Price   int32 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Rooms   int32 `protobuf:"varint,3,opt,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *CreateFlatRequest) Reset() {
	*x = CreateFlatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_flat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlatRequest) ProtoMessage() {}

func (x *CreateFlatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_create_flat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlatRequest.ProtoReflect.Descriptor instead.
func (*CreateFlatRequest) Descriptor() ([]byte, []int) {
	return file_create_flat_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlatRequest) GetHouseId() int32 {
	if x != nil {
		return x.HouseId
	}
	return 0
}

func (x *CreateFlatRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateFlatRequest) GetRooms() int32 {
	if x != nil {
		return x.Rooms
	}
	return 0
}

type CreateFlatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseId int32  `protobuf:"varint,2,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	Price   int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Rooms   int32  `protobuf:"varint,4,opt,name=rooms,proto3" json:"rooms,omitempty"`
	Status  Status `protobuf:"varint,5,opt,name=status,proto3,enum=home.Status" json:"status,omitempty"`
}

func (x *CreateFlatResponse) Reset() {
	*x = CreateFlatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_flat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlatResponse) ProtoMessage() {}

func (x *CreateFlatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_create_flat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlatResponse.ProtoReflect.Descriptor instead.
func (*CreateFlatResponse) Descriptor() ([]byte, []int) {
	return file_create_flat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFlatResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateFlatResponse) GetHouseId() int32 {
	if x != nil {
		return x.HouseId
	}
	return 0
}

func (x *CreateFlatResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateFlatResponse) GetRooms() int32 {
	if x != nil {
		return x.Rooms
	}
	return 0
}

func (x *CreateFlatResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_CREATED
}

var File_create_flat_proto protoreflect.FileDescriptor

var file_create_flat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6c, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x1a, 0x14, 0x66, 0x6c, 0x61, 0x74, 0x73,
	0x2d, 0x69, 0x6e, 0x2d, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_create_flat_proto_rawDescOnce sync.Once
	file_create_flat_proto_rawDescData = file_create_flat_proto_rawDesc
)

func file_create_flat_proto_rawDescGZIP() []byte {
	file_create_flat_proto_rawDescOnce.Do(func() {
		file_create_flat_proto_rawDescData = protoimpl.X.CompressGZIP(file_create_flat_proto_rawDescData)
	})
	return file_create_flat_proto_rawDescData
}

var file_create_flat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_create_flat_proto_goTypes = []any{
	(*CreateFlatRequest)(nil),  // 0: home.CreateFlatRequest
	(*CreateFlatResponse)(nil), // 1: home.CreateFlatResponse
	(Status)(0),                // 2: home.Status
}
var file_create_flat_proto_depIdxs = []int32{
	2, // 0: home.CreateFlatResponse.status:type_name -> home.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_create_flat_proto_init() }
func file_create_flat_proto_init() {
	if File_create_flat_proto != nil {
		return
	}
	file_flats_in_house_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_create_flat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFlatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_create_flat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFlatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_create_flat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_create_flat_proto_goTypes,
		DependencyIndexes: file_create_flat_proto_depIdxs,
		MessageInfos:      file_create_flat_proto_msgTypes,
	}.Build()
	File_create_flat_proto = out.File
	file_create_flat_proto_rawDesc = nil
	file_create_flat_proto_goTypes = nil
	file_create_flat_proto_depIdxs = nil
}
