// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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

type Request struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_39ab6cb263caac41, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Reply struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_39ab6cb263caac41, []int{1}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "example.Request")
	proto.RegisterType((*Reply)(nil), "example.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	DisplayLines(ctx context.Context, in *Request, opts ...grpc.CallOption) (ExampleService_DisplayLinesClient, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) DisplayLines(ctx context.Context, in *Request, opts ...grpc.CallOption) (ExampleService_DisplayLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExampleService_serviceDesc.Streams[0], "/example.ExampleService/DisplayLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceDisplayLinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_DisplayLinesClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type exampleServiceDisplayLinesClient struct {
	grpc.ClientStream
}

func (x *exampleServiceDisplayLinesClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	DisplayLines(*Request, ExampleService_DisplayLinesServer) error
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_DisplayLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).DisplayLines(m, &exampleServiceDisplayLinesServer{stream})
}

type ExampleService_DisplayLinesServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type exampleServiceDisplayLinesServer struct {
	grpc.ServerStream
}

func (x *exampleServiceDisplayLinesServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DisplayLines",
			Handler:       _ExampleService_DisplayLines_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_example_39ab6cb263caac41) }

var fileDescriptor_example_39ab6cb263caac41 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb9,
	0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0x4a, 0x73, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x40, 0x4c, 0x25, 0x69, 0x2e, 0xd6, 0xa0, 0xd4, 0x82, 0x9c,
	0x4a, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0xb0, 0x1c, 0x67, 0x10, 0x98, 0x6d, 0xe4,
	0xc6, 0xc5, 0xe7, 0x0a, 0x31, 0x24, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x84, 0x8b,
	0xc7, 0x25, 0xb3, 0xb8, 0x20, 0x27, 0xb1, 0xd2, 0x27, 0x33, 0x2f, 0xb5, 0x58, 0x48, 0x40, 0x0f,
	0x66, 0x29, 0xd4, 0x0a, 0x29, 0x3e, 0x24, 0x91, 0x82, 0x9c, 0x4a, 0x25, 0x06, 0x03, 0xc6, 0x24,
	0x36, 0xb0, 0x8b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x8f, 0x40, 0x77, 0xa2, 0x00,
	0x00, 0x00,
}