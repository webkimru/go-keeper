// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.2
// source: gokeeper_data_keyvalue.proto

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

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type AddKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *KeyValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddKeyValueRequest) Reset() {
	*x = AddKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyValueRequest) ProtoMessage() {}

func (x *AddKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyValueRequest.ProtoReflect.Descriptor instead.
func (*AddKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{1}
}

func (x *AddKeyValueRequest) GetData() *KeyValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddKeyValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddKeyValueResponse) Reset() {
	*x = AddKeyValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyValueResponse) ProtoMessage() {}

func (x *AddKeyValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyValueResponse.ProtoReflect.Descriptor instead.
func (*AddKeyValueResponse) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{2}
}

func (x *AddKeyValueResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetKeyValueRequest) Reset() {
	*x = GetKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyValueRequest) ProtoMessage() {}

func (x *GetKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyValueRequest.ProtoReflect.Descriptor instead.
func (*GetKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{3}
}

func (x *GetKeyValueRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetKeyValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *KeyValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetKeyValueResponse) Reset() {
	*x = GetKeyValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyValueResponse) ProtoMessage() {}

func (x *GetKeyValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyValueResponse.ProtoReflect.Descriptor instead.
func (*GetKeyValueResponse) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{4}
}

func (x *GetKeyValueResponse) GetData() *KeyValue {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetKeyValueResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListKeyValueRequest) Reset() {
	*x = ListKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyValueRequest) ProtoMessage() {}

func (x *ListKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyValueRequest.ProtoReflect.Descriptor instead.
func (*ListKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{5}
}

func (x *ListKeyValueRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListKeyValueRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListKeyValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Data  []*KeyValue `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListKeyValueResponse) Reset() {
	*x = ListKeyValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeyValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeyValueResponse) ProtoMessage() {}

func (x *ListKeyValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeyValueResponse.ProtoReflect.Descriptor instead.
func (*ListKeyValueResponse) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{6}
}

func (x *ListKeyValueResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListKeyValueResponse) GetData() []*KeyValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type DelKeyValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelKeyValueRequest) Reset() {
	*x = DelKeyValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelKeyValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelKeyValueRequest) ProtoMessage() {}

func (x *DelKeyValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelKeyValueRequest.ProtoReflect.Descriptor instead.
func (*DelKeyValueRequest) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{7}
}

func (x *DelKeyValueRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelKeyValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DelKeyValueResponse) Reset() {
	*x = DelKeyValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gokeeper_data_keyvalue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelKeyValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelKeyValueResponse) ProtoMessage() {}

func (x *DelKeyValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gokeeper_data_keyvalue_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelKeyValueResponse.ProtoReflect.Descriptor instead.
func (*DelKeyValueResponse) Descriptor() ([]byte, []int) {
	return file_gokeeper_data_keyvalue_proto_rawDescGZIP(), []int{8}
}

func (x *DelKeyValueResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_gokeeper_data_keyvalue_proto protoreflect.FileDescriptor

var file_gokeeper_data_keyvalue_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6b, 0x69, 0x6d, 0x2e, 0x67, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x08,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x69, 0x6d,
	0x2e, 0x67, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6b, 0x69, 0x6d, 0x2e, 0x67, 0x6f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x69, 0x6d, 0x2e, 0x67, 0x6f, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gokeeper_data_keyvalue_proto_rawDescOnce sync.Once
	file_gokeeper_data_keyvalue_proto_rawDescData = file_gokeeper_data_keyvalue_proto_rawDesc
)

func file_gokeeper_data_keyvalue_proto_rawDescGZIP() []byte {
	file_gokeeper_data_keyvalue_proto_rawDescOnce.Do(func() {
		file_gokeeper_data_keyvalue_proto_rawDescData = protoimpl.X.CompressGZIP(file_gokeeper_data_keyvalue_proto_rawDescData)
	})
	return file_gokeeper_data_keyvalue_proto_rawDescData
}

var file_gokeeper_data_keyvalue_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gokeeper_data_keyvalue_proto_goTypes = []interface{}{
	(*KeyValue)(nil),             // 0: kim.gokeeper.KeyValue
	(*AddKeyValueRequest)(nil),   // 1: kim.gokeeper.AddKeyValueRequest
	(*AddKeyValueResponse)(nil),  // 2: kim.gokeeper.AddKeyValueResponse
	(*GetKeyValueRequest)(nil),   // 3: kim.gokeeper.GetKeyValueRequest
	(*GetKeyValueResponse)(nil),  // 4: kim.gokeeper.GetKeyValueResponse
	(*ListKeyValueRequest)(nil),  // 5: kim.gokeeper.ListKeyValueRequest
	(*ListKeyValueResponse)(nil), // 6: kim.gokeeper.ListKeyValueResponse
	(*DelKeyValueRequest)(nil),   // 7: kim.gokeeper.DelKeyValueRequest
	(*DelKeyValueResponse)(nil),  // 8: kim.gokeeper.DelKeyValueResponse
}
var file_gokeeper_data_keyvalue_proto_depIdxs = []int32{
	0, // 0: kim.gokeeper.AddKeyValueRequest.data:type_name -> kim.gokeeper.KeyValue
	0, // 1: kim.gokeeper.GetKeyValueResponse.data:type_name -> kim.gokeeper.KeyValue
	0, // 2: kim.gokeeper.ListKeyValueResponse.data:type_name -> kim.gokeeper.KeyValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gokeeper_data_keyvalue_proto_init() }
func file_gokeeper_data_keyvalue_proto_init() {
	if File_gokeeper_data_keyvalue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gokeeper_data_keyvalue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyValueRequest); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyValueResponse); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyValueRequest); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyValueResponse); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyValueRequest); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeyValueResponse); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelKeyValueRequest); i {
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
		file_gokeeper_data_keyvalue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelKeyValueResponse); i {
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
			RawDescriptor: file_gokeeper_data_keyvalue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gokeeper_data_keyvalue_proto_goTypes,
		DependencyIndexes: file_gokeeper_data_keyvalue_proto_depIdxs,
		MessageInfos:      file_gokeeper_data_keyvalue_proto_msgTypes,
	}.Build()
	File_gokeeper_data_keyvalue_proto = out.File
	file_gokeeper_data_keyvalue_proto_rawDesc = nil
	file_gokeeper_data_keyvalue_proto_goTypes = nil
	file_gokeeper_data_keyvalue_proto_depIdxs = nil
}
