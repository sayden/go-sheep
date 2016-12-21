// Code generated by protoc-gen-go.
// source: state.proto
// DO NOT EDIT!

/*
Package transport is a generated protocol buffer package.

It is generated from these files:
	state.proto

It has these top-level messages:
	ProtoState
*/
package transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

// The request message containing the user's name.
type ProtoState struct {
	Uuid     string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Address  string                     `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	LastSeen *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=lastSeen" json:"lastSeen,omitempty"`
}

func (m *ProtoState) Reset()                    { *m = ProtoState{} }
func (m *ProtoState) String() string            { return proto.CompactTextString(m) }
func (*ProtoState) ProtoMessage()               {}
func (*ProtoState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoState) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ProtoState) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProtoState) GetLastSeen() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoState)(nil), "transport.ProtoState")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pinger service

type PingerClient interface {
	// Sends a greeting
	Ping(ctx context.Context, in *ProtoState, opts ...grpc.CallOption) (*ProtoState, error)
}

type pingerClient struct {
	cc *grpc.ClientConn
}

func NewPingerClient(cc *grpc.ClientConn) PingerClient {
	return &pingerClient{cc}
}

func (c *pingerClient) Ping(ctx context.Context, in *ProtoState, opts ...grpc.CallOption) (*ProtoState, error) {
	out := new(ProtoState)
	err := grpc.Invoke(ctx, "/transport.Pinger/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pinger service

type PingerServer interface {
	// Sends a greeting
	Ping(context.Context, *ProtoState) (*ProtoState, error)
}

func RegisterPingerServer(s *grpc.Server, srv PingerServer) {
	s.RegisterService(&_Pinger_serviceDesc, srv)
}

func _Pinger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.Pinger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerServer).Ping(ctx, req.(*ProtoState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pinger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.Pinger",
	HandlerType: (*PingerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pinger_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "state.proto",
}

func init() { proto.RegisterFile("state.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x6a, 0xc4, 0x20,
	0x14, 0x45, 0x9b, 0x76, 0x98, 0x76, 0xde, 0xec, 0x84, 0x82, 0x64, 0x35, 0x64, 0x95, 0x95, 0x42,
	0x0a, 0xd9, 0x74, 0xd3, 0x4f, 0x08, 0x49, 0x7f, 0x40, 0x1b, 0x63, 0x85, 0x44, 0x45, 0x9f, 0x8b,
	0xfe, 0x7d, 0x89, 0x6d, 0x92, 0xcd, 0xec, 0x8e, 0x97, 0xcb, 0xf1, 0x3e, 0xb8, 0x46, 0x14, 0xa8,
	0x98, 0x0f, 0x0e, 0x1d, 0xb9, 0x60, 0x10, 0x36, 0x7a, 0x17, 0xb0, 0x7c, 0xd7, 0x06, 0xbf, 0x93,
	0x64, 0x5f, 0x6e, 0xe1, 0xda, 0xcd, 0xc2, 0x6a, 0x9e, 0x3b, 0x32, 0x4d, 0xdc, 0xe3, 0x8f, 0x57,
	0x91, 0xa3, 0x59, 0x54, 0x44, 0xb1, 0xf8, 0x83, 0xfe, 0x3c, 0x55, 0x00, 0xe8, 0x56, 0x18, 0x56,
	0x37, 0x21, 0x70, 0x4a, 0xc9, 0x8c, 0xb4, 0xb8, 0x15, 0xf5, 0xa5, 0xcf, 0x4c, 0x28, 0x3c, 0x8b,
	0x71, 0x0c, 0x2a, 0x46, 0xfa, 0x98, 0xe3, 0xed, 0x49, 0x5a, 0x78, 0x99, 0x45, 0xc4, 0x41, 0x29,
	0x4b, 0x9f, 0x6e, 0x45, 0x7d, 0x6d, 0x4a, 0xa6, 0x9d, 0xd3, 0xf3, 0xff, 0x48, 0x99, 0x26, 0xf6,
	0xb9, 0xfd, 0xd7, 0xef, 0xdd, 0xe6, 0x03, 0xce, 0x9d, 0xb1, 0x5a, 0x05, 0xd2, 0xc2, 0x69, 0x25,
	0xf2, 0xca, 0xf6, 0x73, 0xd8, 0x31, 0xa7, 0xbc, 0x1f, 0x57, 0x0f, 0xf2, 0x9c, 0xfd, 0x6f, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x5a, 0x66, 0x24, 0x13, 0x01, 0x00, 0x00,
}
