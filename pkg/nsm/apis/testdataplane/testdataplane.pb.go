// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdataplane.proto

package testdataplane

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

type NSMPodType int32

const (
	NSMPodType_DEFAULT_POD NSMPodType = 0
	NSMPodType_NSMCLIENT   NSMPodType = 1
	NSMPodType_NSE         NSMPodType = 2
)

var NSMPodType_name = map[int32]string{
	0: "DEFAULT_POD",
	1: "NSMCLIENT",
	2: "NSE",
}
var NSMPodType_value = map[string]int32{
	"DEFAULT_POD": 0,
	"NSMCLIENT":   1,
	"NSE":         2,
}

func (x NSMPodType) String() string {
	return proto.EnumName(NSMPodType_name, int32(x))
}
func (NSMPodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{0}
}

type Metadata struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Pod struct {
	Metadata             *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	IpAddress            string    `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{1}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (dst *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(dst, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pod) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type BuildConnectRequest struct {
	SourcePod            *Pod     `protobuf:"bytes,1,opt,name=source_pod,json=sourcePod,proto3" json:"source_pod,omitempty"`
	DestinationPod       *Pod     `protobuf:"bytes,2,opt,name=destination_pod,json=destinationPod,proto3" json:"destination_pod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildConnectRequest) Reset()         { *m = BuildConnectRequest{} }
func (m *BuildConnectRequest) String() string { return proto.CompactTextString(m) }
func (*BuildConnectRequest) ProtoMessage()    {}
func (*BuildConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{2}
}
func (m *BuildConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConnectRequest.Unmarshal(m, b)
}
func (m *BuildConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConnectRequest.Marshal(b, m, deterministic)
}
func (dst *BuildConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConnectRequest.Merge(dst, src)
}
func (m *BuildConnectRequest) XXX_Size() int {
	return xxx_messageInfo_BuildConnectRequest.Size(m)
}
func (m *BuildConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConnectRequest proto.InternalMessageInfo

func (m *BuildConnectRequest) GetSourcePod() *Pod {
	if m != nil {
		return m.SourcePod
	}
	return nil
}

func (m *BuildConnectRequest) GetDestinationPod() *Pod {
	if m != nil {
		return m.DestinationPod
	}
	return nil
}

type BuildConnectReply struct {
	Built                bool     `protobuf:"varint,1,opt,name=built,proto3" json:"built,omitempty"`
	BuildError           string   `protobuf:"bytes,2,opt,name=build_error,json=buildError,proto3" json:"build_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildConnectReply) Reset()         { *m = BuildConnectReply{} }
func (m *BuildConnectReply) String() string { return proto.CompactTextString(m) }
func (*BuildConnectReply) ProtoMessage()    {}
func (*BuildConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{3}
}
func (m *BuildConnectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildConnectReply.Unmarshal(m, b)
}
func (m *BuildConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildConnectReply.Marshal(b, m, deterministic)
}
func (dst *BuildConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildConnectReply.Merge(dst, src)
}
func (m *BuildConnectReply) XXX_Size() int {
	return xxx_messageInfo_BuildConnectReply.Size(m)
}
func (m *BuildConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_BuildConnectReply proto.InternalMessageInfo

func (m *BuildConnectReply) GetBuilt() bool {
	if m != nil {
		return m.Built
	}
	return false
}

func (m *BuildConnectReply) GetBuildError() string {
	if m != nil {
		return m.BuildError
	}
	return ""
}

type DeleteConnectRequest struct {
	Pod                  *Pod       `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	PodType              NSMPodType `protobuf:"varint,2,opt,name=pod_type,json=podType,proto3,enum=testdataplane.NSMPodType" json:"pod_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteConnectRequest) Reset()         { *m = DeleteConnectRequest{} }
func (m *DeleteConnectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectRequest) ProtoMessage()    {}
func (*DeleteConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{4}
}
func (m *DeleteConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectRequest.Unmarshal(m, b)
}
func (m *DeleteConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectRequest.Merge(dst, src)
}
func (m *DeleteConnectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectRequest.Size(m)
}
func (m *DeleteConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectRequest proto.InternalMessageInfo

func (m *DeleteConnectRequest) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *DeleteConnectRequest) GetPodType() NSMPodType {
	if m != nil {
		return m.PodType
	}
	return NSMPodType_DEFAULT_POD
}

type DeleteConnectReply struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	DeleteError          string   `protobuf:"bytes,2,opt,name=delete_error,json=deleteError,proto3" json:"delete_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectReply) Reset()         { *m = DeleteConnectReply{} }
func (m *DeleteConnectReply) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectReply) ProtoMessage()    {}
func (*DeleteConnectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdataplane_c25ea7476a78660e, []int{5}
}
func (m *DeleteConnectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectReply.Unmarshal(m, b)
}
func (m *DeleteConnectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectReply.Marshal(b, m, deterministic)
}
func (dst *DeleteConnectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectReply.Merge(dst, src)
}
func (m *DeleteConnectReply) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectReply.Size(m)
}
func (m *DeleteConnectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectReply proto.InternalMessageInfo

func (m *DeleteConnectReply) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *DeleteConnectReply) GetDeleteError() string {
	if m != nil {
		return m.DeleteError
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "testdataplane.Metadata")
	proto.RegisterType((*Pod)(nil), "testdataplane.Pod")
	proto.RegisterType((*BuildConnectRequest)(nil), "testdataplane.BuildConnectRequest")
	proto.RegisterType((*BuildConnectReply)(nil), "testdataplane.BuildConnectReply")
	proto.RegisterType((*DeleteConnectRequest)(nil), "testdataplane.DeleteConnectRequest")
	proto.RegisterType((*DeleteConnectReply)(nil), "testdataplane.DeleteConnectReply")
	proto.RegisterEnum("testdataplane.NSMPodType", NSMPodType_name, NSMPodType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildConnectClient is the client API for BuildConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildConnectClient interface {
	RequestBuildConnect(ctx context.Context, in *BuildConnectRequest, opts ...grpc.CallOption) (*BuildConnectReply, error)
}

type buildConnectClient struct {
	cc *grpc.ClientConn
}

func NewBuildConnectClient(cc *grpc.ClientConn) BuildConnectClient {
	return &buildConnectClient{cc}
}

func (c *buildConnectClient) RequestBuildConnect(ctx context.Context, in *BuildConnectRequest, opts ...grpc.CallOption) (*BuildConnectReply, error) {
	out := new(BuildConnectReply)
	err := c.cc.Invoke(ctx, "/testdataplane.BuildConnect/RequestBuildConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildConnectServer is the server API for BuildConnect service.
type BuildConnectServer interface {
	RequestBuildConnect(context.Context, *BuildConnectRequest) (*BuildConnectReply, error)
}

func RegisterBuildConnectServer(s *grpc.Server, srv BuildConnectServer) {
	s.RegisterService(&_BuildConnect_serviceDesc, srv)
}

func _BuildConnect_RequestBuildConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildConnectServer).RequestBuildConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdataplane.BuildConnect/RequestBuildConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildConnectServer).RequestBuildConnect(ctx, req.(*BuildConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdataplane.BuildConnect",
	HandlerType: (*BuildConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestBuildConnect",
			Handler:    _BuildConnect_RequestBuildConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdataplane.proto",
}

// DeleteConnectClient is the client API for DeleteConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeleteConnectClient interface {
	RequestDeleteConnect(ctx context.Context, in *DeleteConnectRequest, opts ...grpc.CallOption) (*DeleteConnectReply, error)
}

type deleteConnectClient struct {
	cc *grpc.ClientConn
}

func NewDeleteConnectClient(cc *grpc.ClientConn) DeleteConnectClient {
	return &deleteConnectClient{cc}
}

func (c *deleteConnectClient) RequestDeleteConnect(ctx context.Context, in *DeleteConnectRequest, opts ...grpc.CallOption) (*DeleteConnectReply, error) {
	out := new(DeleteConnectReply)
	err := c.cc.Invoke(ctx, "/testdataplane.DeleteConnect/RequestDeleteConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteConnectServer is the server API for DeleteConnect service.
type DeleteConnectServer interface {
	RequestDeleteConnect(context.Context, *DeleteConnectRequest) (*DeleteConnectReply, error)
}

func RegisterDeleteConnectServer(s *grpc.Server, srv DeleteConnectServer) {
	s.RegisterService(&_DeleteConnect_serviceDesc, srv)
}

func _DeleteConnect_RequestDeleteConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteConnectServer).RequestDeleteConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdataplane.DeleteConnect/RequestDeleteConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteConnectServer).RequestDeleteConnect(ctx, req.(*DeleteConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeleteConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdataplane.DeleteConnect",
	HandlerType: (*DeleteConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDeleteConnect",
			Handler:    _DeleteConnect_RequestDeleteConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdataplane.proto",
}

func init() { proto.RegisterFile("testdataplane.proto", fileDescriptor_testdataplane_c25ea7476a78660e) }

var fileDescriptor_testdataplane_c25ea7476a78660e = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0x71, 0x02, 0xc4, 0x1e, 0x37, 0x6d, 0x98, 0x44, 0x22, 0x54, 0x20, 0x5a, 0xc3, 0x01,
	0x71, 0xa8, 0x84, 0x0b, 0x27, 0xb8, 0x94, 0xc6, 0x48, 0xa0, 0xc6, 0x18, 0x27, 0x1c, 0x10, 0x12,
	0xd6, 0x36, 0x3b, 0x07, 0x0b, 0xc7, 0xbb, 0xd8, 0xeb, 0x83, 0x7f, 0x00, 0xff, 0x1b, 0x79, 0xed,
	0x90, 0xd8, 0x4a, 0x73, 0xca, 0xec, 0xb7, 0x6f, 0xe6, 0xed, 0xdb, 0x8d, 0x61, 0xac, 0x28, 0x57,
	0x9c, 0x29, 0x26, 0x13, 0x96, 0xd2, 0x85, 0xcc, 0x84, 0x12, 0x38, 0x6c, 0x41, 0xe7, 0x03, 0x98,
	0x73, 0x52, 0xac, 0x02, 0x88, 0x70, 0x3f, 0x65, 0x6b, 0x9a, 0x1a, 0x67, 0xc6, 0x2b, 0x2b, 0xd4,
	0x35, 0x3e, 0x05, 0xab, 0xfa, 0xcd, 0x25, 0x5b, 0xd1, 0xb4, 0xa7, 0x37, 0xb6, 0xc0, 0xf9, 0x01,
	0xfd, 0x40, 0x70, 0xbc, 0x04, 0x73, 0xdd, 0x0c, 0xd1, 0xcd, 0xb6, 0xfb, 0xf8, 0xa2, 0xed, 0xbd,
	0xf1, 0x08, 0xff, 0x0b, 0xf1, 0x19, 0x40, 0x2c, 0x23, 0xc6, 0x79, 0x46, 0x79, 0xbe, 0x19, 0x1d,
	0xcb, 0xab, 0x1a, 0x38, 0x7f, 0x0d, 0x18, 0x7f, 0x2c, 0xe2, 0x84, 0x5f, 0x8b, 0x34, 0xa5, 0x95,
	0x0a, 0xe9, 0x4f, 0x41, 0xb9, 0xc2, 0x37, 0x00, 0xb9, 0x28, 0xb2, 0x15, 0x45, 0x52, 0xf0, 0xc6,
	0x0d, 0x3b, 0x6e, 0x81, 0xe0, 0xa1, 0x55, 0xab, 0xaa, 0xe3, 0xbd, 0x87, 0x13, 0x4e, 0xb9, 0x8a,
	0x53, 0xa6, 0x62, 0x91, 0xea, 0xbe, 0xde, 0x9d, 0x7d, 0xc7, 0x3b, 0xd2, 0x40, 0x70, 0xe7, 0x0b,
	0x3c, 0x6a, 0x1f, 0x43, 0x26, 0x25, 0x4e, 0xe0, 0xc1, 0x6d, 0x11, 0x27, 0x4a, 0xfb, 0x9b, 0x61,
	0xbd, 0xc0, 0xe7, 0x60, 0x57, 0x05, 0x8f, 0x28, 0xcb, 0x44, 0xd6, 0x44, 0x02, 0x8d, 0xbc, 0x8a,
	0x38, 0x19, 0x4c, 0x66, 0x94, 0x90, 0xa2, 0x4e, 0xa6, 0x97, 0xd0, 0x3f, 0x1c, 0xa6, 0xda, 0xc6,
	0xb7, 0x60, 0x4a, 0xc1, 0x23, 0x55, 0xca, 0xfa, 0x25, 0x8e, 0xdd, 0x27, 0x1d, 0xa9, 0xbf, 0x98,
	0x07, 0x82, 0x2f, 0x4b, 0x49, 0xe1, 0x40, 0xd6, 0x85, 0xf3, 0x0d, 0xb0, 0xe3, 0x59, 0x05, 0x98,
	0xc2, 0x80, 0x6b, 0xca, 0x9b, 0x08, 0x9b, 0x25, 0x9e, 0xc3, 0x51, 0x5d, 0xb6, 0x52, 0xd8, 0x35,
	0xd3, 0x31, 0x5e, 0xbf, 0x03, 0xd8, 0x3a, 0xe1, 0x09, 0xd8, 0x33, 0xef, 0xd3, 0xd5, 0xf7, 0x9b,
	0x65, 0x14, 0x7c, 0x9d, 0x8d, 0xee, 0xe1, 0x10, 0x2c, 0x7f, 0x31, 0xbf, 0xbe, 0xf9, 0xec, 0xf9,
	0xcb, 0x91, 0x81, 0x03, 0xe8, 0xfb, 0x0b, 0x6f, 0xd4, 0x73, 0x7f, 0xc3, 0xd1, 0xee, 0x4d, 0xe2,
	0x4f, 0x18, 0x37, 0x17, 0xd0, 0xc2, 0x4e, 0x27, 0xd4, 0x9e, 0x3f, 0xc1, 0xe9, 0xd9, 0x41, 0x8d,
	0x4c, 0x4a, 0x57, 0xc0, 0xb0, 0x15, 0x1b, 0x7f, 0xc1, 0xa4, 0xe9, 0x6e, 0xf3, 0x17, 0x9d, 0x51,
	0xfb, 0x1e, 0xe8, 0xf4, 0xfc, 0xb0, 0x48, 0x26, 0xe5, 0xed, 0x43, 0xfd, 0x79, 0x5d, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0xb5, 0xa8, 0x80, 0x75, 0x03, 0x00, 0x00,
}
