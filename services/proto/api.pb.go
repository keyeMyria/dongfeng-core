// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-core/services/proto/api.proto

package dongfeng_svc_core_server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ilovelili/dongfeng-protobuf"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/ilovelili/dongfeng-core/services/proto/api.proto", fileDescriptor_704a4f15475a1c86)
}

var fileDescriptor_704a4f15475a1c86 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x05, 0x51, 0xf4, 0x29, 0xfe, 0xc8, 0x49, 0x76, 0x51, 0xd4, 0x4d, 0xe7, 0x5c, 0x0a,
	0x8a, 0x27, 0x8f, 0x0a, 0xbb, 0x8c, 0x1d, 0xd4, 0x81, 0x20, 0x08, 0x69, 0xf7, 0xda, 0x45, 0xba,
	0xa6, 0x26, 0x69, 0xff, 0x61, 0xff, 0x11, 0x59, 0x6b, 0xb2, 0x6e, 0xb6, 0x6b, 0x6f, 0x25, 0xef,
	0xf3, 0xbe, 0x9f, 0x2f, 0x69, 0xe0, 0x31, 0xe0, 0x7a, 0x9a, 0xb8, 0xd4, 0x13, 0x33, 0x87, 0x87,
	0x22, 0xc5, 0x90, 0x87, 0xdc, 0x99, 0x88, 0x28, 0xf0, 0x31, 0x0a, 0xfa, 0x9e, 0x90, 0xe8, 0x28,
	0x94, 0x29, 0xf7, 0x50, 0x39, 0xb1, 0x14, 0x5a, 0x38, 0x2c, 0xe6, 0x34, 0xfb, 0x22, 0x27, 0x86,
	0xa3, 0x2a, 0xf5, 0xe8, 0x9c, 0xa5, 0x73, 0x16, 0x65, 0xeb, 0x61, 0x7d, 0x6c, 0xb6, 0xee, 0x26,
	0xbe, 0x3d, 0xc9, 0x03, 0xef, 0x7e, 0x76, 0x60, 0x93, 0xc5, 0x9c, 0x0c, 0x61, 0x6b, 0x28, 0x02,
	0x1e, 0x91, 0x53, 0xba, 0x4c, 0xb8, 0x89, 0x4f, 0xb3, 0xc9, 0x0b, 0x7e, 0x27, 0xa8, 0x74, 0xeb,
	0xac, 0x1a, 0x50, 0xb1, 0x88, 0x14, 0x9e, 0x6f, 0x90, 0x77, 0xd8, 0x7d, 0x66, 0x6a, 0xea, 0x0a,
	0x26, 0x27, 0xe4, 0xa2, 0x64, 0xc1, 0x4e, 0x4d, 0xea, 0xe5, 0x7a, 0xc8, 0x26, 0x7f, 0x00, 0x8c,
	0xe3, 0x09, 0xd3, 0x38, 0x56, 0x28, 0x49, 0xd9, 0xd6, 0x62, 0x6c, 0xb2, 0xdb, 0x35, 0x94, 0x0d,
	0x4f, 0x80, 0xe4, 0xe7, 0x23, 0xa1, 0xb9, 0xcf, 0x3d, 0xa6, 0xb9, 0x88, 0x48, 0xbf, 0x72, 0xbd,
	0x88, 0x29, 0x63, 0xa3, 0x4d, 0x71, 0xab, 0x65, 0xb0, 0x3f, 0x40, 0xfd, 0x14, 0x32, 0xa5, 0x42,
	0xae, 0x34, 0xe9, 0x94, 0x24, 0x14, 0x01, 0x63, 0xba, 0xaa, 0xe5, 0xac, 0xe2, 0x0b, 0x0e, 0xf3,
	0x0e, 0x0b, 0x4b, 0xb7, 0xb2, 0xe7, 0x3f, 0xd1, 0x4d, 0x13, 0xd4, 0xba, 0x3e, 0x61, 0x6f, 0x80,
	0x7a, 0xc4, 0x66, 0x98, 0x79, 0xda, 0xe5, 0x2d, 0xcd, 0xdc, 0x38, 0x3a, 0x75, 0x98, 0xcd, 0x0f,
	0xe0, 0xe0, 0xef, 0x3e, 0x8d, 0xe2, 0xba, 0xfa, 0xca, 0x57, 0x2c, 0xdd, 0x06, 0x64, 0x51, 0x34,
	0x40, 0xfd, 0x86, 0xcc, 0x9b, 0xa2, 0xac, 0x14, 0x2d, 0x23, 0xeb, 0x44, 0xab, 0xa4, 0x15, 0xc5,
	0x70, 0x9c, 0x97, 0x28, 0xba, 0x7a, 0x95, 0x55, 0x4b, 0x74, 0xb7, 0xcd, 0x60, 0x6b, 0x9c, 0xc1,
	0xd1, 0x2b, 0x4b, 0x97, 0xdf, 0x79, 0xd9, 0x5f, 0x5e, 0x85, 0x8c, 0xaf, 0xd7, 0x88, 0x35, 0x3a,
	0x77, 0x3b, 0x83, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0x82, 0x8a, 0x96, 0xfc, 0x04,
	0x00, 0x00,
}
