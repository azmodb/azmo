// Code generated by protoc-gen-go.
// source: azmo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	azmo.proto

It has these top-level messages:
	TxnRequest
	TxnResponse
	GenericRequest
	GenericResponse
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	GetRequest
	GetResponse
	RangeRequest
	RangeResponse
*/
package pb

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

type GenericRequest_Type int32

const (
	GenericRequest_InvalidRequest GenericRequest_Type = 0
	GenericRequest_DeleteRequest  GenericRequest_Type = 1
	GenericRequest_PutRequest     GenericRequest_Type = 2
)

var GenericRequest_Type_name = map[int32]string{
	0: "InvalidRequest",
	1: "DeleteRequest",
	2: "PutRequest",
}
var GenericRequest_Type_value = map[string]int32{
	"InvalidRequest": 0,
	"DeleteRequest":  1,
	"PutRequest":     2,
}

func (x GenericRequest_Type) String() string {
	return proto.EnumName(GenericRequest_Type_name, int32(x))
}
func (GenericRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type TxnRequest struct {
	Requests []*GenericRequest `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *TxnRequest) Reset()                    { *m = TxnRequest{} }
func (m *TxnRequest) String() string            { return proto.CompactTextString(m) }
func (*TxnRequest) ProtoMessage()               {}
func (*TxnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxnRequest) GetRequests() []*GenericRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

type TxnResponse struct {
	Responses []*GenericResponse `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
}

func (m *TxnResponse) Reset()                    { *m = TxnResponse{} }
func (m *TxnResponse) String() string            { return proto.CompactTextString(m) }
func (*TxnResponse) ProtoMessage()               {}
func (*TxnResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TxnResponse) GetResponses() []*GenericResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

type GenericRequest struct {
	Type      GenericRequest_Type `protobuf:"varint,1,opt,name=type,enum=azmo.GenericRequest_Type" json:"type,omitempty"`
	Num       int32               `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	Key       []byte              `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte              `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Tombstone bool                `protobuf:"varint,5,opt,name=tombstone" json:"tombstone,omitempty"`
}

func (m *GenericRequest) Reset()                    { *m = GenericRequest{} }
func (m *GenericRequest) String() string            { return proto.CompactTextString(m) }
func (*GenericRequest) ProtoMessage()               {}
func (*GenericRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GenericResponse struct {
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Rev int64 `protobuf:"varint,2,opt,name=rev" json:"rev,omitempty"`
}

func (m *GenericResponse) Reset()                    { *m = GenericResponse{} }
func (m *GenericResponse) String() string            { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()               {}
func (*GenericResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PutRequest struct {
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Tombstone bool   `protobuf:"varint,4,opt,name=tombstone" json:"tombstone,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PutResponse struct {
	Rev int64 `protobuf:"varint,1,opt,name=rev" json:"rev,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteResponse struct {
	Rev int64 `protobuf:"varint,1,opt,name=rev" json:"rev,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Rev int64  `protobuf:"varint,2,opt,name=rev" json:"rev,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetResponse struct {
	Value []byte  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Revs  []int64 `protobuf:"varint,2,rep,name=revs" json:"revs,omitempty"`
	Rev   int64   `protobuf:"varint,3,opt,name=rev" json:"rev,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type RangeRequest struct {
	From []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Rev  int64  `protobuf:"varint,3,opt,name=rev" json:"rev,omitempty"`
}

func (m *RangeRequest) Reset()                    { *m = RangeRequest{} }
func (m *RangeRequest) String() string            { return proto.CompactTextString(m) }
func (*RangeRequest) ProtoMessage()               {}
func (*RangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type RangeResponse struct {
	Key   []byte  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Revs  []int64 `protobuf:"varint,2,rep,name=revs" json:"revs,omitempty"`
	Value []byte  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Rev   int64   `protobuf:"varint,4,opt,name=rev" json:"rev,omitempty"`
}

func (m *RangeResponse) Reset()                    { *m = RangeResponse{} }
func (m *RangeResponse) String() string            { return proto.CompactTextString(m) }
func (*RangeResponse) ProtoMessage()               {}
func (*RangeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*TxnRequest)(nil), "azmo.TxnRequest")
	proto.RegisterType((*TxnResponse)(nil), "azmo.TxnResponse")
	proto.RegisterType((*GenericRequest)(nil), "azmo.GenericRequest")
	proto.RegisterType((*GenericResponse)(nil), "azmo.GenericResponse")
	proto.RegisterType((*PutRequest)(nil), "azmo.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "azmo.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "azmo.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "azmo.DeleteResponse")
	proto.RegisterType((*GetRequest)(nil), "azmo.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "azmo.GetResponse")
	proto.RegisterType((*RangeRequest)(nil), "azmo.RangeRequest")
	proto.RegisterType((*RangeResponse)(nil), "azmo.RangeResponse")
	proto.RegisterEnum("azmo.GenericRequest_Type", GenericRequest_Type_name, GenericRequest_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for DB service

type DBClient interface {
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (DB_RangeClient, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error)
}

type dBClient struct {
	cc *grpc.ClientConn
}

func NewDBClient(cc *grpc.ClientConn) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (DB_RangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[0], c.cc, "/azmo.DB/Range", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_RangeClient interface {
	Recv() (*RangeResponse, error)
	grpc.ClientStream
}

type dBRangeClient struct {
	grpc.ClientStream
}

func (x *dBRangeClient) Recv() (*RangeResponse, error) {
	m := new(RangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/azmo.DB/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/azmo.DB/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/azmo.DB/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error) {
	out := new(TxnResponse)
	err := grpc.Invoke(ctx, "/azmo.DB/Txn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DB service

type DBServer interface {
	Range(*RangeRequest, DB_RangeServer) error
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Txn(context.Context, *TxnRequest) (*TxnResponse, error)
}

func RegisterDBServer(s *grpc.Server, srv DBServer) {
	s.RegisterService(&_DB_serviceDesc, srv)
}

func _DB_Range_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Range(m, &dBRangeServer{stream})
}

type DB_RangeServer interface {
	Send(*RangeResponse) error
	grpc.ServerStream
}

type dBRangeServer struct {
	grpc.ServerStream
}

func (x *dBRangeServer) Send(m *RangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azmo.DB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azmo.DB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azmo.DB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Txn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Txn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azmo.DB/Txn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Txn(ctx, req.(*TxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "azmo.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DB_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DB_Delete_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DB_Put_Handler,
		},
		{
			MethodName: "Txn",
			Handler:    _DB_Txn_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Range",
			Handler:       _DB_Range_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x65, 0xed, 0x4d, 0xd5, 0x4e, 0x5a, 0x93, 0x0e, 0x41, 0x32, 0x11, 0x12, 0xb0, 0xa7, 0x0a,
	0x89, 0x28, 0x4a, 0xc5, 0x11, 0x0e, 0x51, 0xa4, 0xaa, 0x37, 0xb4, 0xea, 0x09, 0x89, 0x43, 0x02,
	0x0b, 0xaa, 0x68, 0xbc, 0xc1, 0x5e, 0x57, 0x84, 0x1f, 0xca, 0xff, 0xe0, 0x1f, 0x30, 0x5e, 0xef,
	0x7a, 0x6d, 0xcb, 0xb9, 0x3d, 0x8f, 0xdf, 0xbc, 0x37, 0x6f, 0xbf, 0x00, 0x36, 0x7f, 0x76, 0x7a,
	0xbe, 0xcf, 0xb5, 0xd1, 0xc8, 0x2b, 0x2c, 0x3e, 0x02, 0xdc, 0xfd, 0xce, 0xa4, 0xfa, 0x55, 0xaa,
	0xc2, 0xe0, 0x02, 0x4e, 0xf3, 0x1a, 0x16, 0x29, 0x7b, 0x1d, 0x5f, 0x8d, 0x97, 0xd3, 0xb9, 0x6d,
	0xb9, 0x51, 0x99, 0xca, 0xef, 0xbf, 0x3a, 0x9e, 0x6c, 0x58, 0x62, 0x05, 0x63, 0xdb, 0x5f, 0xec,
	0x75, 0x56, 0x28, 0xbc, 0x86, 0xb3, 0xdc, 0x61, 0xaf, 0xf0, 0xbc, 0xa7, 0x50, 0xff, 0x95, 0x81,
	0x27, 0xfe, 0x32, 0x48, 0xba, 0x06, 0xf8, 0x0e, 0xb8, 0x39, 0xec, 0x15, 0x49, 0xb0, 0xab, 0x64,
	0xf9, 0x62, 0x68, 0x88, 0xf9, 0x1d, 0x11, 0xa4, 0xa5, 0xe1, 0x04, 0xe2, 0xac, 0xdc, 0xa5, 0x11,
	0xb1, 0x47, 0xb2, 0x82, 0x55, 0xe5, 0xa7, 0x3a, 0xa4, 0x31, 0x55, 0xce, 0x65, 0x05, 0x71, 0x0a,
	0xa3, 0xc7, 0xcd, 0x43, 0xa9, 0x52, 0x6e, 0x6b, 0xf5, 0x07, 0xbe, 0x84, 0x33, 0xa3, 0x77, 0xdb,
	0xc2, 0xe8, 0x4c, 0xa5, 0x23, 0xfa, 0x73, 0x2a, 0x43, 0x41, 0x7c, 0x00, 0x5e, 0xb9, 0x20, 0x42,
	0x72, 0x9b, 0x51, 0xc3, 0xfd, 0x37, 0x67, 0x3e, 0x79, 0x82, 0x97, 0x70, 0xb1, 0x56, 0x0f, 0xca,
	0x28, 0x5f, 0x62, 0x98, 0x00, 0x7c, 0x2a, 0x8d, 0xff, 0x8e, 0xc4, 0x7b, 0x78, 0xda, 0x8b, 0xed,
	0x27, 0x65, 0x9d, 0x49, 0x73, 0xf5, 0x68, 0x67, 0x8f, 0x65, 0x05, 0x85, 0x6c, 0xcb, 0xf8, 0x24,
	0xd1, 0x40, 0x92, 0xf8, 0x68, 0x12, 0xde, 0x4f, 0xf2, 0x0a, 0xc6, 0x56, 0x33, 0x8c, 0x51, 0x99,
	0xb2, 0x60, 0xfa, 0xa6, 0x17, 0xc7, 0xfb, 0xb2, 0xc6, 0x57, 0x08, 0x48, 0x3c, 0xe5, 0xa8, 0xcc,
	0x02, 0xe0, 0x46, 0x99, 0xa3, 0x1a, 0x03, 0x69, 0x6f, 0x61, 0x6c, 0x3b, 0x9c, 0x64, 0x13, 0x8e,
	0xb5, 0xc3, 0x21, 0x70, 0xe2, 0x16, 0xd4, 0x17, 0x53, 0x9f, 0xc5, 0x5e, 0x2a, 0x0e, 0x52, 0x6b,
	0x38, 0x97, 0x9b, 0xec, 0x47, 0x13, 0x81, 0xba, 0xbe, 0xe7, 0x7a, 0xe7, 0xa4, 0x2c, 0xa6, 0x3d,
	0x8a, 0x8c, 0x76, 0xab, 0x49, 0x68, 0x40, 0xe5, 0x0b, 0x5c, 0x38, 0x95, 0x90, 0xb2, 0x97, 0x62,
	0x68, 0x9c, 0xe1, 0x5d, 0x71, 0xf2, 0xbc, 0x91, 0x5f, 0xfe, 0x63, 0x10, 0xad, 0x57, 0xb8, 0x84,
	0x91, 0x75, 0x41, 0xac, 0x0f, 0x77, 0x7b, 0xf0, 0xd9, 0xb3, 0x4e, 0xad, 0x1e, 0x63, 0xc1, 0xf0,
	0x2d, 0xc4, 0xb4, 0x54, 0x38, 0xf1, 0xd7, 0xc1, 0xaf, 0xf3, 0xec, 0xb2, 0x55, 0x69, 0x6e, 0xe2,
	0x49, 0xbd, 0x59, 0xe8, 0xc4, 0x3a, 0xbb, 0x3b, 0x9b, 0x76, 0x8b, 0xae, 0x89, 0x0c, 0xe8, 0x94,
	0x78, 0x83, 0x70, 0x08, 0xbd, 0x41, 0xfb, 0x08, 0x11, 0x97, 0x6e, 0xbe, 0xe7, 0x86, 0x47, 0xc4,
	0x73, 0x5b, 0xcf, 0xc2, 0x8a, 0x7f, 0x8e, 0xf6, 0xdb, 0xed, 0x89, 0x7d, 0x78, 0xae, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x63, 0xca, 0x0d, 0x86, 0x04, 0x00, 0x00,
}
