// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculatorpb

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorReponse, error)
	GetPrime(ctx context.Context, in *GetPrimeRequest, opts ...grpc.CallOption) (CalculateService_GetPrimeClient, error)
	GetAvg(ctx context.Context, opts ...grpc.CallOption) (CalculateService_GetAvgClient, error)
	GetMax(ctx context.Context, opts ...grpc.CallOption) (CalculateService_GetMaxClient, error)
	// Error handling
	// This RPC will throw an exception if the number is negative
	// The error being sent is of of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorReponse, error) {
	out := new(CalculatorReponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) GetPrime(ctx context.Context, in *GetPrimeRequest, opts ...grpc.CallOption) (CalculateService_GetPrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[0], "/calculator.CalculateService/GetPrime", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceGetPrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_GetPrimeClient interface {
	Recv() (*GetPrimeResponse, error)
	grpc.ClientStream
}

type calculateServiceGetPrimeClient struct {
	grpc.ClientStream
}

func (x *calculateServiceGetPrimeClient) Recv() (*GetPrimeResponse, error) {
	m := new(GetPrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) GetAvg(ctx context.Context, opts ...grpc.CallOption) (CalculateService_GetAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[1], "/calculator.CalculateService/GetAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceGetAvgClient{stream}
	return x, nil
}

type CalculateService_GetAvgClient interface {
	Send(*GetAvgRequest) error
	CloseAndRecv() (*GetAvgResponse, error)
	grpc.ClientStream
}

type calculateServiceGetAvgClient struct {
	grpc.ClientStream
}

func (x *calculateServiceGetAvgClient) Send(m *GetAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceGetAvgClient) CloseAndRecv() (*GetAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) GetMax(ctx context.Context, opts ...grpc.CallOption) (CalculateService_GetMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[2], "/calculator.CalculateService/GetMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceGetMaxClient{stream}
	return x, nil
}

type CalculateService_GetMaxClient interface {
	Send(*GetMaxRequest) error
	Recv() (*GetMaxResponse, error)
	grpc.ClientStream
}

type calculateServiceGetMaxClient struct {
	grpc.ClientStream
}

func (x *calculateServiceGetMaxClient) Send(m *GetMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceGetMaxClient) Recv() (*GetMaxResponse, error) {
	m := new(GetMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations must embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	Calculate(context.Context, *CalculatorRequest) (*CalculatorReponse, error)
	GetPrime(*GetPrimeRequest, CalculateService_GetPrimeServer) error
	GetAvg(CalculateService_GetAvgServer) error
	GetMax(CalculateService_GetMaxServer) error
	// Error handling
	// This RPC will throw an exception if the number is negative
	// The error being sent is of of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
	mustEmbedUnimplementedCalculateServiceServer()
}

// UnimplementedCalculateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) Calculate(context.Context, *CalculatorRequest) (*CalculatorReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculateServiceServer) GetPrime(*GetPrimeRequest, CalculateService_GetPrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrime not implemented")
}
func (UnimplementedCalculateServiceServer) GetAvg(CalculateService_GetAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAvg not implemented")
}
func (UnimplementedCalculateServiceServer) GetMax(CalculateService_GetMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMax not implemented")
}
func (UnimplementedCalculateServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}
func (UnimplementedCalculateServiceServer) mustEmbedUnimplementedCalculateServiceServer() {}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Calculate(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_GetPrime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).GetPrime(m, &calculateServiceGetPrimeServer{stream})
}

type CalculateService_GetPrimeServer interface {
	Send(*GetPrimeResponse) error
	grpc.ServerStream
}

type calculateServiceGetPrimeServer struct {
	grpc.ServerStream
}

func (x *calculateServiceGetPrimeServer) Send(m *GetPrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_GetAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).GetAvg(&calculateServiceGetAvgServer{stream})
}

type CalculateService_GetAvgServer interface {
	SendAndClose(*GetAvgResponse) error
	Recv() (*GetAvgRequest, error)
	grpc.ServerStream
}

type calculateServiceGetAvgServer struct {
	grpc.ServerStream
}

func (x *calculateServiceGetAvgServer) SendAndClose(m *GetAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceGetAvgServer) Recv() (*GetAvgRequest, error) {
	m := new(GetAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_GetMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).GetMax(&calculateServiceGetMaxServer{stream})
}

type CalculateService_GetMaxServer interface {
	Send(*GetMaxResponse) error
	Recv() (*GetMaxRequest, error)
	grpc.ServerStream
}

type calculateServiceGetMaxServer struct {
	grpc.ServerStream
}

func (x *calculateServiceGetMaxServer) Send(m *GetMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceGetMaxServer) Recv() (*GetMaxRequest, error) {
	m := new(GetMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculateService_Calculate_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrime",
			Handler:       _CalculateService_GetPrime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAvg",
			Handler:       _CalculateService_GetAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMax",
			Handler:       _CalculateService_GetMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
