// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reverse.proto

package msgreverse

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ReplyInteger struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyInteger) Reset()         { *m = ReplyInteger{} }
func (m *ReplyInteger) String() string { return proto.CompactTextString(m) }
func (*ReplyInteger) ProtoMessage()    {}
func (*ReplyInteger) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{0}
}

func (m *ReplyInteger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInteger.Unmarshal(m, b)
}
func (m *ReplyInteger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInteger.Marshal(b, m, deterministic)
}
func (m *ReplyInteger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInteger.Merge(m, src)
}
func (m *ReplyInteger) XXX_Size() int {
	return xxx_messageInfo_ReplyInteger.Size(m)
}
func (m *ReplyInteger) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInteger.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInteger proto.InternalMessageInfo

func (m *ReplyInteger) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type RequestInteger struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInteger) Reset()         { *m = RequestInteger{} }
func (m *RequestInteger) String() string { return proto.CompactTextString(m) }
func (*RequestInteger) ProtoMessage()    {}
func (*RequestInteger) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{1}
}

func (m *RequestInteger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInteger.Unmarshal(m, b)
}
func (m *RequestInteger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInteger.Marshal(b, m, deterministic)
}
func (m *RequestInteger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInteger.Merge(m, src)
}
func (m *RequestInteger) XXX_Size() int {
	return xxx_messageInfo_RequestInteger.Size(m)
}
func (m *RequestInteger) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInteger.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInteger proto.InternalMessageInfo

func (m *RequestInteger) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Input struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{2}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type InputReply struct {
	RetMsg               string   `protobuf:"bytes,1,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputReply) Reset()         { *m = InputReply{} }
func (m *InputReply) String() string { return proto.CompactTextString(m) }
func (*InputReply) ProtoMessage()    {}
func (*InputReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_142d12653d52cef8, []int{3}
}

func (m *InputReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputReply.Unmarshal(m, b)
}
func (m *InputReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputReply.Marshal(b, m, deterministic)
}
func (m *InputReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputReply.Merge(m, src)
}
func (m *InputReply) XXX_Size() int {
	return xxx_messageInfo_InputReply.Size(m)
}
func (m *InputReply) XXX_DiscardUnknown() {
	xxx_messageInfo_InputReply.DiscardUnknown(m)
}

var xxx_messageInfo_InputReply proto.InternalMessageInfo

func (m *InputReply) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReplyInteger)(nil), "msgreverse.replyInteger")
	proto.RegisterType((*RequestInteger)(nil), "msgreverse.requestInteger")
	proto.RegisterType((*Input)(nil), "msgreverse.input")
	proto.RegisterType((*InputReply)(nil), "msgreverse.input_reply")
}

func init() { proto.RegisterFile("reverse.proto", fileDescriptor_142d12653d52cef8) }

var fileDescriptor_142d12653d52cef8 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x4b,
	0x2d, 0x2a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2d, 0x4e, 0x87, 0x8a,
	0x28, 0xa9, 0x70, 0xf1, 0x14, 0xa5, 0x16, 0xe4, 0x54, 0x7a, 0xe6, 0x95, 0xa4, 0xa6, 0xa7, 0x16,
	0x09, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06,
	0x41, 0x38, 0x4a, 0x6a, 0x5c, 0x7c, 0x45, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0xf8, 0xd5, 0x49,
	0x72, 0xb1, 0x66, 0xe6, 0x15, 0x94, 0x96, 0x08, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x83, 0x25,
	0x39, 0x83, 0x40, 0x4c, 0x25, 0x55, 0x2e, 0x6e, 0xb0, 0x54, 0x3c, 0xd8, 0x3a, 0x21, 0x31, 0x2e,
	0xb6, 0xa2, 0xd4, 0x12, 0x5f, 0xb8, 0x1a, 0x28, 0xcf, 0xc8, 0x83, 0x8b, 0x33, 0x37, 0xb5, 0xb8,
	0x38, 0x35, 0x0f, 0x64, 0x89, 0x35, 0x17, 0x77, 0x70, 0x6a, 0x5e, 0x8a, 0x6f, 0x6a, 0x71, 0x71,
	0x62, 0x7a, 0xaa, 0x90, 0xa0, 0x1e, 0xc2, 0xe1, 0x7a, 0x60, 0xc3, 0xa4, 0xc4, 0x31, 0x84, 0x20,
	0xe6, 0x2b, 0x31, 0x18, 0x45, 0x70, 0xf1, 0x67, 0x42, 0x1c, 0x1b, 0x9f, 0x0b, 0x35, 0xc0, 0x15,
	0x62, 0x1e, 0xcc, 0x0f, 0x52, 0xc8, 0x9a, 0x51, 0xfd, 0x27, 0x25, 0x81, 0x2a, 0x87, 0x08, 0x21,
	0x25, 0x06, 0x27, 0x01, 0x27, 0x9e, 0x20, 0x88, 0x4c, 0x00, 0x28, 0x3c, 0x03, 0x18, 0x93, 0xd8,
	0xc0, 0x01, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2a, 0xe9, 0xca, 0x69, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessengerClient interface {
	SendMessage(ctx context.Context, in *Input, opts ...grpc.CallOption) (*InputReply, error)
}

type messengerClient struct {
	cc *grpc.ClientConn
}

func NewMessengerClient(cc *grpc.ClientConn) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) SendMessage(ctx context.Context, in *Input, opts ...grpc.CallOption) (*InputReply, error) {
	out := new(InputReply)
	err := c.cc.Invoke(ctx, "/msgreverse.messenger/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
type MessengerServer interface {
	SendMessage(context.Context, *Input) (*InputReply, error)
}

func RegisterMessengerServer(s *grpc.Server, srv MessengerServer) {
	s.RegisterService(&_Messenger_serviceDesc, srv)
}

func _Messenger_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgreverse.messenger/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).SendMessage(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msgreverse.messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messenger_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reverse.proto",
}

// IntegerMessageClient is the client API for IntegerMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IntegerMessageClient interface {
	SendInteger(ctx context.Context, in *RequestInteger, opts ...grpc.CallOption) (*ReplyInteger, error)
}

type integerMessageClient struct {
	cc *grpc.ClientConn
}

func NewIntegerMessageClient(cc *grpc.ClientConn) IntegerMessageClient {
	return &integerMessageClient{cc}
}

func (c *integerMessageClient) SendInteger(ctx context.Context, in *RequestInteger, opts ...grpc.CallOption) (*ReplyInteger, error) {
	out := new(ReplyInteger)
	err := c.cc.Invoke(ctx, "/msgreverse.integer_message/SendInteger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegerMessageServer is the server API for IntegerMessage service.
type IntegerMessageServer interface {
	SendInteger(context.Context, *RequestInteger) (*ReplyInteger, error)
}

func RegisterIntegerMessageServer(s *grpc.Server, srv IntegerMessageServer) {
	s.RegisterService(&_IntegerMessage_serviceDesc, srv)
}

func _IntegerMessage_SendInteger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInteger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegerMessageServer).SendInteger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgreverse.integer_message/SendInteger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegerMessageServer).SendInteger(ctx, req.(*RequestInteger))
	}
	return interceptor(ctx, in, info, handler)
}

var _IntegerMessage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msgreverse.integer_message",
	HandlerType: (*IntegerMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendInteger",
			Handler:    _IntegerMessage_SendInteger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reverse.proto",
}
