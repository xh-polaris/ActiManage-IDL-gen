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
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x08, 0x0a, 0x0d, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x47, 0x0a, 0x13, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x54, 0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x41, 0x63,
	0x74, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d, 0x49, 0x44, 0x4c, 0x2d, 0x67, 0x65, 0x6e,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ActiManage_system_system_proto_goTypes = []interface{}{
	(*MerchantLoginReq)(nil),       // 0: system.MerchantLoginReq
	(*MerchantSetPasswordReq)(nil), // 1: system.MerchantSetPasswordReq
	(*UpdateHeaderReq)(nil),        // 2: system.UpdateHeaderReq
	(*UpdateCoverReq)(nil),         // 3: system.UpdateCoverReq
	(*UpdateFooterReq)(nil),        // 4: system.UpdateFooterReq
	(*CreateActivityReq)(nil),      // 5: system.CreateActivityReq
	(*TopActivityReq)(nil),         // 6: system.TopActivityReq
	(*DeleteActivityReq)(nil),      // 7: system.DeleteActivityReq
	(*UpdateActivityReq)(nil),      // 8: system.UpdateActivityReq
	(*ListActivitiesReq)(nil),      // 9: system.ListActivitiesReq
	(*GetActivityReq)(nil),         // 10: system.GetActivityReq
	(*AdminLoginReq)(nil),          // 11: system.AdminLoginReq
	(*AdminSetPasswordReq)(nil),    // 12: system.AdminSetPasswordReq
	(*CreateMerchantReq)(nil),      // 13: system.CreateMerchantReq
	(*UpdateMerchantReq)(nil),      // 14: system.UpdateMerchantReq
	(*DeleteMerchantReq)(nil),      // 15: system.DeleteMerchantReq
	(*MerchantLoginResp)(nil),      // 16: system.MerchantLoginResp
	(*Response)(nil),               // 17: system.Response
	(*ListActivitiesResp)(nil),     // 18: system.ListActivitiesResp
	(*GetActivityResp)(nil),        // 19: system.GetActivityResp
	(*AdminLoginResp)(nil),         // 20: system.AdminLoginResp
}
var file_ActiManage_system_system_proto_depIdxs = []int32{
	0,  // 0: system.SystemService.MerchantLogin:input_type -> system.MerchantLoginReq
	1,  // 1: system.SystemService.MerchantSetPassword:input_type -> system.MerchantSetPasswordReq
	2,  // 2: system.SystemService.UpdateHeader:input_type -> system.UpdateHeaderReq
	3,  // 3: system.SystemService.UpdateCover:input_type -> system.UpdateCoverReq
	4,  // 4: system.SystemService.UpdateFooter:input_type -> system.UpdateFooterReq
	5,  // 5: system.SystemService.CreateActivity:input_type -> system.CreateActivityReq
	6,  // 6: system.SystemService.TopActivity:input_type -> system.TopActivityReq
	7,  // 7: system.SystemService.DeleteActivity:input_type -> system.DeleteActivityReq
	8,  // 8: system.SystemService.UpdateActivity:input_type -> system.UpdateActivityReq
	9,  // 9: system.SystemService.ListActivities:input_type -> system.ListActivitiesReq
	10, // 10: system.SystemService.GetActivity:input_type -> system.GetActivityReq
	11, // 11: system.SystemService.AdminLogin:input_type -> system.AdminLoginReq
	12, // 12: system.SystemService.AdminSetPassword:input_type -> system.AdminSetPasswordReq
	13, // 13: system.SystemService.CreateMerchant:input_type -> system.CreateMerchantReq
	14, // 14: system.SystemService.UpdateMerchant:input_type -> system.UpdateMerchantReq
	15, // 15: system.SystemService.DeleteMerchant:input_type -> system.DeleteMerchantReq
	16, // 16: system.SystemService.MerchantLogin:output_type -> system.MerchantLoginResp
	17, // 17: system.SystemService.MerchantSetPassword:output_type -> system.Response
	17, // 18: system.SystemService.UpdateHeader:output_type -> system.Response
	17, // 19: system.SystemService.UpdateCover:output_type -> system.Response
	17, // 20: system.SystemService.UpdateFooter:output_type -> system.Response
	17, // 21: system.SystemService.CreateActivity:output_type -> system.Response
	17, // 22: system.SystemService.TopActivity:output_type -> system.Response
	17, // 23: system.SystemService.DeleteActivity:output_type -> system.Response
	17, // 24: system.SystemService.UpdateActivity:output_type -> system.Response
	18, // 25: system.SystemService.ListActivities:output_type -> system.ListActivitiesResp
	19, // 26: system.SystemService.GetActivity:output_type -> system.GetActivityResp
	20, // 27: system.SystemService.AdminLogin:output_type -> system.AdminLoginResp
	17, // 28: system.SystemService.AdminSetPassword:output_type -> system.Response
	17, // 29: system.SystemService.CreateMerchant:output_type -> system.Response
	17, // 30: system.SystemService.UpdateMerchant:output_type -> system.Response
	17, // 31: system.SystemService.DeleteMerchant:output_type -> system.Response
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
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

// Code generated by Kitex v0.12.1. DO NOT EDIT.

type SystemService interface {
	MerchantLogin(ctx context.Context, req *MerchantLoginReq) (res *MerchantLoginResp, err error)
	MerchantSetPassword(ctx context.Context, req *MerchantSetPasswordReq) (res *Response, err error)
	UpdateHeader(ctx context.Context, req *UpdateHeaderReq) (res *Response, err error)
	UpdateCover(ctx context.Context, req *UpdateCoverReq) (res *Response, err error)
	UpdateFooter(ctx context.Context, req *UpdateFooterReq) (res *Response, err error)
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
}
