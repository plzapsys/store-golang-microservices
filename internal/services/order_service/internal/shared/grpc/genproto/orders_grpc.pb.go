// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: order_service/orders.proto

package orders_service

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
	OrdersService_CreateOrder_FullMethodName        = "/orders_service.OrdersService/CreateOrder"
	OrdersService_SubmitOrder_FullMethodName        = "/orders_service.OrdersService/SubmitOrder"
	OrdersService_UpdateShoppingCart_FullMethodName = "/orders_service.OrdersService/UpdateShoppingCart"
	OrdersService_GetOrderByID_FullMethodName       = "/orders_service.OrdersService/GetOrderByID"
	OrdersService_GetOrders_FullMethodName          = "/orders_service.OrdersService/GetOrders"
)

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error)
	SubmitOrder(ctx context.Context, in *SubmitOrderReq, opts ...grpc.CallOption) (*SubmitOrderRes, error)
	UpdateShoppingCart(ctx context.Context, in *UpdateShoppingCartReq, opts ...grpc.CallOption) (*UpdateShoppingCartRes, error)
	GetOrderByID(ctx context.Context, in *GetOrderByIDReq, opts ...grpc.CallOption) (*GetOrderByIDRes, error)
	GetOrders(ctx context.Context, in *GetOrdersReq, opts ...grpc.CallOption) (*GetOrdersRes, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRes, error) {
	out := new(CreateOrderRes)
	err := c.cc.Invoke(ctx, OrdersService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) SubmitOrder(ctx context.Context, in *SubmitOrderReq, opts ...grpc.CallOption) (*SubmitOrderRes, error) {
	out := new(SubmitOrderRes)
	err := c.cc.Invoke(ctx, OrdersService_SubmitOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) UpdateShoppingCart(ctx context.Context, in *UpdateShoppingCartReq, opts ...grpc.CallOption) (*UpdateShoppingCartRes, error) {
	out := new(UpdateShoppingCartRes)
	err := c.cc.Invoke(ctx, OrdersService_UpdateShoppingCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrderByID(ctx context.Context, in *GetOrderByIDReq, opts ...grpc.CallOption) (*GetOrderByIDRes, error) {
	out := new(GetOrderByIDRes)
	err := c.cc.Invoke(ctx, OrdersService_GetOrderByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrders(ctx context.Context, in *GetOrdersReq, opts ...grpc.CallOption) (*GetOrdersRes, error) {
	out := new(GetOrdersRes)
	err := c.cc.Invoke(ctx, OrdersService_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations should embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error)
	SubmitOrder(context.Context, *SubmitOrderReq) (*SubmitOrderRes, error)
	UpdateShoppingCart(context.Context, *UpdateShoppingCartReq) (*UpdateShoppingCartRes, error)
	GetOrderByID(context.Context, *GetOrderByIDReq) (*GetOrderByIDRes, error)
	GetOrders(context.Context, *GetOrdersReq) (*GetOrdersRes, error)
}

// UnimplementedOrdersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrdersServiceServer) SubmitOrder(context.Context, *SubmitOrderReq) (*SubmitOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrdersServiceServer) UpdateShoppingCart(context.Context, *UpdateShoppingCartReq) (*UpdateShoppingCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShoppingCart not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrderByID(context.Context, *GetOrderByIDReq) (*GetOrderByIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrders(context.Context, *GetOrdersReq) (*GetOrdersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_SubmitOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).SubmitOrder(ctx, req.(*SubmitOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_UpdateShoppingCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShoppingCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).UpdateShoppingCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_UpdateShoppingCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).UpdateShoppingCart(ctx, req.(*UpdateShoppingCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_GetOrderByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrderByID(ctx, req.(*GetOrderByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrders(ctx, req.(*GetOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders_service.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrdersService_CreateOrder_Handler,
		},
		{
			MethodName: "SubmitOrder",
			Handler:    _OrdersService_SubmitOrder_Handler,
		},
		{
			MethodName: "UpdateShoppingCart",
			Handler:    _OrdersService_UpdateShoppingCart_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _OrdersService_GetOrderByID_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrdersService_GetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service/orders.proto",
}
