// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/http/stateful_session/header/v3/header.proto

package headerv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// This extension allows the session state to be tracked via request headers.
//
// This extension encodes the address of the upstream host selected by the load balancer
// into a response header with the :ref:`header configuration
// <envoy_v3_api_msg_extensions.http.stateful_session.header.v3.HeaderBasedSessionState>`.
// When new requests are incoming, this extension will try to parse the specific upstream host
// address by header name. If the address parsed from the header corresponds to a valid
// upstream host, this upstream host will be selected first. See :ref:`stateful session filter
// <envoy_v3_api_msg_extensions.filters.http.stateful_session.v3.StatefulSession>`.
//
// For example, if the header name is set to “session-header“, envoy will prefer “1.2.3.4:80“
// as the upstream host when the request contains the following header:
//
// .. code-block:: none
//
//	session-header: "MS4yLjMuNDo4MA=="
//
// When processing the upstream response, if “1.2.3.4:80“ is indeed the final choice the extension
// does nothing. If “1.2.3.4:80“ is not the final choice, the new selected host will be set to
// response headers (via the “session-header“ response header).
//
// [#extension: envoy.http.stateful_session.header]
type HeaderBasedSessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name that will be used to obtain header value from downstream HTTP request or generate
	// new header for downstream.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HeaderBasedSessionState) Reset() {
	*x = HeaderBasedSessionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_stateful_session_header_v3_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderBasedSessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderBasedSessionState) ProtoMessage() {}

func (x *HeaderBasedSessionState) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_stateful_session_header_v3_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderBasedSessionState.ProtoReflect.Descriptor instead.
func (*HeaderBasedSessionState) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderBasedSessionState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_envoy_extensions_http_stateful_session_header_v3_header_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x17, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0xb9, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3e, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x33, 0x3b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescData = file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDesc
)

func file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescData)
	})
	return file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDescData
}

var file_envoy_extensions_http_stateful_session_header_v3_header_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_stateful_session_header_v3_header_proto_goTypes = []interface{}{
	(*HeaderBasedSessionState)(nil), // 0: envoy.extensions.http.stateful_session.header.v3.HeaderBasedSessionState
}
var file_envoy_extensions_http_stateful_session_header_v3_header_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_http_stateful_session_header_v3_header_proto_init() }
func file_envoy_extensions_http_stateful_session_header_v3_header_proto_init() {
	if File_envoy_extensions_http_stateful_session_header_v3_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_stateful_session_header_v3_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderBasedSessionState); i {
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
			RawDescriptor: file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_stateful_session_header_v3_header_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_stateful_session_header_v3_header_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_stateful_session_header_v3_header_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_stateful_session_header_v3_header_proto = out.File
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_rawDesc = nil
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_goTypes = nil
	file_envoy_extensions_http_stateful_session_header_v3_header_proto_depIdxs = nil
}
