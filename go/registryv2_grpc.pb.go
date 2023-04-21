// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReplicaRegistry2Client is the client API for ReplicaRegistry2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicaRegistry2Client interface {
	// Update is used for forward registry updates to replica nodes.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Sync sends the known versions for the members in the registry (or a
	// subset if too large) and receives any missed updates for those members
	// in response.
	Sync(ctx context.Context, in *ReplicaSyncRequest, opts ...grpc.CallOption) (*ReplicaSyncResponse, error)
}

type replicaRegistry2Client struct {
	cc grpc.ClientConnInterface
}

func NewReplicaRegistry2Client(cc grpc.ClientConnInterface) ReplicaRegistry2Client {
	return &replicaRegistry2Client{cc}
}

func (c *replicaRegistry2Client) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/registryv2.ReplicaRegistry2/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicaRegistry2Client) Sync(ctx context.Context, in *ReplicaSyncRequest, opts ...grpc.CallOption) (*ReplicaSyncResponse, error) {
	out := new(ReplicaSyncResponse)
	err := c.cc.Invoke(ctx, "/registryv2.ReplicaRegistry2/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicaRegistry2Server is the server API for ReplicaRegistry2 service.
// All implementations must embed UnimplementedReplicaRegistry2Server
// for forward compatibility
type ReplicaRegistry2Server interface {
	// Update is used for forward registry updates to replica nodes.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Sync sends the known versions for the members in the registry (or a
	// subset if too large) and receives any missed updates for those members
	// in response.
	Sync(context.Context, *ReplicaSyncRequest) (*ReplicaSyncResponse, error)
	mustEmbedUnimplementedReplicaRegistry2Server()
}

// UnimplementedReplicaRegistry2Server must be embedded to have forward compatible implementations.
type UnimplementedReplicaRegistry2Server struct {
}

func (UnimplementedReplicaRegistry2Server) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReplicaRegistry2Server) Sync(context.Context, *ReplicaSyncRequest) (*ReplicaSyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedReplicaRegistry2Server) mustEmbedUnimplementedReplicaRegistry2Server() {}

// UnsafeReplicaRegistry2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicaRegistry2Server will
// result in compilation errors.
type UnsafeReplicaRegistry2Server interface {
	mustEmbedUnimplementedReplicaRegistry2Server()
}

func RegisterReplicaRegistry2Server(s grpc.ServiceRegistrar, srv ReplicaRegistry2Server) {
	s.RegisterService(&ReplicaRegistry2_ServiceDesc, srv)
}

func _ReplicaRegistry2_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaRegistry2Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryv2.ReplicaRegistry2/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaRegistry2Server).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicaRegistry2_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicaSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaRegistry2Server).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryv2.ReplicaRegistry2/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaRegistry2Server).Sync(ctx, req.(*ReplicaSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicaRegistry2_ServiceDesc is the grpc.ServiceDesc for ReplicaRegistry2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicaRegistry2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registryv2.ReplicaRegistry2",
	HandlerType: (*ReplicaRegistry2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _ReplicaRegistry2_Update_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _ReplicaRegistry2_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registryv2.proto",
}

// ClientReadRegistry2Client is the client API for ClientReadRegistry2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientReadRegistry2Client interface {
	Sync(ctx context.Context, in *ClientSyncRequest, opts ...grpc.CallOption) (ClientReadRegistry2_SyncClient, error)
}

type clientReadRegistry2Client struct {
	cc grpc.ClientConnInterface
}

func NewClientReadRegistry2Client(cc grpc.ClientConnInterface) ClientReadRegistry2Client {
	return &clientReadRegistry2Client{cc}
}

func (c *clientReadRegistry2Client) Sync(ctx context.Context, in *ClientSyncRequest, opts ...grpc.CallOption) (ClientReadRegistry2_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientReadRegistry2_ServiceDesc.Streams[0], "/registryv2.ClientReadRegistry2/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientReadRegistry2SyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientReadRegistry2_SyncClient interface {
	Recv() (*ClientSyncUpdate, error)
	grpc.ClientStream
}

type clientReadRegistry2SyncClient struct {
	grpc.ClientStream
}

func (x *clientReadRegistry2SyncClient) Recv() (*ClientSyncUpdate, error) {
	m := new(ClientSyncUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientReadRegistry2Server is the server API for ClientReadRegistry2 service.
// All implementations must embed UnimplementedClientReadRegistry2Server
// for forward compatibility
type ClientReadRegistry2Server interface {
	Sync(*ClientSyncRequest, ClientReadRegistry2_SyncServer) error
	mustEmbedUnimplementedClientReadRegistry2Server()
}

// UnimplementedClientReadRegistry2Server must be embedded to have forward compatible implementations.
type UnimplementedClientReadRegistry2Server struct {
}

func (UnimplementedClientReadRegistry2Server) Sync(*ClientSyncRequest, ClientReadRegistry2_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedClientReadRegistry2Server) mustEmbedUnimplementedClientReadRegistry2Server() {}

// UnsafeClientReadRegistry2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientReadRegistry2Server will
// result in compilation errors.
type UnsafeClientReadRegistry2Server interface {
	mustEmbedUnimplementedClientReadRegistry2Server()
}

func RegisterClientReadRegistry2Server(s grpc.ServiceRegistrar, srv ClientReadRegistry2Server) {
	s.RegisterService(&ClientReadRegistry2_ServiceDesc, srv)
}

func _ClientReadRegistry2_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientSyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientReadRegistry2Server).Sync(m, &clientReadRegistry2SyncServer{stream})
}

type ClientReadRegistry2_SyncServer interface {
	Send(*ClientSyncUpdate) error
	grpc.ServerStream
}

type clientReadRegistry2SyncServer struct {
	grpc.ServerStream
}

func (x *clientReadRegistry2SyncServer) Send(m *ClientSyncUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// ClientReadRegistry2_ServiceDesc is the grpc.ServiceDesc for ClientReadRegistry2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientReadRegistry2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registryv2.ClientReadRegistry2",
	HandlerType: (*ClientReadRegistry2Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _ClientReadRegistry2_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "registryv2.proto",
}

// ClientWriteRegistry2Client is the client API for ClientWriteRegistry2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientWriteRegistry2Client interface {
	MemberJoin(ctx context.Context, in *ClientMemberJoinRequest, opts ...grpc.CallOption) (*ClientMemberJoinResponse, error)
	MemberLeave(ctx context.Context, in *ClientMemberLeaveRequest, opts ...grpc.CallOption) (*ClientMemberLeaveResponse, error)
	MemberHeartbeat(ctx context.Context, in *ClientMemberHeartbeatRequest, opts ...grpc.CallOption) (*ClientMemberHeartbeatResponse, error)
}

type clientWriteRegistry2Client struct {
	cc grpc.ClientConnInterface
}

func NewClientWriteRegistry2Client(cc grpc.ClientConnInterface) ClientWriteRegistry2Client {
	return &clientWriteRegistry2Client{cc}
}

func (c *clientWriteRegistry2Client) MemberJoin(ctx context.Context, in *ClientMemberJoinRequest, opts ...grpc.CallOption) (*ClientMemberJoinResponse, error) {
	out := new(ClientMemberJoinResponse)
	err := c.cc.Invoke(ctx, "/registryv2.ClientWriteRegistry2/MemberJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientWriteRegistry2Client) MemberLeave(ctx context.Context, in *ClientMemberLeaveRequest, opts ...grpc.CallOption) (*ClientMemberLeaveResponse, error) {
	out := new(ClientMemberLeaveResponse)
	err := c.cc.Invoke(ctx, "/registryv2.ClientWriteRegistry2/MemberLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientWriteRegistry2Client) MemberHeartbeat(ctx context.Context, in *ClientMemberHeartbeatRequest, opts ...grpc.CallOption) (*ClientMemberHeartbeatResponse, error) {
	out := new(ClientMemberHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/registryv2.ClientWriteRegistry2/MemberHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientWriteRegistry2Server is the server API for ClientWriteRegistry2 service.
// All implementations must embed UnimplementedClientWriteRegistry2Server
// for forward compatibility
type ClientWriteRegistry2Server interface {
	MemberJoin(context.Context, *ClientMemberJoinRequest) (*ClientMemberJoinResponse, error)
	MemberLeave(context.Context, *ClientMemberLeaveRequest) (*ClientMemberLeaveResponse, error)
	MemberHeartbeat(context.Context, *ClientMemberHeartbeatRequest) (*ClientMemberHeartbeatResponse, error)
	mustEmbedUnimplementedClientWriteRegistry2Server()
}

// UnimplementedClientWriteRegistry2Server must be embedded to have forward compatible implementations.
type UnimplementedClientWriteRegistry2Server struct {
}

func (UnimplementedClientWriteRegistry2Server) MemberJoin(context.Context, *ClientMemberJoinRequest) (*ClientMemberJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberJoin not implemented")
}
func (UnimplementedClientWriteRegistry2Server) MemberLeave(context.Context, *ClientMemberLeaveRequest) (*ClientMemberLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberLeave not implemented")
}
func (UnimplementedClientWriteRegistry2Server) MemberHeartbeat(context.Context, *ClientMemberHeartbeatRequest) (*ClientMemberHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberHeartbeat not implemented")
}
func (UnimplementedClientWriteRegistry2Server) mustEmbedUnimplementedClientWriteRegistry2Server() {}

// UnsafeClientWriteRegistry2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientWriteRegistry2Server will
// result in compilation errors.
type UnsafeClientWriteRegistry2Server interface {
	mustEmbedUnimplementedClientWriteRegistry2Server()
}

func RegisterClientWriteRegistry2Server(s grpc.ServiceRegistrar, srv ClientWriteRegistry2Server) {
	s.RegisterService(&ClientWriteRegistry2_ServiceDesc, srv)
}

func _ClientWriteRegistry2_MemberJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMemberJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientWriteRegistry2Server).MemberJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryv2.ClientWriteRegistry2/MemberJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientWriteRegistry2Server).MemberJoin(ctx, req.(*ClientMemberJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientWriteRegistry2_MemberLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMemberLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientWriteRegistry2Server).MemberLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryv2.ClientWriteRegistry2/MemberLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientWriteRegistry2Server).MemberLeave(ctx, req.(*ClientMemberLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientWriteRegistry2_MemberHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMemberHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientWriteRegistry2Server).MemberHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryv2.ClientWriteRegistry2/MemberHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientWriteRegistry2Server).MemberHeartbeat(ctx, req.(*ClientMemberHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientWriteRegistry2_ServiceDesc is the grpc.ServiceDesc for ClientWriteRegistry2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientWriteRegistry2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registryv2.ClientWriteRegistry2",
	HandlerType: (*ClientWriteRegistry2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MemberJoin",
			Handler:    _ClientWriteRegistry2_MemberJoin_Handler,
		},
		{
			MethodName: "MemberLeave",
			Handler:    _ClientWriteRegistry2_MemberLeave_Handler,
		},
		{
			MethodName: "MemberHeartbeat",
			Handler:    _ClientWriteRegistry2_MemberHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registryv2.proto",
}
