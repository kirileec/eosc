// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: master.proto

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

type ListenerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme string `protobuf:"bytes,1,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
	Port   int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *ListenerRequest) Reset() {
	*x = ListenerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerRequest) ProtoMessage() {}

func (x *ListenerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerRequest.ProtoReflect.Descriptor instead.
func (*ListenerRequest) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *ListenerRequest) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ListenerRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type AcceptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	FD      []byte `protobuf:"bytes,2,opt,name=FD,proto3" json:"FD,omitempty"`
	Scheme  string `protobuf:"bytes,3,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
	Port    int32  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Message string `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AcceptResponse) Reset() {
	*x = AcceptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptResponse) ProtoMessage() {}

func (x *AcceptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptResponse.ProtoReflect.Descriptor instead.
func (*AcceptResponse) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AcceptResponse) GetFD() []byte {
	if x != nil {
		return x.FD
	}
	return nil
}

func (x *AcceptResponse) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AcceptResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AcceptResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FileHandler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	FD    uint64 `protobuf:"varint,2,opt,name=FD,proto3" json:"FD,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Path  string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
	Error string `protobuf:"bytes,5,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *FileHandler) Reset() {
	*x = FileHandler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHandler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHandler) ProtoMessage() {}

func (x *FileHandler) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHandler.ProtoReflect.Descriptor instead.
func (*FileHandler) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{2}
}

func (x *FileHandler) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FileHandler) GetFD() uint64 {
	if x != nil {
		return x.FD
	}
	return 0
}

func (x *FileHandler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileHandler) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileHandler) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FileOpen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Flag     int32  `protobuf:"varint,2,opt,name=flag,proto3" json:"flag,omitempty"`
	FileMode uint32 `protobuf:"varint,3,opt,name=FileMode,proto3" json:"FileMode,omitempty"`
}

func (x *FileOpen) Reset() {
	*x = FileOpen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOpen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOpen) ProtoMessage() {}

func (x *FileOpen) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOpen.ProtoReflect.Descriptor instead.
func (*FileOpen) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{3}
}

func (x *FileOpen) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileOpen) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *FileOpen) GetFileMode() uint32 {
	if x != nil {
		return x.FileMode
	}
	return 0
}

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x7e, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x46, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x46, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x46, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x56, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x32,
	0x7c, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x04, 0x4f, 0x70,
	0x65, 0x6e, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x65, 0x6e, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x00, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x69,
	0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData = file_master_proto_rawDesc
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_proto_rawDescData)
	})
	return file_master_proto_rawDescData
}

var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_master_proto_goTypes = []interface{}{
	(*ListenerRequest)(nil), // 0: service.ListenerRequest
	(*AcceptResponse)(nil),  // 1: service.AcceptResponse
	(*FileHandler)(nil),     // 2: service.FileHandler
	(*FileOpen)(nil),        // 3: service.FileOpen
}
var file_master_proto_depIdxs = []int32{
	0, // 0: service.Master.Accept:input_type -> service.ListenerRequest
	3, // 1: service.Master.Open:input_type -> service.FileOpen
	1, // 2: service.Master.Accept:output_type -> service.AcceptResponse
	2, // 3: service.Master.Open:output_type -> service.FileHandler
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerRequest); i {
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
		file_master_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptResponse); i {
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
		file_master_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHandler); i {
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
		file_master_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOpen); i {
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
			RawDescriptor: file_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_rawDesc = nil
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}
