// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: ActiManage/core_api/common.proto

package core_api

import (
	context "context"
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

// 注册
type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActiManage_core_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_ActiManage_core_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_ActiManage_core_api_common_proto_rawDescGZIP(), []int{0}
}

type SignUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpResp) Reset() {
	*x = SignUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActiManage_core_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResp) ProtoMessage() {}

func (x *SignUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_ActiManage_core_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResp.ProtoReflect.Descriptor instead.
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return file_ActiManage_core_api_common_proto_rawDescGZIP(), []int{1}
}

// 登录
type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActiManage_core_api_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_ActiManage_core_api_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_ActiManage_core_api_common_proto_rawDescGZIP(), []int{2}
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActiManage_core_api_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_ActiManage_core_api_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_ActiManage_core_api_common_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignInResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_ActiManage_core_api_common_proto protoreflect.FileDescriptor

var file_ActiManage_core_api_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x0b, 0x0a, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x71, 0x22, 0x32, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d, 0x49, 0x44, 0x4c,
	0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActiManage_core_api_common_proto_rawDescOnce sync.Once
	file_ActiManage_core_api_common_proto_rawDescData = file_ActiManage_core_api_common_proto_rawDesc
)

func file_ActiManage_core_api_common_proto_rawDescGZIP() []byte {
	file_ActiManage_core_api_common_proto_rawDescOnce.Do(func() {
		file_ActiManage_core_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActiManage_core_api_common_proto_rawDescData)
	})
	return file_ActiManage_core_api_common_proto_rawDescData
}

var file_ActiManage_core_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ActiManage_core_api_common_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),  // 0: core_api.SignUpReq
	(*SignUpResp)(nil), // 1: core_api.SignUpResp
	(*SignInReq)(nil),  // 2: core_api.SignInReq
	(*SignInResp)(nil), // 3: core_api.SignInResp
}
var file_ActiManage_core_api_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActiManage_core_api_common_proto_init() }
func file_ActiManage_core_api_common_proto_init() {
	if File_ActiManage_core_api_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActiManage_core_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_ActiManage_core_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResp); i {
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
		file_ActiManage_core_api_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_ActiManage_core_api_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
			RawDescriptor: file_ActiManage_core_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActiManage_core_api_common_proto_goTypes,
		DependencyIndexes: file_ActiManage_core_api_common_proto_depIdxs,
		MessageInfos:      file_ActiManage_core_api_common_proto_msgTypes,
	}.Build()
	File_ActiManage_core_api_common_proto = out.File
	file_ActiManage_core_api_common_proto_rawDesc = nil
	file_ActiManage_core_api_common_proto_goTypes = nil
	file_ActiManage_core_api_common_proto_depIdxs = nil
}

var _ context.Context
