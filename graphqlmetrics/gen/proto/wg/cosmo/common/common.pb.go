// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: wg/cosmo/common/common.proto

package common

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

type EnumStatusCode int32

const (
	EnumStatusCode_OK                              EnumStatusCode = 0
	EnumStatusCode_ERR                             EnumStatusCode = 1
	EnumStatusCode_ERR_NOT_FOUND                   EnumStatusCode = 2
	EnumStatusCode_ERR_ALREADY_EXISTS              EnumStatusCode = 3
	EnumStatusCode_ERR_INVALID_SUBGRAPH_SCHEMA     EnumStatusCode = 4
	EnumStatusCode_ERR_SUBGRAPH_COMPOSITION_FAILED EnumStatusCode = 5
	EnumStatusCode_ERR_SUBGRAPH_CHECK_FAILED       EnumStatusCode = 6
	EnumStatusCode_ERR_INVALID_LABELS              EnumStatusCode = 7
	EnumStatusCode_ERR_ANALYTICS_DISABLED          EnumStatusCode = 8
	EnumStatusCode_ERROR_NOT_AUTHENTICATED         EnumStatusCode = 9
	EnumStatusCode_ERR_OPENAI_DISABLED             EnumStatusCode = 10
	EnumStatusCode_ERR_FREE_TRIAL_EXPIRED          EnumStatusCode = 11
)

// Enum value maps for EnumStatusCode.
var (
	EnumStatusCode_name = map[int32]string{
		0:  "OK",
		1:  "ERR",
		2:  "ERR_NOT_FOUND",
		3:  "ERR_ALREADY_EXISTS",
		4:  "ERR_INVALID_SUBGRAPH_SCHEMA",
		5:  "ERR_SUBGRAPH_COMPOSITION_FAILED",
		6:  "ERR_SUBGRAPH_CHECK_FAILED",
		7:  "ERR_INVALID_LABELS",
		8:  "ERR_ANALYTICS_DISABLED",
		9:  "ERROR_NOT_AUTHENTICATED",
		10: "ERR_OPENAI_DISABLED",
		11: "ERR_FREE_TRIAL_EXPIRED",
	}
	EnumStatusCode_value = map[string]int32{
		"OK":                              0,
		"ERR":                             1,
		"ERR_NOT_FOUND":                   2,
		"ERR_ALREADY_EXISTS":              3,
		"ERR_INVALID_SUBGRAPH_SCHEMA":     4,
		"ERR_SUBGRAPH_COMPOSITION_FAILED": 5,
		"ERR_SUBGRAPH_CHECK_FAILED":       6,
		"ERR_INVALID_LABELS":              7,
		"ERR_ANALYTICS_DISABLED":          8,
		"ERROR_NOT_AUTHENTICATED":         9,
		"ERR_OPENAI_DISABLED":             10,
		"ERR_FREE_TRIAL_EXPIRED":          11,
	}
)

func (x EnumStatusCode) Enum() *EnumStatusCode {
	p := new(EnumStatusCode)
	*p = x
	return p
}

func (x EnumStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_wg_cosmo_common_common_proto_enumTypes[0].Descriptor()
}

func (EnumStatusCode) Type() protoreflect.EnumType {
	return &file_wg_cosmo_common_common_proto_enumTypes[0]
}

func (x EnumStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumStatusCode.Descriptor instead.
func (EnumStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_wg_cosmo_common_common_proto_rawDescGZIP(), []int{0}
}

var File_wg_cosmo_common_common_proto protoreflect.FileDescriptor

var file_wg_cosmo_common_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a,
	0xb7, 0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x52,
	0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x55,
	0x42, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x04, 0x12,
	0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x47,
	0x52, 0x41, 0x50, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x52, 0x52, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x50, 0x45, 0x4e,
	0x41, 0x49, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0b, 0x42, 0xc7, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x6e,
	0x64, 0x65, 0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x57, 0x43, 0x43, 0xaa, 0x02, 0x0f,
	0x57, 0x67, 0x2e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca,
	0x02, 0x0f, 0x57, 0x67, 0x5c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xe2, 0x02, 0x1b, 0x57, 0x67, 0x5c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x57, 0x67, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wg_cosmo_common_common_proto_rawDescOnce sync.Once
	file_wg_cosmo_common_common_proto_rawDescData = file_wg_cosmo_common_common_proto_rawDesc
)

func file_wg_cosmo_common_common_proto_rawDescGZIP() []byte {
	file_wg_cosmo_common_common_proto_rawDescOnce.Do(func() {
		file_wg_cosmo_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_wg_cosmo_common_common_proto_rawDescData)
	})
	return file_wg_cosmo_common_common_proto_rawDescData
}

var file_wg_cosmo_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wg_cosmo_common_common_proto_goTypes = []interface{}{
	(EnumStatusCode)(0), // 0: wg.cosmo.common.EnumStatusCode
}
var file_wg_cosmo_common_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wg_cosmo_common_common_proto_init() }
func file_wg_cosmo_common_common_proto_init() {
	if File_wg_cosmo_common_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wg_cosmo_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wg_cosmo_common_common_proto_goTypes,
		DependencyIndexes: file_wg_cosmo_common_common_proto_depIdxs,
		EnumInfos:         file_wg_cosmo_common_common_proto_enumTypes,
	}.Build()
	File_wg_cosmo_common_common_proto = out.File
	file_wg_cosmo_common_common_proto_rawDesc = nil
	file_wg_cosmo_common_common_proto_goTypes = nil
	file_wg_cosmo_common_common_proto_depIdxs = nil
}
