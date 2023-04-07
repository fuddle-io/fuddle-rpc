// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: registry.proto

package rpc

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

type MemberUpdateType int32

const (
	MemberUpdateType_REGISTER   MemberUpdateType = 0
	MemberUpdateType_UNREGISTER MemberUpdateType = 1
)

// Enum value maps for MemberUpdateType.
var (
	MemberUpdateType_name = map[int32]string{
		0: "REGISTER",
		1: "UNREGISTER",
	}
	MemberUpdateType_value = map[string]int32{
		"REGISTER":   0,
		"UNREGISTER": 1,
	}
)

func (x MemberUpdateType) Enum() *MemberUpdateType {
	p := new(MemberUpdateType)
	*p = x
	return p
}

func (x MemberUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[0].Descriptor()
}

func (MemberUpdateType) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[0]
}

func (x MemberUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberUpdateType.Descriptor instead.
func (MemberUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

type MemberStatus int32

const (
	MemberStatus_UP   MemberStatus = 0
	MemberStatus_DOWN MemberStatus = 1
)

// Enum value maps for MemberStatus.
var (
	MemberStatus_name = map[int32]string{
		0: "UP",
		1: "DOWN",
	}
	MemberStatus_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
	}
)

func (x MemberStatus) Enum() *MemberStatus {
	p := new(MemberStatus)
	*p = x
	return p
}

func (x MemberStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[1].Descriptor()
}

func (MemberStatus) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[1]
}

func (x MemberStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberStatus.Descriptor instead.
func (MemberStatus) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

type ClientUpdateType int32

const (
	ClientUpdateType_CLIENT_REGISTER   ClientUpdateType = 0
	ClientUpdateType_CLIENT_UNREGISTER ClientUpdateType = 1
	ClientUpdateType_CLIENT_HEARTBEAT  ClientUpdateType = 2
)

// Enum value maps for ClientUpdateType.
var (
	ClientUpdateType_name = map[int32]string{
		0: "CLIENT_REGISTER",
		1: "CLIENT_UNREGISTER",
		2: "CLIENT_HEARTBEAT",
	}
	ClientUpdateType_value = map[string]int32{
		"CLIENT_REGISTER":   0,
		"CLIENT_UNREGISTER": 1,
		"CLIENT_HEARTBEAT":  2,
	}
)

func (x ClientUpdateType) Enum() *ClientUpdateType {
	p := new(ClientUpdateType)
	*p = x
	return p
}

func (x ClientUpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientUpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[2].Descriptor()
}

func (ClientUpdateType) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[2]
}

func (x ClientUpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientUpdateType.Descriptor instead.
func (ClientUpdateType) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{2}
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner     string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Counter   uint64 `protobuf:"varint,3,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Version) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Version) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status   MemberStatus      `protobuf:"varint,2,opt,name=status,proto3,enum=registry.MemberStatus" json:"status,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetStatus() MemberStatus {
	if x != nil {
		return x.Status
	}
	return MemberStatus_UP
}

func (x *Member) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type LocalMemberUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType MemberUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=registry.MemberUpdateType" json:"update_type,omitempty"`
	Member     *Member          `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *LocalMemberUpdate) Reset() {
	*x = LocalMemberUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalMemberUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalMemberUpdate) ProtoMessage() {}

func (x *LocalMemberUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalMemberUpdate.ProtoReflect.Descriptor instead.
func (*LocalMemberUpdate) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{2}
}

func (x *LocalMemberUpdate) GetUpdateType() MemberUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return MemberUpdateType_REGISTER
}

func (x *LocalMemberUpdate) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type RemoteMemberUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType MemberUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=registry.MemberUpdateType" json:"update_type,omitempty"`
	Member     *Member          `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	Version    *Version         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *RemoteMemberUpdate) Reset() {
	*x = RemoteMemberUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteMemberUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteMemberUpdate) ProtoMessage() {}

func (x *RemoteMemberUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteMemberUpdate.ProtoReflect.Descriptor instead.
func (*RemoteMemberUpdate) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteMemberUpdate) GetUpdateType() MemberUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return MemberUpdateType_REGISTER
}

func (x *RemoteMemberUpdate) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *RemoteMemberUpdate) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KnownMembers map[string]*Version `protobuf:"bytes,1,rep,name=known_members,json=knownMembers,proto3" json:"known_members,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OwnerOnly    bool                `protobuf:"varint,2,opt,name=owner_only,json=ownerOnly,proto3" json:"owner_only,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequest) GetKnownMembers() map[string]*Version {
	if x != nil {
		return x.KnownMembers
	}
	return nil
}

func (x *SubscribeRequest) GetOwnerOnly() bool {
	if x != nil {
		return x.OwnerOnly
	}
	return false
}

type ClientUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType ClientUpdateType `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=registry.ClientUpdateType" json:"update_type,omitempty"`
	Member     *Member          `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	SeqId      uint64           `protobuf:"varint,3,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
}

func (x *ClientUpdate) Reset() {
	*x = ClientUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientUpdate) ProtoMessage() {}

func (x *ClientUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientUpdate.ProtoReflect.Descriptor instead.
func (*ClientUpdate) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{5}
}

func (x *ClientUpdate) GetUpdateType() ClientUpdateType {
	if x != nil {
		return x.UpdateType
	}
	return ClientUpdateType_CLIENT_REGISTER
}

func (x *ClientUpdate) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *ClientUpdate) GetSeqId() uint64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

type ClientAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqId uint64 `protobuf:"varint,1,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
}

func (x *ClientAck) Reset() {
	*x = ClientAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAck) ProtoMessage() {}

func (x *ClientAck) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAck.ProtoReflect.Descriptor instead.
func (*ClientAck) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{6}
}

func (x *ClientAck) GetSeqId() uint64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

var File_registry_proto protoreflect.FileDescriptor

var file_registry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x57, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7a, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd8,
	0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x0d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x1a, 0x52, 0x0a, 0x11, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x2a, 0x30, 0x0a, 0x10,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x20,
	0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06,
	0x0a, 0x02, 0x55, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x2a, 0x54, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54,
	0x42, 0x45, 0x41, 0x54, 0x10, 0x02, 0x32, 0x90, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x47, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x13, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x64, 0x64, 0x6c, 0x65, 0x2d, 0x69,
	0x6f, 0x2f, 0x66, 0x75, 0x64, 0x64, 0x6c, 0x65, 0x2d, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_proto_rawDescOnce sync.Once
	file_registry_proto_rawDescData = file_registry_proto_rawDesc
)

func file_registry_proto_rawDescGZIP() []byte {
	file_registry_proto_rawDescOnce.Do(func() {
		file_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_proto_rawDescData)
	})
	return file_registry_proto_rawDescData
}

var file_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_registry_proto_goTypes = []interface{}{
	(MemberUpdateType)(0),      // 0: registry.MemberUpdateType
	(MemberStatus)(0),          // 1: registry.MemberStatus
	(ClientUpdateType)(0),      // 2: registry.ClientUpdateType
	(*Version)(nil),            // 3: registry.Version
	(*Member)(nil),             // 4: registry.Member
	(*LocalMemberUpdate)(nil),  // 5: registry.LocalMemberUpdate
	(*RemoteMemberUpdate)(nil), // 6: registry.RemoteMemberUpdate
	(*SubscribeRequest)(nil),   // 7: registry.SubscribeRequest
	(*ClientUpdate)(nil),       // 8: registry.ClientUpdate
	(*ClientAck)(nil),          // 9: registry.ClientAck
	nil,                        // 10: registry.Member.MetadataEntry
	nil,                        // 11: registry.SubscribeRequest.KnownMembersEntry
}
var file_registry_proto_depIdxs = []int32{
	1,  // 0: registry.Member.status:type_name -> registry.MemberStatus
	10, // 1: registry.Member.metadata:type_name -> registry.Member.MetadataEntry
	0,  // 2: registry.LocalMemberUpdate.update_type:type_name -> registry.MemberUpdateType
	4,  // 3: registry.LocalMemberUpdate.member:type_name -> registry.Member
	0,  // 4: registry.RemoteMemberUpdate.update_type:type_name -> registry.MemberUpdateType
	4,  // 5: registry.RemoteMemberUpdate.member:type_name -> registry.Member
	3,  // 6: registry.RemoteMemberUpdate.version:type_name -> registry.Version
	11, // 7: registry.SubscribeRequest.known_members:type_name -> registry.SubscribeRequest.KnownMembersEntry
	2,  // 8: registry.ClientUpdate.update_type:type_name -> registry.ClientUpdateType
	4,  // 9: registry.ClientUpdate.member:type_name -> registry.Member
	3,  // 10: registry.SubscribeRequest.KnownMembersEntry.value:type_name -> registry.Version
	7,  // 11: registry.Registry.Subscribe:input_type -> registry.SubscribeRequest
	8,  // 12: registry.Registry.Register:input_type -> registry.ClientUpdate
	6,  // 13: registry.Registry.Subscribe:output_type -> registry.RemoteMemberUpdate
	9,  // 14: registry.Registry.Register:output_type -> registry.ClientAck
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_registry_proto_init() }
func file_registry_proto_init() {
	if File_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalMemberUpdate); i {
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
		file_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteMemberUpdate); i {
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
		file_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientUpdate); i {
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
		file_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAck); i {
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
			RawDescriptor: file_registry_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_proto_goTypes,
		DependencyIndexes: file_registry_proto_depIdxs,
		EnumInfos:         file_registry_proto_enumTypes,
		MessageInfos:      file_registry_proto_msgTypes,
	}.Build()
	File_registry_proto = out.File
	file_registry_proto_rawDesc = nil
	file_registry_proto_goTypes = nil
	file_registry_proto_depIdxs = nil
}
