// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: WeatherService.proto

package weatherService

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

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetCurrentWeather(ctx context.Context, in *Place, opts ...grpc.CallOption) (*GetCurrentWeatherResponse, error)
	GetWeatherForecast(ctx context.Context, in *Place, opts ...grpc.CallOption) (*GetWeatherForecastResponse, error)
	ReportWeatherCondition(ctx context.Context, in *ReportWeatherConditionRequest, opts ...grpc.CallOption) (*ReportWeatherConditionResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetCurrentWeather(ctx context.Context, in *Place, opts ...grpc.CallOption) (*GetCurrentWeatherResponse, error) {
	out := new(GetCurrentWeatherResponse)
	err := c.cc.Invoke(ctx, "/weatherService.WeatherService/GetCurrentWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetWeatherForecast(ctx context.Context, in *Place, opts ...grpc.CallOption) (*GetWeatherForecastResponse, error) {
	out := new(GetWeatherForecastResponse)
	err := c.cc.Invoke(ctx, "/weatherService.WeatherService/GetWeatherForecast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) ReportWeatherCondition(ctx context.Context, in *ReportWeatherConditionRequest, opts ...grpc.CallOption) (*ReportWeatherConditionResponse, error) {
	out := new(ReportWeatherConditionResponse)
	err := c.cc.Invoke(ctx, "/weatherService.WeatherService/ReportWeatherCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetCurrentWeather(context.Context, *Place) (*GetCurrentWeatherResponse, error)
	GetWeatherForecast(context.Context, *Place) (*GetWeatherForecastResponse, error)
	ReportWeatherCondition(context.Context, *ReportWeatherConditionRequest) (*ReportWeatherConditionResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetCurrentWeather(context.Context, *Place) (*GetCurrentWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentWeather not implemented")
}
func (UnimplementedWeatherServiceServer) GetWeatherForecast(context.Context, *Place) (*GetWeatherForecastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherForecast not implemented")
}
func (UnimplementedWeatherServiceServer) ReportWeatherCondition(context.Context, *ReportWeatherConditionRequest) (*ReportWeatherConditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportWeatherCondition not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetCurrentWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Place)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherService.WeatherService/GetCurrentWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, req.(*Place))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetWeatherForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Place)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherService.WeatherService/GetWeatherForecast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, req.(*Place))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_ReportWeatherCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportWeatherConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherService.WeatherService/ReportWeatherCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, req.(*ReportWeatherConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weatherService.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentWeather",
			Handler:    _WeatherService_GetCurrentWeather_Handler,
		},
		{
			MethodName: "GetWeatherForecast",
			Handler:    _WeatherService_GetWeatherForecast_Handler,
		},
		{
			MethodName: "ReportWeatherCondition",
			Handler:    _WeatherService_ReportWeatherCondition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "WeatherService.proto",
}
