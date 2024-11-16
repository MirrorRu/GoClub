// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: rooms.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{0}
}

func (x *RoomInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Creating
type RoomCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *RoomInfo `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *RoomCreateRequest) Reset() {
	*x = RoomCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateRequest) ProtoMessage() {}

func (x *RoomCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateRequest.ProtoReflect.Descriptor instead.
func (*RoomCreateRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{1}
}

func (x *RoomCreateRequest) GetObject() *RoomInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type RoomCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crud_result,json=crudResult,proto3" json:"crud_result,omitempty"`
	Object     *RoomInfo   `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *RoomCreateResponse) Reset() {
	*x = RoomCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateResponse) ProtoMessage() {}

func (x *RoomCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateResponse.ProtoReflect.Descriptor instead.
func (*RoomCreateResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{2}
}

func (x *RoomCreateResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *RoomCreateResponse) GetObject() *RoomInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

// Updating
type RoomUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *RoomInfo `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *RoomUpdateRequest) Reset() {
	*x = RoomUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomUpdateRequest) ProtoMessage() {}

func (x *RoomUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoomUpdateRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{3}
}

func (x *RoomUpdateRequest) GetObject() *RoomInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type RoomUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crud_result,json=crudResult,proto3" json:"crud_result,omitempty"`
	Object     *RoomInfo   `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *RoomUpdateResponse) Reset() {
	*x = RoomUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomUpdateResponse) ProtoMessage() {}

func (x *RoomUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomUpdateResponse.ProtoReflect.Descriptor instead.
func (*RoomUpdateResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{4}
}

func (x *RoomUpdateResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *RoomUpdateResponse) GetObject() *RoomInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

// Deleting
type RoomDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoomDeleteRequest) Reset() {
	*x = RoomDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDeleteRequest) ProtoMessage() {}

func (x *RoomDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDeleteRequest.ProtoReflect.Descriptor instead.
func (*RoomDeleteRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{5}
}

func (x *RoomDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RoomDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crud_result,json=crudResult,proto3" json:"crud_result,omitempty"`
}

func (x *RoomDeleteResponse) Reset() {
	*x = RoomDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDeleteResponse) ProtoMessage() {}

func (x *RoomDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDeleteResponse.ProtoReflect.Descriptor instead.
func (*RoomDeleteResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{6}
}

func (x *RoomDeleteResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

type RoomReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoomReadRequest) Reset() {
	*x = RoomReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReadRequest) ProtoMessage() {}

func (x *RoomReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReadRequest.ProtoReflect.Descriptor instead.
func (*RoomReadRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{7}
}

func (x *RoomReadRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RoomReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crud_result,json=crudResult,proto3" json:"crud_result,omitempty"`
	Object     *RoomInfo   `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *RoomReadResponse) Reset() {
	*x = RoomReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReadResponse) ProtoMessage() {}

func (x *RoomReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReadResponse.ProtoReflect.Descriptor instead.
func (*RoomReadResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{8}
}

func (x *RoomReadResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *RoomReadResponse) GetObject() *RoomInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type RoomListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomListingRequest) Reset() {
	*x = RoomListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListingRequest) ProtoMessage() {}

func (x *RoomListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListingRequest.ProtoReflect.Descriptor instead.
func (*RoomListingRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{9}
}

type RoomListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crud_result,json=crudResult,proto3" json:"crud_result,omitempty"`
	Objects    []*RoomInfo `protobuf:"bytes,2,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *RoomListingResponse) Reset() {
	*x = RoomListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListingResponse) ProtoMessage() {}

func (x *RoomListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListingResponse.ProtoReflect.Descriptor instead.
func (*RoomListingResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{10}
}

func (x *RoomListingResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *RoomListingResponse) GetObjects() []*RoomInfo {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_rooms_proto protoreflect.FileDescriptor

var file_rooms_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67,
	0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x39, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x11,
	0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x63, 0x72, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x47, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x6f,
	0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x0a, 0x0f,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x85, 0x01, 0x0a, 0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55,
	0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8a, 0x01,
	0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x75, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52,
	0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0xd0, 0x04, 0x0a, 0x05, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x71, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x7d, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x71, 0x0a, 0x0b, 0x52, 0x6f,
	0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x16, 0x5a,
	0x14, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rooms_proto_rawDescOnce sync.Once
	file_rooms_proto_rawDescData = file_rooms_proto_rawDesc
)

func file_rooms_proto_rawDescGZIP() []byte {
	file_rooms_proto_rawDescOnce.Do(func() {
		file_rooms_proto_rawDescData = protoimpl.X.CompressGZIP(file_rooms_proto_rawDescData)
	})
	return file_rooms_proto_rawDescData
}

var file_rooms_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rooms_proto_goTypes = []any{
	(*RoomInfo)(nil),            // 0: goclub.engine.v1.RoomInfo
	(*RoomCreateRequest)(nil),   // 1: goclub.engine.v1.RoomCreateRequest
	(*RoomCreateResponse)(nil),  // 2: goclub.engine.v1.RoomCreateResponse
	(*RoomUpdateRequest)(nil),   // 3: goclub.engine.v1.RoomUpdateRequest
	(*RoomUpdateResponse)(nil),  // 4: goclub.engine.v1.RoomUpdateResponse
	(*RoomDeleteRequest)(nil),   // 5: goclub.engine.v1.RoomDeleteRequest
	(*RoomDeleteResponse)(nil),  // 6: goclub.engine.v1.RoomDeleteResponse
	(*RoomReadRequest)(nil),     // 7: goclub.engine.v1.RoomReadRequest
	(*RoomReadResponse)(nil),    // 8: goclub.engine.v1.RoomReadResponse
	(*RoomListingRequest)(nil),  // 9: goclub.engine.v1.RoomListingRequest
	(*RoomListingResponse)(nil), // 10: goclub.engine.v1.RoomListingResponse
	(*CRUDResult)(nil),          // 11: goclub.engine.v1.CRUDResult
}
var file_rooms_proto_depIdxs = []int32{
	0,  // 0: goclub.engine.v1.RoomCreateRequest.object:type_name -> goclub.engine.v1.RoomInfo
	11, // 1: goclub.engine.v1.RoomCreateResponse.crud_result:type_name -> goclub.engine.v1.CRUDResult
	0,  // 2: goclub.engine.v1.RoomCreateResponse.object:type_name -> goclub.engine.v1.RoomInfo
	0,  // 3: goclub.engine.v1.RoomUpdateRequest.object:type_name -> goclub.engine.v1.RoomInfo
	11, // 4: goclub.engine.v1.RoomUpdateResponse.crud_result:type_name -> goclub.engine.v1.CRUDResult
	0,  // 5: goclub.engine.v1.RoomUpdateResponse.object:type_name -> goclub.engine.v1.RoomInfo
	11, // 6: goclub.engine.v1.RoomDeleteResponse.crud_result:type_name -> goclub.engine.v1.CRUDResult
	11, // 7: goclub.engine.v1.RoomReadResponse.crud_result:type_name -> goclub.engine.v1.CRUDResult
	0,  // 8: goclub.engine.v1.RoomReadResponse.object:type_name -> goclub.engine.v1.RoomInfo
	11, // 9: goclub.engine.v1.RoomListingResponse.crud_result:type_name -> goclub.engine.v1.CRUDResult
	0,  // 10: goclub.engine.v1.RoomListingResponse.objects:type_name -> goclub.engine.v1.RoomInfo
	1,  // 11: goclub.engine.v1.Rooms.RoomCreate:input_type -> goclub.engine.v1.RoomCreateRequest
	3,  // 12: goclub.engine.v1.Rooms.RoomUpdate:input_type -> goclub.engine.v1.RoomUpdateRequest
	5,  // 13: goclub.engine.v1.Rooms.RoomDelete:input_type -> goclub.engine.v1.RoomDeleteRequest
	7,  // 14: goclub.engine.v1.Rooms.RoomRead:input_type -> goclub.engine.v1.RoomReadRequest
	9,  // 15: goclub.engine.v1.Rooms.RoomListing:input_type -> goclub.engine.v1.RoomListingRequest
	2,  // 16: goclub.engine.v1.Rooms.RoomCreate:output_type -> goclub.engine.v1.RoomCreateResponse
	4,  // 17: goclub.engine.v1.Rooms.RoomUpdate:output_type -> goclub.engine.v1.RoomUpdateResponse
	6,  // 18: goclub.engine.v1.Rooms.RoomDelete:output_type -> goclub.engine.v1.RoomDeleteResponse
	8,  // 19: goclub.engine.v1.Rooms.RoomRead:output_type -> goclub.engine.v1.RoomReadResponse
	10, // 20: goclub.engine.v1.Rooms.RoomListing:output_type -> goclub.engine.v1.RoomListingResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_rooms_proto_init() }
func file_rooms_proto_init() {
	if File_rooms_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rooms_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RoomInfo); i {
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
		file_rooms_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RoomCreateRequest); i {
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
		file_rooms_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RoomCreateResponse); i {
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
		file_rooms_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RoomUpdateRequest); i {
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
		file_rooms_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RoomUpdateResponse); i {
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
		file_rooms_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RoomDeleteRequest); i {
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
		file_rooms_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RoomDeleteResponse); i {
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
		file_rooms_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RoomReadRequest); i {
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
		file_rooms_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RoomReadResponse); i {
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
		file_rooms_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RoomListingRequest); i {
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
		file_rooms_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*RoomListingResponse); i {
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
			RawDescriptor: file_rooms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rooms_proto_goTypes,
		DependencyIndexes: file_rooms_proto_depIdxs,
		MessageInfos:      file_rooms_proto_msgTypes,
	}.Build()
	File_rooms_proto = out.File
	file_rooms_proto_rawDesc = nil
	file_rooms_proto_goTypes = nil
	file_rooms_proto_depIdxs = nil
}
