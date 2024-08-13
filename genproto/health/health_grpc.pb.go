// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: health.proto

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
	HealthCheck_GenerateHealthRecommendations_FullMethodName = "/health.HealthCheck/GenerateHealthRecommendations"
	HealthCheck_GetRealtimeHealthMonitoring_FullMethodName   = "/health.HealthCheck/GetRealtimeHealthMonitoring"
	HealthCheck_GetDailyHealthSummary_FullMethodName         = "/health.HealthCheck/GetDailyHealthSummary"
	HealthCheck_GetWeeklyHealthSummary_FullMethodName        = "/health.HealthCheck/GetWeeklyHealthSummary"
	HealthCheck_AddRealTimeData_FullMethodName               = "/health.HealthCheck/AddRealTimeData"
	HealthCheck_GetRealTimeData_FullMethodName               = "/health.HealthCheck/GetRealTimeData"
)

// HealthCheckClient is the client API for HealthCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthCheckClient interface {
	GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsReq, opts ...grpc.CallOption) (*GenerateHealthRecommendationsRes, error)
	GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringReq, opts ...grpc.CallOption) (*GetRealtimeHealthMonitoringRes, error)
	GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryReq, opts ...grpc.CallOption) (*GetDailyHealthSummaryRes, error)
	GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryReq, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryRes, error)
	AddRealTimeData(ctx context.Context, in *AddRealTimeDataReq, opts ...grpc.CallOption) (*AddRealTimeDataRes, error)
	GetRealTimeData(ctx context.Context, in *GetRealTimeDataReq, opts ...grpc.CallOption) (*GetRealTimeDataRes, error)
}

type healthCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthCheckClient(cc grpc.ClientConnInterface) HealthCheckClient {
	return &healthCheckClient{cc}
}

func (c *healthCheckClient) GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsReq, opts ...grpc.CallOption) (*GenerateHealthRecommendationsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateHealthRecommendationsRes)
	err := c.cc.Invoke(ctx, HealthCheck_GenerateHealthRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckClient) GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringReq, opts ...grpc.CallOption) (*GetRealtimeHealthMonitoringRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRealtimeHealthMonitoringRes)
	err := c.cc.Invoke(ctx, HealthCheck_GetRealtimeHealthMonitoring_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckClient) GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryReq, opts ...grpc.CallOption) (*GetDailyHealthSummaryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDailyHealthSummaryRes)
	err := c.cc.Invoke(ctx, HealthCheck_GetDailyHealthSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckClient) GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryReq, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeeklyHealthSummaryRes)
	err := c.cc.Invoke(ctx, HealthCheck_GetWeeklyHealthSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckClient) AddRealTimeData(ctx context.Context, in *AddRealTimeDataReq, opts ...grpc.CallOption) (*AddRealTimeDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddRealTimeDataRes)
	err := c.cc.Invoke(ctx, HealthCheck_AddRealTimeData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckClient) GetRealTimeData(ctx context.Context, in *GetRealTimeDataReq, opts ...grpc.CallOption) (*GetRealTimeDataRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRealTimeDataRes)
	err := c.cc.Invoke(ctx, HealthCheck_GetRealTimeData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthCheckServer is the server API for HealthCheck service.
// All implementations must embed UnimplementedHealthCheckServer
// for forward compatibility
type HealthCheckServer interface {
	GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsReq) (*GenerateHealthRecommendationsRes, error)
	GetRealtimeHealthMonitoring(context.Context, *GetRealtimeHealthMonitoringReq) (*GetRealtimeHealthMonitoringRes, error)
	GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryReq) (*GetDailyHealthSummaryRes, error)
	GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryReq) (*GetWeeklyHealthSummaryRes, error)
	AddRealTimeData(context.Context, *AddRealTimeDataReq) (*AddRealTimeDataRes, error)
	GetRealTimeData(context.Context, *GetRealTimeDataReq) (*GetRealTimeDataRes, error)
	mustEmbedUnimplementedHealthCheckServer()
}

// UnimplementedHealthCheckServer must be embedded to have forward compatible implementations.
type UnimplementedHealthCheckServer struct {
}

func (UnimplementedHealthCheckServer) GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsReq) (*GenerateHealthRecommendationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHealthRecommendations not implemented")
}
func (UnimplementedHealthCheckServer) GetRealtimeHealthMonitoring(context.Context, *GetRealtimeHealthMonitoringReq) (*GetRealtimeHealthMonitoringRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealtimeHealthMonitoring not implemented")
}
func (UnimplementedHealthCheckServer) GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryReq) (*GetDailyHealthSummaryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyHealthSummary not implemented")
}
func (UnimplementedHealthCheckServer) GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryReq) (*GetWeeklyHealthSummaryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyHealthSummary not implemented")
}
func (UnimplementedHealthCheckServer) AddRealTimeData(context.Context, *AddRealTimeDataReq) (*AddRealTimeDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRealTimeData not implemented")
}
func (UnimplementedHealthCheckServer) GetRealTimeData(context.Context, *GetRealTimeDataReq) (*GetRealTimeDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealTimeData not implemented")
}
func (UnimplementedHealthCheckServer) mustEmbedUnimplementedHealthCheckServer() {}

// UnsafeHealthCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthCheckServer will
// result in compilation errors.
type UnsafeHealthCheckServer interface {
	mustEmbedUnimplementedHealthCheckServer()
}

func RegisterHealthCheckServer(s grpc.ServiceRegistrar, srv HealthCheckServer) {
	s.RegisterService(&HealthCheck_ServiceDesc, srv)
}

func _HealthCheck_GenerateHealthRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHealthRecommendationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).GenerateHealthRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_GenerateHealthRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).GenerateHealthRecommendations(ctx, req.(*GenerateHealthRecommendationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheck_GetRealtimeHealthMonitoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealtimeHealthMonitoringReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).GetRealtimeHealthMonitoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_GetRealtimeHealthMonitoring_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).GetRealtimeHealthMonitoring(ctx, req.(*GetRealtimeHealthMonitoringReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheck_GetDailyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDailyHealthSummaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).GetDailyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_GetDailyHealthSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).GetDailyHealthSummary(ctx, req.(*GetDailyHealthSummaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheck_GetWeeklyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeeklyHealthSummaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).GetWeeklyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_GetWeeklyHealthSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).GetWeeklyHealthSummary(ctx, req.(*GetWeeklyHealthSummaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheck_AddRealTimeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRealTimeDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).AddRealTimeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_AddRealTimeData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).AddRealTimeData(ctx, req.(*AddRealTimeDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheck_GetRealTimeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealTimeDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).GetRealTimeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheck_GetRealTimeData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).GetRealTimeData(ctx, req.(*GetRealTimeDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthCheck_ServiceDesc is the grpc.ServiceDesc for HealthCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health.HealthCheck",
	HandlerType: (*HealthCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateHealthRecommendations",
			Handler:    _HealthCheck_GenerateHealthRecommendations_Handler,
		},
		{
			MethodName: "GetRealtimeHealthMonitoring",
			Handler:    _HealthCheck_GetRealtimeHealthMonitoring_Handler,
		},
		{
			MethodName: "GetDailyHealthSummary",
			Handler:    _HealthCheck_GetDailyHealthSummary_Handler,
		},
		{
			MethodName: "GetWeeklyHealthSummary",
			Handler:    _HealthCheck_GetWeeklyHealthSummary_Handler,
		},
		{
			MethodName: "AddRealTimeData",
			Handler:    _HealthCheck_AddRealTimeData_Handler,
		},
		{
			MethodName: "GetRealTimeData",
			Handler:    _HealthCheck_GetRealTimeData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}
