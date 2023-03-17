// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: blockchain.proto

package blockchain

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx       string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx" form:"tx"`
	From     string `protobuf:"bytes,2,opt,name=from,proto3" json:"from" form:"from"`
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to" form:"to"`
	Amount   string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount" form:"amount"`
	GasFee   string `protobuf:"bytes,5,opt,name=gas_fee,json=gasFee,proto3" json:"gas_fee" form:"gas_fee"`
	GasPrice string `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price" form:"gas_price"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetGasFee() string {
	if x != nil {
		return x.GasFee
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

// HTTP POST request: /blockchain/transaction/detail
type GetTransactionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx" form:"tx"`
}

func (x *GetTransactionReq) Reset() {
	*x = GetTransactionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionReq) ProtoMessage() {}

func (x *GetTransactionReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionReq.ProtoReflect.Descriptor instead.
func (*GetTransactionReq) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransactionReq) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

type GetTransactionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	Data    *Transaction `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetTransactionRsp) Reset() {
	*x = GetTransactionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRsp) ProtoMessage() {}

func (x *GetTransactionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRsp.ProtoReflect.Descriptor instead.
func (*GetTransactionRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTransactionRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTransactionRsp) GetData() *Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x8f,
	0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61,
	0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x73,
	0x46, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x78, 0x22, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x16, 0x5a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchain_proto_goTypes = []interface{}{
	(*Transaction)(nil),       // 0: blockchain.Transaction
	(*GetTransactionReq)(nil), // 1: blockchain.GetTransactionReq
	(*GetTransactionRsp)(nil), // 2: blockchain.GetTransactionRsp
}
var file_blockchain_proto_depIdxs = []int32{
	0, // 0: blockchain.GetTransactionRsp.data:type_name -> blockchain.Transaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionReq); i {
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
		file_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRsp); i {
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
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}