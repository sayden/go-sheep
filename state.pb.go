// Code generated by protoc-gen-go.
// source: state.proto
// DO NOT EDIT!

/*
Package go_sheep is a generated protocol buffer package.

It is generated from these files:
	state.proto

It has these top-level messages:
	Node
	State
*/
package go_sheep

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
type Node struct {
	Uuid     string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Address  string                     `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	LastSeen *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=lastSeen" json:"lastSeen,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Node) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetLastSeen() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

type State struct {
	Nodes []*Node `protobuf:"bytes,4,rep,name=Nodes,json=nodes" json:"Nodes,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *State) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "go_sheep.Node")
	proto.RegisterType((*State)(nil), "go_sheep.State")
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
	Ping(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
}

type pingerClient struct {
	cc *grpc.ClientConn
}

func NewPingerClient(cc *grpc.ClientConn) PingerClient {
	return &pingerClient{cc}
}

func (c *pingerClient) Ping(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := grpc.Invoke(ctx, "/go_sheep.Pinger/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pinger service

type PingerServer interface {
	// Sends a greeting
	Ping(context.Context, *State) (*State, error)
}

func RegisterPingerServer(s *grpc.Server, srv PingerServer) {
	s.RegisterService(&_Pinger_serviceDesc, srv)
}

func _Pinger_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_sheep.Pinger/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingerServer).Ping(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pinger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go_sheep.Pinger",
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
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x4f, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0xb5, 0x6e, 0xbb, 0xae, 0x53, 0x50, 0xc8, 0x29, 0xf4, 0x54, 0x8a, 0x87, 0x22, 0x98, 0x42,
	0x15, 0x2f, 0xfe, 0x07, 0x91, 0xae, 0x77, 0x49, 0xed, 0x6c, 0xb6, 0xd0, 0x26, 0xa1, 0x99, 0x1c,
	0xfc, 0xf7, 0x92, 0x68, 0x2d, 0x78, 0x7b, 0xf3, 0xe6, 0xf1, 0x3e, 0x20, 0x77, 0x24, 0x09, 0x85,
	0x5d, 0x0c, 0x19, 0x76, 0x50, 0xe6, 0xc3, 0x9d, 0x11, 0x6d, 0xf1, 0xa2, 0x46, 0x3a, 0xfb, 0x5e,
	0x7c, 0x9a, 0xb9, 0x51, 0x66, 0x92, 0x5a, 0x35, 0x51, 0xd2, 0xfb, 0x53, 0x63, 0xe9, 0xcb, 0xa2,
	0x6b, 0x68, 0x9c, 0xd1, 0x91, 0x9c, 0xed, 0x86, 0x7e, 0x6c, 0xaa, 0x09, 0xd2, 0x57, 0x33, 0x20,
	0x63, 0x90, 0x7a, 0x3f, 0x0e, 0x3c, 0x29, 0x93, 0xfa, 0xba, 0x8b, 0x98, 0x71, 0xb8, 0x92, 0xc3,
	0xb0, 0xa0, 0x73, 0xfc, 0x32, 0xd2, 0xeb, 0xc9, 0x9e, 0xe1, 0x30, 0x49, 0x47, 0x47, 0x44, 0xcd,
	0x77, 0x65, 0x52, 0xe7, 0x6d, 0x21, 0x94, 0x31, 0x6a, 0xfa, 0x6d, 0xd7, 0xfb, 0x93, 0x78, 0x5f,
	0x93, 0xba, 0x3f, 0x6d, 0xf5, 0x00, 0xd9, 0x31, 0x6c, 0x60, 0x77, 0x90, 0x85, 0x58, 0xc7, 0xd3,
	0x72, 0x57, 0xe7, 0xed, 0x8d, 0x58, 0xd7, 0x88, 0x40, 0x77, 0x99, 0x0e, 0xcf, 0xf6, 0x09, 0xf6,
	0x6f, 0xa3, 0x56, 0xb8, 0xb0, 0x7b, 0x48, 0x03, 0x62, 0xb7, 0x9b, 0x30, 0x1a, 0x15, 0xff, 0x89,
	0xea, 0xa2, 0xdf, 0xc7, 0x0a, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x12, 0xdb, 0x58,
	0x2f, 0x01, 0x00, 0x00,
}
