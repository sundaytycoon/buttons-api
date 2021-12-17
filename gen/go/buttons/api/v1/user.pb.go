// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: buttons/api/v1/user.proto

package _go

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserServiceGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // the id of this message.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // @exclude the name of this message
}

func (x *UserServiceGetRequest) Reset() {
	*x = UserServiceGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buttons_api_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceGetRequest) ProtoMessage() {}

func (x *UserServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buttons_api_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceGetRequest.ProtoReflect.Descriptor instead.
func (*UserServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_buttons_api_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserServiceGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserServiceGetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserServiceGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // the id of this message.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // @exclude the name of this message
}

func (x *UserServiceGetResponse) Reset() {
	*x = UserServiceGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buttons_api_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceGetResponse) ProtoMessage() {}

func (x *UserServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buttons_api_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceGetResponse.ProtoReflect.Descriptor instead.
func (*UserServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_buttons_api_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserServiceGetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserServiceGetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_buttons_api_v1_user_proto protoreflect.FileDescriptor

var file_buttons_api_v1_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x15, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x63, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x62,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x0a, 0x14, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x42, 0x04, 0x55, 0x73, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x74, 0x79,
	0x63, 0x6f, 0x6f, 0x6e, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buttons_api_v1_user_proto_rawDescOnce sync.Once
	file_buttons_api_v1_user_proto_rawDescData = file_buttons_api_v1_user_proto_rawDesc
)

func file_buttons_api_v1_user_proto_rawDescGZIP() []byte {
	file_buttons_api_v1_user_proto_rawDescOnce.Do(func() {
		file_buttons_api_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_buttons_api_v1_user_proto_rawDescData)
	})
	return file_buttons_api_v1_user_proto_rawDescData
}

var file_buttons_api_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buttons_api_v1_user_proto_goTypes = []interface{}{
	(*UserServiceGetRequest)(nil),  // 0: buttons.api.v1.UserServiceGetRequest
	(*UserServiceGetResponse)(nil), // 1: buttons.api.v1.UserServiceGetResponse
}
var file_buttons_api_v1_user_proto_depIdxs = []int32{
	0, // 0: buttons.api.v1.UserService.Get:input_type -> buttons.api.v1.UserServiceGetRequest
	1, // 1: buttons.api.v1.UserService.Get:output_type -> buttons.api.v1.UserServiceGetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buttons_api_v1_user_proto_init() }
func file_buttons_api_v1_user_proto_init() {
	if File_buttons_api_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buttons_api_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceGetRequest); i {
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
		file_buttons_api_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceGetResponse); i {
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
			RawDescriptor: file_buttons_api_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buttons_api_v1_user_proto_goTypes,
		DependencyIndexes: file_buttons_api_v1_user_proto_depIdxs,
		MessageInfos:      file_buttons_api_v1_user_proto_msgTypes,
	}.Build()
	File_buttons_api_v1_user_proto = out.File
	file_buttons_api_v1_user_proto_rawDesc = nil
	file_buttons_api_v1_user_proto_goTypes = nil
	file_buttons_api_v1_user_proto_depIdxs = nil
}