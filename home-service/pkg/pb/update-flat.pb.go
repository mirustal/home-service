// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: update-flat.proto

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

type UpdateFlatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateFlatRequest) Reset() {
	*x = UpdateFlatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_flat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlatRequest) ProtoMessage() {}

func (x *UpdateFlatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_flat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlatRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlatRequest) Descriptor() ([]byte, []int) {
	return file_update_flat_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateFlatRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFlatRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateFlatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	HouseId int32  `protobuf:"varint,2,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	Price   int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Rooms   int32  `protobuf:"varint,4,opt,name=rooms,proto3" json:"rooms,omitempty"`
	Status  string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateFlatResponse) Reset() {
	*x = UpdateFlatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_flat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlatResponse) ProtoMessage() {}

func (x *UpdateFlatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_update_flat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlatResponse.ProtoReflect.Descriptor instead.
func (*UpdateFlatResponse) Descriptor() ([]byte, []int) {
	return file_update_flat_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateFlatResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFlatResponse) GetHouseId() int32 {
	if x != nil {
		return x.HouseId
	}
	return 0
}

func (x *UpdateFlatResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateFlatResponse) GetRooms() int32 {
	if x != nil {
		return x.Rooms
	}
	return 0
}

func (x *UpdateFlatResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_update_flat_proto protoreflect.FileDescriptor

var file_update_flat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6c, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_update_flat_proto_rawDescOnce sync.Once
	file_update_flat_proto_rawDescData = file_update_flat_proto_rawDesc
)

func file_update_flat_proto_rawDescGZIP() []byte {
	file_update_flat_proto_rawDescOnce.Do(func() {
		file_update_flat_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_flat_proto_rawDescData)
	})
	return file_update_flat_proto_rawDescData
}

var file_update_flat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_update_flat_proto_goTypes = []any{
	(*UpdateFlatRequest)(nil),  // 0: home.UpdateFlatRequest
	(*UpdateFlatResponse)(nil), // 1: home.UpdateFlatResponse
}
var file_update_flat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_update_flat_proto_init() }
func file_update_flat_proto_init() {
	if File_update_flat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_update_flat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateFlatRequest); i {
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
		file_update_flat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateFlatResponse); i {
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
			RawDescriptor: file_update_flat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_update_flat_proto_goTypes,
		DependencyIndexes: file_update_flat_proto_depIdxs,
		MessageInfos:      file_update_flat_proto_msgTypes,
	}.Build()
	File_update_flat_proto = out.File
	file_update_flat_proto_rawDesc = nil
	file_update_flat_proto_goTypes = nil
	file_update_flat_proto_depIdxs = nil
}
