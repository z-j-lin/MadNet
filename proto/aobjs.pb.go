// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: aobjs.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Protobuf message implementation for struct Tx
type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vin  []*TXIn  `protobuf:"bytes,1,rep,name=Vin,proto3" json:"Vin,omitempty"`
	Vout []*TXOut `protobuf:"bytes,2,rep,name=Vout,proto3" json:"Vout,omitempty"`
	Fee  string   `protobuf:"bytes,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{0}
}

func (x *Tx) GetVin() []*TXIn {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *Tx) GetVout() []*TXOut {
	if x != nil {
		return x.Vout
	}
	return nil
}

func (x *Tx) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

// Protobuf message implementation for struct TXOut
type TXOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Utxo:
	//	*TXOut_AtomicSwap
	//	*TXOut_ValueStore
	//	*TXOut_DataStore
	Utxo isTXOut_Utxo `protobuf_oneof:"utxo"`
}

func (x *TXOut) Reset() {
	*x = TXOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXOut) ProtoMessage() {}

func (x *TXOut) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXOut.ProtoReflect.Descriptor instead.
func (*TXOut) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{1}
}

func (m *TXOut) GetUtxo() isTXOut_Utxo {
	if m != nil {
		return m.Utxo
	}
	return nil
}

func (x *TXOut) GetAtomicSwap() *AtomicSwap {
	if x, ok := x.GetUtxo().(*TXOut_AtomicSwap); ok {
		return x.AtomicSwap
	}
	return nil
}

func (x *TXOut) GetValueStore() *ValueStore {
	if x, ok := x.GetUtxo().(*TXOut_ValueStore); ok {
		return x.ValueStore
	}
	return nil
}

func (x *TXOut) GetDataStore() *DataStore {
	if x, ok := x.GetUtxo().(*TXOut_DataStore); ok {
		return x.DataStore
	}
	return nil
}

type isTXOut_Utxo interface {
	isTXOut_Utxo()
}

type TXOut_AtomicSwap struct {
	AtomicSwap *AtomicSwap `protobuf:"bytes,1,opt,name=AtomicSwap,proto3,oneof"`
}

type TXOut_ValueStore struct {
	ValueStore *ValueStore `protobuf:"bytes,2,opt,name=ValueStore,proto3,oneof"`
}

type TXOut_DataStore struct {
	DataStore *DataStore `protobuf:"bytes,3,opt,name=DataStore,proto3,oneof"`
}

func (*TXOut_AtomicSwap) isTXOut_Utxo() {}

func (*TXOut_ValueStore) isTXOut_Utxo() {}

func (*TXOut_DataStore) isTXOut_Utxo() {}

// Protobuf message implementation for struct TXIn
type TXIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TXInLinker *TXInLinker `protobuf:"bytes,1,opt,name=TXInLinker,proto3" json:"TXInLinker,omitempty"`
	Signature  string      `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *TXIn) Reset() {
	*x = TXIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXIn) ProtoMessage() {}

func (x *TXIn) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXIn.ProtoReflect.Descriptor instead.
func (*TXIn) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{2}
}

func (x *TXIn) GetTXInLinker() *TXInLinker {
	if x != nil {
		return x.TXInLinker
	}
	return nil
}

func (x *TXIn) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// Protobuf message implementation for struct TXInLinker
type TXInLinker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TXInPreImage *TXInPreImage `protobuf:"bytes,1,opt,name=TXInPreImage,proto3" json:"TXInPreImage,omitempty"`
	TxHash       string        `protobuf:"bytes,2,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *TXInLinker) Reset() {
	*x = TXInLinker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXInLinker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXInLinker) ProtoMessage() {}

func (x *TXInLinker) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXInLinker.ProtoReflect.Descriptor instead.
func (*TXInLinker) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{3}
}

func (x *TXInLinker) GetTXInPreImage() *TXInPreImage {
	if x != nil {
		return x.TXInPreImage
	}
	return nil
}

func (x *TXInLinker) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// Protobuf message implementation for struct TXInPreImage
type TXInPreImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID        uint32 `protobuf:"varint,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	ConsumedTxIdx  uint32 `protobuf:"varint,2,opt,name=ConsumedTxIdx,proto3" json:"ConsumedTxIdx,omitempty"`
	ConsumedTxHash string `protobuf:"bytes,3,opt,name=ConsumedTxHash,proto3" json:"ConsumedTxHash,omitempty"`
}

func (x *TXInPreImage) Reset() {
	*x = TXInPreImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXInPreImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXInPreImage) ProtoMessage() {}

func (x *TXInPreImage) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXInPreImage.ProtoReflect.Descriptor instead.
func (*TXInPreImage) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{4}
}

func (x *TXInPreImage) GetChainID() uint32 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *TXInPreImage) GetConsumedTxIdx() uint32 {
	if x != nil {
		return x.ConsumedTxIdx
	}
	return 0
}

func (x *TXInPreImage) GetConsumedTxHash() string {
	if x != nil {
		return x.ConsumedTxHash
	}
	return ""
}

// Protobuf message implementation for struct AtomicSwap
type AtomicSwap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ASPreImage *ASPreImage `protobuf:"bytes,1,opt,name=ASPreImage,proto3" json:"ASPreImage,omitempty"`
	TxHash     string      `protobuf:"bytes,2,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *AtomicSwap) Reset() {
	*x = AtomicSwap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicSwap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicSwap) ProtoMessage() {}

func (x *AtomicSwap) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicSwap.ProtoReflect.Descriptor instead.
func (*AtomicSwap) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{5}
}

func (x *AtomicSwap) GetASPreImage() *ASPreImage {
	if x != nil {
		return x.ASPreImage
	}
	return nil
}

func (x *AtomicSwap) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// Protobuf message implementation for struct ASPreImage
type ASPreImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID  uint32 `protobuf:"varint,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	TXOutIdx uint32 `protobuf:"varint,3,opt,name=TXOutIdx,proto3" json:"TXOutIdx,omitempty"`
	IssuedAt uint32 `protobuf:"varint,4,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Exp      uint32 `protobuf:"varint,5,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Owner    string `protobuf:"bytes,6,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Fee      string `protobuf:"bytes,7,opt,name=Fee,proto3" json:"Fee,omitempty"`
}

func (x *ASPreImage) Reset() {
	*x = ASPreImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ASPreImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ASPreImage) ProtoMessage() {}

func (x *ASPreImage) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ASPreImage.ProtoReflect.Descriptor instead.
func (*ASPreImage) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{6}
}

func (x *ASPreImage) GetChainID() uint32 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *ASPreImage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ASPreImage) GetTXOutIdx() uint32 {
	if x != nil {
		return x.TXOutIdx
	}
	return 0
}

func (x *ASPreImage) GetIssuedAt() uint32 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *ASPreImage) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *ASPreImage) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ASPreImage) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

// Protobuf message implementation for struct ValueStore
type ValueStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VSPreImage *VSPreImage `protobuf:"bytes,1,opt,name=VSPreImage,proto3" json:"VSPreImage,omitempty"`
	TxHash     string      `protobuf:"bytes,2,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *ValueStore) Reset() {
	*x = ValueStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueStore) ProtoMessage() {}

func (x *ValueStore) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueStore.ProtoReflect.Descriptor instead.
func (*ValueStore) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{7}
}

func (x *ValueStore) GetVSPreImage() *VSPreImage {
	if x != nil {
		return x.VSPreImage
	}
	return nil
}

func (x *ValueStore) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// Protobuf message implementation for struct VSPreImage
type VSPreImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID  uint32 `protobuf:"varint,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	TXOutIdx uint32 `protobuf:"varint,3,opt,name=TXOutIdx,proto3" json:"TXOutIdx,omitempty"`
	Owner    string `protobuf:"bytes,4,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Fee      string `protobuf:"bytes,5,opt,name=Fee,proto3" json:"Fee,omitempty"`
}

func (x *VSPreImage) Reset() {
	*x = VSPreImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VSPreImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VSPreImage) ProtoMessage() {}

func (x *VSPreImage) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VSPreImage.ProtoReflect.Descriptor instead.
func (*VSPreImage) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{8}
}

func (x *VSPreImage) GetChainID() uint32 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *VSPreImage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *VSPreImage) GetTXOutIdx() uint32 {
	if x != nil {
		return x.TXOutIdx
	}
	return 0
}

func (x *VSPreImage) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *VSPreImage) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

// Protobuf message implementation for struct DataStore
type DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DSLinker  *DSLinker `protobuf:"bytes,1,opt,name=DSLinker,proto3" json:"DSLinker,omitempty"`
	Signature string    `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *DataStore) Reset() {
	*x = DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStore) ProtoMessage() {}

func (x *DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStore.ProtoReflect.Descriptor instead.
func (*DataStore) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{9}
}

func (x *DataStore) GetDSLinker() *DSLinker {
	if x != nil {
		return x.DSLinker
	}
	return nil
}

func (x *DataStore) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// Protobuf message implementation for struct DSLinker
type DSLinker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DSPreImage *DSPreImage `protobuf:"bytes,1,opt,name=DSPreImage,proto3" json:"DSPreImage,omitempty"`
	TxHash     string      `protobuf:"bytes,2,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *DSLinker) Reset() {
	*x = DSLinker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSLinker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSLinker) ProtoMessage() {}

func (x *DSLinker) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSLinker.ProtoReflect.Descriptor instead.
func (*DSLinker) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{10}
}

func (x *DSLinker) GetDSPreImage() *DSPreImage {
	if x != nil {
		return x.DSPreImage
	}
	return nil
}

func (x *DSLinker) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// Protobuf message implementation for struct DSPreImage
type DSPreImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID  uint32 `protobuf:"varint,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Index    string `protobuf:"bytes,2,opt,name=Index,proto3" json:"Index,omitempty"`
	IssuedAt uint32 `protobuf:"varint,3,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Deposit  string `protobuf:"bytes,4,opt,name=Deposit,proto3" json:"Deposit,omitempty"`
	RawData  string `protobuf:"bytes,5,opt,name=RawData,proto3" json:"RawData,omitempty"`
	TXOutIdx uint32 `protobuf:"varint,6,opt,name=TXOutIdx,proto3" json:"TXOutIdx,omitempty"`
	Owner    string `protobuf:"bytes,7,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Fee      string `protobuf:"bytes,8,opt,name=Fee,proto3" json:"Fee,omitempty"`
}

func (x *DSPreImage) Reset() {
	*x = DSPreImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aobjs_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSPreImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSPreImage) ProtoMessage() {}

func (x *DSPreImage) ProtoReflect() protoreflect.Message {
	mi := &file_aobjs_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSPreImage.ProtoReflect.Descriptor instead.
func (*DSPreImage) Descriptor() ([]byte, []int) {
	return file_aobjs_proto_rawDescGZIP(), []int{11}
}

func (x *DSPreImage) GetChainID() uint32 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *DSPreImage) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *DSPreImage) GetIssuedAt() uint32 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *DSPreImage) GetDeposit() string {
	if x != nil {
		return x.Deposit
	}
	return ""
}

func (x *DSPreImage) GetRawData() string {
	if x != nil {
		return x.RawData
	}
	return ""
}

func (x *DSPreImage) GetTXOutIdx() uint32 {
	if x != nil {
		return x.TXOutIdx
	}
	return 0
}

func (x *DSPreImage) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *DSPreImage) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

var File_aobjs_proto protoreflect.FileDescriptor

var file_aobjs_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6f, 0x62, 0x6a, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x1d, 0x0a, 0x03, 0x56, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x58, 0x49, 0x6e, 0x52, 0x03, 0x56, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x56, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x58, 0x4f, 0x75, 0x74, 0x52, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x46,
	0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x65, 0x65, 0x22, 0xab, 0x01,
	0x0a, 0x05, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x41, 0x74, 0x6f, 0x6d, 0x69,
	0x63, 0x53, 0x77, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x77, 0x61, 0x70, 0x48, 0x00,
	0x52, 0x0a, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x77, 0x61, 0x70, 0x12, 0x33, 0x0a, 0x0a,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x30, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x00, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x75, 0x74, 0x78, 0x6f, 0x22, 0x57, 0x0a, 0x04, 0x54,
	0x58, 0x49, 0x6e, 0x12, 0x31, 0x0a, 0x0a, 0x54, 0x58, 0x49, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x58, 0x49, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x52, 0x0a, 0x54, 0x58, 0x49, 0x6e,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x5d, 0x0a, 0x0a, 0x54, 0x58, 0x49, 0x6e, 0x4c, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x54, 0x58, 0x49, 0x6e, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x58, 0x49, 0x6e, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x54,
	0x58, 0x49, 0x6e, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x76, 0x0a, 0x0c, 0x54, 0x58, 0x49, 0x6e, 0x50, 0x72, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x54, 0x78, 0x49, 0x64, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x54, 0x78,
	0x49, 0x64, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x54,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x64, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x57, 0x0a, 0x0a, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x77, 0x61, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x41, 0x53, 0x50,
	0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x0a, 0x41, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x41, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45,
	0x78, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x46, 0x65, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x56, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x56, 0x53, 0x50, 0x72,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x80,
	0x01, 0x0a, 0x0a, 0x56, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x65,
	0x65, 0x22, 0x56, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x44, 0x53, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x4c, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x52, 0x08, 0x44, 0x53, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x55, 0x0a, 0x08, 0x44, 0x53, 0x4c,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x44, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x44, 0x53,
	0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x22, 0xd0, 0x01, 0x0a, 0x0a, 0x44, 0x53, 0x50, 0x72, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x49, 0x64, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x46, 0x65, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aobjs_proto_rawDescOnce sync.Once
	file_aobjs_proto_rawDescData = file_aobjs_proto_rawDesc
)

func file_aobjs_proto_rawDescGZIP() []byte {
	file_aobjs_proto_rawDescOnce.Do(func() {
		file_aobjs_proto_rawDescData = protoimpl.X.CompressGZIP(file_aobjs_proto_rawDescData)
	})
	return file_aobjs_proto_rawDescData
}

var file_aobjs_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_aobjs_proto_goTypes = []interface{}{
	(*Tx)(nil),           // 0: proto.Tx
	(*TXOut)(nil),        // 1: proto.TXOut
	(*TXIn)(nil),         // 2: proto.TXIn
	(*TXInLinker)(nil),   // 3: proto.TXInLinker
	(*TXInPreImage)(nil), // 4: proto.TXInPreImage
	(*AtomicSwap)(nil),   // 5: proto.AtomicSwap
	(*ASPreImage)(nil),   // 6: proto.ASPreImage
	(*ValueStore)(nil),   // 7: proto.ValueStore
	(*VSPreImage)(nil),   // 8: proto.VSPreImage
	(*DataStore)(nil),    // 9: proto.DataStore
	(*DSLinker)(nil),     // 10: proto.DSLinker
	(*DSPreImage)(nil),   // 11: proto.DSPreImage
}
var file_aobjs_proto_depIdxs = []int32{
	2,  // 0: proto.Tx.Vin:type_name -> proto.TXIn
	1,  // 1: proto.Tx.Vout:type_name -> proto.TXOut
	5,  // 2: proto.TXOut.AtomicSwap:type_name -> proto.AtomicSwap
	7,  // 3: proto.TXOut.ValueStore:type_name -> proto.ValueStore
	9,  // 4: proto.TXOut.DataStore:type_name -> proto.DataStore
	3,  // 5: proto.TXIn.TXInLinker:type_name -> proto.TXInLinker
	4,  // 6: proto.TXInLinker.TXInPreImage:type_name -> proto.TXInPreImage
	6,  // 7: proto.AtomicSwap.ASPreImage:type_name -> proto.ASPreImage
	8,  // 8: proto.ValueStore.VSPreImage:type_name -> proto.VSPreImage
	10, // 9: proto.DataStore.DSLinker:type_name -> proto.DSLinker
	11, // 10: proto.DSLinker.DSPreImage:type_name -> proto.DSPreImage
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_aobjs_proto_init() }
func file_aobjs_proto_init() {
	if File_aobjs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aobjs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_aobjs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXOut); i {
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
		file_aobjs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXIn); i {
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
		file_aobjs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXInLinker); i {
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
		file_aobjs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXInPreImage); i {
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
		file_aobjs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtomicSwap); i {
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
		file_aobjs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ASPreImage); i {
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
		file_aobjs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueStore); i {
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
		file_aobjs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VSPreImage); i {
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
		file_aobjs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStore); i {
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
		file_aobjs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSLinker); i {
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
		file_aobjs_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSPreImage); i {
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
	file_aobjs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TXOut_AtomicSwap)(nil),
		(*TXOut_ValueStore)(nil),
		(*TXOut_DataStore)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aobjs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aobjs_proto_goTypes,
		DependencyIndexes: file_aobjs_proto_depIdxs,
		MessageInfos:      file_aobjs_proto_msgTypes,
	}.Build()
	File_aobjs_proto = out.File
	file_aobjs_proto_rawDesc = nil
	file_aobjs_proto_goTypes = nil
	file_aobjs_proto_depIdxs = nil
}
