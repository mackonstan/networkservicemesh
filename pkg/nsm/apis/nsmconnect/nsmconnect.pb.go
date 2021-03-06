// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nsmconnect.proto

package nsmconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"

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

// ConnectionRequest is sent by a NSM client to build a connection.
type ConnectionRequest struct {
	// Since connection request will trigger certain actions
	// executed by NSM for a client to address idempotency, request_id
	// will be tracked.
	RequestId            string                   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkServiceName   string                   `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	LinuxNamespace       string                   `protobuf:"bytes,3,opt,name=linux_namespace,json=linuxNamespace,proto3" json:"linux_namespace,omitempty"`
	LocalMechanisms      []*common.LocalMechanism `protobuf:"bytes,4,rep,name=local_mechanisms,json=localMechanisms,proto3" json:"local_mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_1036a755e1aa2d72, []int{0}
}
func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(dst, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ConnectionRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *ConnectionRequest) GetLinuxNamespace() string {
	if m != nil {
		return m.LinuxNamespace
	}
	return ""
}

func (m *ConnectionRequest) GetLocalMechanisms() []*common.LocalMechanism {
	if m != nil {
		return m.LocalMechanisms
	}
	return nil
}

type ConnectionParameters struct {
	ConnectionParameters map[string]string `protobuf:"bytes,1,rep,name=connection_parameters,json=connectionParameters,proto3" json:"connection_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConnectionParameters) Reset()         { *m = ConnectionParameters{} }
func (m *ConnectionParameters) String() string { return proto.CompactTextString(m) }
func (*ConnectionParameters) ProtoMessage()    {}
func (*ConnectionParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_1036a755e1aa2d72, []int{1}
}
func (m *ConnectionParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionParameters.Unmarshal(m, b)
}
func (m *ConnectionParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionParameters.Marshal(b, m, deterministic)
}
func (dst *ConnectionParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionParameters.Merge(dst, src)
}
func (m *ConnectionParameters) XXX_Size() int {
	return xxx_messageInfo_ConnectionParameters.Size(m)
}
func (m *ConnectionParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionParameters proto.InternalMessageInfo

func (m *ConnectionParameters) GetConnectionParameters() map[string]string {
	if m != nil {
		return m.ConnectionParameters
	}
	return nil
}

// ConnectionReply is sent back by NSM as a reply to ConnectionRequest
// accepted true will indicate that the connection is accepted, otherwise false
// indicates that connection was refused and admission_error will provide details
// why connection was refused.
type ConnectionReply struct {
	Accepted             bool                   `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string                 `protobuf:"bytes,2,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	ConnectionParameters *ConnectionParameters  `protobuf:"bytes,3,opt,name=connection_parameters,json=connectionParameters,proto3" json:"connection_parameters,omitempty"`
	LocalMechanism       *common.LocalMechanism `protobuf:"bytes,4,opt,name=local_mechanism,json=localMechanism,proto3" json:"local_mechanism,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectionReply) Reset()         { *m = ConnectionReply{} }
func (m *ConnectionReply) String() string { return proto.CompactTextString(m) }
func (*ConnectionReply) ProtoMessage()    {}
func (*ConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_1036a755e1aa2d72, []int{2}
}
func (m *ConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionReply.Unmarshal(m, b)
}
func (m *ConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionReply.Marshal(b, m, deterministic)
}
func (dst *ConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionReply.Merge(dst, src)
}
func (m *ConnectionReply) XXX_Size() int {
	return xxx_messageInfo_ConnectionReply.Size(m)
}
func (m *ConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionReply proto.InternalMessageInfo

func (m *ConnectionReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *ConnectionReply) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

func (m *ConnectionReply) GetConnectionParameters() *ConnectionParameters {
	if m != nil {
		return m.ConnectionParameters
	}
	return nil
}

func (m *ConnectionReply) GetLocalMechanism() *common.LocalMechanism {
	if m != nil {
		return m.LocalMechanism
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionRequest)(nil), "nsmconnect.ConnectionRequest")
	proto.RegisterType((*ConnectionParameters)(nil), "nsmconnect.ConnectionParameters")
	proto.RegisterMapType((map[string]string)(nil), "nsmconnect.ConnectionParameters.ConnectionParametersEntry")
	proto.RegisterType((*ConnectionReply)(nil), "nsmconnect.ConnectionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientConnectionClient is the client API for ClientConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientConnectionClient interface {
	RequestConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionReply, error)
}

type clientConnectionClient struct {
	cc *grpc.ClientConn
}

func NewClientConnectionClient(cc *grpc.ClientConn) ClientConnectionClient {
	return &clientConnectionClient{cc}
}

func (c *clientConnectionClient) RequestConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionReply, error) {
	out := new(ConnectionReply)
	err := c.cc.Invoke(ctx, "/nsmconnect.ClientConnection/RequestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientConnectionServer is the server API for ClientConnection service.
type ClientConnectionServer interface {
	RequestConnection(context.Context, *ConnectionRequest) (*ConnectionReply, error)
}

func RegisterClientConnectionServer(s *grpc.Server, srv ClientConnectionServer) {
	s.RegisterService(&_ClientConnection_serviceDesc, srv)
}

func _ClientConnection_RequestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientConnectionServer).RequestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmconnect.ClientConnection/RequestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientConnectionServer).RequestConnection(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nsmconnect.ClientConnection",
	HandlerType: (*ClientConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestConnection",
			Handler:    _ClientConnection_RequestConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nsmconnect.proto",
}

func init() { proto.RegisterFile("nsmconnect.proto", fileDescriptor_nsmconnect_1036a755e1aa2d72) }

var fileDescriptor_nsmconnect_1036a755e1aa2d72 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xdb, 0x55, 0x76, 0xef, 0x42, 0xdb, 0x1d, 0xaa, 0xc4, 0xc8, 0x42, 0xd9, 0x17, 0xf7,
	0x29, 0x91, 0xfa, 0x22, 0xfb, 0x22, 0xb2, 0x14, 0x11, 0x74, 0x95, 0x88, 0xcf, 0x61, 0x76, 0x7a,
	0x69, 0x87, 0xce, 0x97, 0x33, 0x93, 0xd5, 0xfc, 0x4b, 0x7f, 0x89, 0xbf, 0x41, 0x92, 0x4c, 0xd3,
	0x56, 0x12, 0x7c, 0xca, 0xdc, 0x73, 0x6e, 0xce, 0x9d, 0x7b, 0x4e, 0x02, 0x53, 0xe5, 0x24, 0xd3,
	0x4a, 0x21, 0xf3, 0xa9, 0xb1, 0xda, 0x6b, 0x02, 0x7b, 0x24, 0x59, 0xae, 0xb9, 0xdf, 0x94, 0x0f,
	0x29, 0xd3, 0x32, 0x13, 0x7c, 0x4d, 0xbd, 0xce, 0x14, 0xfa, 0x9f, 0xda, 0x6e, 0x1d, 0xda, 0x47,
	0xce, 0x50, 0xa2, 0xdb, 0x64, 0x66, 0xbb, 0xce, 0x94, 0x93, 0x19, 0x35, 0xdc, 0x65, 0x4c, 0x4b,
	0xa9, 0x55, 0x78, 0xb4, 0x92, 0xd7, 0xbf, 0x23, 0xb8, 0xbc, 0x6b, 0x25, 0xb9, 0x56, 0x39, 0xfe,
	0x28, 0xd1, 0x79, 0x72, 0x05, 0x60, 0xdb, 0x63, 0xc1, 0x57, 0x71, 0x34, 0x8f, 0x6e, 0xce, 0xf3,
	0xf3, 0x80, 0x7c, 0x5c, 0x91, 0xd7, 0x30, 0x0b, 0xb3, 0x8a, 0x30, 0xac, 0x50, 0x54, 0x62, 0x7c,
	0xd2, 0x34, 0x92, 0xc0, 0x7d, 0x6b, 0xa9, 0x7b, 0x2a, 0x91, 0xbc, 0x82, 0x89, 0xe0, 0xaa, 0xfc,
	0xd5, 0xf4, 0x39, 0x43, 0x19, 0xc6, 0xa3, 0xa6, 0x79, 0xdc, 0xc0, 0xf7, 0x3b, 0x94, 0xbc, 0x87,
	0xa9, 0xd0, 0x8c, 0x8a, 0x42, 0x22, 0xdb, 0x50, 0xc5, 0x9d, 0x74, 0xf1, 0xe9, 0x7c, 0x74, 0x73,
	0xb1, 0x78, 0x9e, 0x86, 0x8b, 0x7f, 0xaa, 0xf9, 0xcf, 0x3b, 0x3a, 0x9f, 0x88, 0xa3, 0xda, 0xd5,
	0x2b, 0xcd, 0xf6, 0x2b, 0x7d, 0xa5, 0x96, 0x4a, 0xf4, 0x68, 0x1d, 0xd1, 0xf0, 0x8c, 0x75, 0x78,
	0x61, 0x3a, 0x22, 0x8e, 0x9a, 0x01, 0xb7, 0xe9, 0x81, 0xe1, 0x7d, 0x02, 0xbd, 0xe0, 0x52, 0x79,
	0x5b, 0xe5, 0x33, 0xd6, 0x43, 0x25, 0x1f, 0xe0, 0xc5, 0xe0, 0x2b, 0x64, 0x0a, 0xa3, 0x2d, 0x56,
	0xc1, 0xdc, 0xfa, 0x48, 0x66, 0xf0, 0xe4, 0x91, 0x8a, 0x72, 0xe7, 0x63, 0x5b, 0xdc, 0x9e, 0xbc,
	0x8d, 0xae, 0xff, 0x44, 0x30, 0x39, 0x4c, 0xc9, 0x88, 0x8a, 0x24, 0x70, 0x46, 0x19, 0x43, 0xe3,
	0xb1, 0x4d, 0xe8, 0x2c, 0xef, 0xea, 0xda, 0x6e, 0xba, 0x92, 0xdc, 0xb9, 0x7a, 0x51, 0xb4, 0x56,
	0xdb, 0xa0, 0x39, 0xee, 0xe0, 0x65, 0x8d, 0x92, 0xef, 0x43, 0x96, 0xd4, 0xe9, 0x5c, 0x2c, 0xe6,
	0xff, 0xb3, 0xa4, 0x7f, 0x71, 0xf2, 0x0e, 0x26, 0xff, 0xa4, 0x18, 0x9f, 0x36, 0x82, 0x43, 0x21,
	0x8e, 0x8f, 0x43, 0x5c, 0x30, 0x98, 0xde, 0x09, 0x8e, 0xca, 0xef, 0x87, 0x92, 0x2f, 0x70, 0x19,
	0xbe, 0xcf, 0x03, 0xf0, 0xaa, 0xff, 0x86, 0xa1, 0x31, 0x79, 0x39, 0x44, 0x1b, 0x51, 0x3d, 0x3c,
	0x6d, 0x7e, 0x81, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x25, 0xa0, 0x74, 0xbd, 0x69, 0x03,
	0x00, 0x00,
}
