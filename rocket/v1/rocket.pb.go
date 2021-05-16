// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rocket/v1/rocket.proto

package rocket

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

// Rocket represents the model of our rocket
type Rocket struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rocket) Reset()         { *m = Rocket{} }
func (m *Rocket) String() string { return proto.CompactTextString(m) }
func (*Rocket) ProtoMessage()    {}
func (*Rocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{0}
}
func (m *Rocket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rocket.Unmarshal(m, b)
}
func (m *Rocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rocket.Marshal(b, m, deterministic)
}
func (dst *Rocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rocket.Merge(dst, src)
}
func (m *Rocket) XXX_Size() int {
	return xxx_messageInfo_Rocket.Size(m)
}
func (m *Rocket) XXX_DiscardUnknown() {
	xxx_messageInfo_Rocket.DiscardUnknown(m)
}

var xxx_messageInfo_Rocket proto.InternalMessageInfo

func (m *Rocket) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// -- Get Rocket Handler --
type GetRocketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRocketRequest) Reset()         { *m = GetRocketRequest{} }
func (m *GetRocketRequest) String() string { return proto.CompactTextString(m) }
func (*GetRocketRequest) ProtoMessage()    {}
func (*GetRocketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{1}
}
func (m *GetRocketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRocketRequest.Unmarshal(m, b)
}
func (m *GetRocketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRocketRequest.Marshal(b, m, deterministic)
}
func (dst *GetRocketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRocketRequest.Merge(dst, src)
}
func (m *GetRocketRequest) XXX_Size() int {
	return xxx_messageInfo_GetRocketRequest.Size(m)
}
func (m *GetRocketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRocketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRocketRequest proto.InternalMessageInfo

func (m *GetRocketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRocketResponse struct {
	Rocket               *Rocket  `protobuf:"bytes,1,opt,name=rocket,proto3" json:"rocket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRocketResponse) Reset()         { *m = GetRocketResponse{} }
func (m *GetRocketResponse) String() string { return proto.CompactTextString(m) }
func (*GetRocketResponse) ProtoMessage()    {}
func (*GetRocketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{2}
}
func (m *GetRocketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRocketResponse.Unmarshal(m, b)
}
func (m *GetRocketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRocketResponse.Marshal(b, m, deterministic)
}
func (dst *GetRocketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRocketResponse.Merge(dst, src)
}
func (m *GetRocketResponse) XXX_Size() int {
	return xxx_messageInfo_GetRocketResponse.Size(m)
}
func (m *GetRocketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRocketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRocketResponse proto.InternalMessageInfo

func (m *GetRocketResponse) GetRocket() *Rocket {
	if m != nil {
		return m.Rocket
	}
	return nil
}

// -- Add Rocket Handler --
type AddRocketRequest struct {
	Rocket               *Rocket  `protobuf:"bytes,1,opt,name=rocket,proto3" json:"rocket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRocketRequest) Reset()         { *m = AddRocketRequest{} }
func (m *AddRocketRequest) String() string { return proto.CompactTextString(m) }
func (*AddRocketRequest) ProtoMessage()    {}
func (*AddRocketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{3}
}
func (m *AddRocketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRocketRequest.Unmarshal(m, b)
}
func (m *AddRocketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRocketRequest.Marshal(b, m, deterministic)
}
func (dst *AddRocketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRocketRequest.Merge(dst, src)
}
func (m *AddRocketRequest) XXX_Size() int {
	return xxx_messageInfo_AddRocketRequest.Size(m)
}
func (m *AddRocketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRocketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRocketRequest proto.InternalMessageInfo

func (m *AddRocketRequest) GetRocket() *Rocket {
	if m != nil {
		return m.Rocket
	}
	return nil
}

type AddRocketResponse struct {
	Rocket               *Rocket  `protobuf:"bytes,1,opt,name=rocket,proto3" json:"rocket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRocketResponse) Reset()         { *m = AddRocketResponse{} }
func (m *AddRocketResponse) String() string { return proto.CompactTextString(m) }
func (*AddRocketResponse) ProtoMessage()    {}
func (*AddRocketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{4}
}
func (m *AddRocketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRocketResponse.Unmarshal(m, b)
}
func (m *AddRocketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRocketResponse.Marshal(b, m, deterministic)
}
func (dst *AddRocketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRocketResponse.Merge(dst, src)
}
func (m *AddRocketResponse) XXX_Size() int {
	return xxx_messageInfo_AddRocketResponse.Size(m)
}
func (m *AddRocketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRocketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRocketResponse proto.InternalMessageInfo

func (m *AddRocketResponse) GetRocket() *Rocket {
	if m != nil {
		return m.Rocket
	}
	return nil
}

// -- Add Rocket Handler --
type DeleteRocketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRocketRequest) Reset()         { *m = DeleteRocketRequest{} }
func (m *DeleteRocketRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRocketRequest) ProtoMessage()    {}
func (*DeleteRocketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{5}
}
func (m *DeleteRocketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRocketRequest.Unmarshal(m, b)
}
func (m *DeleteRocketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRocketRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRocketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRocketRequest.Merge(dst, src)
}
func (m *DeleteRocketRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRocketRequest.Size(m)
}
func (m *DeleteRocketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRocketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRocketRequest proto.InternalMessageInfo

func (m *DeleteRocketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRocketResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRocketResponse) Reset()         { *m = DeleteRocketResponse{} }
func (m *DeleteRocketResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRocketResponse) ProtoMessage()    {}
func (*DeleteRocketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocket_122f075222d9fc46, []int{6}
}
func (m *DeleteRocketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRocketResponse.Unmarshal(m, b)
}
func (m *DeleteRocketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRocketResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteRocketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRocketResponse.Merge(dst, src)
}
func (m *DeleteRocketResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRocketResponse.Size(m)
}
func (m *DeleteRocketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRocketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRocketResponse proto.InternalMessageInfo

func (m *DeleteRocketResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Rocket)(nil), "rocket.Rocket")
	proto.RegisterType((*GetRocketRequest)(nil), "rocket.GetRocketRequest")
	proto.RegisterType((*GetRocketResponse)(nil), "rocket.GetRocketResponse")
	proto.RegisterType((*AddRocketRequest)(nil), "rocket.AddRocketRequest")
	proto.RegisterType((*AddRocketResponse)(nil), "rocket.AddRocketResponse")
	proto.RegisterType((*DeleteRocketRequest)(nil), "rocket.DeleteRocketRequest")
	proto.RegisterType((*DeleteRocketResponse)(nil), "rocket.DeleteRocketResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RocketServiceClient is the client API for RocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RocketServiceClient interface {
	GetRocket(ctx context.Context, in *GetRocketRequest, opts ...grpc.CallOption) (*GetRocketResponse, error)
	AddRocket(ctx context.Context, in *AddRocketRequest, opts ...grpc.CallOption) (*AddRocketResponse, error)
	DeleteRocket(ctx context.Context, in *DeleteRocketRequest, opts ...grpc.CallOption) (*DeleteRocketResponse, error)
}

type rocketServiceClient struct {
	cc *grpc.ClientConn
}

func NewRocketServiceClient(cc *grpc.ClientConn) RocketServiceClient {
	return &rocketServiceClient{cc}
}

func (c *rocketServiceClient) GetRocket(ctx context.Context, in *GetRocketRequest, opts ...grpc.CallOption) (*GetRocketResponse, error) {
	out := new(GetRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/GetRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rocketServiceClient) AddRocket(ctx context.Context, in *AddRocketRequest, opts ...grpc.CallOption) (*AddRocketResponse, error) {
	out := new(AddRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/AddRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rocketServiceClient) DeleteRocket(ctx context.Context, in *DeleteRocketRequest, opts ...grpc.CallOption) (*DeleteRocketResponse, error) {
	out := new(DeleteRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/DeleteRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RocketServiceServer is the server API for RocketService service.
type RocketServiceServer interface {
	GetRocket(context.Context, *GetRocketRequest) (*GetRocketResponse, error)
	AddRocket(context.Context, *AddRocketRequest) (*AddRocketResponse, error)
	DeleteRocket(context.Context, *DeleteRocketRequest) (*DeleteRocketResponse, error)
}

func RegisterRocketServiceServer(s *grpc.Server, srv RocketServiceServer) {
	s.RegisterService(&_RocketService_serviceDesc, srv)
}

func _RocketService_GetRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).GetRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/GetRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).GetRocket(ctx, req.(*GetRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RocketService_AddRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).AddRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/AddRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).AddRocket(ctx, req.(*AddRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RocketService_DeleteRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).DeleteRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/DeleteRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).DeleteRocket(ctx, req.(*DeleteRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RocketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rocket.RocketService",
	HandlerType: (*RocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRocket",
			Handler:    _RocketService_GetRocket_Handler,
		},
		{
			MethodName: "AddRocket",
			Handler:    _RocketService_AddRocket_Handler,
		},
		{
			MethodName: "DeleteRocket",
			Handler:    _RocketService_DeleteRocket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rocket/v1/rocket.proto",
}

func init() { proto.RegisterFile("rocket/v1/rocket.proto", fileDescriptor_rocket_122f075222d9fc46) }

var fileDescriptor_rocket_122f075222d9fc46 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xca, 0x4f, 0xce,
	0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd8,
	0x20, 0x3c, 0x25, 0x09, 0x2e, 0xb6, 0x20, 0x30, 0x4b, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x49, 0x89, 0x4b, 0xc0, 0x3d, 0xb5, 0x04,
	0x22, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x8c, 0xac, 0x86, 0x19, 0xac, 0xc6, 0x9a, 0x4b, 0x10,
	0x49, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1a, 0x17, 0xd4, 0x70, 0xb0, 0x42, 0x6e,
	0x23, 0x3e, 0x3d, 0xa8, 0xcd, 0x50, 0x75, 0x30, 0xab, 0xad, 0xb8, 0x04, 0x1c, 0x53, 0x52, 0x50,
	0x2d, 0x20, 0x56, 0xaf, 0x35, 0x97, 0x20, 0x92, 0x5e, 0x12, 0x2d, 0x56, 0xe5, 0x12, 0x76, 0x49,
	0xcd, 0x49, 0x2d, 0x49, 0xc5, 0xef, 0x39, 0x3d, 0x2e, 0x11, 0x54, 0x65, 0x50, 0x6b, 0xc4, 0xb8,
	0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xa1, 0x81, 0x05, 0xe5, 0x19, 0xdd, 0x63, 0xe4, 0xe2,
	0x85, 0x28, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe0, 0xe2, 0x84, 0x07, 0x8f,
	0x90, 0x04, 0xcc, 0x35, 0xe8, 0xa1, 0x2a, 0x25, 0x89, 0x45, 0x06, 0x6a, 0x97, 0x03, 0x17, 0x27,
	0xdc, 0x9f, 0x08, 0x13, 0xd0, 0x83, 0x0d, 0x61, 0x02, 0x66, 0xa0, 0x78, 0x72, 0xf1, 0x20, 0xfb,
	0x42, 0x48, 0x1a, 0xa6, 0x14, 0x4b, 0x10, 0x48, 0xc9, 0x60, 0x97, 0x84, 0x18, 0x95, 0xc4, 0x06,
	0x4e, 0x3a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x6a, 0xc0, 0xaf, 0x54, 0x02, 0x00,
	0x00,
}
