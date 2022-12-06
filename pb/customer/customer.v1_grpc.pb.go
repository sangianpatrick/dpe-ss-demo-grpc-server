// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: customer.v1.proto

package customerpb

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

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	Register(ctx context.Context, in *AccountRegistrationRequest, opts ...grpc.CallOption) (*AccountRegistrationResponse, error)
	SubscribeNotification(ctx context.Context, in *SubscribeNotificationRequest, opts ...grpc.CallOption) (Customer_SubscribeNotificationClient, error)
	SumNumbers(ctx context.Context, opts ...grpc.CallOption) (Customer_SumNumbersClient, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (Customer_ChatClient, error)
	MakePayment(ctx context.Context, in *MakePaymentRequest, opts ...grpc.CallOption) (*MakePaymentResponse, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Register(ctx context.Context, in *AccountRegistrationRequest, opts ...grpc.CallOption) (*AccountRegistrationResponse, error) {
	out := new(AccountRegistrationResponse)
	err := c.cc.Invoke(ctx, "/customer.v1.Customer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) SubscribeNotification(ctx context.Context, in *SubscribeNotificationRequest, opts ...grpc.CallOption) (Customer_SubscribeNotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Customer_ServiceDesc.Streams[0], "/customer.v1.Customer/SubscribeNotification", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerSubscribeNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_SubscribeNotificationClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type customerSubscribeNotificationClient struct {
	grpc.ClientStream
}

func (x *customerSubscribeNotificationClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) SumNumbers(ctx context.Context, opts ...grpc.CallOption) (Customer_SumNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Customer_ServiceDesc.Streams[1], "/customer.v1.Customer/SumNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerSumNumbersClient{stream}
	return x, nil
}

type Customer_SumNumbersClient interface {
	Send(*SumNumbersRequest) error
	CloseAndRecv() (*SumNumbersResponse, error)
	grpc.ClientStream
}

type customerSumNumbersClient struct {
	grpc.ClientStream
}

func (x *customerSumNumbersClient) Send(m *SumNumbersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *customerSumNumbersClient) CloseAndRecv() (*SumNumbersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumNumbersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Customer_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Customer_ServiceDesc.Streams[2], "/customer.v1.Customer/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerChatClient{stream}
	return x, nil
}

type Customer_ChatClient interface {
	Send(*ChatRequest) error
	Recv() (*ChatResponse, error)
	grpc.ClientStream
}

type customerChatClient struct {
	grpc.ClientStream
}

func (x *customerChatClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *customerChatClient) Recv() (*ChatResponse, error) {
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) MakePayment(ctx context.Context, in *MakePaymentRequest, opts ...grpc.CallOption) (*MakePaymentResponse, error) {
	out := new(MakePaymentResponse)
	err := c.cc.Invoke(ctx, "/customer.v1.Customer/MakePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility
type CustomerServer interface {
	Register(context.Context, *AccountRegistrationRequest) (*AccountRegistrationResponse, error)
	SubscribeNotification(*SubscribeNotificationRequest, Customer_SubscribeNotificationServer) error
	SumNumbers(Customer_SumNumbersServer) error
	Chat(Customer_ChatServer) error
	MakePayment(context.Context, *MakePaymentRequest) (*MakePaymentResponse, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (UnimplementedCustomerServer) Register(context.Context, *AccountRegistrationRequest) (*AccountRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCustomerServer) SubscribeNotification(*SubscribeNotificationRequest, Customer_SubscribeNotificationServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeNotification not implemented")
}
func (UnimplementedCustomerServer) SumNumbers(Customer_SumNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method SumNumbers not implemented")
}
func (UnimplementedCustomerServer) Chat(Customer_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedCustomerServer) MakePayment(context.Context, *MakePaymentRequest) (*MakePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePayment not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.v1.Customer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Register(ctx, req.(*AccountRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_SubscribeNotification_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeNotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).SubscribeNotification(m, &customerSubscribeNotificationServer{stream})
}

type Customer_SubscribeNotificationServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type customerSubscribeNotificationServer struct {
	grpc.ServerStream
}

func (x *customerSubscribeNotificationServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_SumNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CustomerServer).SumNumbers(&customerSumNumbersServer{stream})
}

type Customer_SumNumbersServer interface {
	SendAndClose(*SumNumbersResponse) error
	Recv() (*SumNumbersRequest, error)
	grpc.ServerStream
}

type customerSumNumbersServer struct {
	grpc.ServerStream
}

func (x *customerSumNumbersServer) SendAndClose(m *SumNumbersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *customerSumNumbersServer) Recv() (*SumNumbersRequest, error) {
	m := new(SumNumbersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Customer_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CustomerServer).Chat(&customerChatServer{stream})
}

type Customer_ChatServer interface {
	Send(*ChatResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type customerChatServer struct {
	grpc.ServerStream
}

func (x *customerChatServer) Send(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *customerChatServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Customer_MakePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).MakePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.v1.Customer/MakePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).MakePayment(ctx, req.(*MakePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer.v1.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Customer_Register_Handler,
		},
		{
			MethodName: "MakePayment",
			Handler:    _Customer_MakePayment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeNotification",
			Handler:       _Customer_SubscribeNotification_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SumNumbers",
			Handler:       _Customer_SumNumbers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _Customer_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "customer.v1.proto",
}
