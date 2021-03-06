// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloudprovidervsphere.proto

package cloudprovidervsphere

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

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

type ListNodesRequest struct {
	Vcenter              string   `protobuf:"bytes,1,opt,name=vcenter,proto3" json:"vcenter,omitempty"`
	Datacenter           string   `protobuf:"bytes,2,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodesRequest) Reset()         { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()    {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b637d4c33cef7514, []int{0}
}

func (m *ListNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodesRequest.Unmarshal(m, b)
}
func (m *ListNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodesRequest.Marshal(b, m, deterministic)
}
func (m *ListNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesRequest.Merge(m, src)
}
func (m *ListNodesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNodesRequest.Size(m)
}
func (m *ListNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesRequest proto.InternalMessageInfo

func (m *ListNodesRequest) GetVcenter() string {
	if m != nil {
		return m.Vcenter
	}
	return ""
}

func (m *ListNodesRequest) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

type Node struct {
	Vcenter              string   `protobuf:"bytes,1,opt,name=vcenter,proto3" json:"vcenter,omitempty"`
	Datacenter           string   `protobuf:"bytes,2,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Dnsnames             []string `protobuf:"bytes,4,rep,name=dnsnames,proto3" json:"dnsnames,omitempty"`
	Addresses            []string `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Uuid                 string   `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_b637d4c33cef7514, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetVcenter() string {
	if m != nil {
		return m.Vcenter
	}
	return ""
}

func (m *Node) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetDnsnames() []string {
	if m != nil {
		return m.Dnsnames
	}
	return nil
}

func (m *Node) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Node) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type ListNodesReply struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodesReply) Reset()         { *m = ListNodesReply{} }
func (m *ListNodesReply) String() string { return proto.CompactTextString(m) }
func (*ListNodesReply) ProtoMessage()    {}
func (*ListNodesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b637d4c33cef7514, []int{2}
}

func (m *ListNodesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodesReply.Unmarshal(m, b)
}
func (m *ListNodesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodesReply.Marshal(b, m, deterministic)
}
func (m *ListNodesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesReply.Merge(m, src)
}
func (m *ListNodesReply) XXX_Size() int {
	return xxx_messageInfo_ListNodesReply.Size(m)
}
func (m *ListNodesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesReply proto.InternalMessageInfo

func (m *ListNodesReply) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ListNodesReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b637d4c33cef7514, []int{3}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionReply struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionReply) Reset()         { *m = VersionReply{} }
func (m *VersionReply) String() string { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()    {}
func (*VersionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b637d4c33cef7514, []int{4}
}

func (m *VersionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionReply.Unmarshal(m, b)
}
func (m *VersionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionReply.Marshal(b, m, deterministic)
}
func (m *VersionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionReply.Merge(m, src)
}
func (m *VersionReply) XXX_Size() int {
	return xxx_messageInfo_VersionReply.Size(m)
}
func (m *VersionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionReply.DiscardUnknown(m)
}

var xxx_messageInfo_VersionReply proto.InternalMessageInfo

func (m *VersionReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ListNodesRequest)(nil), "cloudprovidervsphere.ListNodesRequest")
	proto.RegisterType((*Node)(nil), "cloudprovidervsphere.Node")
	proto.RegisterType((*ListNodesReply)(nil), "cloudprovidervsphere.ListNodesReply")
	proto.RegisterType((*VersionRequest)(nil), "cloudprovidervsphere.VersionRequest")
	proto.RegisterType((*VersionReply)(nil), "cloudprovidervsphere.VersionReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CloudProviderVsphereClient is the client API for CloudProviderVsphere service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudProviderVsphereClient interface {
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesReply, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
}

type cloudProviderVsphereClient struct {
	cc *grpc.ClientConn
}

func NewCloudProviderVsphereClient(cc *grpc.ClientConn) CloudProviderVsphereClient {
	return &cloudProviderVsphereClient{cc}
}

func (c *cloudProviderVsphereClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesReply, error) {
	out := new(ListNodesReply)
	err := c.cc.Invoke(ctx, "/cloudprovidervsphere.CloudProviderVsphere/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderVsphereClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := c.cc.Invoke(ctx, "/cloudprovidervsphere.CloudProviderVsphere/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudProviderVsphereServer is the server API for CloudProviderVsphere service.
type CloudProviderVsphereServer interface {
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesReply, error)
	GetVersion(context.Context, *VersionRequest) (*VersionReply, error)
}

func RegisterCloudProviderVsphereServer(s *grpc.Server, srv CloudProviderVsphereServer) {
	s.RegisterService(&_CloudProviderVsphere_serviceDesc, srv)
}

func _CloudProviderVsphere_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderVsphereServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprovidervsphere.CloudProviderVsphere/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderVsphereServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderVsphere_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderVsphereServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprovidervsphere.CloudProviderVsphere/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderVsphereServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudProviderVsphere_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprovidervsphere.CloudProviderVsphere",
	HandlerType: (*CloudProviderVsphereServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNodes",
			Handler:    _CloudProviderVsphere_ListNodes_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _CloudProviderVsphere_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloudprovidervsphere.proto",
}

func init() { proto.RegisterFile("cloudprovidervsphere.proto", fileDescriptor_b637d4c33cef7514) }

var fileDescriptor_b637d4c33cef7514 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0xf6, 0x43, 0x33, 0x4a, 0x29, 0x43, 0x0f, 0x4b, 0x10, 0x09, 0x8b, 0x48, 0x4e, 0x41,
	0xea, 0xcd, 0x63, 0x3d, 0x78, 0x29, 0x52, 0x72, 0x28, 0x05, 0x4f, 0xb1, 0x3b, 0x60, 0xa0, 0xc9,
	0xc6, 0xdd, 0x24, 0xd2, 0xbf, 0xe3, 0xff, 0xf1, 0x3f, 0xc9, 0x66, 0xd3, 0xd8, 0x4a, 0xa4, 0xe0,
	0x6d, 0xde, 0xbc, 0x99, 0x37, 0x33, 0x6f, 0x17, 0xbc, 0xf5, 0x46, 0x96, 0x22, 0x57, 0xb2, 0x4a,
	0x04, 0xa9, 0x4a, 0xe7, 0x6f, 0xa4, 0x28, 0xcc, 0x95, 0x2c, 0x24, 0x4e, 0xba, 0x38, 0x3e, 0x87,
	0xf1, 0x3c, 0xd1, 0xc5, 0xb3, 0x14, 0xa4, 0x23, 0x7a, 0x2f, 0x49, 0x17, 0xc8, 0xe0, 0xac, 0x5a,
	0x53, 0x56, 0x90, 0x62, 0x8e, 0xef, 0x04, 0x6e, 0xb4, 0x83, 0x78, 0x0d, 0x20, 0xe2, 0x22, 0x6e,
	0xc8, 0xd3, 0x9a, 0xdc, 0xcb, 0xf0, 0x4f, 0x07, 0xfa, 0x46, 0xea, 0xff, 0x12, 0x88, 0xd0, 0xcf,
	0xe2, 0x94, 0x58, 0xaf, 0x66, 0xea, 0x18, 0x3d, 0x38, 0x17, 0x99, 0x36, 0xa1, 0x66, 0x7d, 0xbf,
	0x17, 0xb8, 0x51, 0x8b, 0xf1, 0x0a, 0xdc, 0x58, 0x08, 0x45, 0x5a, 0x93, 0x66, 0x83, 0x9a, 0xfc,
	0x49, 0x18, 0xb5, 0xb2, 0x4c, 0x04, 0x1b, 0x5a, 0x35, 0x13, 0xf3, 0x15, 0x8c, 0xf6, 0x4e, 0xce,
	0x37, 0x5b, 0xbc, 0x83, 0x41, 0x66, 0x10, 0x73, 0xfc, 0x5e, 0x70, 0x31, 0xf5, 0xc2, 0x4e, 0x1b,
	0x4d, 0x43, 0x64, 0x0b, 0x71, 0x02, 0x03, 0x52, 0x4a, 0xee, 0x0e, 0xb0, 0x80, 0x8f, 0x61, 0xb4,
	0x24, 0xa5, 0x13, 0x99, 0x35, 0x56, 0xf2, 0x00, 0x2e, 0xdb, 0x8c, 0x99, 0x64, 0x7c, 0xb1, 0xb8,
	0xf5, 0xc5, 0xc2, 0xe9, 0x97, 0x03, 0x93, 0x47, 0x33, 0x76, 0xd1, 0x8c, 0x5d, 0xda, 0xb1, 0xf8,
	0x02, 0x6e, 0xbb, 0x2e, 0xde, 0x76, 0xaf, 0xf6, 0xfb, 0x09, 0xbd, 0x9b, 0xa3, 0x75, 0xf9, 0x66,
	0xcb, 0x4f, 0x70, 0x05, 0xf0, 0x44, 0x45, 0xb3, 0x22, 0xfe, 0xd1, 0x75, 0x78, 0x93, 0xc7, 0x8f,
	0x54, 0xd5, 0xca, 0xb3, 0x07, 0xf0, 0xd7, 0x32, 0x0d, 0xab, 0xf4, 0x23, 0x56, 0x74, 0xd8, 0x11,
	0x36, 0x2d, 0xb3, 0xce, 0x83, 0x17, 0xce, 0xeb, 0xb0, 0xfe, 0xb1, 0xf7, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0x70, 0xf0, 0x60, 0xcf, 0x02, 0x00, 0x00,
}
