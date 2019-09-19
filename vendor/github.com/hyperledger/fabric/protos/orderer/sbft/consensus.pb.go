// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/sbft/consensus.proto

package sbft // import "github.com/hyperledger/fabric/protos/orderer/sbft"

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

type Handshake struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_7e6bb0f94b2c8614, []int{0}
}
func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (dst *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(dst, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Handshake)(nil), "backend.handshake")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsensusClient is the client API for Consensus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusClient interface {
	Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error)
}

type consensusClient struct {
	cc *grpc.ClientConn
}

func NewConsensusClient(cc *grpc.ClientConn) ConsensusClient {
	return &consensusClient{cc}
}

func (c *consensusClient) Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Consensus_serviceDesc.Streams[0], "/backend.consensus/consensus", opts...)
	if err != nil {
		return nil, err
	}
	x := &consensusConsensusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Consensus_ConsensusClient interface {
	Recv() (*MultiChainMsg, error)
	grpc.ClientStream
}

type consensusConsensusClient struct {
	grpc.ClientStream
}

func (x *consensusConsensusClient) Recv() (*MultiChainMsg, error) {
	m := new(MultiChainMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsensusServer is the server API for Consensus service.
type ConsensusServer interface {
	Consensus(*Handshake, Consensus_ConsensusServer) error
}

func RegisterConsensusServer(s *grpc.Server, srv ConsensusServer) {
	s.RegisterService(&_Consensus_serviceDesc, srv)
}

func _Consensus_Consensus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Handshake)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsensusServer).Consensus(m, &consensusConsensusServer{stream})
}

type Consensus_ConsensusServer interface {
	Send(*MultiChainMsg) error
	grpc.ServerStream
}

type consensusConsensusServer struct {
	grpc.ServerStream
}

func (x *consensusConsensusServer) Send(m *MultiChainMsg) error {
	return x.ServerStream.SendMsg(m)
}

var _Consensus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.consensus",
	HandlerType: (*ConsensusServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "consensus",
			Handler:       _Consensus_Consensus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderer/sbft/consensus.proto",
}

func init() {
	proto.RegisterFile("orderer/sbft/consensus.proto", fileDescriptor_consensus_7e6bb0f94b2c8614)
}

var fileDescriptor_consensus_7e6bb0f94b2c8614 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xb1, 0xce, 0x82, 0x30,
	0x14, 0x85, 0xff, 0x7f, 0xd1, 0x80, 0x1b, 0x93, 0x21, 0x4e, 0x3c, 0x40, 0xab, 0x32, 0xbb, 0xe8,
	0x66, 0xc2, 0x0b, 0xb8, 0xf5, 0xb6, 0x17, 0xda, 0x00, 0x2d, 0xb9, 0xb7, 0x1d, 0x7c, 0x7b, 0x13,
	0x34, 0x88, 0xe3, 0x39, 0xf9, 0xf2, 0x9d, 0x93, 0x1f, 0x02, 0x19, 0x24, 0x24, 0xc9, 0xd0, 0x46,
	0xa9, 0x83, 0x67, 0xf4, 0x9c, 0x58, 0x4c, 0x14, 0x62, 0x28, 0xb6, 0xa0, 0x74, 0x8f, 0xde, 0x94,
	0xbf, 0x18, 0xbb, 0x71, 0x1a, 0x10, 0xda, 0xf8, 0xc6, 0xaa, 0x5d, 0x9e, 0x59, 0xe5, 0x0d, 0x5b,
	0xd5, 0xe3, 0xf9, 0x9e, 0x67, 0x8b, 0xa6, 0xb8, 0xac, 0x43, 0x21, 0x3e, 0x3a, 0xb1, 0xd0, 0xe5,
	0x5e, 0x7c, 0x65, 0x4d, 0x1a, 0xa2, 0xbb, 0x59, 0xe5, 0x7c, 0xc3, 0x5d, 0xf5, 0x77, 0xfc, 0xbf,
	0xd6, 0x8f, 0x53, 0xe7, 0xa2, 0x4d, 0x20, 0x74, 0x18, 0xa5, 0x7d, 0x4e, 0x48, 0x03, 0x9a, 0x0e,
	0x49, 0xb6, 0x0a, 0xc8, 0x69, 0x39, 0xef, 0xb3, 0x5c, 0xbf, 0x83, 0xcd, 0x5c, 0xd6, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0xca, 0x55, 0x10, 0xdb, 0x00, 0x00, 0x00,
}