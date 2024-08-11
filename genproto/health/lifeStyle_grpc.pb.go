// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: lifeStyle.proto

package health

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LifeStyle_AddLifeStyleData_FullMethodName     = "/lifeStyle.lifeStyle/AddLifeStyleData"
	LifeStyle_GetLifeStyleData_FullMethodName     = "/lifeStyle.lifeStyle/GetLifeStyleData"
	LifeStyle_GetLifeStyleDataById_FullMethodName = "/lifeStyle.lifeStyle/GetLifeStyleDataById"
	LifeStyle_UpdateLifeStyleData_FullMethodName  = "/lifeStyle.lifeStyle/UpdateLifeStyleData"
	LifeStyle_DeleteLifeStyleData_FullMethodName  = "/lifeStyle.lifeStyle/DeleteLifeStyleData"
)

// LifeStyleClient is the client API for LifeStyle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LifeStyleClient interface {
	AddLifeStyleData(ctx context.Context, in *AddLifeStyleDataReq, opts ...grpc.CallOption) (*AddLifeStyleDataRes, error)
	GetLifeStyleData(ctx context.Context, in *GetLifeStyleDataReq, opts ...grpc.CallOption) (*GetLifeStyleDataRes, error)
	GetLifeStyleDataById(ctx context.Context, in *GetLifeStyleDataByIdReq, opts ...grpc.CallOption) (*GetLifeStyleDataRes, error)
	UpdateLifeStyleData(ctx context.Context, in *UpdateLifeStyleDataReq, opts ...grpc.CallOption) (*UpdateLifeStyleDataRes, error)
	DeleteLifeStyleData(ctx context.Context, in *DeleteLifeStyleDataReq, opts ...grpc.CallOption) (*DeleteLifeStyleDataRes, error)
}

type lifeStyleClient struct {
	cc grpc.ClientConnInterface
}

func NewLifeStyleClient(cc grpc.ClientConnInterface) LifeStyleClient {
	return &lifeStyleClient{cc}
}

func (c *lifeStyleClient) AddLifeStyleData(ctx context.Context, in *AddLifeStyleDataReq, opts ...grpc.CallOption) (*AddLifeStyleDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddLifeStyleDataRes)
	err := c.cc.Invoke(ctx, LifeStyle_AddLifeStyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifeStyleClient) GetLifeStyleData(ctx context.Context, in *GetLifeStyleDataReq, opts ...grpc.CallOption) (*GetLifeStyleDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLifeStyleDataRes)
	err := c.cc.Invoke(ctx, LifeStyle_GetLifeStyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifeStyleClient) GetLifeStyleDataById(ctx context.Context, in *GetLifeStyleDataByIdReq, opts ...grpc.CallOption) (*GetLifeStyleDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLifeStyleDataRes)
	err := c.cc.Invoke(ctx, LifeStyle_GetLifeStyleDataById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifeStyleClient) UpdateLifeStyleData(ctx context.Context, in *UpdateLifeStyleDataReq, opts ...grpc.CallOption) (*UpdateLifeStyleDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLifeStyleDataRes)
	err := c.cc.Invoke(ctx, LifeStyle_UpdateLifeStyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifeStyleClient) DeleteLifeStyleData(ctx context.Context, in *DeleteLifeStyleDataReq, opts ...grpc.CallOption) (*DeleteLifeStyleDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLifeStyleDataRes)
	err := c.cc.Invoke(ctx, LifeStyle_DeleteLifeStyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifeStyleServer is the server API for LifeStyle service.
// All implementations must embed UnimplementedLifeStyleServer
// for forward compatibility
type LifeStyleServer interface {
	AddLifeStyleData(context.Context, *AddLifeStyleDataReq) (*AddLifeStyleDataRes, error)
	GetLifeStyleData(context.Context, *GetLifeStyleDataReq) (*GetLifeStyleDataRes, error)
	GetLifeStyleDataById(context.Context, *GetLifeStyleDataByIdReq) (*GetLifeStyleDataRes, error)
	UpdateLifeStyleData(context.Context, *UpdateLifeStyleDataReq) (*UpdateLifeStyleDataRes, error)
	DeleteLifeStyleData(context.Context, *DeleteLifeStyleDataReq) (*DeleteLifeStyleDataRes, error)
	mustEmbedUnimplementedLifeStyleServer()
}

// UnimplementedLifeStyleServer must be embedded to have forward compatible implementations.
type UnimplementedLifeStyleServer struct {
}

func (UnimplementedLifeStyleServer) AddLifeStyleData(context.Context, *AddLifeStyleDataReq) (*AddLifeStyleDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLifeStyleData not implemented")
}
func (UnimplementedLifeStyleServer) GetLifeStyleData(context.Context, *GetLifeStyleDataReq) (*GetLifeStyleDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifeStyleData not implemented")
}
func (UnimplementedLifeStyleServer) GetLifeStyleDataById(context.Context, *GetLifeStyleDataByIdReq) (*GetLifeStyleDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifeStyleDataById not implemented")
}
func (UnimplementedLifeStyleServer) UpdateLifeStyleData(context.Context, *UpdateLifeStyleDataReq) (*UpdateLifeStyleDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLifeStyleData not implemented")
}
func (UnimplementedLifeStyleServer) DeleteLifeStyleData(context.Context, *DeleteLifeStyleDataReq) (*DeleteLifeStyleDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLifeStyleData not implemented")
}
func (UnimplementedLifeStyleServer) mustEmbedUnimplementedLifeStyleServer() {}

// UnsafeLifeStyleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifeStyleServer will
// result in compilation errors.
type UnsafeLifeStyleServer interface {
	mustEmbedUnimplementedLifeStyleServer()
}

func RegisterLifeStyleServer(s grpc.ServiceRegistrar, srv LifeStyleServer) {
	s.RegisterService(&LifeStyle_ServiceDesc, srv)
}

func _LifeStyle_AddLifeStyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLifeStyleDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeStyleServer).AddLifeStyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifeStyle_AddLifeStyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeStyleServer).AddLifeStyleData(ctx, req.(*AddLifeStyleDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifeStyle_GetLifeStyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLifeStyleDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeStyleServer).GetLifeStyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifeStyle_GetLifeStyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeStyleServer).GetLifeStyleData(ctx, req.(*GetLifeStyleDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifeStyle_GetLifeStyleDataById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLifeStyleDataByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeStyleServer).GetLifeStyleDataById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifeStyle_GetLifeStyleDataById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeStyleServer).GetLifeStyleDataById(ctx, req.(*GetLifeStyleDataByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifeStyle_UpdateLifeStyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLifeStyleDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeStyleServer).UpdateLifeStyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifeStyle_UpdateLifeStyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeStyleServer).UpdateLifeStyleData(ctx, req.(*UpdateLifeStyleDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LifeStyle_DeleteLifeStyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLifeStyleDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifeStyleServer).DeleteLifeStyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LifeStyle_DeleteLifeStyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifeStyleServer).DeleteLifeStyleData(ctx, req.(*DeleteLifeStyleDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LifeStyle_ServiceDesc is the grpc.ServiceDesc for LifeStyle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LifeStyle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lifeStyle.lifeStyle",
	HandlerType: (*LifeStyleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLifeStyleData",
			Handler:    _LifeStyle_AddLifeStyleData_Handler,
		},
		{
			MethodName: "GetLifeStyleData",
			Handler:    _LifeStyle_GetLifeStyleData_Handler,
		},
		{
			MethodName: "GetLifeStyleDataById",
			Handler:    _LifeStyle_GetLifeStyleDataById_Handler,
		},
		{
			MethodName: "UpdateLifeStyleData",
			Handler:    _LifeStyle_UpdateLifeStyleData_Handler,
		},
		{
			MethodName: "DeleteLifeStyleData",
			Handler:    _LifeStyle_DeleteLifeStyleData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lifeStyle.proto",
}
