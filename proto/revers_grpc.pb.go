// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: revers.proto

package proto

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

// ReverseClient is the client API for Reverse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReverseClient interface {
	Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type reverseClient struct {
	cc grpc.ClientConnInterface
}

func NewReverseClient(cc grpc.ClientConnInterface) ReverseClient {
	return &reverseClient{cc}
}

func (c *reverseClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/string.Reverse/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverseServer is the server API for Reverse service.
// All implementations must embed UnimplementedReverseServer
// for forward compatibility
type ReverseServer interface {
	Do(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedReverseServer()
}

// UnimplementedReverseServer must be embedded to have forward compatible implementations.
type UnimplementedReverseServer struct {
}

func (UnimplementedReverseServer) Do(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedReverseServer) mustEmbedUnimplementedReverseServer() {}

// UnsafeReverseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReverseServer will
// result in compilation errors.
type UnsafeReverseServer interface {
	mustEmbedUnimplementedReverseServer()
}

func RegisterReverseServer(s grpc.ServiceRegistrar, srv *interface{}) {
	s.RegisterService(&Reverse_ServiceDesc, srv)
}

func _Reverse_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverseServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/string.Reverse/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverseServer).Do(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Reverse_ServiceDesc is the grpc.ServiceDesc for Reverse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reverse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "string.Reverse",
	HandlerType: (*ReverseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Reverse_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "revers.proto",
}
