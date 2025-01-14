// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: danflo27/rps/v1/query.proto

package rps

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryCounterRequest is the request type for the Query/Counter RPC
// method.
type QueryCounterRequest struct {
	// address defines the address to query for the counter.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryCounterRequest) Reset()         { *m = QueryCounterRequest{} }
func (m *QueryCounterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCounterRequest) ProtoMessage()    {}
func (*QueryCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{0}
}
func (m *QueryCounterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterRequest.Merge(m, src)
}
func (m *QueryCounterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterRequest proto.InternalMessageInfo

func (m *QueryCounterRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QueryCounterResponse is the response type for the Query/Counter RPC
// method.
type QueryCounterResponse struct {
	// counter defines the current counter for the sender.
	Counter uint64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *QueryCounterResponse) Reset()         { *m = QueryCounterResponse{} }
func (m *QueryCounterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCounterResponse) ProtoMessage()    {}
func (*QueryCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{1}
}
func (m *QueryCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterResponse.Merge(m, src)
}
func (m *QueryCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterResponse proto.InternalMessageInfo

func (m *QueryCounterResponse) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

// QueryCountersResponse is the request type for the Query/Counters RPC
type QueryCountersRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCountersRequest) Reset()         { *m = QueryCountersRequest{} }
func (m *QueryCountersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCountersRequest) ProtoMessage()    {}
func (*QueryCountersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{2}
}
func (m *QueryCountersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCountersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCountersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCountersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCountersRequest.Merge(m, src)
}
func (m *QueryCountersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCountersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCountersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCountersRequest proto.InternalMessageInfo

func (m *QueryCountersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCountersResponse is the response type for the Query/Counters RPC
// method.
type QueryCountersResponse struct {
	// counters defines all the counters in the store.
	Counters []*Counter `protobuf:"bytes,1,rep,name=counters,proto3" json:"counters,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCountersResponse) Reset()         { *m = QueryCountersResponse{} }
func (m *QueryCountersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCountersResponse) ProtoMessage()    {}
func (*QueryCountersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{3}
}
func (m *QueryCountersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCountersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCountersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCountersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCountersResponse.Merge(m, src)
}
func (m *QueryCountersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCountersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCountersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCountersResponse proto.InternalMessageInfo

func (m *QueryCountersResponse) GetCounters() []*Counter {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *QueryCountersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_144da555a786905c, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryCounterRequest)(nil), "danflo27.rps.v1.QueryCounterRequest")
	proto.RegisterType((*QueryCounterResponse)(nil), "danflo27.rps.v1.QueryCounterResponse")
	proto.RegisterType((*QueryCountersRequest)(nil), "danflo27.rps.v1.QueryCountersRequest")
	proto.RegisterType((*QueryCountersResponse)(nil), "danflo27.rps.v1.QueryCountersResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "danflo27.rps.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "danflo27.rps.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("danflo27/rps/v1/query.proto", fileDescriptor_144da555a786905c) }

var fileDescriptor_144da555a786905c = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x6b, 0x13, 0x4f,
	0x14, 0xcf, 0xb6, 0xdf, 0x6f, 0xd2, 0x4e, 0x0f, 0xe2, 0x18, 0x49, 0xb2, 0x91, 0xad, 0xac, 0x31,
	0x96, 0x82, 0x3b, 0xee, 0x2a, 0x08, 0xde, 0x8c, 0xa0, 0x78, 0x6b, 0xd7, 0x9b, 0x07, 0x65, 0x92,
	0x8c, 0xeb, 0x42, 0x33, 0xb3, 0xdd, 0x99, 0x04, 0x8a, 0x78, 0xd0, 0x8b, 0x82, 0x17, 0x41, 0xf0,
	0x6f, 0xf0, 0xe8, 0xc1, 0x3f, 0xa2, 0xc7, 0xa2, 0x17, 0x4f, 0x22, 0x89, 0xe0, 0xbf, 0x21, 0x3b,
	0xf3, 0xc6, 0x76, 0x93, 0xda, 0x5c, 0xc2, 0xce, 0x7b, 0xef, 0xf3, 0x63, 0x3e, 0x6f, 0x82, 0xda,
	0x43, 0xca, 0x9f, 0xed, 0x89, 0xe8, 0x36, 0xc9, 0x33, 0x49, 0x26, 0x21, 0xd9, 0x1f, 0xb3, 0xfc,
	0x20, 0xc8, 0x72, 0xa1, 0x04, 0x3e, 0x67, 0x9b, 0x41, 0x9e, 0xc9, 0x60, 0x12, 0xba, 0xdb, 0x03,
	0x21, 0x47, 0x42, 0x92, 0x3e, 0x95, 0xcc, 0x4c, 0x92, 0x49, 0xd8, 0x67, 0x8a, 0x86, 0x24, 0xa3,
	0x49, 0xca, 0xa9, 0x4a, 0x05, 0x37, 0x60, 0x77, 0x81, 0x59, 0x1d, 0x64, 0x4c, 0x42, 0xf3, 0x52,
	0x22, 0x44, 0xb2, 0xc7, 0x08, 0xcd, 0x52, 0x42, 0x39, 0x17, 0x4a, 0x23, 0x6d, 0xb7, 0x0d, 0x32,
	0x56, 0xe1, 0xa4, 0x29, 0xf7, 0x3c, 0x1d, 0xa5, 0x5c, 0x10, 0xfd, 0x0b, 0xa5, 0x7a, 0x22, 0x12,
	0xa1, 0x3f, 0x49, 0xf1, 0x05, 0xd5, 0x96, 0x61, 0x79, 0x6a, 0x1a, 0xe6, 0x60, 0x5a, 0xfe, 0x43,
	0x74, 0x61, 0xb7, 0xa0, 0xbc, 0x27, 0xc6, 0x5c, 0xb1, 0x3c, 0x66, 0xfb, 0x63, 0x26, 0x15, 0x8e,
	0x50, 0x8d, 0x0e, 0x87, 0x39, 0x93, 0xb2, 0xe9, 0x5c, 0x76, 0xb6, 0xd6, 0x7b, 0xcd, 0xaf, 0x5f,
	0xae, 0xd7, 0x01, 0x79, 0xd7, 0x74, 0x1e, 0xa9, 0x3c, 0xe5, 0x49, 0x6c, 0x07, 0xfd, 0x1b, 0xa8,
	0x5e, 0xa6, 0x92, 0x99, 0xe0, 0x92, 0xe1, 0x26, 0xaa, 0x0d, 0x4c, 0x49, 0x73, 0xfd, 0x17, 0xdb,
	0xa3, 0xff, 0xa4, 0x8c, 0x90, 0x56, 0xfd, 0x3e, 0x42, 0xc7, 0x21, 0x6a, 0xd0, 0x46, 0xd4, 0x0d,
	0x40, 0xbd, 0x48, 0x3c, 0x30, 0x31, 0x40, 0xe2, 0xc1, 0x0e, 0x4d, 0x18, 0x60, 0xe3, 0x13, 0x48,
	0xff, 0xa3, 0x83, 0x2e, 0xce, 0x09, 0x80, 0xa7, 0x5b, 0x68, 0x0d, 0x4c, 0x14, 0x17, 0x5c, 0xdd,
	0xda, 0x88, 0x9a, 0xc1, 0xdc, 0x8a, 0x03, 0x7b, 0x8f, 0xbf, 0x93, 0xf8, 0x41, 0xc9, 0xd7, 0x8a,
	0xf6, 0x75, 0x6d, 0xa9, 0x2f, 0x23, 0x59, 0x32, 0x56, 0x47, 0x58, 0xfb, 0xda, 0xa1, 0x39, 0x1d,
	0xd9, 0x6b, 0xfb, 0xbb, 0xb0, 0x0b, 0x5b, 0x05, 0xaf, 0x77, 0x50, 0x35, 0xd3, 0x15, 0x48, 0xa2,
	0xb1, 0xe0, 0xd4, 0x00, 0x7a, 0xeb, 0x87, 0x3f, 0x36, 0x2b, 0x9f, 0x7e, 0x7f, 0xde, 0x76, 0x62,
	0x40, 0x44, 0xef, 0x56, 0xd1, 0xff, 0x9a, 0x13, 0xbf, 0x71, 0x50, 0x0d, 0x6e, 0x84, 0x3b, 0x0b,
	0x0c, 0xa7, 0xbc, 0x01, 0xf7, 0xea, 0x92, 0x29, 0x63, 0xcf, 0x27, 0x6f, 0x0b, 0xc5, 0xd7, 0xdf,
	0x7e, 0x7d, 0x58, 0xe9, 0x60, 0x9f, 0xcc, 0xbf, 0x75, 0x08, 0x8f, 0xbc, 0x80, 0x67, 0xf2, 0x12,
	0xbf, 0x72, 0xd0, 0x9a, 0x5d, 0x08, 0x3e, 0x5b, 0xc4, 0x46, 0xe3, 0x76, 0x97, 0x8d, 0x81, 0x99,
	0xee, 0xb1, 0x99, 0x36, 0x6e, 0xfd, 0xcb, 0x8c, 0xc4, 0x0a, 0x55, 0x4d, 0x68, 0xf8, 0xca, 0xe9,
	0xcc, 0xa5, 0xcd, 0xb8, 0x9d, 0xb3, 0x87, 0x40, 0x7c, 0x53, 0xeb, 0xb6, 0x70, 0x63, 0x41, 0xd7,
	0x6c, 0xa3, 0x17, 0x1e, 0x4e, 0x3d, 0xe7, 0x68, 0xea, 0x39, 0x3f, 0xa7, 0x9e, 0xf3, 0x7e, 0xe6,
	0x55, 0x8e, 0x66, 0x5e, 0xe5, 0xfb, 0xcc, 0xab, 0x3c, 0x6e, 0x24, 0xa9, 0x7a, 0x3e, 0xee, 0x07,
	0x03, 0x31, 0x2a, 0x81, 0xfb, 0x55, 0xfd, 0x37, 0xbd, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x34, 0x6c, 0xd1, 0x9e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Counter returns the current counter value.
	Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error)
	// Counters returns all the counter values.
	Counters(ctx context.Context, in *QueryCountersRequest, opts ...grpc.CallOption) (*QueryCountersResponse, error)
	// Params returns the module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Counter(ctx context.Context, in *QueryCounterRequest, opts ...grpc.CallOption) (*QueryCounterResponse, error) {
	out := new(QueryCounterResponse)
	err := c.cc.Invoke(ctx, "/danflo27.rps.v1.Query/Counter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Counters(ctx context.Context, in *QueryCountersRequest, opts ...grpc.CallOption) (*QueryCountersResponse, error) {
	out := new(QueryCountersResponse)
	err := c.cc.Invoke(ctx, "/danflo27.rps.v1.Query/Counters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/danflo27.rps.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Counter returns the current counter value.
	Counter(context.Context, *QueryCounterRequest) (*QueryCounterResponse, error)
	// Counters returns all the counter values.
	Counters(context.Context, *QueryCountersRequest) (*QueryCountersResponse, error)
	// Params returns the module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Counter(ctx context.Context, req *QueryCounterRequest) (*QueryCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counter not implemented")
}
func (*UnimplementedQueryServer) Counters(ctx context.Context, req *QueryCountersRequest) (*QueryCountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counters not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Counter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/danflo27.rps.v1.Query/Counter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Counter(ctx, req.(*QueryCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Counters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/danflo27.rps.v1.Query/Counters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Counters(ctx, req.(*QueryCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/danflo27.rps.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "danflo27.rps.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Counter",
			Handler:    _Query_Counter_Handler,
		},
		{
			MethodName: "Counters",
			Handler:    _Query_Counters_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "danflo27/rps/v1/query.proto",
}

func (m *QueryCounterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Counter != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCountersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCountersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCountersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCountersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCountersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCountersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Counters) > 0 {
		for iNdEx := len(m.Counters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Counters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCounterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Counter != 0 {
		n += 1 + sovQuery(uint64(m.Counter))
	}
	return n
}

func (m *QueryCountersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCountersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Counters) > 0 {
		for _, e := range m.Counters {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCounterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCounterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCountersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCountersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCountersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCountersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCountersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCountersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Counters = append(m.Counters, &Counter{})
			if err := m.Counters[len(m.Counters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
