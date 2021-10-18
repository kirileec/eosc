// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	DeleteCheck(ctx context.Context, in *WorkerDeleteRequest, opts ...grpc.CallOption) (*WorkerDeleteResponse, error)
	SetCheck(ctx context.Context, in *WorkerSetRequest, opts ...grpc.CallOption) (*WorkerSetResponse, error)
	Delete(ctx context.Context, in *WorkerDeleteRequest, opts ...grpc.CallOption) (*WorkerDeleteResponse, error)
	Set(ctx context.Context, in *WorkerSetRequest, opts ...grpc.CallOption) (*WorkerSetResponse, error)
	Ping(ctx context.Context, in *WorkerHelloRequest, opts ...grpc.CallOption) (*WorkerHelloResponse, error)
	Refresh(ctx context.Context, in *WorkerRefreshRequest, opts ...grpc.CallOption) (*WorkerRefreshResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) DeleteCheck(ctx context.Context, in *WorkerDeleteRequest, opts ...grpc.CallOption) (*WorkerDeleteResponse, error) {
	out := new(WorkerDeleteResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/deleteCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) SetCheck(ctx context.Context, in *WorkerSetRequest, opts ...grpc.CallOption) (*WorkerSetResponse, error) {
	out := new(WorkerSetResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/setCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Delete(ctx context.Context, in *WorkerDeleteRequest, opts ...grpc.CallOption) (*WorkerDeleteResponse, error) {
	out := new(WorkerDeleteResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Set(ctx context.Context, in *WorkerSetRequest, opts ...grpc.CallOption) (*WorkerSetResponse, error) {
	out := new(WorkerSetResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Ping(ctx context.Context, in *WorkerHelloRequest, opts ...grpc.CallOption) (*WorkerHelloResponse, error) {
	out := new(WorkerHelloResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) Refresh(ctx context.Context, in *WorkerRefreshRequest, opts ...grpc.CallOption) (*WorkerRefreshResponse, error) {
	out := new(WorkerRefreshResponse)
	err := c.cc.Invoke(ctx, "/service.WorkerService/refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations must embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	DeleteCheck(context.Context, *WorkerDeleteRequest) (*WorkerDeleteResponse, error)
	SetCheck(context.Context, *WorkerSetRequest) (*WorkerSetResponse, error)
	Delete(context.Context, *WorkerDeleteRequest) (*WorkerDeleteResponse, error)
	Set(context.Context, *WorkerSetRequest) (*WorkerSetResponse, error)
	Ping(context.Context, *WorkerHelloRequest) (*WorkerHelloResponse, error)
	Refresh(context.Context, *WorkerRefreshRequest) (*WorkerRefreshResponse, error)
	mustEmbedUnimplementedWorkerServiceServer()
}

// UnimplementedWorkerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) DeleteCheck(context.Context, *WorkerDeleteRequest) (*WorkerDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCheck not implemented")
}
func (UnimplementedWorkerServiceServer) SetCheck(context.Context, *WorkerSetRequest) (*WorkerSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCheck not implemented")
}
func (UnimplementedWorkerServiceServer) Delete(context.Context, *WorkerDeleteRequest) (*WorkerDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWorkerServiceServer) Set(context.Context, *WorkerSetRequest) (*WorkerSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedWorkerServiceServer) Ping(context.Context, *WorkerHelloRequest) (*WorkerHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWorkerServiceServer) Refresh(context.Context, *WorkerRefreshRequest) (*WorkerRefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedWorkerServiceServer) mustEmbedUnimplementedWorkerServiceServer() {}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_DeleteCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).DeleteCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/deleteCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).DeleteCheck(ctx, req.(*WorkerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_SetCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).SetCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/setCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).SetCheck(ctx, req.(*WorkerSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Delete(ctx, req.(*WorkerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Set(ctx, req.(*WorkerSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Ping(ctx, req.(*WorkerHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.WorkerService/refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).Refresh(ctx, req.(*WorkerRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "deleteCheck",
			Handler:    _WorkerService_DeleteCheck_Handler,
		},
		{
			MethodName: "setCheck",
			Handler:    _WorkerService_SetCheck_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _WorkerService_Delete_Handler,
		},
		{
			MethodName: "set",
			Handler:    _WorkerService_Set_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _WorkerService_Ping_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _WorkerService_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
