// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: ActiManage/system/system.proto

package system

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

var File_ActiManage_system_system_proto protoreflect.FileDescriptor

var file_ActiManage_system_system_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1e, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb3, 0x0f, 0x0a, 0x0d, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x47, 0x0a, 0x13, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55,
	0x72, 0x69, 0x12, 0x1f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x72, 0x69,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x55, 0x72,
	0x69, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x41, 0x64, 0x12, 0x10,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x05, 0x53, 0x65, 0x74, 0x41, 0x64, 0x12, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x54, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x10,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12,
	0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x59,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x67, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x73, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x28, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x69, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x43, 0x0a, 0x11, 0x53, 0x74, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x53, 0x74, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x53, 0x74, 0x73, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d,
	0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2d, 0x49, 0x44, 0x4c, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_ActiManage_system_system_proto_goTypes = []interface{}{
	(*MerchantLoginReq)(nil),                  // 0: system.MerchantLoginReq
	(*MerchantSetPasswordReq)(nil),            // 1: system.MerchantSetPasswordReq
	(*UpdateSettingReq)(nil),                  // 2: system.UpdateSettingReq
	(*GetMerchantSettingReq)(nil),             // 3: system.GetMerchantSettingReq
	(*UpdateMerchantInfoReq)(nil),             // 4: system.UpdateMerchantInfoReq
	(*GetMerchantInfoByUriReq)(nil),           // 5: system.GetMerchantInfoByUriReq
	(*GetAdReq)(nil),                          // 6: system.GetAdReq
	(*SetAdReq)(nil),                          // 7: system.SetAdReq
	(*CreateActivityReq)(nil),                 // 8: system.CreateActivityReq
	(*TopActivityReq)(nil),                    // 9: system.TopActivityReq
	(*DeleteActivityReq)(nil),                 // 10: system.DeleteActivityReq
	(*UpdateActivityReq)(nil),                 // 11: system.UpdateActivityReq
	(*ListActivitiesReq)(nil),                 // 12: system.ListActivitiesReq
	(*GetActivityReq)(nil),                    // 13: system.GetActivityReq
	(*AdminLoginReq)(nil),                     // 14: system.AdminLoginReq
	(*AdminSetPasswordReq)(nil),               // 15: system.AdminSetPasswordReq
	(*CreateMerchantReq)(nil),                 // 16: system.CreateMerchantReq
	(*UpdateMerchantReq)(nil),                 // 17: system.UpdateMerchantReq
	(*DeleteMerchantReq)(nil),                 // 18: system.DeleteMerchantReq
	(*GetMerchantInfoReq)(nil),                // 19: system.GetMerchantInfoReq
	(*ListMerchantsReq)(nil),                  // 20: system.ListMerchantsReq
	(*GetMerchantTotalDataReq)(nil),           // 21: system.GetMerchantTotalDataReq
	(*ListMerchantsByMerchantIdReq)(nil),      // 22: system.ListMerchantsByMerchantIdReq
	(*ListMerchantsByActivityNumberReq)(nil),  // 23: system.ListMerchantsByActivityNumberReq
	(*ListActivitiesByActivityIdReq)(nil),     // 24: system.ListActivitiesByActivityIdReq
	(*StsSendVerifyCodeReq)(nil),              // 25: system.StsSendVerifyCodeReq
	(*StsCheckVerifyCodeReq)(nil),             // 26: system.StsCheckVerifyCodeReq
	(*MerchantLoginResp)(nil),                 // 27: system.MerchantLoginResp
	(*Response)(nil),                          // 28: system.Response
	(*GetMerchantSettingResp)(nil),            // 29: system.GetMerchantSettingResp
	(*GetMerchantInfoByUriResp)(nil),          // 30: system.GetMerchantInfoByUriResp
	(*GetAdResp)(nil),                         // 31: system.GetAdResp
	(*ListActivitiesResp)(nil),                // 32: system.ListActivitiesResp
	(*GetActivityResp)(nil),                   // 33: system.GetActivityResp
	(*AdminLoginResp)(nil),                    // 34: system.AdminLoginResp
	(*GetMerchantInfoResp)(nil),               // 35: system.GetMerchantInfoResp
	(*ListMerchantsResp)(nil),                 // 36: system.ListMerchantsResp
	(*GetMerchantTotalDataResp)(nil),          // 37: system.GetMerchantTotalDataResp
	(*ListMerchantsByMerchantIdResp)(nil),     // 38: system.ListMerchantsByMerchantIdResp
	(*ListMerchantsByActivityNumberResp)(nil), // 39: system.ListMerchantsByActivityNumberResp
	(*ListActivitiesByActivityIdResp)(nil),    // 40: system.ListActivitiesByActivityIdResp
}
var file_ActiManage_system_system_proto_depIdxs = []int32{
	0,  // 0: system.SystemService.MerchantLogin:input_type -> system.MerchantLoginReq
	1,  // 1: system.SystemService.MerchantSetPassword:input_type -> system.MerchantSetPasswordReq
	2,  // 2: system.SystemService.UpdateMerchantSetting:input_type -> system.UpdateSettingReq
	3,  // 3: system.SystemService.GetMerchantSetting:input_type -> system.GetMerchantSettingReq
	4,  // 4: system.SystemService.UpdateMerchantInfo:input_type -> system.UpdateMerchantInfoReq
	5,  // 5: system.SystemService.GetMerchantInfoByUri:input_type -> system.GetMerchantInfoByUriReq
	6,  // 6: system.SystemService.GetAd:input_type -> system.GetAdReq
	7,  // 7: system.SystemService.SetAd:input_type -> system.SetAdReq
	8,  // 8: system.SystemService.CreateActivity:input_type -> system.CreateActivityReq
	9,  // 9: system.SystemService.TopActivity:input_type -> system.TopActivityReq
	10, // 10: system.SystemService.DeleteActivity:input_type -> system.DeleteActivityReq
	11, // 11: system.SystemService.UpdateActivity:input_type -> system.UpdateActivityReq
	12, // 12: system.SystemService.ListActivities:input_type -> system.ListActivitiesReq
	13, // 13: system.SystemService.GetActivity:input_type -> system.GetActivityReq
	14, // 14: system.SystemService.AdminLogin:input_type -> system.AdminLoginReq
	15, // 15: system.SystemService.AdminSetPassword:input_type -> system.AdminSetPasswordReq
	16, // 16: system.SystemService.CreateMerchant:input_type -> system.CreateMerchantReq
	17, // 17: system.SystemService.UpdateMerchant:input_type -> system.UpdateMerchantReq
	18, // 18: system.SystemService.DeleteMerchant:input_type -> system.DeleteMerchantReq
	19, // 19: system.SystemService.GetMerchantInfo:input_type -> system.GetMerchantInfoReq
	20, // 20: system.SystemService.ListMerchants:input_type -> system.ListMerchantsReq
	21, // 21: system.SystemService.GetMerchantTotalData:input_type -> system.GetMerchantTotalDataReq
	22, // 22: system.SystemService.ListMerchantByMerchantId:input_type -> system.ListMerchantsByMerchantIdReq
	23, // 23: system.SystemService.ListMerchantByActivityNumber:input_type -> system.ListMerchantsByActivityNumberReq
	24, // 24: system.SystemService.ListActivityByActivityId:input_type -> system.ListActivitiesByActivityIdReq
	25, // 25: system.SystemService.StsSendVerifyCode:input_type -> system.StsSendVerifyCodeReq
	26, // 26: system.SystemService.StsCheckVerifyCode:input_type -> system.StsCheckVerifyCodeReq
	27, // 27: system.SystemService.MerchantLogin:output_type -> system.MerchantLoginResp
	28, // 28: system.SystemService.MerchantSetPassword:output_type -> system.Response
	28, // 29: system.SystemService.UpdateMerchantSetting:output_type -> system.Response
	29, // 30: system.SystemService.GetMerchantSetting:output_type -> system.GetMerchantSettingResp
	28, // 31: system.SystemService.UpdateMerchantInfo:output_type -> system.Response
	30, // 32: system.SystemService.GetMerchantInfoByUri:output_type -> system.GetMerchantInfoByUriResp
	31, // 33: system.SystemService.GetAd:output_type -> system.GetAdResp
	28, // 34: system.SystemService.SetAd:output_type -> system.Response
	28, // 35: system.SystemService.CreateActivity:output_type -> system.Response
	28, // 36: system.SystemService.TopActivity:output_type -> system.Response
	28, // 37: system.SystemService.DeleteActivity:output_type -> system.Response
	28, // 38: system.SystemService.UpdateActivity:output_type -> system.Response
	32, // 39: system.SystemService.ListActivities:output_type -> system.ListActivitiesResp
	33, // 40: system.SystemService.GetActivity:output_type -> system.GetActivityResp
	34, // 41: system.SystemService.AdminLogin:output_type -> system.AdminLoginResp
	28, // 42: system.SystemService.AdminSetPassword:output_type -> system.Response
	28, // 43: system.SystemService.CreateMerchant:output_type -> system.Response
	28, // 44: system.SystemService.UpdateMerchant:output_type -> system.Response
	28, // 45: system.SystemService.DeleteMerchant:output_type -> system.Response
	35, // 46: system.SystemService.GetMerchantInfo:output_type -> system.GetMerchantInfoResp
	36, // 47: system.SystemService.ListMerchants:output_type -> system.ListMerchantsResp
	37, // 48: system.SystemService.GetMerchantTotalData:output_type -> system.GetMerchantTotalDataResp
	38, // 49: system.SystemService.ListMerchantByMerchantId:output_type -> system.ListMerchantsByMerchantIdResp
	39, // 50: system.SystemService.ListMerchantByActivityNumber:output_type -> system.ListMerchantsByActivityNumberResp
	40, // 51: system.SystemService.ListActivityByActivityId:output_type -> system.ListActivitiesByActivityIdResp
	28, // 52: system.SystemService.StsSendVerifyCode:output_type -> system.Response
	28, // 53: system.SystemService.StsCheckVerifyCode:output_type -> system.Response
	27, // [27:54] is the sub-list for method output_type
	0,  // [0:27] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ActiManage_system_system_proto_init() }
func file_ActiManage_system_system_proto_init() {
	if File_ActiManage_system_system_proto != nil {
		return
	}
	file_ActiManage_system_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ActiManage_system_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ActiManage_system_system_proto_goTypes,
		DependencyIndexes: file_ActiManage_system_system_proto_depIdxs,
	}.Build()
	File_ActiManage_system_system_proto = out.File
	file_ActiManage_system_system_proto_rawDesc = nil
	file_ActiManage_system_system_proto_goTypes = nil
	file_ActiManage_system_system_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.2. DO NOT EDIT.

type SystemService interface {
	MerchantLogin(ctx context.Context, req *MerchantLoginReq) (res *MerchantLoginResp, err error)
	MerchantSetPassword(ctx context.Context, req *MerchantSetPasswordReq) (res *Response, err error)
	UpdateMerchantSetting(ctx context.Context, req *UpdateSettingReq) (res *Response, err error)
	GetMerchantSetting(ctx context.Context, req *GetMerchantSettingReq) (res *GetMerchantSettingResp, err error)
	UpdateMerchantInfo(ctx context.Context, req *UpdateMerchantInfoReq) (res *Response, err error)
	GetMerchantInfoByUri(ctx context.Context, req *GetMerchantInfoByUriReq) (res *GetMerchantInfoByUriResp, err error)
	GetAd(ctx context.Context, req *GetAdReq) (res *GetAdResp, err error)
	SetAd(ctx context.Context, req *SetAdReq) (res *Response, err error)
	CreateActivity(ctx context.Context, req *CreateActivityReq) (res *Response, err error)
	TopActivity(ctx context.Context, req *TopActivityReq) (res *Response, err error)
	DeleteActivity(ctx context.Context, req *DeleteActivityReq) (res *Response, err error)
	UpdateActivity(ctx context.Context, req *UpdateActivityReq) (res *Response, err error)
	ListActivities(ctx context.Context, req *ListActivitiesReq) (res *ListActivitiesResp, err error)
	GetActivity(ctx context.Context, req *GetActivityReq) (res *GetActivityResp, err error)
	AdminLogin(ctx context.Context, req *AdminLoginReq) (res *AdminLoginResp, err error)
	AdminSetPassword(ctx context.Context, req *AdminSetPasswordReq) (res *Response, err error)
	CreateMerchant(ctx context.Context, req *CreateMerchantReq) (res *Response, err error)
	UpdateMerchant(ctx context.Context, req *UpdateMerchantReq) (res *Response, err error)
	DeleteMerchant(ctx context.Context, req *DeleteMerchantReq) (res *Response, err error)
	GetMerchantInfo(ctx context.Context, req *GetMerchantInfoReq) (res *GetMerchantInfoResp, err error)
	ListMerchants(ctx context.Context, req *ListMerchantsReq) (res *ListMerchantsResp, err error)
	GetMerchantTotalData(ctx context.Context, req *GetMerchantTotalDataReq) (res *GetMerchantTotalDataResp, err error)
	ListMerchantByMerchantId(ctx context.Context, req *ListMerchantsByMerchantIdReq) (res *ListMerchantsByMerchantIdResp, err error)
	ListMerchantByActivityNumber(ctx context.Context, req *ListMerchantsByActivityNumberReq) (res *ListMerchantsByActivityNumberResp, err error)
	ListActivityByActivityId(ctx context.Context, req *ListActivitiesByActivityIdReq) (res *ListActivitiesByActivityIdResp, err error)
	StsSendVerifyCode(ctx context.Context, req *StsSendVerifyCodeReq) (res *Response, err error)
	StsCheckVerifyCode(ctx context.Context, req *StsCheckVerifyCodeReq) (res *Response, err error)
}
