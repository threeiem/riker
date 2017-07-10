// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bot.proto

/*
Package botpb is a generated protocol buffer package.

It is generated from these files:
	bot.proto

It has these top-level messages:
	Error
	ChatMessage
	Command
	CommandAuth
*/
package botpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ChatMessage struct {
	Channel string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	Message string `protobuf:"bytes,50,opt,name=message" json:"message,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChatMessage) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Command struct {
	Name        string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Usage       string       `protobuf:"bytes,2,opt,name=usage" json:"usage,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Auth        *CommandAuth `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *Command) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Command) GetAuth() *CommandAuth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type CommandAuth struct {
	Users  []string `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Groups []string `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty"`
}

func (m *CommandAuth) Reset()                    { *m = CommandAuth{} }
func (m *CommandAuth) String() string            { return proto.CompactTextString(m) }
func (*CommandAuth) ProtoMessage()               {}
func (*CommandAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CommandAuth) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *CommandAuth) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "botpb.Error")
	proto.RegisterType((*ChatMessage)(nil), "botpb.ChatMessage")
	proto.RegisterType((*Command)(nil), "botpb.Command")
	proto.RegisterType((*CommandAuth)(nil), "botpb.CommandAuth")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RedShirt service

type RedShirtClient interface {
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (RedShirt_ClientStreamClient, error)
	RegisterCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (RedShirt_RegisterCommandClient, error)
}

type redShirtClient struct {
	cc *grpc.ClientConn
}

func NewRedShirtClient(cc *grpc.ClientConn) RedShirtClient {
	return &redShirtClient{cc}
}

func (c *redShirtClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (RedShirt_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RedShirt_serviceDesc.Streams[0], c.cc, "/botpb.RedShirt/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &redShirtClientStreamClient{stream}
	return x, nil
}

type RedShirt_ClientStreamClient interface {
	Send(*ChatMessage) error
	CloseAndRecv() (*Error, error)
	grpc.ClientStream
}

type redShirtClientStreamClient struct {
	grpc.ClientStream
}

func (x *redShirtClientStreamClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *redShirtClientStreamClient) CloseAndRecv() (*Error, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Error)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *redShirtClient) RegisterCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (RedShirt_RegisterCommandClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RedShirt_serviceDesc.Streams[1], c.cc, "/botpb.RedShirt/RegisterCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &redShirtRegisterCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RedShirt_RegisterCommandClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type redShirtRegisterCommandClient struct {
	grpc.ClientStream
}

func (x *redShirtRegisterCommandClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RedShirt service

type RedShirtServer interface {
	ClientStream(RedShirt_ClientStreamServer) error
	RegisterCommand(*Command, RedShirt_RegisterCommandServer) error
}

func RegisterRedShirtServer(s *grpc.Server, srv RedShirtServer) {
	s.RegisterService(&_RedShirt_serviceDesc, srv)
}

func _RedShirt_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RedShirtServer).ClientStream(&redShirtClientStreamServer{stream})
}

type RedShirt_ClientStreamServer interface {
	SendAndClose(*Error) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type redShirtClientStreamServer struct {
	grpc.ServerStream
}

func (x *redShirtClientStreamServer) SendAndClose(m *Error) error {
	return x.ServerStream.SendMsg(m)
}

func (x *redShirtClientStreamServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RedShirt_RegisterCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Command)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RedShirtServer).RegisterCommand(m, &redShirtRegisterCommandServer{stream})
}

type RedShirt_RegisterCommandServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type redShirtRegisterCommandServer struct {
	grpc.ServerStream
}

func (x *redShirtRegisterCommandServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _RedShirt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "botpb.RedShirt",
	HandlerType: (*RedShirtServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _RedShirt_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RegisterCommand",
			Handler:       _RedShirt_RegisterCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bot.proto",
}

func init() { proto.RegisterFile("bot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x6a, 0xeb, 0x30,
	0x10, 0x85, 0xa3, 0xfc, 0xde, 0x8c, 0xc3, 0x2d, 0x0c, 0xa5, 0x88, 0x40, 0xc1, 0x68, 0x51, 0xb2,
	0x0a, 0xc5, 0xed, 0xa6, 0x74, 0x15, 0x4c, 0x97, 0xdd, 0x38, 0x4f, 0x20, 0xc7, 0x83, 0x6d, 0x88,
	0x25, 0x23, 0xc9, 0xab, 0xd2, 0x77, 0x2f, 0x96, 0x94, 0x36, 0x85, 0xee, 0xe6, 0x9c, 0x33, 0xa3,
	0x4f, 0x1a, 0xc1, 0xba, 0xd4, 0x6e, 0xdf, 0x1b, 0xed, 0x34, 0x2e, 0x4a, 0xed, 0xfa, 0x52, 0xdc,
	0xc3, 0xe2, 0xcd, 0x18, 0x6d, 0xf0, 0x16, 0x16, 0x34, 0x16, 0x9c, 0xa5, 0x6c, 0xb7, 0x2e, 0x82,
	0x10, 0x07, 0x48, 0xf2, 0x46, 0xba, 0x77, 0xb2, 0x56, 0xd6, 0x84, 0x1c, 0x56, 0xa7, 0x46, 0x2a,
	0x45, 0xe7, 0xd8, 0x76, 0x91, 0x63, 0xd2, 0x85, 0x26, 0x9e, 0x85, 0x24, 0x4a, 0xf1, 0x09, 0xab,
	0x5c, 0x77, 0x9d, 0x54, 0x15, 0x22, 0xcc, 0x95, 0xec, 0x28, 0xce, 0xfa, 0x7a, 0xe4, 0x0e, 0x7e,
	0x6c, 0x1a, 0xb8, 0x5e, 0x60, 0x0a, 0x49, 0x45, 0xf6, 0x64, 0xda, 0xde, 0xb5, 0x5a, 0xf1, 0x99,
	0xcf, 0xae, 0x2d, 0x7c, 0x80, 0xb9, 0x1c, 0x5c, 0xc3, 0xe7, 0x29, 0xdb, 0x25, 0x19, 0xee, 0xfd,
	0x73, 0xf6, 0x91, 0x74, 0x18, 0x5c, 0x53, 0xf8, 0x5c, 0xbc, 0x42, 0x72, 0x65, 0x06, 0x1c, 0x19,
	0xcb, 0x59, 0x3a, 0x0b, 0x38, 0x32, 0x16, 0xef, 0x60, 0x59, 0x1b, 0x3d, 0xf4, 0x96, 0x4f, 0xbd,
	0x1d, 0x55, 0xf6, 0x01, 0xff, 0x0a, 0xaa, 0x8e, 0x4d, 0x6b, 0x1c, 0x3e, 0xc3, 0x26, 0x3f, 0xb7,
	0xa4, 0xdc, 0xd1, 0x19, 0x92, 0x1d, 0x7e, 0x23, 0x7f, 0xf6, 0xb3, 0xdd, 0x44, 0xcf, 0xaf, 0x54,
	0x4c, 0x76, 0x0c, 0x5f, 0xe0, 0xa6, 0xa0, 0xba, 0xb5, 0x8e, 0xcc, 0x65, 0x0b, 0xff, 0x7f, 0xdf,
	0x75, 0xfb, 0xc7, 0x41, 0x62, 0xf2, 0xc8, 0xca, 0xa5, 0xff, 0xa8, 0xa7, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0xef, 0xef, 0xce, 0xb5, 0x01, 0x00, 0x00,
}