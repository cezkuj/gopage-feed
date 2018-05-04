// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gopage.proto

package gopage

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

type Number struct {
	Value                int32    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Number) Reset()         { *m = Number{} }
func (m *Number) String() string { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()    {}
func (*Number) Descriptor() ([]byte, []int) {
	return fileDescriptor_gopage_d56a91834654a8e8, []int{0}
}
func (m *Number) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number.Unmarshal(m, b)
}
func (m *Number) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number.Marshal(b, m, deterministic)
}
func (dst *Number) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number.Merge(dst, src)
}
func (m *Number) XXX_Size() int {
	return xxx_messageInfo_Number.Size(m)
}
func (m *Number) XXX_DiscardUnknown() {
	xxx_messageInfo_Number.DiscardUnknown(m)
}

var xxx_messageInfo_Number proto.InternalMessageInfo

func (m *Number) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParams) Reset()         { *m = GetParams{} }
func (m *GetParams) String() string { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()    {}
func (*GetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_gopage_d56a91834654a8e8, []int{1}
}
func (m *GetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParams.Unmarshal(m, b)
}
func (m *GetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParams.Marshal(b, m, deterministic)
}
func (dst *GetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParams.Merge(dst, src)
}
func (m *GetParams) XXX_Size() int {
	return xxx_messageInfo_GetParams.Size(m)
}
func (m *GetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Number)(nil), "gopage.Number")
	proto.RegisterType((*GetParams)(nil), "gopage.GetParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gopage service

type GopageClient interface {
	Get1(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Number, error)
	Get2(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Number, error)
}

type gopageClient struct {
	cc *grpc.ClientConn
}

func NewGopageClient(cc *grpc.ClientConn) GopageClient {
	return &gopageClient{cc}
}

func (c *gopageClient) Get1(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := grpc.Invoke(ctx, "/gopage.Gopage/get1", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gopageClient) Get2(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := grpc.Invoke(ctx, "/gopage.Gopage/get2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gopage service

type GopageServer interface {
	Get1(context.Context, *GetParams) (*Number, error)
	Get2(context.Context, *GetParams) (*Number, error)
}

func RegisterGopageServer(s *grpc.Server, srv GopageServer) {
	s.RegisterService(&_Gopage_serviceDesc, srv)
}

func _Gopage_Get1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopageServer).Get1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopage.Gopage/Get1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopageServer).Get1(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gopage_Get2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopageServer).Get2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopage.Gopage/Get2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopageServer).Get2(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gopage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gopage.Gopage",
	HandlerType: (*GopageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get1",
			Handler:    _Gopage_Get1_Handler,
		},
		{
			MethodName: "get2",
			Handler:    _Gopage_Get2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gopage.proto",
}

func init() { proto.RegisterFile("gopage.proto", fileDescriptor_gopage_d56a91834654a8e8) }

var fileDescriptor_gopage_d56a91834654a8e8 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcf, 0x2f, 0x48,
	0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xe4, 0xb8, 0xd8,
	0xfc, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0x25, 0x6e, 0x2e, 0x4e, 0xf7, 0xd4, 0x92, 0x80,
	0xc4, 0xa2, 0xc4, 0xdc, 0x62, 0xa3, 0x24, 0x2e, 0x36, 0x77, 0xb0, 0x36, 0x21, 0x6d, 0x2e, 0x96,
	0xf4, 0xd4, 0x12, 0x43, 0x21, 0x41, 0x3d, 0xa8, 0xa9, 0x70, 0x45, 0x52, 0x7c, 0x30, 0x21, 0x88,
	0xb9, 0x4a, 0x0c, 0x50, 0xc5, 0x46, 0x44, 0x29, 0x4e, 0x62, 0x03, 0xbb, 0xcf, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x7b, 0xa2, 0xe1, 0xaf, 0x00, 0x00, 0x00,
}