// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: hobby_service.proto

package user_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

var File_hobby_service_proto protoreflect.FileDescriptor

var file_hobby_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x03,
	0x0a, 0x0c, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x62, 0x62, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x62,
	0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x62, 0x62, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x62, 0x62, 0x79,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hobby_service_proto_goTypes = []interface{}{
	(*CreateHobby)(nil),          // 0: user_service.CreateHobby
	(*HobbyPrimaryKey)(nil),      // 1: user_service.HobbyPrimaryKey
	(*GetListHobbyRequest)(nil),  // 2: user_service.GetListHobbyRequest
	(*UpdateHobby)(nil),          // 3: user_service.UpdateHobby
	(*UpdatePatchHobby)(nil),     // 4: user_service.UpdatePatchHobby
	(*Hobby)(nil),                // 5: user_service.Hobby
	(*GetListHobbyResponse)(nil), // 6: user_service.GetListHobbyResponse
	(*empty.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_hobby_service_proto_depIdxs = []int32{
	0, // 0: user_service.HobbyService.Create:input_type -> user_service.CreateHobby
	1, // 1: user_service.HobbyService.GetByID:input_type -> user_service.HobbyPrimaryKey
	2, // 2: user_service.HobbyService.GetList:input_type -> user_service.GetListHobbyRequest
	3, // 3: user_service.HobbyService.Update:input_type -> user_service.UpdateHobby
	4, // 4: user_service.HobbyService.UpdatePatch:input_type -> user_service.UpdatePatchHobby
	1, // 5: user_service.HobbyService.Delete:input_type -> user_service.HobbyPrimaryKey
	5, // 6: user_service.HobbyService.Create:output_type -> user_service.Hobby
	5, // 7: user_service.HobbyService.GetByID:output_type -> user_service.Hobby
	6, // 8: user_service.HobbyService.GetList:output_type -> user_service.GetListHobbyResponse
	5, // 9: user_service.HobbyService.Update:output_type -> user_service.Hobby
	5, // 10: user_service.HobbyService.UpdatePatch:output_type -> user_service.Hobby
	7, // 11: user_service.HobbyService.Delete:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hobby_service_proto_init() }
func file_hobby_service_proto_init() {
	if File_hobby_service_proto != nil {
		return
	}
	file_hobby_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hobby_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hobby_service_proto_goTypes,
		DependencyIndexes: file_hobby_service_proto_depIdxs,
	}.Build()
	File_hobby_service_proto = out.File
	file_hobby_service_proto_rawDesc = nil
	file_hobby_service_proto_goTypes = nil
	file_hobby_service_proto_depIdxs = nil
}
