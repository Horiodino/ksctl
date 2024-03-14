// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: agent.proto

package pb

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

type ScaleOperation int32

const (
	ScaleOperation_UP   ScaleOperation = 0
	ScaleOperation_DOWN ScaleOperation = 1
)

// Enum value maps for ScaleOperation.
var (
	ScaleOperation_name = map[int32]string{
		0: "UP",
		1: "DOWN",
	}
	ScaleOperation_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
	}
)

func (x ScaleOperation) Enum() *ScaleOperation {
	p := new(ScaleOperation)
	*p = x
	return p
}

func (x ScaleOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScaleOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_proto_enumTypes[0].Descriptor()
}

func (ScaleOperation) Type() protoreflect.EnumType {
	return &file_agent_proto_enumTypes[0]
}

func (x ScaleOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScaleOperation.Descriptor instead.
func (ScaleOperation) EnumDescriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

type ApplicationOperation int32

const (
	ApplicationOperation_CREATE ApplicationOperation = 0
	ApplicationOperation_DELETE ApplicationOperation = 1
	ApplicationOperation_UPDATE ApplicationOperation = 2
)

// Enum value maps for ApplicationOperation.
var (
	ApplicationOperation_name = map[int32]string{
		0: "CREATE",
		1: "DELETE",
		2: "UPDATE",
	}
	ApplicationOperation_value = map[string]int32{
		"CREATE": 0,
		"DELETE": 1,
		"UPDATE": 2,
	}
)

func (x ApplicationOperation) Enum() *ApplicationOperation {
	p := new(ApplicationOperation)
	*p = x
	return p
}

func (x ApplicationOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_proto_enumTypes[1].Descriptor()
}

func (ApplicationOperation) Type() protoreflect.EnumType {
	return &file_agent_proto_enumTypes[1]
}

func (x ApplicationOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationOperation.Descriptor instead.
func (ApplicationOperation) EnumDescriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

// TODO: ReqScale we need to think of were should the ksctl agent gets its data about the cluster?
//
//	from
//	 1. inside this will read from a configmap
//	 2. get from the caller aka grpc client
type ReqScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation ScaleOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=ksctlAgent.ScaleOperation" json:"operation,omitempty"`
	ScaleTo   uint32         `protobuf:"varint,2,opt,name=scaleTo,proto3" json:"scaleTo,omitempty"`
}

func (x *ReqScale) Reset() {
	*x = ReqScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqScale) ProtoMessage() {}

func (x *ReqScale) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqScale.ProtoReflect.Descriptor instead.
func (*ReqScale) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *ReqScale) GetOperation() ScaleOperation {
	if x != nil {
		return x.Operation
	}
	return ScaleOperation_UP
}

func (x *ReqScale) GetScaleTo() uint32 {
	if x != nil {
		return x.ScaleTo
	}
	return 0
}

type ResScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedWP uint32 `protobuf:"varint,1,opt,name=updatedWP,proto3" json:"updatedWP,omitempty"`
}

func (x *ResScale) Reset() {
	*x = ResScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResScale) ProtoMessage() {}

func (x *ResScale) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResScale.ProtoReflect.Descriptor instead.
func (*ResScale) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *ResScale) GetUpdatedWP() uint32 {
	if x != nil {
		return x.UpdatedWP
	}
	return 0
}

type ReqApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation ApplicationOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=ksctlAgent.ApplicationOperation" json:"operation,omitempty"`
	AppName   string               `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
}

func (x *ReqApplication) Reset() {
	*x = ReqApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqApplication) ProtoMessage() {}

func (x *ReqApplication) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqApplication.ProtoReflect.Descriptor instead.
func (*ReqApplication) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *ReqApplication) GetOperation() ApplicationOperation {
	if x != nil {
		return x.Operation
	}
	return ApplicationOperation_CREATE
}

func (x *ReqApplication) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type ResApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResApplication) Reset() {
	*x = ResApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResApplication) ProtoMessage() {}

func (x *ResApplication) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResApplication.ProtoReflect.Descriptor instead.
func (*ResApplication) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

type PortMappings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To       uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Protocol uint64 `protobuf:"varint,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *PortMappings) Reset() {
	*x = PortMappings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMappings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMappings) ProtoMessage() {}

func (x *PortMappings) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMappings.ProtoReflect.Descriptor instead.
func (*PortMappings) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

func (x *PortMappings) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PortMappings) GetTo() uint64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *PortMappings) GetProtocol() uint64 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

type ReqLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// private Ip of controlplanes
	// nodePorts and corresponding port for Loadbalancer svc
	// Protocol
	PrivateIPs []string        `protobuf:"bytes,1,rep,name=privateIPs,proto3" json:"privateIPs,omitempty"`
	Ports      []*PortMappings `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"` // need network, region and apikeys
}

func (x *ReqLB) Reset() {
	*x = ReqLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqLB) ProtoMessage() {}

func (x *ReqLB) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqLB.ProtoReflect.Descriptor instead.
func (*ReqLB) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{5}
}

func (x *ReqLB) GetPrivateIPs() []string {
	if x != nil {
		return x.PrivateIPs
	}
	return nil
}

func (x *ReqLB) GetPorts() []*PortMappings {
	if x != nil {
		return x.Ports
	}
	return nil
}

type ResLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadBalancerPublicIP string `protobuf:"bytes,1,opt,name=loadBalancerPublicIP,proto3" json:"loadBalancerPublicIP,omitempty"`
	CreatedResourceId    string `protobuf:"bytes,2,opt,name=createdResourceId,proto3" json:"createdResourceId,omitempty"`
}

func (x *ResLB) Reset() {
	*x = ResLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResLB) ProtoMessage() {}

func (x *ResLB) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResLB.ProtoReflect.Descriptor instead.
func (*ResLB) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{6}
}

func (x *ResLB) GetLoadBalancerPublicIP() string {
	if x != nil {
		return x.LoadBalancerPublicIP
	}
	return ""
}

func (x *ResLB) GetCreatedResourceId() string {
	if x != nil {
		return x.CreatedResourceId
	}
	return ""
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b,
	0x73, 0x63, 0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x08, 0x52, 0x65, 0x71,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6b, 0x73, 0x63, 0x74, 0x6c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x54, 0x6f, 0x22, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x57, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x57, 0x50, 0x22, 0x6a, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6b, 0x73, 0x63, 0x74, 0x6c,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x10, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x4e, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x22, 0x57, 0x0a, 0x05, 0x52, 0x65, 0x71, 0x4c, 0x42, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49, 0x50, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x73, 0x63, 0x74,
	0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x05, 0x52, 0x65,
	0x73, 0x4c, 0x42, 0x12, 0x32, 0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x22, 0x0a, 0x0e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x2a, 0x3a, 0x0a, 0x14, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0xbe, 0x01, 0x0a, 0x0a, 0x4b, 0x73, 0x63, 0x74, 0x6c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x14, 0x2e,
	0x6b, 0x73, 0x63, 0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x53, 0x63,
	0x61, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x6b, 0x73, 0x63, 0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x4c, 0x6f, 0x61,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x6b, 0x73, 0x63, 0x74,
	0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x42, 0x1a, 0x11, 0x2e, 0x6b,
	0x73, 0x63, 0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x4c, 0x42, 0x12,
	0x45, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x2e, 0x6b, 0x73, 0x63, 0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x6b, 0x73, 0x63,
	0x74, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_agent_proto_goTypes = []interface{}{
	(ScaleOperation)(0),       // 0: ksctlAgent.ScaleOperation
	(ApplicationOperation)(0), // 1: ksctlAgent.ApplicationOperation
	(*ReqScale)(nil),          // 2: ksctlAgent.ReqScale
	(*ResScale)(nil),          // 3: ksctlAgent.ResScale
	(*ReqApplication)(nil),    // 4: ksctlAgent.ReqApplication
	(*ResApplication)(nil),    // 5: ksctlAgent.ResApplication
	(*PortMappings)(nil),      // 6: ksctlAgent.PortMappings
	(*ReqLB)(nil),             // 7: ksctlAgent.ReqLB
	(*ResLB)(nil),             // 8: ksctlAgent.ResLB
}
var file_agent_proto_depIdxs = []int32{
	0, // 0: ksctlAgent.ReqScale.operation:type_name -> ksctlAgent.ScaleOperation
	1, // 1: ksctlAgent.ReqApplication.operation:type_name -> ksctlAgent.ApplicationOperation
	6, // 2: ksctlAgent.ReqLB.ports:type_name -> ksctlAgent.PortMappings
	2, // 3: ksctlAgent.KsctlAgent.Scale:input_type -> ksctlAgent.ReqScale
	7, // 4: ksctlAgent.KsctlAgent.LoadBalancer:input_type -> ksctlAgent.ReqLB
	4, // 5: ksctlAgent.KsctlAgent.Application:input_type -> ksctlAgent.ReqApplication
	3, // 6: ksctlAgent.KsctlAgent.Scale:output_type -> ksctlAgent.ResScale
	8, // 7: ksctlAgent.KsctlAgent.LoadBalancer:output_type -> ksctlAgent.ResLB
	5, // 8: ksctlAgent.KsctlAgent.Application:output_type -> ksctlAgent.ResApplication
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqScale); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResScale); i {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqApplication); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResApplication); i {
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
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMappings); i {
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
		file_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqLB); i {
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
		file_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResLB); i {
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
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		EnumInfos:         file_agent_proto_enumTypes,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
