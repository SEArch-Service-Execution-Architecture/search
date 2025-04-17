// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: search/v1/contracts.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GlobalContractFormat int32

const (
	GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_UNSPECIFIED GlobalContractFormat = 0
	GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA         GlobalContractFormat = 1 // System of CFSMs in FSA file format.
	GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_GC          GlobalContractFormat = 2 // Global Choreography.
)

// Enum value maps for GlobalContractFormat.
var (
	GlobalContractFormat_name = map[int32]string{
		0: "GLOBAL_CONTRACT_FORMAT_UNSPECIFIED",
		1: "GLOBAL_CONTRACT_FORMAT_FSA",
		2: "GLOBAL_CONTRACT_FORMAT_GC",
	}
	GlobalContractFormat_value = map[string]int32{
		"GLOBAL_CONTRACT_FORMAT_UNSPECIFIED": 0,
		"GLOBAL_CONTRACT_FORMAT_FSA":         1,
		"GLOBAL_CONTRACT_FORMAT_GC":          2,
	}
)

func (x GlobalContractFormat) Enum() *GlobalContractFormat {
	p := new(GlobalContractFormat)
	*p = x
	return p
}

func (x GlobalContractFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GlobalContractFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_search_v1_contracts_proto_enumTypes[0].Descriptor()
}

func (GlobalContractFormat) Type() protoreflect.EnumType {
	return &file_search_v1_contracts_proto_enumTypes[0]
}

func (x GlobalContractFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GlobalContractFormat.Descriptor instead.
func (GlobalContractFormat) EnumDescriptor() ([]byte, []int) {
	return file_search_v1_contracts_proto_rawDescGZIP(), []int{0}
}

type LocalContractFormat int32

const (
	LocalContractFormat_LOCAL_CONTRACT_FORMAT_UNSPECIFIED              LocalContractFormat = 0
	LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA                      LocalContractFormat = 1 // Single CFSM in FSA file format (for Service Providers).
	LocalContractFormat_LOCAL_CONTRACT_FORMAT_PYTHON_BISIMULATION_CODE LocalContractFormat = 2 // Python code to construct CFSM for https://github.com/diegosenarruzza/bisimulation/
)

// Enum value maps for LocalContractFormat.
var (
	LocalContractFormat_name = map[int32]string{
		0: "LOCAL_CONTRACT_FORMAT_UNSPECIFIED",
		1: "LOCAL_CONTRACT_FORMAT_FSA",
		2: "LOCAL_CONTRACT_FORMAT_PYTHON_BISIMULATION_CODE",
	}
	LocalContractFormat_value = map[string]int32{
		"LOCAL_CONTRACT_FORMAT_UNSPECIFIED":              0,
		"LOCAL_CONTRACT_FORMAT_FSA":                      1,
		"LOCAL_CONTRACT_FORMAT_PYTHON_BISIMULATION_CODE": 2,
	}
)

func (x LocalContractFormat) Enum() *LocalContractFormat {
	p := new(LocalContractFormat)
	*p = x
	return p
}

func (x LocalContractFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalContractFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_search_v1_contracts_proto_enumTypes[1].Descriptor()
}

func (LocalContractFormat) Type() protoreflect.EnumType {
	return &file_search_v1_contracts_proto_enumTypes[1]
}

func (x LocalContractFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalContractFormat.Descriptor instead.
func (LocalContractFormat) EnumDescriptor() ([]byte, []int) {
	return file_search_v1_contracts_proto_rawDescGZIP(), []int{1}
}

type GlobalContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contract      []byte                 `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Format        GlobalContractFormat   `protobuf:"varint,2,opt,name=format,proto3,enum=search.v1.GlobalContractFormat" json:"format,omitempty"`
	InitiatorName string                 `protobuf:"bytes,3,opt,name=initiator_name,json=initiatorName,proto3" json:"initiator_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GlobalContract) Reset() {
	*x = GlobalContract{}
	mi := &file_search_v1_contracts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlobalContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalContract) ProtoMessage() {}

func (x *GlobalContract) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_contracts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalContract.ProtoReflect.Descriptor instead.
func (*GlobalContract) Descriptor() ([]byte, []int) {
	return file_search_v1_contracts_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalContract) GetContract() []byte {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *GlobalContract) GetFormat() GlobalContractFormat {
	if x != nil {
		return x.Format
	}
	return GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_UNSPECIFIED
}

func (x *GlobalContract) GetInitiatorName() string {
	if x != nil {
		return x.InitiatorName
	}
	return ""
}

type LocalContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Contract      []byte                 `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Format        LocalContractFormat    `protobuf:"varint,2,opt,name=format,proto3,enum=search.v1.LocalContractFormat" json:"format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocalContract) Reset() {
	*x = LocalContract{}
	mi := &file_search_v1_contracts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalContract) ProtoMessage() {}

func (x *LocalContract) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_contracts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalContract.ProtoReflect.Descriptor instead.
func (*LocalContract) Descriptor() ([]byte, []int) {
	return file_search_v1_contracts_proto_rawDescGZIP(), []int{1}
}

func (x *LocalContract) GetContract() []byte {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *LocalContract) GetFormat() LocalContractFormat {
	if x != nil {
		return x.Format
	}
	return LocalContractFormat_LOCAL_CONTRACT_FORMAT_UNSPECIFIED
}

type MessageTranslations struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Different contracts can be compatible while using different names for messages that are equivalent.
	// This data structure contains the mapping of each message name to the name used by the other participant.
	// The keys are the message names in the contract of who receives this message, and the values are the message
	// names according to the other participant's contract.
	Participant   string            `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	Translations  map[string]string `protobuf:"bytes,2,rep,name=translations,proto3" json:"translations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageTranslations) Reset() {
	*x = MessageTranslations{}
	mi := &file_search_v1_contracts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageTranslations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTranslations) ProtoMessage() {}

func (x *MessageTranslations) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_contracts_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTranslations.ProtoReflect.Descriptor instead.
func (*MessageTranslations) Descriptor() ([]byte, []int) {
	return file_search_v1_contracts_proto_rawDescGZIP(), []int{2}
}

func (x *MessageTranslations) GetParticipant() string {
	if x != nil {
		return x.Participant
	}
	return ""
}

func (x *MessageTranslations) GetTranslations() map[string]string {
	if x != nil {
		return x.Translations
	}
	return nil
}

var File_search_v1_contracts_proto protoreflect.FileDescriptor

const file_search_v1_contracts_proto_rawDesc = "" +
	"\n" +
	"\x19search/v1/contracts.proto\x12\tsearch.v1\"\x8c\x01\n" +
	"\x0eGlobalContract\x12\x1a\n" +
	"\bcontract\x18\x01 \x01(\fR\bcontract\x127\n" +
	"\x06format\x18\x02 \x01(\x0e2\x1f.search.v1.GlobalContractFormatR\x06format\x12%\n" +
	"\x0einitiator_name\x18\x03 \x01(\tR\rinitiatorName\"c\n" +
	"\rLocalContract\x12\x1a\n" +
	"\bcontract\x18\x01 \x01(\fR\bcontract\x126\n" +
	"\x06format\x18\x02 \x01(\x0e2\x1e.search.v1.LocalContractFormatR\x06format\"\xce\x01\n" +
	"\x13MessageTranslations\x12 \n" +
	"\vparticipant\x18\x01 \x01(\tR\vparticipant\x12T\n" +
	"\ftranslations\x18\x02 \x03(\v20.search.v1.MessageTranslations.TranslationsEntryR\ftranslations\x1a?\n" +
	"\x11TranslationsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01*}\n" +
	"\x14GlobalContractFormat\x12&\n" +
	"\"GLOBAL_CONTRACT_FORMAT_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aGLOBAL_CONTRACT_FORMAT_FSA\x10\x01\x12\x1d\n" +
	"\x19GLOBAL_CONTRACT_FORMAT_GC\x10\x02*\x8f\x01\n" +
	"\x13LocalContractFormat\x12%\n" +
	"!LOCAL_CONTRACT_FORMAT_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19LOCAL_CONTRACT_FORMAT_FSA\x10\x01\x122\n" +
	".LOCAL_CONTRACT_FORMAT_PYTHON_BISIMULATION_CODE\x10\x02Ba\n" +
	"\x1car.com.montepagano.search.v1ZAgithub.com/SEArch-Service-Execution-Architecture/search/gen/go/v1b\x06proto3"

var (
	file_search_v1_contracts_proto_rawDescOnce sync.Once
	file_search_v1_contracts_proto_rawDescData []byte
)

func file_search_v1_contracts_proto_rawDescGZIP() []byte {
	file_search_v1_contracts_proto_rawDescOnce.Do(func() {
		file_search_v1_contracts_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_search_v1_contracts_proto_rawDesc), len(file_search_v1_contracts_proto_rawDesc)))
	})
	return file_search_v1_contracts_proto_rawDescData
}

var file_search_v1_contracts_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_search_v1_contracts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_search_v1_contracts_proto_goTypes = []any{
	(GlobalContractFormat)(0),   // 0: search.v1.GlobalContractFormat
	(LocalContractFormat)(0),    // 1: search.v1.LocalContractFormat
	(*GlobalContract)(nil),      // 2: search.v1.GlobalContract
	(*LocalContract)(nil),       // 3: search.v1.LocalContract
	(*MessageTranslations)(nil), // 4: search.v1.MessageTranslations
	nil,                         // 5: search.v1.MessageTranslations.TranslationsEntry
}
var file_search_v1_contracts_proto_depIdxs = []int32{
	0, // 0: search.v1.GlobalContract.format:type_name -> search.v1.GlobalContractFormat
	1, // 1: search.v1.LocalContract.format:type_name -> search.v1.LocalContractFormat
	5, // 2: search.v1.MessageTranslations.translations:type_name -> search.v1.MessageTranslations.TranslationsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_search_v1_contracts_proto_init() }
func file_search_v1_contracts_proto_init() {
	if File_search_v1_contracts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_search_v1_contracts_proto_rawDesc), len(file_search_v1_contracts_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_search_v1_contracts_proto_goTypes,
		DependencyIndexes: file_search_v1_contracts_proto_depIdxs,
		EnumInfos:         file_search_v1_contracts_proto_enumTypes,
		MessageInfos:      file_search_v1_contracts_proto_msgTypes,
	}.Build()
	File_search_v1_contracts_proto = out.File
	file_search_v1_contracts_proto_goTypes = nil
	file_search_v1_contracts_proto_depIdxs = nil
}
