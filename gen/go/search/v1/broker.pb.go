// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: search/v1/broker.proto

package v1

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

type BrokerChannelResponse_Result int32

const (
	BrokerChannelResponse_RESULT_UNSPECIFIED BrokerChannelResponse_Result = 0
	BrokerChannelResponse_RESULT_ACK         BrokerChannelResponse_Result = 1
	BrokerChannelResponse_RESULT_ERR         BrokerChannelResponse_Result = 2
)

// Enum value maps for BrokerChannelResponse_Result.
var (
	BrokerChannelResponse_Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "RESULT_ACK",
		2: "RESULT_ERR",
	}
	BrokerChannelResponse_Result_value = map[string]int32{
		"RESULT_UNSPECIFIED": 0,
		"RESULT_ACK":         1,
		"RESULT_ERR":         2,
	}
)

func (x BrokerChannelResponse_Result) Enum() *BrokerChannelResponse_Result {
	p := new(BrokerChannelResponse_Result)
	*p = x
	return p
}

func (x BrokerChannelResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrokerChannelResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_search_v1_broker_proto_enumTypes[0].Descriptor()
}

func (BrokerChannelResponse_Result) Type() protoreflect.EnumType {
	return &file_search_v1_broker_proto_enumTypes[0]
}

func (x BrokerChannelResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrokerChannelResponse_Result.Descriptor instead.
func (BrokerChannelResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{1, 0}
}

type BrokerChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *GlobalContract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"` // requirements contract
	// subset of contract's participants that are already decided. This should at least
	// include the initiator's RemoteParticpant data
	PresetParticipants map[string]*RemoteParticipant `protobuf:"bytes,2,rep,name=preset_participants,json=presetParticipants,proto3" json:"preset_participants,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BrokerChannelRequest) Reset() {
	*x = BrokerChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_v1_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerChannelRequest) ProtoMessage() {}

func (x *BrokerChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerChannelRequest.ProtoReflect.Descriptor instead.
func (*BrokerChannelRequest) Descriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{0}
}

func (x *BrokerChannelRequest) GetContract() *GlobalContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *BrokerChannelRequest) GetPresetParticipants() map[string]*RemoteParticipant {
	if x != nil {
		return x.PresetParticipants
	}
	return nil
}

type BrokerChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       BrokerChannelResponse_Result  `protobuf:"varint,1,opt,name=result,proto3,enum=search.v1.BrokerChannelResponse_Result" json:"result,omitempty"`
	ChannelId    string                        `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`                                                                              // uuidv4
	Participants map[string]*RemoteParticipant `protobuf:"bytes,3,rep,name=participants,proto3" json:"participants,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // preset + brokered participants
}

func (x *BrokerChannelResponse) Reset() {
	*x = BrokerChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_v1_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerChannelResponse) ProtoMessage() {}

func (x *BrokerChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerChannelResponse.ProtoReflect.Descriptor instead.
func (*BrokerChannelResponse) Descriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{1}
}

func (x *BrokerChannelResponse) GetResult() BrokerChannelResponse_Result {
	if x != nil {
		return x.Result
	}
	return BrokerChannelResponse_RESULT_UNSPECIFIED
}

func (x *BrokerChannelResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *BrokerChannelResponse) GetParticipants() map[string]*RemoteParticipant {
	if x != nil {
		return x.Participants
	}
	return nil
}

type RegisterProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *LocalContract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Url      string         `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RegisterProviderRequest) Reset() {
	*x = RegisterProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_v1_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProviderRequest) ProtoMessage() {}

func (x *RegisterProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProviderRequest.ProtoReflect.Descriptor instead.
func (*RegisterProviderRequest) Descriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterProviderRequest) GetContract() *LocalContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *RegisterProviderRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// The registry assigns the provider an ID
type RegisterProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *RegisterProviderResponse) Reset() {
	*x = RegisterProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_v1_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProviderResponse) ProtoMessage() {}

func (x *RegisterProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_broker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProviderResponse.ProtoReflect.Descriptor instead.
func (*RegisterProviderResponse) Descriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterProviderResponse) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type RemoteParticipant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`                  // points to the middleware for this participant
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"` // points to the specific app that is served by the middleware
}

func (x *RemoteParticipant) Reset() {
	*x = RemoteParticipant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_v1_broker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteParticipant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteParticipant) ProtoMessage() {}

func (x *RemoteParticipant) ProtoReflect() protoreflect.Message {
	mi := &file_search_v1_broker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteParticipant.ProtoReflect.Descriptor instead.
func (*RemoteParticipant) Descriptor() ([]byte, []int) {
	return file_search_v1_broker_proto_rawDescGZIP(), []int{4}
}

func (x *RemoteParticipant) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RemoteParticipant) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

var File_search_v1_broker_proto protoreflect.FileDescriptor

var file_search_v1_broker_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x02, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x68,
	0x0a, 0x13, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x1a, 0x63, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf0, 0x02,
	0x0a, 0x15, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x1a,
	0x5d, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x02,
	0x22, 0x61, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x32, 0xc4, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6d, 0x6f, 0x6e, 0x74, 0x65,
	0x70, 0x61, 0x67, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_v1_broker_proto_rawDescOnce sync.Once
	file_search_v1_broker_proto_rawDescData = file_search_v1_broker_proto_rawDesc
)

func file_search_v1_broker_proto_rawDescGZIP() []byte {
	file_search_v1_broker_proto_rawDescOnce.Do(func() {
		file_search_v1_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_v1_broker_proto_rawDescData)
	})
	return file_search_v1_broker_proto_rawDescData
}

var file_search_v1_broker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_search_v1_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_search_v1_broker_proto_goTypes = []interface{}{
	(BrokerChannelResponse_Result)(0), // 0: search.v1.BrokerChannelResponse.Result
	(*BrokerChannelRequest)(nil),      // 1: search.v1.BrokerChannelRequest
	(*BrokerChannelResponse)(nil),     // 2: search.v1.BrokerChannelResponse
	(*RegisterProviderRequest)(nil),   // 3: search.v1.RegisterProviderRequest
	(*RegisterProviderResponse)(nil),  // 4: search.v1.RegisterProviderResponse
	(*RemoteParticipant)(nil),         // 5: search.v1.RemoteParticipant
	nil,                               // 6: search.v1.BrokerChannelRequest.PresetParticipantsEntry
	nil,                               // 7: search.v1.BrokerChannelResponse.ParticipantsEntry
	(*GlobalContract)(nil),            // 8: search.v1.GlobalContract
	(*LocalContract)(nil),             // 9: search.v1.LocalContract
}
var file_search_v1_broker_proto_depIdxs = []int32{
	8, // 0: search.v1.BrokerChannelRequest.contract:type_name -> search.v1.GlobalContract
	6, // 1: search.v1.BrokerChannelRequest.preset_participants:type_name -> search.v1.BrokerChannelRequest.PresetParticipantsEntry
	0, // 2: search.v1.BrokerChannelResponse.result:type_name -> search.v1.BrokerChannelResponse.Result
	7, // 3: search.v1.BrokerChannelResponse.participants:type_name -> search.v1.BrokerChannelResponse.ParticipantsEntry
	9, // 4: search.v1.RegisterProviderRequest.contract:type_name -> search.v1.LocalContract
	5, // 5: search.v1.BrokerChannelRequest.PresetParticipantsEntry.value:type_name -> search.v1.RemoteParticipant
	5, // 6: search.v1.BrokerChannelResponse.ParticipantsEntry.value:type_name -> search.v1.RemoteParticipant
	1, // 7: search.v1.BrokerService.BrokerChannel:input_type -> search.v1.BrokerChannelRequest
	3, // 8: search.v1.BrokerService.RegisterProvider:input_type -> search.v1.RegisterProviderRequest
	2, // 9: search.v1.BrokerService.BrokerChannel:output_type -> search.v1.BrokerChannelResponse
	4, // 10: search.v1.BrokerService.RegisterProvider:output_type -> search.v1.RegisterProviderResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_search_v1_broker_proto_init() }
func file_search_v1_broker_proto_init() {
	if File_search_v1_broker_proto != nil {
		return
	}
	file_search_v1_contracts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_search_v1_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerChannelRequest); i {
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
		file_search_v1_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerChannelResponse); i {
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
		file_search_v1_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProviderRequest); i {
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
		file_search_v1_broker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProviderResponse); i {
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
		file_search_v1_broker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteParticipant); i {
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
			RawDescriptor: file_search_v1_broker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_v1_broker_proto_goTypes,
		DependencyIndexes: file_search_v1_broker_proto_depIdxs,
		EnumInfos:         file_search_v1_broker_proto_enumTypes,
		MessageInfos:      file_search_v1_broker_proto_msgTypes,
	}.Build()
	File_search_v1_broker_proto = out.File
	file_search_v1_broker_proto_rawDesc = nil
	file_search_v1_broker_proto_goTypes = nil
	file_search_v1_broker_proto_depIdxs = nil
}
