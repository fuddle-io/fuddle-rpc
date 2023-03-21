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
	Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesRespose, error)
	Node(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, opts ...grpc.CallOption) (Registry_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Registry_ServiceDesc.Streams[0], "/registry.Registry/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryRegisterClient{stream}
	return x, nil
}

type Registry_RegisterClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type registryRegisterClient struct {
	grpc.ClientStream
}

func (x *registryRegisterClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registryRegisterClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *registryClient) Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesRespose, error) {
	out := new(NodesRespose)
	err := c.cc.Invoke(ctx, "/registry.Registry/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Node(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error) {
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, "/registry.Registry/Node", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	Register(Registry_RegisterServer) error
	Nodes(context.Context, *NodesRequest) (*NodesRespose, error)
	Node(context.Context, *NodeRequest) (*NodeResponse, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) Register(Registry_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistryServer) Nodes(context.Context, *NodesRequest) (*NodesRespose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nodes not implemented")
}
func (UnimplementedRegistryServer) Node(context.Context, *NodeRequest) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node not implemented")
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

func _Registry_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistryServer).Register(&registryRegisterServer{stream})
}

type Registry_RegisterServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type registryRegisterServer struct {
	grpc.ServerStream
}

func (x *registryRegisterServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registryRegisterServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Registry_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Nodes(ctx, req.(*NodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Node_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Node(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/Node",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Node(ctx, req.(*NodeRequest))
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
			MethodName: "Nodes",
			Handler:    _Registry_Nodes_Handler,
		},
		{
			MethodName: "Node",
			Handler:    _Registry_Node_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Registry_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "registry.proto",
}