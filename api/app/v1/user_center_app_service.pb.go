// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: user_center_app_service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_center_app_service_proto protoreflect.FileDescriptor

var file_user_center_app_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x70, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd4, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x50,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x12, 0x73, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1b, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42, 0x79, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x5a, 0x1c,
	0x12, 0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x13, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x12, 0x4e, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x12, 0x12, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x12, 0x5a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x15,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x1a, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x57, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x2a, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_center_app_service_proto_goTypes = []interface{}{
	(*PageAppRequest)(nil),         // 0: app.PageAppRequest
	(*GetAppByKeywordRequest)(nil), // 1: app.GetAppByKeywordRequest
	(*AddAppRequest)(nil),          // 2: app.AddAppRequest
	(*UpdateAppRequest)(nil),       // 3: app.UpdateAppRequest
	(*DeleteAppRequest)(nil),       // 4: app.DeleteAppRequest
	(*PageAppReply)(nil),           // 5: app.PageAppReply
	(*App)(nil),                    // 6: app.App
	(*AddAppReply)(nil),            // 7: app.AddAppReply
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_user_center_app_service_proto_depIdxs = []int32{
	0, // 0: app.Service.PageApp:input_type -> app.PageAppRequest
	1, // 1: app.Service.GetAppByKeyword:input_type -> app.GetAppByKeywordRequest
	2, // 2: app.Service.AddApp:input_type -> app.AddAppRequest
	3, // 3: app.Service.UpdateApp:input_type -> app.UpdateAppRequest
	4, // 4: app.Service.DeleteApp:input_type -> app.DeleteAppRequest
	5, // 5: app.Service.PageApp:output_type -> app.PageAppReply
	6, // 6: app.Service.GetAppByKeyword:output_type -> app.App
	7, // 7: app.Service.AddApp:output_type -> app.AddAppReply
	8, // 8: app.Service.UpdateApp:output_type -> google.protobuf.Empty
	8, // 9: app.Service.DeleteApp:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_center_app_service_proto_init() }
func file_user_center_app_service_proto_init() {
	if File_user_center_app_service_proto != nil {
		return
	}
	file_user_center_app_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_center_app_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_center_app_service_proto_goTypes,
		DependencyIndexes: file_user_center_app_service_proto_depIdxs,
	}.Build()
	File_user_center_app_service_proto = out.File
	file_user_center_app_service_proto_rawDesc = nil
	file_user_center_app_service_proto_goTypes = nil
	file_user_center_app_service_proto_depIdxs = nil
}
