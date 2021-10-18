// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: worker.proto

package service

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

type WorkerStatusCode int32

const (
	WorkerStatusCode_SUCCESS WorkerStatusCode = 0
	WorkerStatusCode_FAIL    WorkerStatusCode = 1
	WorkerStatusCode_SKILL   WorkerStatusCode = 2
)

// Enum value maps for WorkerStatusCode.
var (
	WorkerStatusCode_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
		2: "SKILL",
	}
	WorkerStatusCode_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
		"SKILL":   2,
	}
)

func (x WorkerStatusCode) Enum() *WorkerStatusCode {
	p := new(WorkerStatusCode)
	*p = x
	return p
}

func (x WorkerStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_worker_proto_enumTypes[0].Descriptor()
}

func (WorkerStatusCode) Type() protoreflect.EnumType {
	return &file_worker_proto_enumTypes[0]
}

func (x WorkerStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerStatusCode.Descriptor instead.
func (WorkerStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

type WorkerHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *WorkerHelloRequest) Reset() {
	*x = WorkerHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerHelloRequest) ProtoMessage() {}

func (x *WorkerHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerHelloRequest.ProtoReflect.Descriptor instead.
func (*WorkerHelloRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerHelloRequest) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type WorkerHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello    string          `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	Resource *WorkerResource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *WorkerHelloResponse) Reset() {
	*x = WorkerHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerHelloResponse) ProtoMessage() {}

func (x *WorkerHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerHelloResponse.ProtoReflect.Descriptor instead.
func (*WorkerHelloResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerHelloResponse) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

func (x *WorkerHelloResponse) GetResource() *WorkerResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type WorkerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkerDeleteRequest) Reset() {
	*x = WorkerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeleteRequest) ProtoMessage() {}

func (x *WorkerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeleteRequest.ProtoReflect.Descriptor instead.
func (*WorkerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WorkerSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Profession string `protobuf:"bytes,2,opt,name=profession,proto3" json:"profession,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Driver     string `protobuf:"bytes,4,opt,name=driver,proto3" json:"driver,omitempty"`
	Body       []byte `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *WorkerSetRequest) Reset() {
	*x = WorkerSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSetRequest) ProtoMessage() {}

func (x *WorkerSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerSetRequest.ProtoReflect.Descriptor instead.
func (*WorkerSetRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *WorkerSetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkerSetRequest) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *WorkerSetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkerSetRequest) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *WorkerSetRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type WorkerResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port []int32 `protobuf:"varint,1,rep,packed,name=port,proto3" json:"port,omitempty"`
}

func (x *WorkerResource) Reset() {
	*x = WorkerResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerResource) ProtoMessage() {}

func (x *WorkerResource) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerResource.ProtoReflect.Descriptor instead.
func (*WorkerResource) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *WorkerResource) GetPort() []int32 {
	if x != nil {
		return x.Port
	}
	return nil
}

type WorkerDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   WorkerStatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=service.WorkerStatusCode" json:"status,omitempty"`
	Message  string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Resource *WorkerResource  `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *WorkerDeleteResponse) Reset() {
	*x = WorkerDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeleteResponse) ProtoMessage() {}

func (x *WorkerDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeleteResponse.ProtoReflect.Descriptor instead.
func (*WorkerDeleteResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

func (x *WorkerDeleteResponse) GetStatus() WorkerStatusCode {
	if x != nil {
		return x.Status
	}
	return WorkerStatusCode_SUCCESS
}

func (x *WorkerDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *WorkerDeleteResponse) GetResource() *WorkerResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type WorkerSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   WorkerStatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=service.WorkerStatusCode" json:"status,omitempty"`
	Message  string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Resource *WorkerResource  `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *WorkerSetResponse) Reset() {
	*x = WorkerSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSetResponse) ProtoMessage() {}

func (x *WorkerSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerSetResponse.ProtoReflect.Descriptor instead.
func (*WorkerSetResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

func (x *WorkerSetResponse) GetStatus() WorkerStatusCode {
	if x != nil {
		return x.Status
	}
	return WorkerStatusCode_SUCCESS
}

func (x *WorkerSetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *WorkerSetResponse) GetResource() *WorkerResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type WorkerRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []int32 `protobuf:"varint,1,rep,packed,name=ports,proto3" json:"ports,omitempty"`
}

func (x *WorkerRefreshRequest) Reset() {
	*x = WorkerRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerRefreshRequest) ProtoMessage() {}

func (x *WorkerRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerRefreshRequest.ProtoReflect.Descriptor instead.
func (*WorkerRefreshRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

func (x *WorkerRefreshRequest) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

type WorkerRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  WorkerStatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=service.WorkerStatusCode" json:"status,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WorkerRefreshResponse) Reset() {
	*x = WorkerRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerRefreshResponse) ProtoMessage() {}

func (x *WorkerRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerRefreshResponse.ProtoReflect.Descriptor instead.
func (*WorkerRefreshResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *WorkerRefreshResponse) GetStatus() WorkerStatusCode {
	if x != nil {
		return x.Status
	}
	return WorkerStatusCode_SUCCESS
}

func (x *WorkerRefreshResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x22, 0x60, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x24, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x64, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x34,
	0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4b, 0x49,
	0x4c, 0x4c, 0x10, 0x02, 0x32, 0xbc, 0x03, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_worker_proto_goTypes = []interface{}{
	(WorkerStatusCode)(0),         // 0: service.WorkerStatusCode
	(*WorkerHelloRequest)(nil),    // 1: service.WorkerHelloRequest
	(*WorkerHelloResponse)(nil),   // 2: service.WorkerHelloResponse
	(*WorkerDeleteRequest)(nil),   // 3: service.WorkerDeleteRequest
	(*WorkerSetRequest)(nil),      // 4: service.WorkerSetRequest
	(*WorkerResource)(nil),        // 5: service.WorkerResource
	(*WorkerDeleteResponse)(nil),  // 6: service.WorkerDeleteResponse
	(*WorkerSetResponse)(nil),     // 7: service.WorkerSetResponse
	(*WorkerRefreshRequest)(nil),  // 8: service.WorkerRefreshRequest
	(*WorkerRefreshResponse)(nil), // 9: service.WorkerRefreshResponse
}
var file_worker_proto_depIdxs = []int32{
	5,  // 0: service.WorkerHelloResponse.resource:type_name -> service.WorkerResource
	0,  // 1: service.WorkerDeleteResponse.status:type_name -> service.WorkerStatusCode
	5,  // 2: service.WorkerDeleteResponse.resource:type_name -> service.WorkerResource
	0,  // 3: service.WorkerSetResponse.status:type_name -> service.WorkerStatusCode
	5,  // 4: service.WorkerSetResponse.resource:type_name -> service.WorkerResource
	0,  // 5: service.WorkerRefreshResponse.status:type_name -> service.WorkerStatusCode
	3,  // 6: service.WorkerService.deleteCheck:input_type -> service.WorkerDeleteRequest
	4,  // 7: service.WorkerService.setCheck:input_type -> service.WorkerSetRequest
	3,  // 8: service.WorkerService.delete:input_type -> service.WorkerDeleteRequest
	4,  // 9: service.WorkerService.set:input_type -> service.WorkerSetRequest
	1,  // 10: service.WorkerService.ping:input_type -> service.WorkerHelloRequest
	8,  // 11: service.WorkerService.refresh:input_type -> service.WorkerRefreshRequest
	6,  // 12: service.WorkerService.deleteCheck:output_type -> service.WorkerDeleteResponse
	7,  // 13: service.WorkerService.setCheck:output_type -> service.WorkerSetResponse
	6,  // 14: service.WorkerService.delete:output_type -> service.WorkerDeleteResponse
	7,  // 15: service.WorkerService.set:output_type -> service.WorkerSetResponse
	2,  // 16: service.WorkerService.ping:output_type -> service.WorkerHelloResponse
	9,  // 17: service.WorkerService.refresh:output_type -> service.WorkerRefreshResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerHelloRequest); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerHelloResponse); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDeleteRequest); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerSetRequest); i {
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
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerResource); i {
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
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDeleteResponse); i {
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
		file_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerSetResponse); i {
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
		file_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerRefreshRequest); i {
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
		file_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerRefreshResponse); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		EnumInfos:         file_worker_proto_enumTypes,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
