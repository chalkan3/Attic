// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/physical_environment.proto

package pb

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

// PhysicalEnvironmentClient is the client API for PhysicalEnvironment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhysicalEnvironmentClient interface {
	Create(ctx context.Context, in *PhysicalEnvironmentRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error)
	Update(ctx context.Context, in *PhysicalEnvironmentRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error)
	Get(ctx context.Context, in *PhysicalEnvironmentGetRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error)
	List(ctx context.Context, in *PhysicalEnvironmentRequestEmpty, opts ...grpc.CallOption) (*PhysicalEnvironmentResponseList, error)
}

type physicalEnvironmentClient struct {
	cc grpc.ClientConnInterface
}

func NewPhysicalEnvironmentClient(cc grpc.ClientConnInterface) PhysicalEnvironmentClient {
	return &physicalEnvironmentClient{cc}
}

func (c *physicalEnvironmentClient) Create(ctx context.Context, in *PhysicalEnvironmentRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error) {
	out := new(PhysicalEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/pb.PhysicalEnvironment/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicalEnvironmentClient) Update(ctx context.Context, in *PhysicalEnvironmentRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error) {
	out := new(PhysicalEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/pb.PhysicalEnvironment/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicalEnvironmentClient) Get(ctx context.Context, in *PhysicalEnvironmentGetRequest, opts ...grpc.CallOption) (*PhysicalEnvironmentResponse, error) {
	out := new(PhysicalEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/pb.PhysicalEnvironment/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicalEnvironmentClient) List(ctx context.Context, in *PhysicalEnvironmentRequestEmpty, opts ...grpc.CallOption) (*PhysicalEnvironmentResponseList, error) {
	out := new(PhysicalEnvironmentResponseList)
	err := c.cc.Invoke(ctx, "/pb.PhysicalEnvironment/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhysicalEnvironmentServer is the server API for PhysicalEnvironment service.
// All implementations must embed UnimplementedPhysicalEnvironmentServer
// for forward compatibility
type PhysicalEnvironmentServer interface {
	Create(context.Context, *PhysicalEnvironmentRequest) (*PhysicalEnvironmentResponse, error)
	Update(context.Context, *PhysicalEnvironmentRequest) (*PhysicalEnvironmentResponse, error)
	Get(context.Context, *PhysicalEnvironmentGetRequest) (*PhysicalEnvironmentResponse, error)
	List(context.Context, *PhysicalEnvironmentRequestEmpty) (*PhysicalEnvironmentResponseList, error)
	mustEmbedUnimplementedPhysicalEnvironmentServer()
}

// UnimplementedPhysicalEnvironmentServer must be embedded to have forward compatible implementations.
type UnimplementedPhysicalEnvironmentServer struct {
}

func (UnimplementedPhysicalEnvironmentServer) Create(context.Context, *PhysicalEnvironmentRequest) (*PhysicalEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPhysicalEnvironmentServer) Update(context.Context, *PhysicalEnvironmentRequest) (*PhysicalEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPhysicalEnvironmentServer) Get(context.Context, *PhysicalEnvironmentGetRequest) (*PhysicalEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPhysicalEnvironmentServer) List(context.Context, *PhysicalEnvironmentRequestEmpty) (*PhysicalEnvironmentResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPhysicalEnvironmentServer) mustEmbedUnimplementedPhysicalEnvironmentServer() {}

// UnsafePhysicalEnvironmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhysicalEnvironmentServer will
// result in compilation errors.
type UnsafePhysicalEnvironmentServer interface {
	mustEmbedUnimplementedPhysicalEnvironmentServer()
}

func RegisterPhysicalEnvironmentServer(s grpc.ServiceRegistrar, srv PhysicalEnvironmentServer) {
	s.RegisterService(&PhysicalEnvironment_ServiceDesc, srv)
}

func _PhysicalEnvironment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhysicalEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicalEnvironmentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PhysicalEnvironment/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicalEnvironmentServer).Create(ctx, req.(*PhysicalEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicalEnvironment_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhysicalEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicalEnvironmentServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PhysicalEnvironment/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicalEnvironmentServer).Update(ctx, req.(*PhysicalEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicalEnvironment_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhysicalEnvironmentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicalEnvironmentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PhysicalEnvironment/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicalEnvironmentServer).Get(ctx, req.(*PhysicalEnvironmentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicalEnvironment_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhysicalEnvironmentRequestEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicalEnvironmentServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PhysicalEnvironment/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicalEnvironmentServer).List(ctx, req.(*PhysicalEnvironmentRequestEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

// PhysicalEnvironment_ServiceDesc is the grpc.ServiceDesc for PhysicalEnvironment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhysicalEnvironment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PhysicalEnvironment",
	HandlerType: (*PhysicalEnvironmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PhysicalEnvironment_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PhysicalEnvironment_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PhysicalEnvironment_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PhysicalEnvironment_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/physical_environment.proto",
}
