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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	// Read streams update to the registry on the target instance.
	Read(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Registry_ReadClient, error)
	// Write updates the registry on the target instance.
	Write(ctx context.Context, opts ...grpc.CallOption) (Registry_WriteClient, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Registry_SubscribeClient, error)
	Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error)
	Member(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MemberResponse, error)
	Members(ctx context.Context, in *MembersRequest, opts ...grpc.CallOption) (*MembersResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Read(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Registry_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[0], "/registry.Registry/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_ReadClient interface {
	Recv() (*RemoteMemberUpdate, error)
	grpc.ClientStream
}

type registryReadClient struct {
	grpc.ClientStream
}

func (x *registryReadClient) Recv() (*RemoteMemberUpdate, error) {
	m := new(RemoteMemberUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Write(ctx context.Context, opts ...grpc.CallOption) (Registry_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[1], "/registry.Registry/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryWriteClient{stream}
	return x, nil
}

type Registry_WriteClient interface {
	Send(*ClientUpdate) error
	Recv() (*ClientAck, error)
	grpc.ClientStream
}

type registryWriteClient struct {
	grpc.ClientStream
}

func (x *registryWriteClient) Send(m *ClientUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registryWriteClient) Recv() (*ClientAck, error) {
	m := new(ClientAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Registry_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[2], "/registry.Registry/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &registrySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Registry_SubscribeClient interface {
	Recv() (*RemoteMemberUpdate, error)
	grpc.ClientStream
}

type registrySubscribeClient struct {
	grpc.ClientStream
}

func (x *registrySubscribeClient) Recv() (*RemoteMemberUpdate, error) {
	m := new(RemoteMemberUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[3], "/registry.Registry/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryRegisterClient{stream}
	return x, nil
}

type Registry_RegisterClient interface {
	Send(*ClientUpdate) error
	Recv() (*ClientAck, error)
	grpc.ClientStream
}

type registryRegisterClient struct {
	grpc.ClientStream
}

func (x *registryRegisterClient) Send(m *ClientUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registryRegisterClient) Recv() (*ClientAck, error) {
	m := new(ClientAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Member(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*MemberResponse, error) {
	out := new(MemberResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Member", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Members(ctx context.Context, in *MembersRequest, opts ...grpc.CallOption) (*MembersResponse, error) {
	out := new(MembersResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Members", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	// Read streams update to the registry on the target instance.
	Read(*SubscribeRequest, Registry_ReadServer) error
	// Write updates the registry on the target instance.
	Write(Registry_WriteServer) error
	Subscribe(*SubscribeRequest, Registry_SubscribeServer) error
	Register(Registry_RegisterServer) error
	Member(context.Context, *MemberRequest) (*MemberResponse, error)
	Members(context.Context, *MembersRequest) (*MembersResponse, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) Read(*SubscribeRequest, Registry_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRegistryServer) Write(Registry_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedRegistryServer) Subscribe(*SubscribeRequest, Registry_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedRegistryServer) Register(Registry_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistryServer) Member(context.Context, *MemberRequest) (*MemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Member not implemented")
}
func (UnimplementedRegistryServer) Members(context.Context, *MembersRequest) (*MembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Members not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).Read(m, &registryReadServer{stream})
}

type Registry_ReadServer interface {
	Send(*RemoteMemberUpdate) error
	grpc.ServerStream
}

type registryReadServer struct {
	grpc.ServerStream
}

func (x *registryReadServer) Send(m *RemoteMemberUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _Registry_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Write(&registryWriteServer{stream})
}

type Registry_WriteServer interface {
	Send(*ClientAck) error
	Recv() (*ClientUpdate, error)
	grpc.ServerStream
}

type registryWriteServer struct {
	grpc.ServerStream
}

func (x *registryWriteServer) Send(m *ClientAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registryWriteServer) Recv() (*ClientUpdate, error) {
	m := new(ClientUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Registry_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServer).Subscribe(m, &registrySubscribeServer{stream})
}

type Registry_SubscribeServer interface {
	Send(*RemoteMemberUpdate) error
	grpc.ServerStream
}

type registrySubscribeServer struct {
	grpc.ServerStream
}

func (x *registrySubscribeServer) Send(m *RemoteMemberUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _Registry_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Register(&registryRegisterServer{stream})
}

type Registry_RegisterServer interface {
	Send(*ClientAck) error
	Recv() (*ClientUpdate, error)
	grpc.ServerStream
}

type registryRegisterServer struct {
	grpc.ServerStream
}

func (x *registryRegisterServer) Send(m *ClientAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registryRegisterServer) Recv() (*ClientUpdate, error) {
	m := new(ClientUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Registry_Member_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Member(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Member",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Member(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Members(ctx, req.(*MembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Member",
			Handler:    _Registry_Member_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Registry_Members_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _Registry_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Write",
			Handler:       _Registry_Write_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Registry_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Register",
			Handler:       _Registry_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "registry.proto",
}
