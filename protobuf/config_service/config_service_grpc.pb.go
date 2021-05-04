// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package config_service

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

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	HasConfigs(ctx context.Context, in *HasConfigsRequest, opts ...grpc.CallOption) (*HasConfigsResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/Config/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/Config/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) HasConfigs(ctx context.Context, in *HasConfigsRequest, opts ...grpc.CallOption) (*HasConfigsResponse, error) {
	out := new(HasConfigsResponse)
	err := c.cc.Invoke(ctx, "/Config/HasConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/Config/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
// All implementations must embed UnimplementedConfigServer
// for forward compatibility
type ConfigServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	HasConfigs(context.Context, *HasConfigsRequest) (*HasConfigsResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedConfigServer()
}

// UnimplementedConfigServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServer struct {
}

func (UnimplementedConfigServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedConfigServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfigServer) HasConfigs(context.Context, *HasConfigsRequest) (*HasConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasConfigs not implemented")
}
func (UnimplementedConfigServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedConfigServer) mustEmbedUnimplementedConfigServer() {}

// UnsafeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServer will
// result in compilation errors.
type UnsafeConfigServer interface {
	mustEmbedUnimplementedConfigServer()
}

func RegisterConfigServer(s grpc.ServiceRegistrar, srv ConfigServer) {
	s.RegisterService(&Config_ServiceDesc, srv)
}

func _Config_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Config/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Config/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_HasConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).HasConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Config/HasConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).HasConfigs(ctx, req.(*HasConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Config/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Config_ServiceDesc is the grpc.ServiceDesc for Config service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Config_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Config_Init_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Config_Get_Handler,
		},
		{
			MethodName: "HasConfigs",
			Handler:    _Config_HasConfigs_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Config_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config_service.proto",
}