// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user_center_channel_service.proto

package v1

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

const (
	Service_AllChannel_FullMethodName    = "/channel.Service/AllChannel"
	Service_GetTypes_FullMethodName      = "/channel.Service/GetTypes"
	Service_AddChannel_FullMethodName    = "/channel.Service/AddChannel"
	Service_UpdateChannel_FullMethodName = "/channel.Service/UpdateChannel"
	Service_DeleteChannel_FullMethodName = "/channel.Service/DeleteChannel"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// AllChannel 获取全部渠道信息
	AllChannel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllChannelReply, error)
	// AllChannel 获取全部渠道信息
	GetTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTypesReply, error)
	// AddChannel 添加渠道信息
	AddChannel(ctx context.Context, in *AddChannelRequest, opts ...grpc.CallOption) (*AddChannelReply, error)
	// UpdateChannel 更新渠道信息
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteChannel 删除渠道信息
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AllChannel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AllChannelReply, error) {
	out := new(AllChannelReply)
	err := c.cc.Invoke(ctx, Service_AllChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTypesReply, error) {
	out := new(GetTypesReply)
	err := c.cc.Invoke(ctx, Service_GetTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddChannel(ctx context.Context, in *AddChannelRequest, opts ...grpc.CallOption) (*AddChannelReply, error) {
	out := new(AddChannelReply)
	err := c.cc.Invoke(ctx, Service_AddChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_UpdateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Service_DeleteChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// AllChannel 获取全部渠道信息
	AllChannel(context.Context, *emptypb.Empty) (*AllChannelReply, error)
	// AllChannel 获取全部渠道信息
	GetTypes(context.Context, *emptypb.Empty) (*GetTypesReply, error)
	// AddChannel 添加渠道信息
	AddChannel(context.Context, *AddChannelRequest) (*AddChannelReply, error)
	// UpdateChannel 更新渠道信息
	UpdateChannel(context.Context, *UpdateChannelRequest) (*emptypb.Empty, error)
	// DeleteChannel 删除渠道信息
	DeleteChannel(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) AllChannel(context.Context, *emptypb.Empty) (*AllChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllChannel not implemented")
}
func (UnimplementedServiceServer) GetTypes(context.Context, *emptypb.Empty) (*GetTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypes not implemented")
}
func (UnimplementedServiceServer) AddChannel(context.Context, *AddChannelRequest) (*AddChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChannel not implemented")
}
func (UnimplementedServiceServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedServiceServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AllChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AllChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AllChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AllChannel(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTypes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AddChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddChannel(ctx, req.(*AddChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllChannel",
			Handler:    _Service_AllChannel_Handler,
		},
		{
			MethodName: "GetTypes",
			Handler:    _Service_GetTypes_Handler,
		},
		{
			MethodName: "AddChannel",
			Handler:    _Service_AddChannel_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _Service_UpdateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _Service_DeleteChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_center_channel_service.proto",
}
