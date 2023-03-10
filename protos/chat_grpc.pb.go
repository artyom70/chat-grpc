// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: protos/chat.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Chat_ConnectClient, error)
	JoinGroupChat(ctx context.Context, in *JoinGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LeftGroupChat(ctx context.Context, in *LeftGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Channels, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Chat_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chatserver.Chat/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ConnectClient interface {
	Recv() (*ReplayMessage, error)
	grpc.ClientStream
}

type chatConnectClient struct {
	grpc.ClientStream
}

func (x *chatConnectClient) Recv() (*ReplayMessage, error) {
	m := new(ReplayMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) JoinGroupChat(ctx context.Context, in *JoinGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatserver.Chat/JoinGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) LeftGroupChat(ctx context.Context, in *LeftGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatserver.Chat/LeftGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatserver.Chat/CreateGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chatserver.Chat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ListChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Channels, error) {
	out := new(Channels)
	err := c.cc.Invoke(ctx, "/chatserver.Chat/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Connect(*ConnectRequest, Chat_ConnectServer) error
	JoinGroupChat(context.Context, *JoinGroupChatRequest) (*emptypb.Empty, error)
	LeftGroupChat(context.Context, *LeftGroupChatRequest) (*emptypb.Empty, error)
	CreateGroupChat(context.Context, *CreateGroupChatRequest) (*emptypb.Empty, error)
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
	ListChannels(context.Context, *emptypb.Empty) (*Channels, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Connect(*ConnectRequest, Chat_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatServer) JoinGroupChat(context.Context, *JoinGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroupChat not implemented")
}
func (UnimplementedChatServer) LeftGroupChat(context.Context, *LeftGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeftGroupChat not implemented")
}
func (UnimplementedChatServer) CreateGroupChat(context.Context, *CreateGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChat not implemented")
}
func (UnimplementedChatServer) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) ListChannels(context.Context, *emptypb.Empty) (*Channels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).Connect(m, &chatConnectServer{stream})
}

type Chat_ConnectServer interface {
	Send(*ReplayMessage) error
	grpc.ServerStream
}

type chatConnectServer struct {
	grpc.ServerStream
}

func (x *chatConnectServer) Send(m *ReplayMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_JoinGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).JoinGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.Chat/JoinGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).JoinGroupChat(ctx, req.(*JoinGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_LeftGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeftGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).LeftGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.Chat/LeftGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).LeftGroupChat(ctx, req.(*LeftGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.Chat/CreateGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateGroupChat(ctx, req.(*CreateGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.Chat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatserver.Chat/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ListChannels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatserver.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGroupChat",
			Handler:    _Chat_JoinGroupChat_Handler,
		},
		{
			MethodName: "LeftGroupChat",
			Handler:    _Chat_LeftGroupChat_Handler,
		},
		{
			MethodName: "CreateGroupChat",
			Handler:    _Chat_CreateGroupChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _Chat_ListChannels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Chat_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/chat.proto",
}
