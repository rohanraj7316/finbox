// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: protos/v1/trading/broadcast.proto

package broadcast

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

type BroadcastServiceReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// brodcast's symbol
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *BroadcastServiceReadRequest) Reset() {
	*x = BroadcastServiceReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_trading_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastServiceReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastServiceReadRequest) ProtoMessage() {}

func (x *BroadcastServiceReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_trading_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastServiceReadRequest.ProtoReflect.Descriptor instead.
func (*BroadcastServiceReadRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_trading_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastServiceReadRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type Ticker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol     string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	ClosePrice string `protobuf:"bytes,2,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"`
	Timestamp  string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Ticker) Reset() {
	*x = Ticker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_trading_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticker) ProtoMessage() {}

func (x *Ticker) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_trading_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticker.ProtoReflect.Descriptor instead.
func (*Ticker) Descriptor() ([]byte, []int) {
	return file_protos_v1_trading_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *Ticker) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Ticker) GetClosePrice() string {
	if x != nil {
		return x.ClosePrice
	}
	return ""
}

func (x *Ticker) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type BroadcastServiceReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// blog's unique uuid. which gonna be used across all other systems
	Tickers []*Ticker `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty"`
}

func (x *BroadcastServiceReadResponse) Reset() {
	*x = BroadcastServiceReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_trading_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastServiceReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastServiceReadResponse) ProtoMessage() {}

func (x *BroadcastServiceReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_trading_broadcast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastServiceReadResponse.ProtoReflect.Descriptor instead.
func (*BroadcastServiceReadResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_trading_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastServiceReadResponse) GetTickers() []*Ticker {
	if x != nil {
		return x.Tickers
	}
	return nil
}

var File_protos_v1_trading_broadcast_proto protoreflect.FileDescriptor

var file_protos_v1_trading_broadcast_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x1b, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x04, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x22, 0x5f, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x4b, 0x0a, 0x1c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73,
	0x32, 0x96, 0x01, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x28, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x42, 0x45, 0x54, 0x41, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x6f, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_trading_broadcast_proto_rawDescOnce sync.Once
	file_protos_v1_trading_broadcast_proto_rawDescData = file_protos_v1_trading_broadcast_proto_rawDesc
)

func file_protos_v1_trading_broadcast_proto_rawDescGZIP() []byte {
	file_protos_v1_trading_broadcast_proto_rawDescOnce.Do(func() {
		file_protos_v1_trading_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_trading_broadcast_proto_rawDescData)
	})
	return file_protos_v1_trading_broadcast_proto_rawDescData
}

var file_protos_v1_trading_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_v1_trading_broadcast_proto_goTypes = []interface{}{
	(*BroadcastServiceReadRequest)(nil),  // 0: protos.v1.BroadcastServiceReadRequest
	(*Ticker)(nil),                       // 1: protos.v1.Ticker
	(*BroadcastServiceReadResponse)(nil), // 2: protos.v1.BroadcastServiceReadResponse
}
var file_protos_v1_trading_broadcast_proto_depIdxs = []int32{
	1, // 0: protos.v1.BroadcastServiceReadResponse.tickers:type_name -> protos.v1.Ticker
	0, // 1: protos.v1.BroadcastService.Read:input_type -> protos.v1.BroadcastServiceReadRequest
	2, // 2: protos.v1.BroadcastService.Read:output_type -> protos.v1.BroadcastServiceReadResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_v1_trading_broadcast_proto_init() }
func file_protos_v1_trading_broadcast_proto_init() {
	if File_protos_v1_trading_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_trading_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastServiceReadRequest); i {
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
		file_protos_v1_trading_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticker); i {
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
		file_protos_v1_trading_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastServiceReadResponse); i {
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
			RawDescriptor: file_protos_v1_trading_broadcast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_trading_broadcast_proto_goTypes,
		DependencyIndexes: file_protos_v1_trading_broadcast_proto_depIdxs,
		MessageInfos:      file_protos_v1_trading_broadcast_proto_msgTypes,
	}.Build()
	File_protos_v1_trading_broadcast_proto = out.File
	file_protos_v1_trading_broadcast_proto_rawDesc = nil
	file_protos_v1_trading_broadcast_proto_goTypes = nil
	file_protos_v1_trading_broadcast_proto_depIdxs = nil
}