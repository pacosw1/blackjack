// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Game_ConnectClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Game_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Game_serviceDesc.Streams[0], "/proto.Game/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameConnectClient{stream}
	return x, nil
}

type Game_ConnectClient interface {
	Send(*Action) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type gameConnectClient struct {
	grpc.ClientStream
}

func (x *gameConnectClient) Send(m *Action) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameConnectClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	Connect(Game_ConnectServer) error
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) Connect(Game_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Connect(&gameConnectServer{stream})
}

type Game_ConnectServer interface {
	Send(*Event) error
	Recv() (*Action, error)
	grpc.ServerStream
}

type gameConnectServer struct {
	grpc.ServerStream
}

func (x *gameConnectServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameConnectServer) Recv() (*Action, error) {
	m := new(Action)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Game_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/contract.proto",
}
