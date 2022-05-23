// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: orders.proto

package investapi

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

// OrdersStreamServiceClient is the client API for OrdersStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersStreamServiceClient interface {
	//Stream сделок пользователя
	TradesStream(ctx context.Context, in *TradesStreamRequest, opts ...grpc.CallOption) (OrdersStreamService_TradesStreamClient, error)
}

type ordersStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersStreamServiceClient(cc grpc.ClientConnInterface) OrdersStreamServiceClient {
	return &ordersStreamServiceClient{cc}
}

func (c *ordersStreamServiceClient) TradesStream(ctx context.Context, in *TradesStreamRequest, opts ...grpc.CallOption) (OrdersStreamService_TradesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrdersStreamService_ServiceDesc.Streams[0], "/tinkoff.public.invest.api.contract.v1.OrdersStreamService/TradesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &ordersStreamServiceTradesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrdersStreamService_TradesStreamClient interface {
	Recv() (*TradesStreamResponse, error)
	grpc.ClientStream
}

type ordersStreamServiceTradesStreamClient struct {
	grpc.ClientStream
}

func (x *ordersStreamServiceTradesStreamClient) Recv() (*TradesStreamResponse, error) {
	m := new(TradesStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrdersStreamServiceServer is the server API for OrdersStreamService service.
// All implementations must embed UnimplementedOrdersStreamServiceServer
// for forward compatibility
type OrdersStreamServiceServer interface {
	//Stream сделок пользователя
	TradesStream(*TradesStreamRequest, OrdersStreamService_TradesStreamServer) error
	mustEmbedUnimplementedOrdersStreamServiceServer()
}

// UnimplementedOrdersStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersStreamServiceServer struct {
}

func (UnimplementedOrdersStreamServiceServer) TradesStream(*TradesStreamRequest, OrdersStreamService_TradesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TradesStream not implemented")
}
func (UnimplementedOrdersStreamServiceServer) mustEmbedUnimplementedOrdersStreamServiceServer() {}

// UnsafeOrdersStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersStreamServiceServer will
// result in compilation errors.
type UnsafeOrdersStreamServiceServer interface {
	mustEmbedUnimplementedOrdersStreamServiceServer()
}

func RegisterOrdersStreamServiceServer(s grpc.ServiceRegistrar, srv OrdersStreamServiceServer) {
	s.RegisterService(&OrdersStreamService_ServiceDesc, srv)
}

func _OrdersStreamService_TradesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TradesStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrdersStreamServiceServer).TradesStream(m, &ordersStreamServiceTradesStreamServer{stream})
}

type OrdersStreamService_TradesStreamServer interface {
	Send(*TradesStreamResponse) error
	grpc.ServerStream
}

type ordersStreamServiceTradesStreamServer struct {
	grpc.ServerStream
}

func (x *ordersStreamServiceTradesStreamServer) Send(m *TradesStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OrdersStreamService_ServiceDesc is the grpc.ServiceDesc for OrdersStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.OrdersStreamService",
	HandlerType: (*OrdersStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TradesStream",
			Handler:       _OrdersStreamService_TradesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orders.proto",
}

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	//Метод выставления заявки.
	PostOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error)
	//Метод отмены биржевой заявки.
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	//Метод получения статуса торгового поручения.
	GetOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error)
	//Метод получения списка активных заявок по счёту.
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) PostOrder(ctx context.Context, in *PostOrderRequest, opts ...grpc.CallOption) (*PostOrderResponse, error) {
	out := new(PostOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrderState(ctx context.Context, in *GetOrderStateRequest, opts ...grpc.CallOption) (*OrderState, error) {
	out := new(OrderState)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	//Метод выставления заявки.
	PostOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error)
	//Метод отмены биржевой заявки.
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	//Метод получения статуса торгового поручения.
	GetOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error)
	//Метод получения списка активных заявок по счёту.
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) PostOrder(context.Context, *PostOrderRequest) (*PostOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostOrder not implemented")
}
func (UnimplementedOrdersServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrderState(context.Context, *GetOrderStateRequest) (*OrderState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderState not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_PostOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).PostOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).PostOrder(ctx, req.(*PostOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrderState(ctx, req.(*GetOrderStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostOrder",
			Handler:    _OrdersService_PostOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrdersService_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrderState",
			Handler:    _OrdersService_GetOrderState_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrdersService_GetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders.proto",
}
