// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0568e691a5be4f4, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0568e691a5be4f4, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("messages/service.proto", fileDescriptor_c0568e691a5be4f4) }

var fileDescriptor_c0568e691a5be4f4 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x57, 0xd2, 0xe7, 0x62, 0xf7, 0x85, 0xc8, 0x08, 0x09, 0x71, 0xb1, 0xa4, 0x15, 0xe5,
	0xe7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5,
	0xe9, 0x12, 0x4c, 0x60, 0x21, 0x10, 0x53, 0x89, 0x9d, 0x8b, 0xd5, 0x35, 0xb7, 0xa0, 0xa4, 0xd2,
	0x68, 0x11, 0x23, 0x17, 0x1f, 0x54, 0x6b, 0x30, 0xc4, 0x48, 0x21, 0x05, 0x2e, 0x2e, 0xf7, 0xd4,
	0x12, 0x98, 0x79, 0x6c, 0x7a, 0x60, 0x85, 0x52, 0x1c, 0x7a, 0x50, 0x11, 0x25, 0x06, 0x21, 0x35,
	0x2e, 0xbe, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0xa8, 0x50, 0x31, 0x36, 0x55, 0x06, 0x8c, 0x42,
	0x2a, 0x5c, 0x3c, 0xc1, 0xa9, 0x79, 0x29, 0x70, 0x55, 0x70, 0x59, 0x29, 0xa8, 0x7a, 0x25, 0x06,
	0x0d, 0x46, 0x21, 0x05, 0x2e, 0x16, 0xe7, 0x8c, 0xc4, 0x12, 0x24, 0x59, 0x24, 0x53, 0x34, 0x18,
	0x0d, 0x18, 0x93, 0xd8, 0xc0, 0xbe, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xd2, 0x2d,
	0x32, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	GetMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error)
	StreamMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MessageService_StreamMessagesClient, error)
	SendMessages(ctx context.Context, opts ...grpc.CallOption) (MessageService_SendMessagesClient, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (MessageService_ChatClient, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) GetMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/MessageService/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) StreamMessages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MessageService_StreamMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/MessageService/StreamMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceStreamMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_StreamMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageServiceStreamMessagesClient struct {
	grpc.ClientStream
}

func (x *messageServiceStreamMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) SendMessages(ctx context.Context, opts ...grpc.CallOption) (MessageService_SendMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[1], "/MessageService/SendMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceSendMessagesClient{stream}
	return x, nil
}

type MessageService_SendMessagesClient interface {
	Send(*Message) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type messageServiceSendMessagesClient struct {
	grpc.ClientStream
}

func (x *messageServiceSendMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceSendMessagesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (MessageService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[2], "/MessageService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceChatClient{stream}
	return x, nil
}

type MessageService_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageServiceChatClient struct {
	grpc.ClientStream
}

func (x *messageServiceChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	GetMessage(context.Context, *Empty) (*Message, error)
	StreamMessages(*Empty, MessageService_StreamMessagesServer) error
	SendMessages(MessageService_SendMessagesServer) error
	Chat(MessageService_ChatServer) error
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) GetMessage(ctx context.Context, req *Empty) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (*UnimplementedMessageServiceServer) StreamMessages(req *Empty, srv MessageService_StreamMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessages not implemented")
}
func (*UnimplementedMessageServiceServer) SendMessages(srv MessageService_SendMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessages not implemented")
}
func (*UnimplementedMessageServiceServer) Chat(srv MessageService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_StreamMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).StreamMessages(m, &messageServiceStreamMessagesServer{stream})
}

type MessageService_StreamMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messageServiceStreamMessagesServer struct {
	grpc.ServerStream
}

func (x *messageServiceStreamMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageService_SendMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).SendMessages(&messageServiceSendMessagesServer{stream})
}

type MessageService_SendMessagesServer interface {
	SendAndClose(*Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageServiceSendMessagesServer struct {
	grpc.ServerStream
}

func (x *messageServiceSendMessagesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceSendMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Chat(&messageServiceChatServer{stream})
}

type MessageService_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messageServiceChatServer struct {
	grpc.ServerStream
}

func (x *messageServiceChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _MessageService_GetMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessages",
			Handler:       _MessageService_StreamMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessages",
			Handler:       _MessageService_SendMessages_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _MessageService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages/service.proto",
}