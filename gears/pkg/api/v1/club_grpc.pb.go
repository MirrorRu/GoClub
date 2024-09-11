// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: club.proto

package gears

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GoClubAPI_Ping_FullMethodName          = "/goclub.gears.v1.GoClubAPI/Ping"
	GoClubAPI_MemberCreate_FullMethodName  = "/goclub.gears.v1.GoClubAPI/MemberCreate"
	GoClubAPI_MemberUpdate_FullMethodName  = "/goclub.gears.v1.GoClubAPI/MemberUpdate"
	GoClubAPI_MemberDelete_FullMethodName  = "/goclub.gears.v1.GoClubAPI/MemberDelete"
	GoClubAPI_MemberRead_FullMethodName    = "/goclub.gears.v1.GoClubAPI/MemberRead"
	GoClubAPI_MemberListing_FullMethodName = "/goclub.gears.v1.GoClubAPI/MemberListing"
)

// GoClubAPIClient is the client API for GoClubAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
type GoClubAPIClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	MemberCreate(ctx context.Context, in *MemberCreateRequest, opts ...grpc.CallOption) (*MemberCreateResponse, error)
	MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error)
	MemberDelete(ctx context.Context, in *MemberDeleteRequest, opts ...grpc.CallOption) (*MemberDeleteResponse, error)
	MemberRead(ctx context.Context, in *MemberReadRequest, opts ...grpc.CallOption) (*MemberReadResponse, error)
	MemberListing(ctx context.Context, in *MemberListingRequest, opts ...grpc.CallOption) (*MemberListingResponse, error)
}

type goClubAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewGoClubAPIClient(cc grpc.ClientConnInterface) GoClubAPIClient {
	return &goClubAPIClient{cc}
}

func (c *goClubAPIClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goClubAPIClient) MemberCreate(ctx context.Context, in *MemberCreateRequest, opts ...grpc.CallOption) (*MemberCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberCreateResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_MemberCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goClubAPIClient) MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberUpdateResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_MemberUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goClubAPIClient) MemberDelete(ctx context.Context, in *MemberDeleteRequest, opts ...grpc.CallOption) (*MemberDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberDeleteResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_MemberDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goClubAPIClient) MemberRead(ctx context.Context, in *MemberReadRequest, opts ...grpc.CallOption) (*MemberReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberReadResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_MemberRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goClubAPIClient) MemberListing(ctx context.Context, in *MemberListingRequest, opts ...grpc.CallOption) (*MemberListingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemberListingResponse)
	err := c.cc.Invoke(ctx, GoClubAPI_MemberListing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoClubAPIServer is the server API for GoClubAPI service.
// All implementations must embed UnimplementedGoClubAPIServer
// for forward compatibility.
//
// https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
type GoClubAPIServer interface {
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	MemberCreate(context.Context, *MemberCreateRequest) (*MemberCreateResponse, error)
	MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error)
	MemberDelete(context.Context, *MemberDeleteRequest) (*MemberDeleteResponse, error)
	MemberRead(context.Context, *MemberReadRequest) (*MemberReadResponse, error)
	MemberListing(context.Context, *MemberListingRequest) (*MemberListingResponse, error)
	mustEmbedUnimplementedGoClubAPIServer()
}

// UnimplementedGoClubAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoClubAPIServer struct{}

func (UnimplementedGoClubAPIServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGoClubAPIServer) MemberCreate(context.Context, *MemberCreateRequest) (*MemberCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberCreate not implemented")
}
func (UnimplementedGoClubAPIServer) MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberUpdate not implemented")
}
func (UnimplementedGoClubAPIServer) MemberDelete(context.Context, *MemberDeleteRequest) (*MemberDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberDelete not implemented")
}
func (UnimplementedGoClubAPIServer) MemberRead(context.Context, *MemberReadRequest) (*MemberReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRead not implemented")
}
func (UnimplementedGoClubAPIServer) MemberListing(context.Context, *MemberListingRequest) (*MemberListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberListing not implemented")
}
func (UnimplementedGoClubAPIServer) mustEmbedUnimplementedGoClubAPIServer() {}
func (UnimplementedGoClubAPIServer) testEmbeddedByValue()                   {}

// UnsafeGoClubAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoClubAPIServer will
// result in compilation errors.
type UnsafeGoClubAPIServer interface {
	mustEmbedUnimplementedGoClubAPIServer()
}

func RegisterGoClubAPIServer(s grpc.ServiceRegistrar, srv GoClubAPIServer) {
	// If the following call pancis, it indicates UnimplementedGoClubAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GoClubAPI_ServiceDesc, srv)
}

func _GoClubAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoClubAPI_MemberCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).MemberCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_MemberCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).MemberCreate(ctx, req.(*MemberCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoClubAPI_MemberUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).MemberUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_MemberUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).MemberUpdate(ctx, req.(*MemberUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoClubAPI_MemberDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).MemberDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_MemberDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).MemberDelete(ctx, req.(*MemberDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoClubAPI_MemberRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).MemberRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_MemberRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).MemberRead(ctx, req.(*MemberReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoClubAPI_MemberListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoClubAPIServer).MemberListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoClubAPI_MemberListing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoClubAPIServer).MemberListing(ctx, req.(*MemberListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoClubAPI_ServiceDesc is the grpc.ServiceDesc for GoClubAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoClubAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goclub.gears.v1.GoClubAPI",
	HandlerType: (*GoClubAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GoClubAPI_Ping_Handler,
		},
		{
			MethodName: "MemberCreate",
			Handler:    _GoClubAPI_MemberCreate_Handler,
		},
		{
			MethodName: "MemberUpdate",
			Handler:    _GoClubAPI_MemberUpdate_Handler,
		},
		{
			MethodName: "MemberDelete",
			Handler:    _GoClubAPI_MemberDelete_Handler,
		},
		{
			MethodName: "MemberRead",
			Handler:    _GoClubAPI_MemberRead_Handler,
		},
		{
			MethodName: "MemberListing",
			Handler:    _GoClubAPI_MemberListing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "club.proto",
}
