// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: router_sidecar.proto

package sidecar

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SliceRouterSidecarServiceClient is the client API for SliceRouterSidecarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SliceRouterSidecarServiceClient interface {
	// Used to add remote cluster subnet routes in the slice router
	UpdateSliceGwConnectionContext(ctx context.Context, in *SliceGwConContext, opts ...grpc.CallOption) (*SidecarResponse, error)
	// Provides connection information of all clients connected to the slice router
	GetSliceRouterClientConnectionInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClientConnectionInfo, error)
}

type sliceRouterSidecarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSliceRouterSidecarServiceClient(cc grpc.ClientConnInterface) SliceRouterSidecarServiceClient {
	return &sliceRouterSidecarServiceClient{cc}
}

func (c *sliceRouterSidecarServiceClient) UpdateSliceGwConnectionContext(ctx context.Context, in *SliceGwConContext, opts ...grpc.CallOption) (*SidecarResponse, error) {
	out := new(SidecarResponse)
	err := c.cc.Invoke(ctx, "/router.SliceRouterSidecarService/UpdateSliceGwConnectionContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliceRouterSidecarServiceClient) GetSliceRouterClientConnectionInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClientConnectionInfo, error) {
	out := new(ClientConnectionInfo)
	err := c.cc.Invoke(ctx, "/router.SliceRouterSidecarService/GetSliceRouterClientConnectionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliceRouterSidecarServiceServer is the server API for SliceRouterSidecarService service.
// All implementations must embed UnimplementedSliceRouterSidecarServiceServer
// for forward compatibility
type SliceRouterSidecarServiceServer interface {
	// Used to add remote cluster subnet routes in the slice router
	UpdateSliceGwConnectionContext(context.Context, *SliceGwConContext) (*SidecarResponse, error)
	// Provides connection information of all clients connected to the slice router
	GetSliceRouterClientConnectionInfo(context.Context, *emptypb.Empty) (*ClientConnectionInfo, error)
	mustEmbedUnimplementedSliceRouterSidecarServiceServer()
}

// UnimplementedSliceRouterSidecarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSliceRouterSidecarServiceServer struct {
}

func (UnimplementedSliceRouterSidecarServiceServer) UpdateSliceGwConnectionContext(context.Context, *SliceGwConContext) (*SidecarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSliceGwConnectionContext not implemented")
}
func (UnimplementedSliceRouterSidecarServiceServer) GetSliceRouterClientConnectionInfo(context.Context, *emptypb.Empty) (*ClientConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliceRouterClientConnectionInfo not implemented")
}
func (UnimplementedSliceRouterSidecarServiceServer) mustEmbedUnimplementedSliceRouterSidecarServiceServer() {
}

// UnsafeSliceRouterSidecarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SliceRouterSidecarServiceServer will
// result in compilation errors.
type UnsafeSliceRouterSidecarServiceServer interface {
	mustEmbedUnimplementedSliceRouterSidecarServiceServer()
}

func RegisterSliceRouterSidecarServiceServer(s grpc.ServiceRegistrar, srv SliceRouterSidecarServiceServer) {
	s.RegisterService(&SliceRouterSidecarService_ServiceDesc, srv)
}

func _SliceRouterSidecarService_UpdateSliceGwConnectionContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliceGwConContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliceRouterSidecarServiceServer).UpdateSliceGwConnectionContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.SliceRouterSidecarService/UpdateSliceGwConnectionContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliceRouterSidecarServiceServer).UpdateSliceGwConnectionContext(ctx, req.(*SliceGwConContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliceRouterSidecarService_GetSliceRouterClientConnectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliceRouterSidecarServiceServer).GetSliceRouterClientConnectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.SliceRouterSidecarService/GetSliceRouterClientConnectionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliceRouterSidecarServiceServer).GetSliceRouterClientConnectionInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SliceRouterSidecarService_ServiceDesc is the grpc.ServiceDesc for SliceRouterSidecarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SliceRouterSidecarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "router.SliceRouterSidecarService",
	HandlerType: (*SliceRouterSidecarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSliceGwConnectionContext",
			Handler:    _SliceRouterSidecarService_UpdateSliceGwConnectionContext_Handler,
		},
		{
			MethodName: "GetSliceRouterClientConnectionInfo",
			Handler:    _SliceRouterSidecarService_GetSliceRouterClientConnectionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router_sidecar.proto",
}