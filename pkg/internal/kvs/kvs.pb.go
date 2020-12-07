// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: kvs.proto

package kvs

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestType int32

const (
	RequestType_GET RequestType = 0
	RequestType_PUT RequestType = 1
	RequestType_DEL RequestType = 2
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "GET",
		1: "PUT",
		2: "DEL",
	}
	RequestType_value = map[string]int32{
		"GET": 0,
		"PUT": 1,
		"DEL": 2,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_kvs_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_kvs_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{0}
}

type ErrorCode int32

const (
	ErrorCode_NOT_FOUND         ErrorCode = 0
	ErrorCode_BAD_REQUEST       ErrorCode = 1
	ErrorCode_TOO_MANY_REQUESTS ErrorCode = 2
	ErrorCode_INTERAL_ERROR     ErrorCode = 3
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "NOT_FOUND",
		1: "BAD_REQUEST",
		2: "TOO_MANY_REQUESTS",
		3: "INTERAL_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"NOT_FOUND":         0,
		"BAD_REQUEST":       1,
		"TOO_MANY_REQUESTS": 2,
		"INTERAL_ERROR":     3,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_kvs_proto_enumTypes[1].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_kvs_proto_enumTypes[1]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{1}
}

type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValuePair) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=kvs.ErrorCode" json:"code,omitempty"`
	Error string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_NOT_FOUND
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{3}
}

func (x *DelRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values [][]byte  `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Error  *Response `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetResponse) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *GetResponse) GetError() *Response {
	if x != nil {
		return x.Error
	}
	return nil
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{5}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`            // Which page number do we want?
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"` // Number of results to return per page.
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kvs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_kvs_proto_rawDescGZIP(), []int{6}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchRequest) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

var File_kvs_proto protoreflect.FileDescriptor

var file_kvs_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6b, 0x76, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6b, 0x76, 0x73,
	0x22, 0x36, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6b, 0x76, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x1e,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1e,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5c,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x76, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x0a,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x6e, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x2a, 0x28, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55,
	0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x4c, 0x10, 0x02, 0x2a, 0x55, 0x0a, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x4f,
	0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x03, 0x32, 0x84, 0x01, 0x0a, 0x0a, 0x4b, 0x56, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x6b, 0x76, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6b, 0x76, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03,
	0x50, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x6b, 0x76, 0x73, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6b, 0x76, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x0f, 0x2e, 0x6b, 0x76, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6b, 0x76,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kvs_proto_rawDescOnce sync.Once
	file_kvs_proto_rawDescData = file_kvs_proto_rawDesc
)

func file_kvs_proto_rawDescGZIP() []byte {
	file_kvs_proto_rawDescOnce.Do(func() {
		file_kvs_proto_rawDescData = protoimpl.X.CompressGZIP(file_kvs_proto_rawDescData)
	})
	return file_kvs_proto_rawDescData
}

var file_kvs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kvs_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_kvs_proto_goTypes = []interface{}{
	(RequestType)(0),      // 0: kvs.RequestType
	(ErrorCode)(0),        // 1: kvs.ErrorCode
	(*KeyValuePair)(nil),  // 2: kvs.KeyValuePair
	(*Response)(nil),      // 3: kvs.Response
	(*GetRequest)(nil),    // 4: kvs.GetRequest
	(*DelRequest)(nil),    // 5: kvs.DelRequest
	(*GetResponse)(nil),   // 6: kvs.GetResponse
	(*PutRequest)(nil),    // 7: kvs.PutRequest
	(*SearchRequest)(nil), // 8: kvs.SearchRequest
}
var file_kvs_proto_depIdxs = []int32{
	1, // 0: kvs.Response.code:type_name -> kvs.ErrorCode
	3, // 1: kvs.GetResponse.error:type_name -> kvs.Response
	4, // 2: kvs.KVSService.Get:input_type -> kvs.GetRequest
	7, // 3: kvs.KVSService.Put:input_type -> kvs.PutRequest
	5, // 4: kvs.KVSService.Del:input_type -> kvs.DelRequest
	6, // 5: kvs.KVSService.Get:output_type -> kvs.GetResponse
	3, // 6: kvs.KVSService.Put:output_type -> kvs.Response
	3, // 7: kvs.KVSService.Del:output_type -> kvs.Response
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kvs_proto_init() }
func file_kvs_proto_init() {
	if File_kvs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kvs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValuePair); i {
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
		file_kvs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_kvs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_kvs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_kvs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_kvs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_kvs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
			RawDescriptor: file_kvs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kvs_proto_goTypes,
		DependencyIndexes: file_kvs_proto_depIdxs,
		EnumInfos:         file_kvs_proto_enumTypes,
		MessageInfos:      file_kvs_proto_msgTypes,
	}.Build()
	File_kvs_proto = out.File
	file_kvs_proto_rawDesc = nil
	file_kvs_proto_goTypes = nil
	file_kvs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVSServiceClient is the client API for KVSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVSServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Response, error)
}

type kVSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKVSServiceClient(cc grpc.ClientConnInterface) KVSServiceClient {
	return &kVSServiceClient{cc}
}

func (c *kVSServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kvs.KVSService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/kvs.KVSService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSServiceClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/kvs.KVSService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVSServiceServer is the server API for KVSService service.
type KVSServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*Response, error)
	Del(context.Context, *DelRequest) (*Response, error)
}

// UnimplementedKVSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKVSServiceServer struct {
}

func (*UnimplementedKVSServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVSServiceServer) Put(context.Context, *PutRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedKVSServiceServer) Del(context.Context, *DelRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterKVSServiceServer(s *grpc.Server, srv KVSServiceServer) {
	s.RegisterService(&_KVSService_serviceDesc, srv)
}

func _KVSService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVSService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVSService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVSService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVSService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVSService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServiceServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvs.KVSService",
	HandlerType: (*KVSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVSService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVSService_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _KVSService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvs.proto",
}
