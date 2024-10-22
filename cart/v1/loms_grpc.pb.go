// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: loms.proto

package grpc_loms

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

const (
	LOMSService_OrderCancel_FullMethodName = "/ecom.loms.LOMSService/OrderCancel"
	LOMSService_OrderCreate_FullMethodName = "/ecom.loms.LOMSService/OrderCreate"
	LOMSService_OrderInfo_FullMethodName   = "/ecom.loms.LOMSService/OrderInfo"
	LOMSService_OrderPay_FullMethodName    = "/ecom.loms.LOMSService/OrderPay"
	LOMSService_StockInfo_FullMethodName   = "/ecom.loms.LOMSService/StockInfo"
)

// LOMSServiceClient is the client API for LOMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LOMSServiceClient interface {
	OrderCancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OrderCancelResponse, error)
	OrderCreate(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error)
	OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error)
	OrderPay(ctx context.Context, in *OrderPayRequest, opts ...grpc.CallOption) (*OrderPayResponse, error)
	StockInfo(ctx context.Context, in *StockInfoRequest, opts ...grpc.CallOption) (*StockInfoResponse, error)
}

type lOMSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLOMSServiceClient(cc grpc.ClientConnInterface) LOMSServiceClient {
	return &lOMSServiceClient{cc}
}

func (c *lOMSServiceClient) OrderCancel(ctx context.Context, in *OrderCancelRequest, opts ...grpc.CallOption) (*OrderCancelResponse, error) {
	out := new(OrderCancelResponse)
	err := c.cc.Invoke(ctx, LOMSService_OrderCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOMSServiceClient) OrderCreate(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error) {
	out := new(OrderCreateResponse)
	err := c.cc.Invoke(ctx, LOMSService_OrderCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOMSServiceClient) OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error) {
	out := new(OrderInfoResponse)
	err := c.cc.Invoke(ctx, LOMSService_OrderInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOMSServiceClient) OrderPay(ctx context.Context, in *OrderPayRequest, opts ...grpc.CallOption) (*OrderPayResponse, error) {
	out := new(OrderPayResponse)
	err := c.cc.Invoke(ctx, LOMSService_OrderPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lOMSServiceClient) StockInfo(ctx context.Context, in *StockInfoRequest, opts ...grpc.CallOption) (*StockInfoResponse, error) {
	out := new(StockInfoResponse)
	err := c.cc.Invoke(ctx, LOMSService_StockInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LOMSServiceServer is the server API for LOMSService service.
// All implementations must embed UnimplementedLOMSServiceServer
// for forward compatibility
type LOMSServiceServer interface {
	OrderCancel(context.Context, *OrderCancelRequest) (*OrderCancelResponse, error)
	OrderCreate(context.Context, *OrderCreateRequest) (*OrderCreateResponse, error)
	OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoResponse, error)
	OrderPay(context.Context, *OrderPayRequest) (*OrderPayResponse, error)
	StockInfo(context.Context, *StockInfoRequest) (*StockInfoResponse, error)
	mustEmbedUnimplementedLOMSServiceServer()
}

// UnimplementedLOMSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLOMSServiceServer struct {
}

func (UnimplementedLOMSServiceServer) OrderCancel(context.Context, *OrderCancelRequest) (*OrderCancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancel not implemented")
}
func (UnimplementedLOMSServiceServer) OrderCreate(context.Context, *OrderCreateRequest) (*OrderCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCreate not implemented")
}
func (UnimplementedLOMSServiceServer) OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderInfo not implemented")
}
func (UnimplementedLOMSServiceServer) OrderPay(context.Context, *OrderPayRequest) (*OrderPayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPay not implemented")
}
func (UnimplementedLOMSServiceServer) StockInfo(context.Context, *StockInfoRequest) (*StockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StockInfo not implemented")
}
func (UnimplementedLOMSServiceServer) mustEmbedUnimplementedLOMSServiceServer() {}

// UnsafeLOMSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LOMSServiceServer will
// result in compilation errors.
type UnsafeLOMSServiceServer interface {
	mustEmbedUnimplementedLOMSServiceServer()
}

func RegisterLOMSServiceServer(s grpc.ServiceRegistrar, srv LOMSServiceServer) {
	s.RegisterService(&LOMSService_ServiceDesc, srv)
}

func _LOMSService_OrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOMSServiceServer).OrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LOMSService_OrderCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOMSServiceServer).OrderCancel(ctx, req.(*OrderCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOMSService_OrderCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOMSServiceServer).OrderCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LOMSService_OrderCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOMSServiceServer).OrderCreate(ctx, req.(*OrderCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOMSService_OrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOMSServiceServer).OrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LOMSService_OrderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOMSServiceServer).OrderInfo(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOMSService_OrderPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOMSServiceServer).OrderPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LOMSService_OrderPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOMSServiceServer).OrderPay(ctx, req.(*OrderPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LOMSService_StockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LOMSServiceServer).StockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LOMSService_StockInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LOMSServiceServer).StockInfo(ctx, req.(*StockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LOMSService_ServiceDesc is the grpc.ServiceDesc for LOMSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LOMSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecom.loms.LOMSService",
	HandlerType: (*LOMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderCancel",
			Handler:    _LOMSService_OrderCancel_Handler,
		},
		{
			MethodName: "OrderCreate",
			Handler:    _LOMSService_OrderCreate_Handler,
		},
		{
			MethodName: "OrderInfo",
			Handler:    _LOMSService_OrderInfo_Handler,
		},
		{
			MethodName: "OrderPay",
			Handler:    _LOMSService_OrderPay_Handler,
		},
		{
			MethodName: "StockInfo",
			Handler:    _LOMSService_StockInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loms.proto",
}