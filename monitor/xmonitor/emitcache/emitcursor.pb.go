// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: monitor/xmonitor/emitcache/emitcursor.proto

package emitcache

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

// EmitCursor stores historical emit cursors for all evm portals.
type EmitCursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                     // Auto-incremented ID
	SrcChainId uint64 `protobuf:"varint,2,opt,name=src_chain_id,json=srcChainId,proto3" json:"src_chain_id,omitempty"` // Source chain ID as per https://chainlist.org
	Height     uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`                             // Block Height
	DstChainId uint64 `protobuf:"varint,4,opt,name=dst_chain_id,json=dstChainId,proto3" json:"dst_chain_id,omitempty"` // Destination Chain ID as per https://chainlist.org
	ShardId    uint64 `protobuf:"varint,5,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`            // Shard ID
	MsgOffset  uint64 `protobuf:"varint,6,opt,name=msg_offset,json=msgOffset,proto3" json:"msg_offset,omitempty"`      // XMsg Stream Offset
}

func (x *EmitCursor) Reset() {
	*x = EmitCursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_xmonitor_emitcache_emitcursor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmitCursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmitCursor) ProtoMessage() {}

func (x *EmitCursor) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_xmonitor_emitcache_emitcursor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmitCursor.ProtoReflect.Descriptor instead.
func (*EmitCursor) Descriptor() ([]byte, []int) {
	return file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescGZIP(), []int{0}
}

func (x *EmitCursor) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmitCursor) GetSrcChainId() uint64 {
	if x != nil {
		return x.SrcChainId
	}
	return 0
}

func (x *EmitCursor) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *EmitCursor) GetDstChainId() uint64 {
	if x != nil {
		return x.DstChainId
	}
	return 0
}

func (x *EmitCursor) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *EmitCursor) GetMsgOffset() uint64 {
	if x != nil {
		return x.MsgOffset
	}
	return 0
}

var File_monitor_xmonitor_emitcache_emitcursor_proto protoreflect.FileDescriptor

var file_monitor_xmonitor_emitcache_emitcursor_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2f, 0x65, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x65, 0x6d, 0x69,
	0x74, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x65, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x0a, 0x45, 0x6d, 0x69, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x72, 0x63, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x64,
	0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x64, 0x73, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x73,
	0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3a, 0x41, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x3b, 0x0a,
	0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x2f, 0x0a, 0x29, 0x73, 0x72, 0x63, 0x5f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x64, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x2c, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x10, 0x02, 0x18, 0x01, 0x18, 0x01, 0x42, 0xf4, 0x01, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x78, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x42, 0x0f, 0x45,
	0x6d, 0x69, 0x74, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e,
	0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x78, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f,
	0x65, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x45, 0xaa,
	0x02, 0x1a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x45, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0xca, 0x02, 0x1a, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c,
	0x45, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0xe2, 0x02, 0x26, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x5c, 0x58, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5c, 0x45, 0x6d, 0x69,
	0x74, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x58, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x45, 0x6d, 0x69, 0x74, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescOnce sync.Once
	file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescData = file_monitor_xmonitor_emitcache_emitcursor_proto_rawDesc
)

func file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescGZIP() []byte {
	file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescOnce.Do(func() {
		file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescData)
	})
	return file_monitor_xmonitor_emitcache_emitcursor_proto_rawDescData
}

var file_monitor_xmonitor_emitcache_emitcursor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_monitor_xmonitor_emitcache_emitcursor_proto_goTypes = []any{
	(*EmitCursor)(nil), // 0: monitor.xmonitor.emitcache.EmitCursor
}
var file_monitor_xmonitor_emitcache_emitcursor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitor_xmonitor_emitcache_emitcursor_proto_init() }
func file_monitor_xmonitor_emitcache_emitcursor_proto_init() {
	if File_monitor_xmonitor_emitcache_emitcursor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_xmonitor_emitcache_emitcursor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmitCursor); i {
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
			RawDescriptor: file_monitor_xmonitor_emitcache_emitcursor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_monitor_xmonitor_emitcache_emitcursor_proto_goTypes,
		DependencyIndexes: file_monitor_xmonitor_emitcache_emitcursor_proto_depIdxs,
		MessageInfos:      file_monitor_xmonitor_emitcache_emitcursor_proto_msgTypes,
	}.Build()
	File_monitor_xmonitor_emitcache_emitcursor_proto = out.File
	file_monitor_xmonitor_emitcache_emitcursor_proto_rawDesc = nil
	file_monitor_xmonitor_emitcache_emitcursor_proto_goTypes = nil
	file_monitor_xmonitor_emitcache_emitcursor_proto_depIdxs = nil
}
