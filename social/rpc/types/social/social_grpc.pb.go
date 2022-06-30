// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: schema/social.proto

package social

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
	FanList(ctx context.Context, in *FanListRequest, opts ...grpc.CallOption) (*FanListResponse, error)
	FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error)
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	out := new(FollowListResponse)
	err := c.cc.Invoke(ctx, "/social.Service/FollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FanList(ctx context.Context, in *FanListRequest, opts ...grpc.CallOption) (*FanListResponse, error) {
	out := new(FanListResponse)
	err := c.cc.Invoke(ctx, "/social.Service/FanList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error) {
	out := new(FollowActionResponse)
	err := c.cc.Invoke(ctx, "/social.Service/FollowAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error) {
	out := new(FavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/social.Service/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	out := new(FavoriteListResponse)
	err := c.cc.Invoke(ctx, "/social.Service/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, "/social.Service/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, "/social.Service/CommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error)
	FanList(context.Context, *FanListRequest) (*FanListResponse, error)
	FollowAction(context.Context, *FollowActionRequest) (*FollowActionResponse, error)
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error)
	FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error)
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowList not implemented")
}
func (UnimplementedServiceServer) FanList(context.Context, *FanListRequest) (*FanListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanList not implemented")
}
func (UnimplementedServiceServer) FollowAction(context.Context, *FollowActionRequest) (*FollowActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowAction not implemented")
}
func (UnimplementedServiceServer) FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedServiceServer) FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedServiceServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedServiceServer) CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
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

func _Service_FollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/FollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FollowList(ctx, req.(*FollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FanList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FanListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FanList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/FanList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FanList(ctx, req.(*FanListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FollowAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FollowAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/FollowAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FollowAction(ctx, req.(*FollowActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FavoriteAction(ctx, req.(*FavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FavoriteList(ctx, req.(*FavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/social.Service/CommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "social.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowList",
			Handler:    _Service_FollowList_Handler,
		},
		{
			MethodName: "FanList",
			Handler:    _Service_FanList_Handler,
		},
		{
			MethodName: "FollowAction",
			Handler:    _Service_FollowAction_Handler,
		},
		{
			MethodName: "FavoriteAction",
			Handler:    _Service_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _Service_FavoriteList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _Service_CommentAction_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _Service_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/social.proto",
}