// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OpenapiClient is the client API for Openapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenapiClient interface {
	// Query identify.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Identify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.IdentifyResponse, error)
	// Post addons identify.
	// TKEEL_COMMENT
	// {
	//  "response" :
	//    {
	//      "raw_data": true
	//    }
	// }
	AddonsIdentify(ctx context.Context, in *v1.AddonsIdentifyRequest, opts ...grpc.CallOption) (*v1.AddonsIdentifyResponse, error)
	// Post tenant enable.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantEnable(ctx context.Context, in *v1.TenantEnableRequest, opts ...grpc.CallOption) (*v1.TenantEnableResponse, error)
	// Post tenant disable.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantDisable(ctx context.Context, in *v1.TenantDisableRequest, opts ...grpc.CallOption) (*v1.TenantDisableResponse, error)
	// Query status.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.StatusResponse, error)
}

type openapiClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenapiClient(cc grpc.ClientConnInterface) OpenapiClient {
	return &openapiClient{cc}
}

func (c *openapiClient) Identify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.IdentifyResponse, error) {
	out := new(v1.IdentifyResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/Identify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) AddonsIdentify(ctx context.Context, in *v1.AddonsIdentifyRequest, opts ...grpc.CallOption) (*v1.AddonsIdentifyResponse, error) {
	out := new(v1.AddonsIdentifyResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/AddonsIdentify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) TenantEnable(ctx context.Context, in *v1.TenantEnableRequest, opts ...grpc.CallOption) (*v1.TenantEnableResponse, error) {
	out := new(v1.TenantEnableResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/TenantEnable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) TenantDisable(ctx context.Context, in *v1.TenantDisableRequest, opts ...grpc.CallOption) (*v1.TenantDisableResponse, error) {
	out := new(v1.TenantDisableResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/TenantDisable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openapiClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.StatusResponse, error) {
	out := new(v1.StatusResponse)
	err := c.cc.Invoke(ctx, "/openapi.v1.Openapi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpenapiServer is the server API for Openapi service.
// All implementations must embed UnimplementedOpenapiServer
// for forward compatibility
type OpenapiServer interface {
	// Query identify.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error)
	// Post addons identify.
	// TKEEL_COMMENT
	// {
	//  "response" :
	//    {
	//      "raw_data": true
	//    }
	// }
	AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error)
	// Post tenant enable.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantEnable(context.Context, *v1.TenantEnableRequest) (*v1.TenantEnableResponse, error)
	// Post tenant disable.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	TenantDisable(context.Context, *v1.TenantDisableRequest) (*v1.TenantDisableResponse, error)
	// Query status.
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	Status(context.Context, *emptypb.Empty) (*v1.StatusResponse, error)
	mustEmbedUnimplementedOpenapiServer()
}

// UnimplementedOpenapiServer must be embedded to have forward compatible implementations.
type UnimplementedOpenapiServer struct {
}

func (UnimplementedOpenapiServer) Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identify not implemented")
}
func (UnimplementedOpenapiServer) AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddonsIdentify not implemented")
}
func (UnimplementedOpenapiServer) TenantEnable(context.Context, *v1.TenantEnableRequest) (*v1.TenantEnableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantEnable not implemented")
}
func (UnimplementedOpenapiServer) TenantDisable(context.Context, *v1.TenantDisableRequest) (*v1.TenantDisableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantDisable not implemented")
}
func (UnimplementedOpenapiServer) Status(context.Context, *emptypb.Empty) (*v1.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedOpenapiServer) mustEmbedUnimplementedOpenapiServer() {}

// UnsafeOpenapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenapiServer will
// result in compilation errors.
type UnsafeOpenapiServer interface {
	mustEmbedUnimplementedOpenapiServer()
}

func RegisterOpenapiServer(s grpc.ServiceRegistrar, srv OpenapiServer) {
	s.RegisterService(&Openapi_ServiceDesc, srv)
}

func _Openapi_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/Identify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).Identify(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_AddonsIdentify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.AddonsIdentifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).AddonsIdentify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/AddonsIdentify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).AddonsIdentify(ctx, req.(*v1.AddonsIdentifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_TenantEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TenantEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).TenantEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/TenantEnable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).TenantEnable(ctx, req.(*v1.TenantEnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_TenantDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TenantDisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).TenantDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/TenantDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).TenantDisable(ctx, req.(*v1.TenantDisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openapi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenapiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openapi.v1.Openapi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenapiServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Openapi_ServiceDesc is the grpc.ServiceDesc for Openapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Openapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openapi.v1.Openapi",
	HandlerType: (*OpenapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Identify",
			Handler:    _Openapi_Identify_Handler,
		},
		{
			MethodName: "AddonsIdentify",
			Handler:    _Openapi_AddonsIdentify_Handler,
		},
		{
			MethodName: "TenantEnable",
			Handler:    _Openapi_TenantEnable_Handler,
		},
		{
			MethodName: "TenantDisable",
			Handler:    _Openapi_TenantDisable_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Openapi_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/openapi/v1/server.proto",
}