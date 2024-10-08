// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: gokeeper_data.proto

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

// KeyValueServiceClient is the client API for KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueServiceClient interface {
	AddKeyValue(ctx context.Context, in *AddKeyValueRequest, opts ...grpc.CallOption) (*AddKeyValueResponse, error)
	GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...grpc.CallOption) (*GetKeyValueResponse, error)
	ListKeyValue(ctx context.Context, in *ListKeyValueRequest, opts ...grpc.CallOption) (*ListKeyValueResponse, error)
	UpdateKeyValue(ctx context.Context, in *UpdateKeyValueRequest, opts ...grpc.CallOption) (*UpdateKeyValueResponse, error)
	DelKeyValue(ctx context.Context, in *DelKeyValueRequest, opts ...grpc.CallOption) (*DelKeyValueResponse, error)
}

type keyValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueServiceClient(cc grpc.ClientConnInterface) KeyValueServiceClient {
	return &keyValueServiceClient{cc}
}

func (c *keyValueServiceClient) AddKeyValue(ctx context.Context, in *AddKeyValueRequest, opts ...grpc.CallOption) (*AddKeyValueResponse, error) {
	out := new(AddKeyValueResponse)
	err := c.cc.Invoke(ctx, "/kim.gokeeper.KeyValueService/AddKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) GetKeyValue(ctx context.Context, in *GetKeyValueRequest, opts ...grpc.CallOption) (*GetKeyValueResponse, error) {
	out := new(GetKeyValueResponse)
	err := c.cc.Invoke(ctx, "/kim.gokeeper.KeyValueService/GetKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) ListKeyValue(ctx context.Context, in *ListKeyValueRequest, opts ...grpc.CallOption) (*ListKeyValueResponse, error) {
	out := new(ListKeyValueResponse)
	err := c.cc.Invoke(ctx, "/kim.gokeeper.KeyValueService/ListKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) UpdateKeyValue(ctx context.Context, in *UpdateKeyValueRequest, opts ...grpc.CallOption) (*UpdateKeyValueResponse, error) {
	out := new(UpdateKeyValueResponse)
	err := c.cc.Invoke(ctx, "/kim.gokeeper.KeyValueService/UpdateKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) DelKeyValue(ctx context.Context, in *DelKeyValueRequest, opts ...grpc.CallOption) (*DelKeyValueResponse, error) {
	out := new(DelKeyValueResponse)
	err := c.cc.Invoke(ctx, "/kim.gokeeper.KeyValueService/DelKeyValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServiceServer is the server API for KeyValueService service.
// All implementations must embed UnimplementedKeyValueServiceServer
// for forward compatibility
type KeyValueServiceServer interface {
	AddKeyValue(context.Context, *AddKeyValueRequest) (*AddKeyValueResponse, error)
	GetKeyValue(context.Context, *GetKeyValueRequest) (*GetKeyValueResponse, error)
	ListKeyValue(context.Context, *ListKeyValueRequest) (*ListKeyValueResponse, error)
	UpdateKeyValue(context.Context, *UpdateKeyValueRequest) (*UpdateKeyValueResponse, error)
	DelKeyValue(context.Context, *DelKeyValueRequest) (*DelKeyValueResponse, error)
	mustEmbedUnimplementedKeyValueServiceServer()
}

// UnimplementedKeyValueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyValueServiceServer struct {
}

func (UnimplementedKeyValueServiceServer) AddKeyValue(context.Context, *AddKeyValueRequest) (*AddKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) GetKeyValue(context.Context, *GetKeyValueRequest) (*GetKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) ListKeyValue(context.Context, *ListKeyValueRequest) (*ListKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) UpdateKeyValue(context.Context, *UpdateKeyValueRequest) (*UpdateKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) DelKeyValue(context.Context, *DelKeyValueRequest) (*DelKeyValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelKeyValue not implemented")
}
func (UnimplementedKeyValueServiceServer) mustEmbedUnimplementedKeyValueServiceServer() {}

// UnsafeKeyValueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueServiceServer will
// result in compilation errors.
type UnsafeKeyValueServiceServer interface {
	mustEmbedUnimplementedKeyValueServiceServer()
}

func RegisterKeyValueServiceServer(s grpc.ServiceRegistrar, srv KeyValueServiceServer) {
	s.RegisterService(&KeyValueService_ServiceDesc, srv)
}

func _KeyValueService_AddKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).AddKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kim.gokeeper.KeyValueService/AddKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).AddKeyValue(ctx, req.(*AddKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_GetKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).GetKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kim.gokeeper.KeyValueService/GetKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).GetKeyValue(ctx, req.(*GetKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_ListKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).ListKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kim.gokeeper.KeyValueService/ListKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).ListKeyValue(ctx, req.(*ListKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_UpdateKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).UpdateKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kim.gokeeper.KeyValueService/UpdateKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).UpdateKeyValue(ctx, req.(*UpdateKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_DelKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).DelKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kim.gokeeper.KeyValueService/DelKeyValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).DelKeyValue(ctx, req.(*DelKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueService_ServiceDesc is the grpc.ServiceDesc for KeyValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kim.gokeeper.KeyValueService",
	HandlerType: (*KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddKeyValue",
			Handler:    _KeyValueService_AddKeyValue_Handler,
		},
		{
			MethodName: "GetKeyValue",
			Handler:    _KeyValueService_GetKeyValue_Handler,
		},
		{
			MethodName: "ListKeyValue",
			Handler:    _KeyValueService_ListKeyValue_Handler,
		},
		{
			MethodName: "UpdateKeyValue",
			Handler:    _KeyValueService_UpdateKeyValue_Handler,
		},
		{
			MethodName: "DelKeyValue",
			Handler:    _KeyValueService_DelKeyValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gokeeper_data.proto",
}
