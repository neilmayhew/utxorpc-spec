// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: utxorpc/ledger/v1/scripts.proto

package ledgerv1

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

// Represents a native script in Cardano.
type NativeScript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NativeScript:
	//
	//	*NativeScript_ScriptPubkey
	//	*NativeScript_ScriptAll
	//	*NativeScript_ScriptAny
	//	*NativeScript_ScriptNOfK
	//	*NativeScript_InvalidBefore
	//	*NativeScript_InvalidHereafter
	NativeScript isNativeScript_NativeScript `protobuf_oneof:"native_script"`
}

func (x *NativeScript) Reset() {
	*x = NativeScript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NativeScript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NativeScript) ProtoMessage() {}

func (x *NativeScript) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NativeScript.ProtoReflect.Descriptor instead.
func (*NativeScript) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{0}
}

func (m *NativeScript) GetNativeScript() isNativeScript_NativeScript {
	if m != nil {
		return m.NativeScript
	}
	return nil
}

func (x *NativeScript) GetScriptPubkey() []byte {
	if x, ok := x.GetNativeScript().(*NativeScript_ScriptPubkey); ok {
		return x.ScriptPubkey
	}
	return nil
}

func (x *NativeScript) GetScriptAll() *NativeScriptList {
	if x, ok := x.GetNativeScript().(*NativeScript_ScriptAll); ok {
		return x.ScriptAll
	}
	return nil
}

func (x *NativeScript) GetScriptAny() *NativeScriptList {
	if x, ok := x.GetNativeScript().(*NativeScript_ScriptAny); ok {
		return x.ScriptAny
	}
	return nil
}

func (x *NativeScript) GetScriptNOfK() *ScriptNOfK {
	if x, ok := x.GetNativeScript().(*NativeScript_ScriptNOfK); ok {
		return x.ScriptNOfK
	}
	return nil
}

func (x *NativeScript) GetInvalidBefore() uint64 {
	if x, ok := x.GetNativeScript().(*NativeScript_InvalidBefore); ok {
		return x.InvalidBefore
	}
	return 0
}

func (x *NativeScript) GetInvalidHereafter() uint64 {
	if x, ok := x.GetNativeScript().(*NativeScript_InvalidHereafter); ok {
		return x.InvalidHereafter
	}
	return 0
}

type isNativeScript_NativeScript interface {
	isNativeScript_NativeScript()
}

type NativeScript_ScriptPubkey struct {
	ScriptPubkey []byte `protobuf:"bytes,1,opt,name=script_pubkey,json=scriptPubkey,proto3,oneof"` // Script based on an address key hash.
}

type NativeScript_ScriptAll struct {
	ScriptAll *NativeScriptList `protobuf:"bytes,2,opt,name=script_all,json=scriptAll,proto3,oneof"` // Script that requires all nested scripts to be satisfied.
}

type NativeScript_ScriptAny struct {
	ScriptAny *NativeScriptList `protobuf:"bytes,3,opt,name=script_any,json=scriptAny,proto3,oneof"` // Script that requires any of the nested scripts to be satisfied.
}

type NativeScript_ScriptNOfK struct {
	ScriptNOfK *ScriptNOfK `protobuf:"bytes,4,opt,name=script_n_of_k,json=scriptNOfK,proto3,oneof"` // Script that requires k out of n nested scripts to be satisfied.
}

type NativeScript_InvalidBefore struct {
	InvalidBefore uint64 `protobuf:"varint,5,opt,name=invalid_before,json=invalidBefore,proto3,oneof"` // Slot number before which the script is invalid.
}

type NativeScript_InvalidHereafter struct {
	InvalidHereafter uint64 `protobuf:"varint,6,opt,name=invalid_hereafter,json=invalidHereafter,proto3,oneof"` // Slot number after which the script is invalid.
}

func (*NativeScript_ScriptPubkey) isNativeScript_NativeScript() {}

func (*NativeScript_ScriptAll) isNativeScript_NativeScript() {}

func (*NativeScript_ScriptAny) isNativeScript_NativeScript() {}

func (*NativeScript_ScriptNOfK) isNativeScript_NativeScript() {}

func (*NativeScript_InvalidBefore) isNativeScript_NativeScript() {}

func (*NativeScript_InvalidHereafter) isNativeScript_NativeScript() {}

// Represents a list of native scripts.
type NativeScriptList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NativeScript `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // List of native scripts.
}

func (x *NativeScriptList) Reset() {
	*x = NativeScriptList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NativeScriptList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NativeScriptList) ProtoMessage() {}

func (x *NativeScriptList) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NativeScriptList.ProtoReflect.Descriptor instead.
func (*NativeScriptList) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{1}
}

func (x *NativeScriptList) GetItems() []*NativeScript {
	if x != nil {
		return x.Items
	}
	return nil
}

// Represents a "k out of n" native script.
type ScriptNOfK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K       uint32          `protobuf:"varint,1,opt,name=k,proto3" json:"k,omitempty"`            // The number of required satisfied scripts.
	Scripts []*NativeScript `protobuf:"bytes,2,rep,name=scripts,proto3" json:"scripts,omitempty"` // List of native scripts.
}

func (x *ScriptNOfK) Reset() {
	*x = ScriptNOfK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptNOfK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptNOfK) ProtoMessage() {}

func (x *ScriptNOfK) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptNOfK.ProtoReflect.Descriptor instead.
func (*ScriptNOfK) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{2}
}

func (x *ScriptNOfK) GetK() uint32 {
	if x != nil {
		return x.K
	}
	return 0
}

func (x *ScriptNOfK) GetScripts() []*NativeScript {
	if x != nil {
		return x.Scripts
	}
	return nil
}

// Represents a constructor for Plutus data in Cardano.
type Constr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Constr) Reset() {
	*x = Constr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constr) ProtoMessage() {}

func (x *Constr) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constr.ProtoReflect.Descriptor instead.
func (*Constr) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{3}
}

// Represents a big integer for Plutus data in Cardano.
type BigInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BigInt) Reset() {
	*x = BigInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigInt) ProtoMessage() {}

func (x *BigInt) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigInt.ProtoReflect.Descriptor instead.
func (*BigInt) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{4}
}

// Represents a bounded bytes for Plutus data in Cardano.
type BoundedBytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BoundedBytes) Reset() {
	*x = BoundedBytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundedBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundedBytes) ProtoMessage() {}

func (x *BoundedBytes) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundedBytes.ProtoReflect.Descriptor instead.
func (*BoundedBytes) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{5}
}

// Represents a key-value pair for Plutus data in Cardano.
type PlutusDataPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *PlutusData `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // Key of the pair.
	Value *PlutusData `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // Value of the pair.
}

func (x *PlutusDataPair) Reset() {
	*x = PlutusDataPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlutusDataPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlutusDataPair) ProtoMessage() {}

func (x *PlutusDataPair) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlutusDataPair.ProtoReflect.Descriptor instead.
func (*PlutusDataPair) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{6}
}

func (x *PlutusDataPair) GetKey() *PlutusData {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PlutusDataPair) GetValue() *PlutusData {
	if x != nil {
		return x.Value
	}
	return nil
}

// Represents a Plutus data item in Cardano.
type PlutusData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PlutusData:
	//
	//	*PlutusData_Constr
	//	*PlutusData_Map
	//	*PlutusData_BigInt
	//	*PlutusData_BoundedBytes
	//	*PlutusData_Array
	PlutusData isPlutusData_PlutusData `protobuf_oneof:"plutus_data"`
}

func (x *PlutusData) Reset() {
	*x = PlutusData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlutusData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlutusData) ProtoMessage() {}

func (x *PlutusData) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlutusData.ProtoReflect.Descriptor instead.
func (*PlutusData) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{7}
}

func (m *PlutusData) GetPlutusData() isPlutusData_PlutusData {
	if m != nil {
		return m.PlutusData
	}
	return nil
}

func (x *PlutusData) GetConstr() *Constr {
	if x, ok := x.GetPlutusData().(*PlutusData_Constr); ok {
		return x.Constr
	}
	return nil
}

func (x *PlutusData) GetMap() *PlutusDataMap {
	if x, ok := x.GetPlutusData().(*PlutusData_Map); ok {
		return x.Map
	}
	return nil
}

func (x *PlutusData) GetBigInt() *BigInt {
	if x, ok := x.GetPlutusData().(*PlutusData_BigInt); ok {
		return x.BigInt
	}
	return nil
}

func (x *PlutusData) GetBoundedBytes() *BoundedBytes {
	if x, ok := x.GetPlutusData().(*PlutusData_BoundedBytes); ok {
		return x.BoundedBytes
	}
	return nil
}

func (x *PlutusData) GetArray() *PlutusDataArray {
	if x, ok := x.GetPlutusData().(*PlutusData_Array); ok {
		return x.Array
	}
	return nil
}

type isPlutusData_PlutusData interface {
	isPlutusData_PlutusData()
}

type PlutusData_Constr struct {
	Constr *Constr `protobuf:"bytes,1,opt,name=Constr,proto3,oneof"` // Constructor.
}

type PlutusData_Map struct {
	Map *PlutusDataMap `protobuf:"bytes,2,opt,name=Map,proto3,oneof"` // Map of Plutus data.
}

type PlutusData_BigInt struct {
	BigInt *BigInt `protobuf:"bytes,3,opt,name=BigInt,proto3,oneof"` // Big integer.
}

type PlutusData_BoundedBytes struct {
	BoundedBytes *BoundedBytes `protobuf:"bytes,4,opt,name=BoundedBytes,proto3,oneof"` // Bounded bytes.
}

type PlutusData_Array struct {
	Array *PlutusDataArray `protobuf:"bytes,5,opt,name=Array,proto3,oneof"` // Array of Plutus data.
}

func (*PlutusData_Constr) isPlutusData_PlutusData() {}

func (*PlutusData_Map) isPlutusData_PlutusData() {}

func (*PlutusData_BigInt) isPlutusData_PlutusData() {}

func (*PlutusData_BoundedBytes) isPlutusData_PlutusData() {}

func (*PlutusData_Array) isPlutusData_PlutusData() {}

// Represents a map of Plutus data in Cardano.
type PlutusDataMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*PlutusDataPair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"` // List of key-value pairs.
}

func (x *PlutusDataMap) Reset() {
	*x = PlutusDataMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlutusDataMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlutusDataMap) ProtoMessage() {}

func (x *PlutusDataMap) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlutusDataMap.ProtoReflect.Descriptor instead.
func (*PlutusDataMap) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{8}
}

func (x *PlutusDataMap) GetPairs() []*PlutusDataPair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// Represents an array of Plutus data in Cardano.
type PlutusDataArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PlutusData `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // List of Plutus data items.
}

func (x *PlutusDataArray) Reset() {
	*x = PlutusDataArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlutusDataArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlutusDataArray) ProtoMessage() {}

func (x *PlutusDataArray) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlutusDataArray.ProtoReflect.Descriptor instead.
func (*PlutusDataArray) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{9}
}

func (x *PlutusDataArray) GetItems() []*PlutusData {
	if x != nil {
		return x.Items
	}
	return nil
}

// Represents a script in Cardano.
type Script struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Script:
	//
	//	*Script_Native
	//	*Script_PlutusV1
	//	*Script_PlutusV2
	Script isScript_Script `protobuf_oneof:"script"`
}

func (x *Script) Reset() {
	*x = Script{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Script) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Script) ProtoMessage() {}

func (x *Script) ProtoReflect() protoreflect.Message {
	mi := &file_utxorpc_ledger_v1_scripts_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Script.ProtoReflect.Descriptor instead.
func (*Script) Descriptor() ([]byte, []int) {
	return file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP(), []int{10}
}

func (m *Script) GetScript() isScript_Script {
	if m != nil {
		return m.Script
	}
	return nil
}

func (x *Script) GetNative() *NativeScript {
	if x, ok := x.GetScript().(*Script_Native); ok {
		return x.Native
	}
	return nil
}

func (x *Script) GetPlutusV1() []byte {
	if x, ok := x.GetScript().(*Script_PlutusV1); ok {
		return x.PlutusV1
	}
	return nil
}

func (x *Script) GetPlutusV2() []byte {
	if x, ok := x.GetScript().(*Script_PlutusV2); ok {
		return x.PlutusV2
	}
	return nil
}

type isScript_Script interface {
	isScript_Script()
}

type Script_Native struct {
	Native *NativeScript `protobuf:"bytes,1,opt,name=Native,proto3,oneof"` // Native script.
}

type Script_PlutusV1 struct {
	PlutusV1 []byte `protobuf:"bytes,2,opt,name=PlutusV1,proto3,oneof"` // Plutus V1 script.
}

type Script_PlutusV2 struct {
	PlutusV2 []byte `protobuf:"bytes,3,opt,name=PlutusV2,proto3,oneof"` // Plutus V2 script.
}

func (*Script_Native) isScript_Script() {}

func (*Script_PlutusV1) isScript_Script() {}

func (*Script_PlutusV2) isScript_Script() {}

var File_utxorpc_ledger_v1_scripts_proto protoreflect.FileDescriptor

var file_utxorpc_ledger_v1_scripts_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0xee, 0x02, 0x0a, 0x0c, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x25, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f,
	0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x0a,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x44, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x61, 0x6e, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x41, 0x6e, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x5f, 0x6e, 0x5f, 0x6f, 0x66, 0x5f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x4f, 0x66, 0x4b, 0x48, 0x00,
	0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x4f, 0x66, 0x4b, 0x12, 0x27, 0x0a, 0x0e,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x68, 0x65, 0x72, 0x65, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x48, 0x65, 0x72, 0x65, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x49, 0x0a, 0x10, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x55, 0x0a, 0x0a, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x4f, 0x66, 0x4b, 0x12, 0x0c,
	0x0a, 0x01, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6b, 0x12, 0x39, 0x0a, 0x07,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x22, 0x08, 0x0a, 0x06, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x0e, 0x50,
	0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2f, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x74, 0x78,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x48, 0x00, 0x52,
	0x06, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x12, 0x34, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x33, 0x0a,
	0x06, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x42, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0c, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x05,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x48, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x4d, 0x61, 0x70, 0x12, 0x37, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x46,
	0x0a, 0x0f, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x39, 0x0a, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x48, 0x00, 0x52, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x08,
	0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x56, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x08, 0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x56, 0x31, 0x12, 0x1c, 0x0a, 0x08, 0x50, 0x6c,
	0x75, 0x74, 0x75, 0x73, 0x56, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08,
	0x50, 0x6c, 0x75, 0x74, 0x75, 0x73, 0x56, 0x32, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x42, 0xc8, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x74, 0x78, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x75,
	0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x4c, 0x58, 0xaa,
	0x02, 0x11, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70, 0x63, 0x5c, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70,
	0x63, 0x5c, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x55, 0x74, 0x78, 0x6f, 0x72, 0x70,
	0x63, 0x3a, 0x3a, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utxorpc_ledger_v1_scripts_proto_rawDescOnce sync.Once
	file_utxorpc_ledger_v1_scripts_proto_rawDescData = file_utxorpc_ledger_v1_scripts_proto_rawDesc
)

func file_utxorpc_ledger_v1_scripts_proto_rawDescGZIP() []byte {
	file_utxorpc_ledger_v1_scripts_proto_rawDescOnce.Do(func() {
		file_utxorpc_ledger_v1_scripts_proto_rawDescData = protoimpl.X.CompressGZIP(file_utxorpc_ledger_v1_scripts_proto_rawDescData)
	})
	return file_utxorpc_ledger_v1_scripts_proto_rawDescData
}

var file_utxorpc_ledger_v1_scripts_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_utxorpc_ledger_v1_scripts_proto_goTypes = []interface{}{
	(*NativeScript)(nil),     // 0: utxorpc.ledger.v1.NativeScript
	(*NativeScriptList)(nil), // 1: utxorpc.ledger.v1.NativeScriptList
	(*ScriptNOfK)(nil),       // 2: utxorpc.ledger.v1.ScriptNOfK
	(*Constr)(nil),           // 3: utxorpc.ledger.v1.Constr
	(*BigInt)(nil),           // 4: utxorpc.ledger.v1.BigInt
	(*BoundedBytes)(nil),     // 5: utxorpc.ledger.v1.BoundedBytes
	(*PlutusDataPair)(nil),   // 6: utxorpc.ledger.v1.PlutusDataPair
	(*PlutusData)(nil),       // 7: utxorpc.ledger.v1.PlutusData
	(*PlutusDataMap)(nil),    // 8: utxorpc.ledger.v1.PlutusDataMap
	(*PlutusDataArray)(nil),  // 9: utxorpc.ledger.v1.PlutusDataArray
	(*Script)(nil),           // 10: utxorpc.ledger.v1.Script
}
var file_utxorpc_ledger_v1_scripts_proto_depIdxs = []int32{
	1,  // 0: utxorpc.ledger.v1.NativeScript.script_all:type_name -> utxorpc.ledger.v1.NativeScriptList
	1,  // 1: utxorpc.ledger.v1.NativeScript.script_any:type_name -> utxorpc.ledger.v1.NativeScriptList
	2,  // 2: utxorpc.ledger.v1.NativeScript.script_n_of_k:type_name -> utxorpc.ledger.v1.ScriptNOfK
	0,  // 3: utxorpc.ledger.v1.NativeScriptList.items:type_name -> utxorpc.ledger.v1.NativeScript
	0,  // 4: utxorpc.ledger.v1.ScriptNOfK.scripts:type_name -> utxorpc.ledger.v1.NativeScript
	7,  // 5: utxorpc.ledger.v1.PlutusDataPair.key:type_name -> utxorpc.ledger.v1.PlutusData
	7,  // 6: utxorpc.ledger.v1.PlutusDataPair.value:type_name -> utxorpc.ledger.v1.PlutusData
	3,  // 7: utxorpc.ledger.v1.PlutusData.Constr:type_name -> utxorpc.ledger.v1.Constr
	8,  // 8: utxorpc.ledger.v1.PlutusData.Map:type_name -> utxorpc.ledger.v1.PlutusDataMap
	4,  // 9: utxorpc.ledger.v1.PlutusData.BigInt:type_name -> utxorpc.ledger.v1.BigInt
	5,  // 10: utxorpc.ledger.v1.PlutusData.BoundedBytes:type_name -> utxorpc.ledger.v1.BoundedBytes
	9,  // 11: utxorpc.ledger.v1.PlutusData.Array:type_name -> utxorpc.ledger.v1.PlutusDataArray
	6,  // 12: utxorpc.ledger.v1.PlutusDataMap.pairs:type_name -> utxorpc.ledger.v1.PlutusDataPair
	7,  // 13: utxorpc.ledger.v1.PlutusDataArray.items:type_name -> utxorpc.ledger.v1.PlutusData
	0,  // 14: utxorpc.ledger.v1.Script.Native:type_name -> utxorpc.ledger.v1.NativeScript
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_utxorpc_ledger_v1_scripts_proto_init() }
func file_utxorpc_ledger_v1_scripts_proto_init() {
	if File_utxorpc_ledger_v1_scripts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NativeScript); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NativeScriptList); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptNOfK); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constr); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigInt); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundedBytes); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlutusDataPair); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlutusData); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlutusDataMap); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlutusDataArray); i {
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
		file_utxorpc_ledger_v1_scripts_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Script); i {
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
	file_utxorpc_ledger_v1_scripts_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NativeScript_ScriptPubkey)(nil),
		(*NativeScript_ScriptAll)(nil),
		(*NativeScript_ScriptAny)(nil),
		(*NativeScript_ScriptNOfK)(nil),
		(*NativeScript_InvalidBefore)(nil),
		(*NativeScript_InvalidHereafter)(nil),
	}
	file_utxorpc_ledger_v1_scripts_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*PlutusData_Constr)(nil),
		(*PlutusData_Map)(nil),
		(*PlutusData_BigInt)(nil),
		(*PlutusData_BoundedBytes)(nil),
		(*PlutusData_Array)(nil),
	}
	file_utxorpc_ledger_v1_scripts_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*Script_Native)(nil),
		(*Script_PlutusV1)(nil),
		(*Script_PlutusV2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utxorpc_ledger_v1_scripts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_utxorpc_ledger_v1_scripts_proto_goTypes,
		DependencyIndexes: file_utxorpc_ledger_v1_scripts_proto_depIdxs,
		MessageInfos:      file_utxorpc_ledger_v1_scripts_proto_msgTypes,
	}.Build()
	File_utxorpc_ledger_v1_scripts_proto = out.File
	file_utxorpc_ledger_v1_scripts_proto_rawDesc = nil
	file_utxorpc_ledger_v1_scripts_proto_goTypes = nil
	file_utxorpc_ledger_v1_scripts_proto_depIdxs = nil
}
