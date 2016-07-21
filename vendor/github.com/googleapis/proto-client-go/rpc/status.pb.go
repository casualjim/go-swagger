// Code generated by protoc-gen-go.
// source: google/rpc/status.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	google/rpc/status.proto

It has these top-level messages:
	Status
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The `Status` type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is used by
// [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error message,
// and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes if needed.  The
// error message should be a developer-facing English message that helps
// developers *understand* and *resolve* the error. If a localized user-facing
// error message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain arbitrary
// information about the error. There is a predefined set of error detail types
// in the package `google.rpc` which can be used for common error conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*google_protobuf.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

func init() { proto.RegisterFile("google/rpc/status.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8f, 0x31, 0xaf, 0x82, 0x30,
	0x14, 0x85, 0xc3, 0xe3, 0x3d, 0xc8, 0x2b, 0x89, 0x43, 0x63, 0x62, 0x75, 0x22, 0x4e, 0x38, 0xd0,
	0x26, 0xfa, 0x0b, 0x64, 0x74, 0x22, 0xb8, 0xb9, 0x95, 0x52, 0x6a, 0x13, 0xa0, 0x84, 0x96, 0x81,
	0x7f, 0x2f, 0xb4, 0x10, 0xb7, 0x7b, 0xcf, 0xf9, 0x6e, 0xce, 0xb9, 0xe0, 0x20, 0x94, 0x12, 0x0d,
	0x27, 0x43, 0xcf, 0x88, 0x36, 0xd4, 0x8c, 0x1a, 0xf7, 0x83, 0x32, 0x0a, 0x02, 0x67, 0xe0, 0xd9,
	0x38, 0x1d, 0x57, 0xc8, 0x3a, 0xe5, 0x58, 0x13, 0xda, 0x4d, 0x0e, 0x3b, 0xd7, 0x20, 0x78, 0xda,
	0x33, 0x08, 0xc1, 0x2f, 0x53, 0x15, 0x47, 0x5e, 0xec, 0x25, 0x7f, 0x85, 0x9d, 0x21, 0x02, 0x61,
	0xcb, 0xb5, 0xa6, 0x82, 0xa3, 0x9f, 0x59, 0xfe, 0x2f, 0xb6, 0x15, 0x62, 0x10, 0x56, 0xdc, 0x50,
	0xd9, 0x68, 0xe4, 0xc7, 0x7e, 0x12, 0x5d, 0xf7, 0x78, 0x0d, 0xdc, 0x42, 0xf0, 0xbd, 0x9b, 0x8a,
	0x0d, 0xca, 0x1e, 0x60, 0xc7, 0x54, 0x8b, 0xbf, 0xa5, 0xb2, 0xc8, 0xe5, 0xe6, 0x0b, 0x9e, 0x7b,
	0xaf, 0x8b, 0x90, 0xe6, 0x3d, 0x96, 0x78, 0xa6, 0x88, 0xa3, 0x68, 0x2f, 0xb5, 0xab, 0x9c, 0xb2,
	0x46, 0xf2, 0xce, 0xa4, 0x42, 0x2d, 0x7f, 0x96, 0x81, 0x15, 0x6f, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x62, 0xd7, 0xe1, 0xfc, 0x00, 0x00, 0x00,
}
