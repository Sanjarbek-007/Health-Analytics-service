// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: medical-record.proto

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
	MedicalRecord_AddMedicalReport_FullMethodName     = "/medical.MedicalRecord/AddMedicalReport"
	MedicalRecord_GetMedicalReport_FullMethodName     = "/medical.MedicalRecord/GetMedicalReport"
	MedicalRecord_GetMedicalReportById_FullMethodName = "/medical.MedicalRecord/GetMedicalReportById"
	MedicalRecord_UpdateMedicalReport_FullMethodName  = "/medical.MedicalRecord/UpdateMedicalReport"
	MedicalRecord_DeleteMedicalReport_FullMethodName  = "/medical.MedicalRecord/DeleteMedicalReport"
)

// MedicalRecordClient is the client API for MedicalRecord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicalRecordClient interface {
	AddMedicalReport(ctx context.Context, in *AddMedicalReportReq, opts ...grpc.CallOption) (*AddMedicalReportRes, error)
	GetMedicalReport(ctx context.Context, in *GetMedicalReportReq, opts ...grpc.CallOption) (*GetMedicalReportRes, error)
	GetMedicalReportById(ctx context.Context, in *GetMedicalReportByIdReq, opts ...grpc.CallOption) (*GetMedicalReportByIdRes, error)
	UpdateMedicalReport(ctx context.Context, in *UpdateMedicalReportReq, opts ...grpc.CallOption) (*UpdateMedicalReportRes, error)
	DeleteMedicalReport(ctx context.Context, in *DeleteMedicalReportReq, opts ...grpc.CallOption) (*DeleteMedicalReportRes, error)
}

type medicalRecordClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalRecordClient(cc grpc.ClientConnInterface) MedicalRecordClient {
	return &medicalRecordClient{cc}
}

func (c *medicalRecordClient) AddMedicalReport(ctx context.Context, in *AddMedicalReportReq, opts ...grpc.CallOption) (*AddMedicalReportRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMedicalReportRes)
	err := c.cc.Invoke(ctx, MedicalRecord_AddMedicalReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordClient) GetMedicalReport(ctx context.Context, in *GetMedicalReportReq, opts ...grpc.CallOption) (*GetMedicalReportRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicalReportRes)
	err := c.cc.Invoke(ctx, MedicalRecord_GetMedicalReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordClient) GetMedicalReportById(ctx context.Context, in *GetMedicalReportByIdReq, opts ...grpc.CallOption) (*GetMedicalReportByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicalReportByIdRes)
	err := c.cc.Invoke(ctx, MedicalRecord_GetMedicalReportById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordClient) UpdateMedicalReport(ctx context.Context, in *UpdateMedicalReportReq, opts ...grpc.CallOption) (*UpdateMedicalReportRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMedicalReportRes)
	err := c.cc.Invoke(ctx, MedicalRecord_UpdateMedicalReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalRecordClient) DeleteMedicalReport(ctx context.Context, in *DeleteMedicalReportReq, opts ...grpc.CallOption) (*DeleteMedicalReportRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMedicalReportRes)
	err := c.cc.Invoke(ctx, MedicalRecord_DeleteMedicalReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalRecordServer is the server API for MedicalRecord service.
// All implementations must embed UnimplementedMedicalRecordServer
// for forward compatibility
type MedicalRecordServer interface {
	AddMedicalReport(context.Context, *AddMedicalReportReq) (*AddMedicalReportRes, error)
	GetMedicalReport(context.Context, *GetMedicalReportReq) (*GetMedicalReportRes, error)
	GetMedicalReportById(context.Context, *GetMedicalReportByIdReq) (*GetMedicalReportByIdRes, error)
	UpdateMedicalReport(context.Context, *UpdateMedicalReportReq) (*UpdateMedicalReportRes, error)
	DeleteMedicalReport(context.Context, *DeleteMedicalReportReq) (*DeleteMedicalReportRes, error)
	mustEmbedUnimplementedMedicalRecordServer()
}

// UnimplementedMedicalRecordServer must be embedded to have forward compatible implementations.
type UnimplementedMedicalRecordServer struct {
}

func (UnimplementedMedicalRecordServer) AddMedicalReport(context.Context, *AddMedicalReportReq) (*AddMedicalReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMedicalReport not implemented")
}
func (UnimplementedMedicalRecordServer) GetMedicalReport(context.Context, *GetMedicalReportReq) (*GetMedicalReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalReport not implemented")
}
func (UnimplementedMedicalRecordServer) GetMedicalReportById(context.Context, *GetMedicalReportByIdReq) (*GetMedicalReportByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalReportById not implemented")
}
func (UnimplementedMedicalRecordServer) UpdateMedicalReport(context.Context, *UpdateMedicalReportReq) (*UpdateMedicalReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMedicalReport not implemented")
}
func (UnimplementedMedicalRecordServer) DeleteMedicalReport(context.Context, *DeleteMedicalReportReq) (*DeleteMedicalReportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedicalReport not implemented")
}
func (UnimplementedMedicalRecordServer) mustEmbedUnimplementedMedicalRecordServer() {}

// UnsafeMedicalRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicalRecordServer will
// result in compilation errors.
type UnsafeMedicalRecordServer interface {
	mustEmbedUnimplementedMedicalRecordServer()
}

func RegisterMedicalRecordServer(s grpc.ServiceRegistrar, srv MedicalRecordServer) {
	s.RegisterService(&MedicalRecord_ServiceDesc, srv)
}

func _MedicalRecord_AddMedicalReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMedicalReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServer).AddMedicalReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecord_AddMedicalReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServer).AddMedicalReport(ctx, req.(*AddMedicalReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecord_GetMedicalReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServer).GetMedicalReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecord_GetMedicalReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServer).GetMedicalReport(ctx, req.(*GetMedicalReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecord_GetMedicalReportById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalReportByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServer).GetMedicalReportById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecord_GetMedicalReportById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServer).GetMedicalReportById(ctx, req.(*GetMedicalReportByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecord_UpdateMedicalReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMedicalReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServer).UpdateMedicalReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecord_UpdateMedicalReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServer).UpdateMedicalReport(ctx, req.(*UpdateMedicalReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalRecord_DeleteMedicalReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMedicalReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalRecordServer).DeleteMedicalReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MedicalRecord_DeleteMedicalReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalRecordServer).DeleteMedicalReport(ctx, req.(*DeleteMedicalReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicalRecord_ServiceDesc is the grpc.ServiceDesc for MedicalRecord service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicalRecord_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medical.MedicalRecord",
	HandlerType: (*MedicalRecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMedicalReport",
			Handler:    _MedicalRecord_AddMedicalReport_Handler,
		},
		{
			MethodName: "GetMedicalReport",
			Handler:    _MedicalRecord_GetMedicalReport_Handler,
		},
		{
			MethodName: "GetMedicalReportById",
			Handler:    _MedicalRecord_GetMedicalReportById_Handler,
		},
		{
			MethodName: "UpdateMedicalReport",
			Handler:    _MedicalRecord_UpdateMedicalReport_Handler,
		},
		{
			MethodName: "DeleteMedicalReport",
			Handler:    _MedicalRecord_DeleteMedicalReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical-record.proto",
}