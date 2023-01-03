// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/shop.proto

package go_gRPC_pg

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

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferClient interface {
	GetProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Product, error)
	StreamProduct(ctx context.Context, in *OrderArray, opts ...grpc.CallOption) (Transfer_StreamProductClient, error)
	StreamOrder(ctx context.Context, opts ...grpc.CallOption) (Transfer_StreamOrderClient, error)
}

type transferClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferClient(cc grpc.ClientConnInterface) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) GetProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/Transfer/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferClient) StreamProduct(ctx context.Context, in *OrderArray, opts ...grpc.CallOption) (Transfer_StreamProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transfer_ServiceDesc.Streams[0], "/Transfer/StreamProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferStreamProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transfer_StreamProductClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type transferStreamProductClient struct {
	grpc.ClientStream
}

func (x *transferStreamProductClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transferClient) StreamOrder(ctx context.Context, opts ...grpc.CallOption) (Transfer_StreamOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transfer_ServiceDesc.Streams[1], "/Transfer/StreamOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferStreamOrderClient{stream}
	return x, nil
}

type Transfer_StreamOrderClient interface {
	Send(*Order) error
	CloseAndRecv() (*Product, error)
	grpc.ClientStream
}

type transferStreamOrderClient struct {
	grpc.ClientStream
}

func (x *transferStreamOrderClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferStreamOrderClient) CloseAndRecv() (*Product, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServer is the server API for Transfer service.
// All implementations must embed UnimplementedTransferServer
// for forward compatibility
type TransferServer interface {
	GetProduct(context.Context, *Order) (*Product, error)
	StreamProduct(*OrderArray, Transfer_StreamProductServer) error
	StreamOrder(Transfer_StreamOrderServer) error
	mustEmbedUnimplementedTransferServer()
}

// UnimplementedTransferServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (UnimplementedTransferServer) GetProduct(context.Context, *Order) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedTransferServer) StreamProduct(*OrderArray, Transfer_StreamProductServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamProduct not implemented")
}
func (UnimplementedTransferServer) StreamOrder(Transfer_StreamOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrder not implemented")
}
func (UnimplementedTransferServer) mustEmbedUnimplementedTransferServer() {}

// UnsafeTransferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServer will
// result in compilation errors.
type UnsafeTransferServer interface {
	mustEmbedUnimplementedTransferServer()
}

func RegisterTransferServer(s grpc.ServiceRegistrar, srv TransferServer) {
	s.RegisterService(&Transfer_ServiceDesc, srv)
}

func _Transfer_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Transfer/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).GetProduct(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfer_StreamProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderArray)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransferServer).StreamProduct(m, &transferStreamProductServer{stream})
}

type Transfer_StreamProductServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type transferStreamProductServer struct {
	grpc.ServerStream
}

func (x *transferStreamProductServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _Transfer_StreamOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServer).StreamOrder(&transferStreamOrderServer{stream})
}

type Transfer_StreamOrderServer interface {
	SendAndClose(*Product) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type transferStreamOrderServer struct {
	grpc.ServerStream
}

func (x *transferStreamOrderServer) SendAndClose(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferStreamOrderServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Transfer_ServiceDesc is the grpc.ServiceDesc for Transfer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transfer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Transfer_GetProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamProduct",
			Handler:       _Transfer_StreamProduct_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamOrder",
			Handler:       _Transfer_StreamOrder_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/shop.proto",
}
