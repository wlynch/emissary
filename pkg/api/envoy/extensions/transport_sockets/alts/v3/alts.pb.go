// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/transport_sockets/alts/v3/alts.proto

package altsv3

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

// Configuration for ALTS transport socket. This provides Google's ALTS protocol to Envoy.
// Store the peer identity in dynamic metadata, namespace is "envoy.transport_socket.peer_information", key is "peer_identity".
// https://cloud.google.com/security/encryption-in-transit/application-layer-transport-security/
type Alts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location of a handshaker service, this is usually 169.254.169.254:8080
	// on GCE.
	HandshakerService string `protobuf:"bytes,1,opt,name=handshaker_service,json=handshakerService,proto3" json:"handshaker_service,omitempty"`
	// The acceptable service accounts from peer, peers not in the list will be rejected in the
	// handshake validation step. If empty, no validation will be performed.
	PeerServiceAccounts []string `protobuf:"bytes,2,rep,name=peer_service_accounts,json=peerServiceAccounts,proto3" json:"peer_service_accounts,omitempty"`
}

func (x *Alts) Reset() {
	*x = Alts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_alts_v3_alts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alts) ProtoMessage() {}

func (x *Alts) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_alts_v3_alts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alts.ProtoReflect.Descriptor instead.
func (*Alts) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescGZIP(), []int{0}
}

func (x *Alts) GetHandshakerService() string {
	if x != nil {
		return x.HandshakerService
	}
	return ""
}

func (x *Alts) GetPeerServiceAccounts() []string {
	if x != nil {
		return x.PeerServiceAccounts
	}
	return nil
}

var File_envoy_extensions_transport_sockets_alts_v3_alts_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x6c, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x6c, 0x74, 0x73,
	0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa,
	0x01, 0x0a, 0x04, 0x41, 0x6c, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x12, 0x68, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13,
	0x70, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x3a, 0x36, 0x9a, 0xc5, 0x88, 0x1e, 0x31, 0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x61, 0x6c, 0x74, 0x73, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x6c, 0x74, 0x73, 0x42, 0xa9, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x61, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x33,
	0x42, 0x09, 0x41, 0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x6c, 0x74, 0x73, 0x2f, 0x76, 0x33,
	0x3b, 0x61, 0x6c, 0x74, 0x73, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescData = file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_alts_v3_alts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_transport_sockets_alts_v3_alts_proto_goTypes = []interface{}{
	(*Alts)(nil), // 0: envoy.extensions.transport_sockets.alts.v3.Alts
}
var file_envoy_extensions_transport_sockets_alts_v3_alts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_alts_v3_alts_proto_init() }
func file_envoy_extensions_transport_sockets_alts_v3_alts_proto_init() {
	if File_envoy_extensions_transport_sockets_alts_v3_alts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_alts_v3_alts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alts); i {
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
			RawDescriptor: file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_alts_v3_alts_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_alts_v3_alts_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_alts_v3_alts_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_alts_v3_alts_proto = out.File
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_alts_v3_alts_proto_depIdxs = nil
}
