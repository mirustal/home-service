// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: register.proto

package authgrpc

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

type UserType int32

const (
	UserType_UNKNOWN   UserType = 0
	UserType_client    UserType = 1
	UserType_moderator UserType = 2
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "UNKNOWN",
		1: "client",
		2: "moderator",
	}
	UserType_value = map[string]int32{
		"UNKNOWN":   0,
		"client":    1,
		"moderator": 2,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_register_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_register_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	UserType UserType `protobuf:"varint,3,opt,name=user_type,json=userType,proto3,enum=auth.UserType" json:"user_type,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_UNKNOWN
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_register_proto protoreflect.FileDescriptor

var file_register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x70, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x02, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2e, 0x2f,
	0x2e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_register_proto_rawDescOnce sync.Once
	file_register_proto_rawDescData = file_register_proto_rawDesc
)

func file_register_proto_rawDescGZIP() []byte {
	file_register_proto_rawDescOnce.Do(func() {
		file_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_proto_rawDescData)
	})
	return file_register_proto_rawDescData
}

var file_register_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_register_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_register_proto_goTypes = []any{
	(UserType)(0),            // 0: auth.UserType
	(*RegisterRequest)(nil),  // 1: auth.RegisterRequest
	(*RegisterResponse)(nil), // 2: auth.RegisterResponse
}
var file_register_proto_depIdxs = []int32{
	0, // 0: auth.RegisterRequest.user_type:type_name -> auth.UserType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_register_proto_init() }
func file_register_proto_init() {
	if File_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterRequest); i {
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
		file_register_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_register_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_proto_goTypes,
		DependencyIndexes: file_register_proto_depIdxs,
		EnumInfos:         file_register_proto_enumTypes,
		MessageInfos:      file_register_proto_msgTypes,
	}.Build()
	File_register_proto = out.File
	file_register_proto_rawDesc = nil
	file_register_proto_goTypes = nil
	file_register_proto_depIdxs = nil
}
