// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: ActiManage/user/user.proto

package user

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ActiManage_user_user_proto protoreflect.FileDescriptor

var file_ActiManage_user_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1c, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xd3, 0x10, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x53,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x58, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x64, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x22, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x85, 0x01, 0x0a, 0x24, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x16,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b,
	0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x44,
	0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x44, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x56, 0x69,
	0x65, 0x77, 0x4f, 0x66, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41,
	0x6e, 0x64, 0x56, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x56, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4f,
	0x66, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4f, 0x66, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x63, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x79, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x42, 0x79, 0x56, 0x69, 0x65, 0x77, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x75,
	0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b,
	0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x75, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x42, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x1a, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x58, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f,
	0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d, 0x49, 0x44, 0x4c, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ActiManage_user_user_proto_goTypes = []interface{}{
	(*UserSignUpReq)(nil),                            // 0: user.UserSignUpReq
	(*UserLoginReq)(nil),                             // 1: user.UserLoginReq
	(*GetUserInfoReq)(nil),                           // 2: user.GetUserInfoReq
	(*UpdateUserInfoReq)(nil),                        // 3: user.UpdateUserInfoReq
	(*SetPasswordReq)(nil),                           // 4: user.SetPasswordReq
	(*SetNoticeReq)(nil),                             // 5: user.SetNoticeReq
	(*CreateReserverReq)(nil),                        // 6: user.CreateReserverReq
	(*DeleteReserverReq)(nil),                        // 7: user.DeleteReserverReq
	(*UpdateReserverReq)(nil),                        // 8: user.UpdateReserverReq
	(*ListReserversReq)(nil),                         // 9: user.ListReserversReq
	(*CreateBookRecordReq)(nil),                      // 10: user.CreateBookRecordReq
	(*CancelBookRecordReq)(nil),                      // 11: user.CancelBookRecordReq
	(*GetBookRecordReq)(nil),                         // 12: user.GetBookRecordReq
	(*ListBookRecordsByUserReq)(nil),                 // 13: user.ListBookRecordsByUserReq
	(*ListBookRecordsByActivityReq)(nil),             // 14: user.ListBookRecordsByActivityReq
	(*CheckBookRecordByUserIdAndActivityIdReq)(nil),  // 15: user.CheckBookRecordByUserIdAndActivityIdReq
	(*CreateReceiptReq)(nil),                         // 16: user.CreateReceiptReq
	(*DeleteReceiptReq)(nil),                         // 17: user.DeleteReceiptReq
	(*ListReceiptsReq)(nil),                          // 18: user.ListReceiptsReq
	(*MarkReceiptReadReq)(nil),                       // 19: user.MarkReceiptReadReq
	(*DoFavoriteReq)(nil),                            // 20: user.DoFavoriteReq
	(*CancelFavoriteReq)(nil),                        // 21: user.CancelFavoriteReq
	(*CreateViewReq)(nil),                            // 22: user.CreateViewReq
	(*GetFavoriteAndViewOfActivityReq)(nil),          // 23: user.GetFavoriteAndViewOfActivityReq
	(*GetViewOfMerchantReq)(nil),                     // 24: user.GetViewOfMerchantReq
	(*ListMerchantIdsByViewRankReq)(nil),             // 25: user.ListMerchantIdsByViewRankReq
	(*ListMerchantIdsByBookRecordRankReq)(nil),       // 26: user.ListMerchantIdsByBookRecordRankReq
	(*ListActivityIdsByBookRecordRankReq)(nil),       // 27: user.ListActivityIdsByBookRecordRankReq
	(*GetViewDataByMerchantReq)(nil),                 // 28: user.GetViewDataByMerchantReq
	(*UserSignUpResp)(nil),                           // 29: user.UserSignUpResp
	(*UserLoginResp)(nil),                            // 30: user.UserLoginResp
	(*GetUserInfoResp)(nil),                          // 31: user.GetUserInfoResp
	(*Response)(nil),                                 // 32: user.Response
	(*ListReserversResp)(nil),                        // 33: user.ListReserversResp
	(*GetBookRecordResp)(nil),                        // 34: user.GetBookRecordResp
	(*ListBookRecordsByUserResp)(nil),                // 35: user.ListBookRecordsByUserResp
	(*ListBookRecordsByActivityResp)(nil),            // 36: user.ListBookRecordsByActivityResp
	(*CheckBookRecordByUserIdAndActivityIdResp)(nil), // 37: user.CheckBookRecordByUserIdAndActivityIdResp
	(*ListReceiptsResp)(nil),                         // 38: user.ListReceiptsResp
	(*GetFavoriteAndViewOfActivityResp)(nil),         // 39: user.GetFavoriteAndViewOfActivityResp
	(*GetViewOfMerchantResp)(nil),                    // 40: user.GetViewOfMerchantResp
	(*ListMerchantIdsByViewRankResp)(nil),            // 41: user.ListMerchantIdsByViewRankResp
	(*ListMerchantIdsByBookRecordRankResp)(nil),      // 42: user.ListMerchantIdsByBookRecordRankResp
	(*ListActivityIdsByBookRecordRankResp)(nil),      // 43: user.ListActivityIdsByBookRecordRankResp
	(*GetViewDataByMerchantResp)(nil),                // 44: user.GetViewDataByMerchantResp
}
var file_ActiManage_user_user_proto_depIdxs = []int32{
	0,  // 0: user.UserService.UserSignUp:input_type -> user.UserSignUpReq
	1,  // 1: user.UserService.UserLogin:input_type -> user.UserLoginReq
	2,  // 2: user.UserService.GetUserInfo:input_type -> user.GetUserInfoReq
	3,  // 3: user.UserService.UpdateUserInfo:input_type -> user.UpdateUserInfoReq
	4,  // 4: user.UserService.SetPassword:input_type -> user.SetPasswordReq
	5,  // 5: user.UserService.SetNotice:input_type -> user.SetNoticeReq
	6,  // 6: user.UserService.CreateReserver:input_type -> user.CreateReserverReq
	7,  // 7: user.UserService.DeleteReserver:input_type -> user.DeleteReserverReq
	8,  // 8: user.UserService.UpdateReserver:input_type -> user.UpdateReserverReq
	9,  // 9: user.UserService.ListReservers:input_type -> user.ListReserversReq
	10, // 10: user.UserService.CreateBookRecord:input_type -> user.CreateBookRecordReq
	11, // 11: user.UserService.CancelBookRecord:input_type -> user.CancelBookRecordReq
	12, // 12: user.UserService.GetBookRecordDetail:input_type -> user.GetBookRecordReq
	13, // 13: user.UserService.ListBookRecordsByUser:input_type -> user.ListBookRecordsByUserReq
	14, // 14: user.UserService.ListBookRecordsByActivity:input_type -> user.ListBookRecordsByActivityReq
	15, // 15: user.UserService.CheckBookRecordByUserIdAndActivityId:input_type -> user.CheckBookRecordByUserIdAndActivityIdReq
	16, // 16: user.UserService.CreateReceipt:input_type -> user.CreateReceiptReq
	17, // 17: user.UserService.DeleteReceipt:input_type -> user.DeleteReceiptReq
	18, // 18: user.UserService.ListReceipts:input_type -> user.ListReceiptsReq
	19, // 19: user.UserService.MarkReceiptRead:input_type -> user.MarkReceiptReadReq
	20, // 20: user.UserService.DoFavorite:input_type -> user.DoFavoriteReq
	21, // 21: user.UserService.CancelFavorite:input_type -> user.CancelFavoriteReq
	22, // 22: user.UserService.CreateView:input_type -> user.CreateViewReq
	23, // 23: user.UserService.GetFavoriteAndViewOfActivity:input_type -> user.GetFavoriteAndViewOfActivityReq
	24, // 24: user.UserService.GetViewOfMerchant:input_type -> user.GetViewOfMerchantReq
	25, // 25: user.UserService.ListMerchantIdByViewRank:input_type -> user.ListMerchantIdsByViewRankReq
	26, // 26: user.UserService.ListMerchantIdByBookRecordRank:input_type -> user.ListMerchantIdsByBookRecordRankReq
	27, // 27: user.UserService.ListActivityIdByBookRecordRank:input_type -> user.ListActivityIdsByBookRecordRankReq
	28, // 28: user.UserService.GetViewDataByMerchant:input_type -> user.GetViewDataByMerchantReq
	29, // 29: user.UserService.UserSignUp:output_type -> user.UserSignUpResp
	30, // 30: user.UserService.UserLogin:output_type -> user.UserLoginResp
	31, // 31: user.UserService.GetUserInfo:output_type -> user.GetUserInfoResp
	32, // 32: user.UserService.UpdateUserInfo:output_type -> user.Response
	32, // 33: user.UserService.SetPassword:output_type -> user.Response
	32, // 34: user.UserService.SetNotice:output_type -> user.Response
	32, // 35: user.UserService.CreateReserver:output_type -> user.Response
	32, // 36: user.UserService.DeleteReserver:output_type -> user.Response
	32, // 37: user.UserService.UpdateReserver:output_type -> user.Response
	33, // 38: user.UserService.ListReservers:output_type -> user.ListReserversResp
	32, // 39: user.UserService.CreateBookRecord:output_type -> user.Response
	32, // 40: user.UserService.CancelBookRecord:output_type -> user.Response
	34, // 41: user.UserService.GetBookRecordDetail:output_type -> user.GetBookRecordResp
	35, // 42: user.UserService.ListBookRecordsByUser:output_type -> user.ListBookRecordsByUserResp
	36, // 43: user.UserService.ListBookRecordsByActivity:output_type -> user.ListBookRecordsByActivityResp
	37, // 44: user.UserService.CheckBookRecordByUserIdAndActivityId:output_type -> user.CheckBookRecordByUserIdAndActivityIdResp
	32, // 45: user.UserService.CreateReceipt:output_type -> user.Response
	32, // 46: user.UserService.DeleteReceipt:output_type -> user.Response
	38, // 47: user.UserService.ListReceipts:output_type -> user.ListReceiptsResp
	32, // 48: user.UserService.MarkReceiptRead:output_type -> user.Response
	32, // 49: user.UserService.DoFavorite:output_type -> user.Response
	32, // 50: user.UserService.CancelFavorite:output_type -> user.Response
	32, // 51: user.UserService.CreateView:output_type -> user.Response
	39, // 52: user.UserService.GetFavoriteAndViewOfActivity:output_type -> user.GetFavoriteAndViewOfActivityResp
	40, // 53: user.UserService.GetViewOfMerchant:output_type -> user.GetViewOfMerchantResp
	41, // 54: user.UserService.ListMerchantIdByViewRank:output_type -> user.ListMerchantIdsByViewRankResp
	42, // 55: user.UserService.ListMerchantIdByBookRecordRank:output_type -> user.ListMerchantIdsByBookRecordRankResp
	43, // 56: user.UserService.ListActivityIdByBookRecordRank:output_type -> user.ListActivityIdsByBookRecordRankResp
	44, // 57: user.UserService.GetViewDataByMerchant:output_type -> user.GetViewDataByMerchantResp
	29, // [29:58] is the sub-list for method output_type
	0,  // [0:29] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ActiManage_user_user_proto_init() }
func file_ActiManage_user_user_proto_init() {
	if File_ActiManage_user_user_proto != nil {
		return
	}
	file_ActiManage_user_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ActiManage_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ActiManage_user_user_proto_goTypes,
		DependencyIndexes: file_ActiManage_user_user_proto_depIdxs,
	}.Build()
	File_ActiManage_user_user_proto = out.File
	file_ActiManage_user_user_proto_rawDesc = nil
	file_ActiManage_user_user_proto_goTypes = nil
	file_ActiManage_user_user_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.2. DO NOT EDIT.

type UserService interface {
	UserSignUp(ctx context.Context, req *UserSignUpReq) (res *UserSignUpResp, err error)
	UserLogin(ctx context.Context, req *UserLoginReq) (res *UserLoginResp, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoReq) (res *GetUserInfoResp, err error)
	UpdateUserInfo(ctx context.Context, req *UpdateUserInfoReq) (res *Response, err error)
	SetPassword(ctx context.Context, req *SetPasswordReq) (res *Response, err error)
	SetNotice(ctx context.Context, req *SetNoticeReq) (res *Response, err error)
	CreateReserver(ctx context.Context, req *CreateReserverReq) (res *Response, err error)
	DeleteReserver(ctx context.Context, req *DeleteReserverReq) (res *Response, err error)
	UpdateReserver(ctx context.Context, req *UpdateReserverReq) (res *Response, err error)
	ListReservers(ctx context.Context, req *ListReserversReq) (res *ListReserversResp, err error)
	CreateBookRecord(ctx context.Context, req *CreateBookRecordReq) (res *Response, err error)
	CancelBookRecord(ctx context.Context, req *CancelBookRecordReq) (res *Response, err error)
	GetBookRecordDetail(ctx context.Context, req *GetBookRecordReq) (res *GetBookRecordResp, err error)
	ListBookRecordsByUser(ctx context.Context, req *ListBookRecordsByUserReq) (res *ListBookRecordsByUserResp, err error)
	ListBookRecordsByActivity(ctx context.Context, req *ListBookRecordsByActivityReq) (res *ListBookRecordsByActivityResp, err error)
	CheckBookRecordByUserIdAndActivityId(ctx context.Context, req *CheckBookRecordByUserIdAndActivityIdReq) (res *CheckBookRecordByUserIdAndActivityIdResp, err error)
	CreateReceipt(ctx context.Context, req *CreateReceiptReq) (res *Response, err error)
	DeleteReceipt(ctx context.Context, req *DeleteReceiptReq) (res *Response, err error)
	ListReceipts(ctx context.Context, req *ListReceiptsReq) (res *ListReceiptsResp, err error)
	MarkReceiptRead(ctx context.Context, req *MarkReceiptReadReq) (res *Response, err error)
	DoFavorite(ctx context.Context, req *DoFavoriteReq) (res *Response, err error)
	CancelFavorite(ctx context.Context, req *CancelFavoriteReq) (res *Response, err error)
	CreateView(ctx context.Context, req *CreateViewReq) (res *Response, err error)
	GetFavoriteAndViewOfActivity(ctx context.Context, req *GetFavoriteAndViewOfActivityReq) (res *GetFavoriteAndViewOfActivityResp, err error)
	GetViewOfMerchant(ctx context.Context, req *GetViewOfMerchantReq) (res *GetViewOfMerchantResp, err error)
	ListMerchantIdByViewRank(ctx context.Context, req *ListMerchantIdsByViewRankReq) (res *ListMerchantIdsByViewRankResp, err error)
	ListMerchantIdByBookRecordRank(ctx context.Context, req *ListMerchantIdsByBookRecordRankReq) (res *ListMerchantIdsByBookRecordRankResp, err error)
	ListActivityIdByBookRecordRank(ctx context.Context, req *ListActivityIdsByBookRecordRankReq) (res *ListActivityIdsByBookRecordRankResp, err error)
	GetViewDataByMerchant(ctx context.Context, req *GetViewDataByMerchantReq) (res *GetViewDataByMerchantResp, err error)
}
