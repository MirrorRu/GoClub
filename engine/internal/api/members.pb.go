// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: members.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CRUDResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64 `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
}

func (x *CRUDResult) Reset() {
	*x = CRUDResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRUDResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRUDResult) ProtoMessage() {}

func (x *CRUDResult) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRUDResult.ProtoReflect.Descriptor instead.
func (*CRUDResult) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{0}
}

func (x *CRUDResult) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type MemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId int64                  `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	FullName string                 `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Birthday *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *MemberInfo) Reset() {
	*x = MemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberInfo) ProtoMessage() {}

func (x *MemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberInfo.ProtoReflect.Descriptor instead.
func (*MemberInfo) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{1}
}

func (x *MemberInfo) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberInfo) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *MemberInfo) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

type MemberCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *MemberInfo `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *MemberCreateRequest) Reset() {
	*x = MemberCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCreateRequest) ProtoMessage() {}

func (x *MemberCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCreateRequest.ProtoReflect.Descriptor instead.
func (*MemberCreateRequest) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{2}
}

func (x *MemberCreateRequest) GetObject() *MemberInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type MemberCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crudResult,proto3" json:"crudResult,omitempty"`
	MemberId   int64       `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *MemberCreateResponse) Reset() {
	*x = MemberCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCreateResponse) ProtoMessage() {}

func (x *MemberCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCreateResponse.ProtoReflect.Descriptor instead.
func (*MemberCreateResponse) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{3}
}

func (x *MemberCreateResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *MemberCreateResponse) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type MemberUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *MemberInfo `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *MemberUpdateRequest) Reset() {
	*x = MemberUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberUpdateRequest) ProtoMessage() {}

func (x *MemberUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberUpdateRequest.ProtoReflect.Descriptor instead.
func (*MemberUpdateRequest) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{4}
}

func (x *MemberUpdateRequest) GetObject() *MemberInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type MemberUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crudResult,proto3" json:"crudResult,omitempty"`
}

func (x *MemberUpdateResponse) Reset() {
	*x = MemberUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberUpdateResponse) ProtoMessage() {}

func (x *MemberUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberUpdateResponse.ProtoReflect.Descriptor instead.
func (*MemberUpdateResponse) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{5}
}

func (x *MemberUpdateResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

type MemberDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId int64 `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *MemberDeleteRequest) Reset() {
	*x = MemberDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberDeleteRequest) ProtoMessage() {}

func (x *MemberDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberDeleteRequest.ProtoReflect.Descriptor instead.
func (*MemberDeleteRequest) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{6}
}

func (x *MemberDeleteRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type MemberDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crudResult,proto3" json:"crudResult,omitempty"`
}

func (x *MemberDeleteResponse) Reset() {
	*x = MemberDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberDeleteResponse) ProtoMessage() {}

func (x *MemberDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberDeleteResponse.ProtoReflect.Descriptor instead.
func (*MemberDeleteResponse) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{7}
}

func (x *MemberDeleteResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

type MemberReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId int64 `protobuf:"varint,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *MemberReadRequest) Reset() {
	*x = MemberReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberReadRequest) ProtoMessage() {}

func (x *MemberReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberReadRequest.ProtoReflect.Descriptor instead.
func (*MemberReadRequest) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{8}
}

func (x *MemberReadRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type MemberReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult `protobuf:"bytes,1,opt,name=crudResult,proto3" json:"crudResult,omitempty"`
	Object     *MemberInfo `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *MemberReadResponse) Reset() {
	*x = MemberReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberReadResponse) ProtoMessage() {}

func (x *MemberReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberReadResponse.ProtoReflect.Descriptor instead.
func (*MemberReadResponse) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{9}
}

func (x *MemberReadResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *MemberReadResponse) GetObject() *MemberInfo {
	if x != nil {
		return x.Object
	}
	return nil
}

type MemberListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *MemberListingRequest_ListingFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *MemberListingRequest) Reset() {
	*x = MemberListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListingRequest) ProtoMessage() {}

func (x *MemberListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListingRequest.ProtoReflect.Descriptor instead.
func (*MemberListingRequest) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{10}
}

func (x *MemberListingRequest) GetFilter() *MemberListingRequest_ListingFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type MemberListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrudResult *CRUDResult   `protobuf:"bytes,1,opt,name=crudResult,proto3" json:"crudResult,omitempty"`
	Objects    []*MemberInfo `protobuf:"bytes,2,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *MemberListingResponse) Reset() {
	*x = MemberListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListingResponse) ProtoMessage() {}

func (x *MemberListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListingResponse.ProtoReflect.Descriptor instead.
func (*MemberListingResponse) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{11}
}

func (x *MemberListingResponse) GetCrudResult() *CRUDResult {
	if x != nil {
		return x.CrudResult
	}
	return nil
}

func (x *MemberListingResponse) GetObjects() []*MemberInfo {
	if x != nil {
		return x.Objects
	}
	return nil
}

type MemberListingRequest_ListingFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullNamePart string `protobuf:"bytes,1,opt,name=fullNamePart,proto3" json:"fullNamePart,omitempty"`
}

func (x *MemberListingRequest_ListingFilter) Reset() {
	*x = MemberListingRequest_ListingFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListingRequest_ListingFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListingRequest_ListingFilter) ProtoMessage() {}

func (x *MemberListingRequest_ListingFilter) ProtoReflect() protoreflect.Message {
	mi := &file_members_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListingRequest_ListingFilter.ProtoReflect.Descriptor instead.
func (*MemberListingRequest_ListingFilter) Descriptor() ([]byte, []int) {
	return file_members_proto_rawDescGZIP(), []int{10, 0}
}

func (x *MemberListingRequest_ListingFilter) GetFullNamePart() string {
	if x != nil {
		return x.FullNamePart
	}
	return ""
}

var File_members_proto protoreflect.FileDescriptor

var file_members_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73,
	0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18,
	0x64, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0xb2, 0x01,
	0x0a, 0x38, 0x01, 0x4a, 0x06, 0x08, 0x80, 0xbc, 0xe0, 0xdf, 0x0b, 0x52, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x4b, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x71, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x54, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72,
	0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x14,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75,
	0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55, 0x44,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x30, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63,
	0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63,
	0x72, 0x75, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x99, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75,
	0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x33, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x15,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x52, 0x55,
	0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0x91, 0x05, 0x0a, 0x07,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x79, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01,
	0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x63,
	0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x1a, 0x22, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x82, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x63, 0x6c,
	0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x79, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42,
	0x16, 0x5a, 0x14, 0x67, 0x6f, 0x63, 0x6c, 0x75, 0x62, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_members_proto_rawDescOnce sync.Once
	file_members_proto_rawDescData = file_members_proto_rawDesc
)

func file_members_proto_rawDescGZIP() []byte {
	file_members_proto_rawDescOnce.Do(func() {
		file_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_members_proto_rawDescData)
	})
	return file_members_proto_rawDescData
}

var file_members_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_members_proto_goTypes = []any{
	(*CRUDResult)(nil),                         // 0: goclub.engine.v1.CRUDResult
	(*MemberInfo)(nil),                         // 1: goclub.engine.v1.MemberInfo
	(*MemberCreateRequest)(nil),                // 2: goclub.engine.v1.MemberCreateRequest
	(*MemberCreateResponse)(nil),               // 3: goclub.engine.v1.MemberCreateResponse
	(*MemberUpdateRequest)(nil),                // 4: goclub.engine.v1.MemberUpdateRequest
	(*MemberUpdateResponse)(nil),               // 5: goclub.engine.v1.MemberUpdateResponse
	(*MemberDeleteRequest)(nil),                // 6: goclub.engine.v1.MemberDeleteRequest
	(*MemberDeleteResponse)(nil),               // 7: goclub.engine.v1.MemberDeleteResponse
	(*MemberReadRequest)(nil),                  // 8: goclub.engine.v1.MemberReadRequest
	(*MemberReadResponse)(nil),                 // 9: goclub.engine.v1.MemberReadResponse
	(*MemberListingRequest)(nil),               // 10: goclub.engine.v1.MemberListingRequest
	(*MemberListingResponse)(nil),              // 11: goclub.engine.v1.MemberListingResponse
	(*MemberListingRequest_ListingFilter)(nil), // 12: goclub.engine.v1.MemberListingRequest.ListingFilter
	(*timestamppb.Timestamp)(nil),              // 13: google.protobuf.Timestamp
}
var file_members_proto_depIdxs = []int32{
	13, // 0: goclub.engine.v1.MemberInfo.birthday:type_name -> google.protobuf.Timestamp
	1,  // 1: goclub.engine.v1.MemberCreateRequest.object:type_name -> goclub.engine.v1.MemberInfo
	0,  // 2: goclub.engine.v1.MemberCreateResponse.crudResult:type_name -> goclub.engine.v1.CRUDResult
	1,  // 3: goclub.engine.v1.MemberUpdateRequest.object:type_name -> goclub.engine.v1.MemberInfo
	0,  // 4: goclub.engine.v1.MemberUpdateResponse.crudResult:type_name -> goclub.engine.v1.CRUDResult
	0,  // 5: goclub.engine.v1.MemberDeleteResponse.crudResult:type_name -> goclub.engine.v1.CRUDResult
	0,  // 6: goclub.engine.v1.MemberReadResponse.crudResult:type_name -> goclub.engine.v1.CRUDResult
	1,  // 7: goclub.engine.v1.MemberReadResponse.object:type_name -> goclub.engine.v1.MemberInfo
	12, // 8: goclub.engine.v1.MemberListingRequest.filter:type_name -> goclub.engine.v1.MemberListingRequest.ListingFilter
	0,  // 9: goclub.engine.v1.MemberListingResponse.crudResult:type_name -> goclub.engine.v1.CRUDResult
	1,  // 10: goclub.engine.v1.MemberListingResponse.objects:type_name -> goclub.engine.v1.MemberInfo
	2,  // 11: goclub.engine.v1.Members.MemberCreate:input_type -> goclub.engine.v1.MemberCreateRequest
	4,  // 12: goclub.engine.v1.Members.MemberUpdate:input_type -> goclub.engine.v1.MemberUpdateRequest
	6,  // 13: goclub.engine.v1.Members.MemberDelete:input_type -> goclub.engine.v1.MemberDeleteRequest
	8,  // 14: goclub.engine.v1.Members.MemberRead:input_type -> goclub.engine.v1.MemberReadRequest
	10, // 15: goclub.engine.v1.Members.MemberListing:input_type -> goclub.engine.v1.MemberListingRequest
	3,  // 16: goclub.engine.v1.Members.MemberCreate:output_type -> goclub.engine.v1.MemberCreateResponse
	5,  // 17: goclub.engine.v1.Members.MemberUpdate:output_type -> goclub.engine.v1.MemberUpdateResponse
	7,  // 18: goclub.engine.v1.Members.MemberDelete:output_type -> goclub.engine.v1.MemberDeleteResponse
	9,  // 19: goclub.engine.v1.Members.MemberRead:output_type -> goclub.engine.v1.MemberReadResponse
	11, // 20: goclub.engine.v1.Members.MemberListing:output_type -> goclub.engine.v1.MemberListingResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_members_proto_init() }
func file_members_proto_init() {
	if File_members_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_members_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CRUDResult); i {
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
		file_members_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MemberInfo); i {
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
		file_members_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MemberCreateRequest); i {
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
		file_members_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MemberCreateResponse); i {
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
		file_members_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MemberUpdateRequest); i {
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
		file_members_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MemberUpdateResponse); i {
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
		file_members_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MemberDeleteRequest); i {
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
		file_members_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MemberDeleteResponse); i {
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
		file_members_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*MemberReadRequest); i {
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
		file_members_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*MemberReadResponse); i {
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
		file_members_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*MemberListingRequest); i {
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
		file_members_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*MemberListingResponse); i {
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
		file_members_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*MemberListingRequest_ListingFilter); i {
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
			RawDescriptor: file_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_members_proto_goTypes,
		DependencyIndexes: file_members_proto_depIdxs,
		MessageInfos:      file_members_proto_msgTypes,
	}.Build()
	File_members_proto = out.File
	file_members_proto_rawDesc = nil
	file_members_proto_goTypes = nil
	file_members_proto_depIdxs = nil
}
