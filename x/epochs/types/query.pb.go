// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qom/epochs/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryEpochsInfoRequest is the request type for the Query/EpochInfos RPC
// method.
type QueryEpochsInfoRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryEpochsInfoRequest) Reset()         { *m = QueryEpochsInfoRequest{} }
func (m *QueryEpochsInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEpochsInfoRequest) ProtoMessage()    {}
func (*QueryEpochsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0605393ea2db71e, []int{0}
}
func (m *QueryEpochsInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochsInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochsInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochsInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochsInfoRequest.Merge(m, src)
}
func (m *QueryEpochsInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochsInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochsInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochsInfoRequest proto.InternalMessageInfo

func (m *QueryEpochsInfoRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryEpochsInfoResponse is the response type for the Query/EpochInfos RPC
// method.
type QueryEpochsInfoResponse struct {
	// epochs is a slice of all EpochInfos
	Epochs []EpochInfo `protobuf:"bytes,1,rep,name=epochs,proto3" json:"epochs"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryEpochsInfoResponse) Reset()         { *m = QueryEpochsInfoResponse{} }
func (m *QueryEpochsInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEpochsInfoResponse) ProtoMessage()    {}
func (*QueryEpochsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0605393ea2db71e, []int{1}
}
func (m *QueryEpochsInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochsInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochsInfoResponse.Merge(m, src)
}
func (m *QueryEpochsInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochsInfoResponse proto.InternalMessageInfo

func (m *QueryEpochsInfoResponse) GetEpochs() []EpochInfo {
	if m != nil {
		return m.Epochs
	}
	return nil
}

func (m *QueryEpochsInfoResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCurrentEpochRequest is the request type for the Query/EpochInfos RPC
// method.
type QueryCurrentEpochRequest struct {
	// identifier of the current epoch
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (m *QueryCurrentEpochRequest) Reset()         { *m = QueryCurrentEpochRequest{} }
func (m *QueryCurrentEpochRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochRequest) ProtoMessage()    {}
func (*QueryCurrentEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0605393ea2db71e, []int{2}
}
func (m *QueryCurrentEpochRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochRequest.Merge(m, src)
}
func (m *QueryCurrentEpochRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochRequest proto.InternalMessageInfo

func (m *QueryCurrentEpochRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

// QueryCurrentEpochResponse is the response type for the Query/EpochInfos RPC
// method.
type QueryCurrentEpochResponse struct {
	// current_epoch is the number of the current epoch
	CurrentEpoch int64 `protobuf:"varint,1,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
}

func (m *QueryCurrentEpochResponse) Reset()         { *m = QueryCurrentEpochResponse{} }
func (m *QueryCurrentEpochResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochResponse) ProtoMessage()    {}
func (*QueryCurrentEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0605393ea2db71e, []int{3}
}
func (m *QueryCurrentEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochResponse.Merge(m, src)
}
func (m *QueryCurrentEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochResponse proto.InternalMessageInfo

func (m *QueryCurrentEpochResponse) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryEpochsInfoRequest)(nil), "qom.epochs.v1.QueryEpochsInfoRequest")
	proto.RegisterType((*QueryEpochsInfoResponse)(nil), "qom.epochs.v1.QueryEpochsInfoResponse")
	proto.RegisterType((*QueryCurrentEpochRequest)(nil), "qom.epochs.v1.QueryCurrentEpochRequest")
	proto.RegisterType((*QueryCurrentEpochResponse)(nil), "qom.epochs.v1.QueryCurrentEpochResponse")
}

func init() { proto.RegisterFile("qom/epochs/v1/query.proto", fileDescriptor_b0605393ea2db71e) }

var fileDescriptor_b0605393ea2db71e = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x6f, 0x13, 0x31,
	0x14, 0xc6, 0x73, 0x29, 0x54, 0xc2, 0x6d, 0x17, 0x0b, 0x68, 0x1a, 0x8a, 0xa9, 0x02, 0xb4, 0x11,
	0x08, 0x5b, 0x09, 0x12, 0x03, 0x13, 0x2a, 0x2a, 0x88, 0x0d, 0x6e, 0x64, 0x01, 0xe7, 0x78, 0x75,
	0x2d, 0x11, 0x3f, 0xe7, 0xec, 0x44, 0x94, 0x91, 0x85, 0x09, 0x09, 0x89, 0x8d, 0xbf, 0xa8, 0x63,
	0x25, 0x16, 0x26, 0x84, 0x12, 0xfe, 0x10, 0x74, 0xb6, 0x0b, 0xb9, 0x36, 0x52, 0xa6, 0xb3, 0xfc,
	0xde, 0xf7, 0xbe, 0xdf, 0xfb, 0xce, 0x64, 0x6b, 0x84, 0x43, 0x01, 0x16, 0x8b, 0x23, 0x27, 0x26,
	0x3d, 0x31, 0x1a, 0x43, 0x79, 0xcc, 0x6d, 0x89, 0x1e, 0xe9, 0xc6, 0x08, 0x87, 0x3c, 0x96, 0xf8,
	0xa4, 0xd7, 0xbe, 0x57, 0xa0, 0x1b, 0xa2, 0x13, 0x03, 0xe9, 0x20, 0xf6, 0x89, 0x49, 0x6f, 0x00,
	0x5e, 0xf6, 0x84, 0x95, 0x4a, 0x1b, 0xe9, 0x35, 0x9a, 0x28, 0x6d, 0x5f, 0x55, 0xa8, 0x30, 0x1c,
	0x45, 0x75, 0x4a, 0xb7, 0xdb, 0x0a, 0x51, 0xbd, 0x07, 0x21, 0xad, 0x16, 0xd2, 0x18, 0xf4, 0x41,
	0xe2, 0x52, 0xf5, 0x46, 0x9d, 0x44, 0x81, 0x01, 0xa7, 0x53, 0xb1, 0xf3, 0x96, 0x5c, 0x7f, 0x55,
	0x59, 0x1e, 0x84, 0xfa, 0x0b, 0x73, 0x88, 0x39, 0x8c, 0xc6, 0xe0, 0x3c, 0x7d, 0x46, 0xc8, 0x7f,
	0xfb, 0x56, 0xb6, 0x93, 0x75, 0xd7, 0xfa, 0xbb, 0x3c, 0xb2, 0xf2, 0x8a, 0x95, 0xc7, 0x9d, 0x12,
	0x2b, 0x7f, 0x29, 0x15, 0x24, 0x6d, 0x3e, 0xa7, 0xec, 0x7c, 0xcf, 0xc8, 0xe6, 0x05, 0x0b, 0x67,
	0xd1, 0x38, 0xa0, 0x8f, 0xc8, 0x6a, 0x04, 0x6b, 0x65, 0x3b, 0x2b, 0xdd, 0xb5, 0x7e, 0x8b, 0xd7,
	0xa2, 0xe1, 0x41, 0x52, 0x29, 0xf6, 0x2f, 0x9d, 0xfc, 0xba, 0xd5, 0xc8, 0x53, 0x37, 0x7d, 0x5e,
	0x63, 0x6b, 0x06, 0xb6, 0xbd, 0xa5, 0x6c, 0xd1, 0xb4, 0x06, 0xf7, 0x98, 0xb4, 0x02, 0xdb, 0xd3,
	0x71, 0x59, 0x82, 0xf1, 0xc1, 0xef, 0x2c, 0x00, 0x46, 0x88, 0x7e, 0x07, 0xc6, 0xeb, 0x43, 0x0d,
	0x65, 0x08, 0xe0, 0x4a, 0x3e, 0x77, 0xd3, 0x79, 0x42, 0xb6, 0x16, 0x68, 0xd3, 0x66, 0xb7, 0xc9,
	0x46, 0x11, 0xef, 0xdf, 0x04, 0xe6, 0xa0, 0x5f, 0xc9, 0xd7, 0x8b, 0xb9, 0xe6, 0xfe, 0x97, 0x26,
	0xb9, 0x1c, 0x46, 0xd0, 0x8f, 0x84, 0xfc, 0xdb, 0xd5, 0xd1, 0xbb, 0xe7, 0x62, 0x58, 0xfc, 0x87,
	0xda, 0xbb, 0xcb, 0xda, 0x22, 0x4b, 0xe7, 0xe6, 0xa7, 0x1f, 0x7f, 0xbe, 0x35, 0x37, 0xe9, 0x35,
	0x51, 0x7f, 0x09, 0x29, 0xcc, 0xcf, 0x19, 0x59, 0x9f, 0xdf, 0x81, 0xee, 0x2d, 0x9a, 0xbb, 0x20,
	0xa1, 0x76, 0x77, 0x79, 0x63, 0x42, 0xb8, 0x13, 0x10, 0x18, 0xdd, 0x3e, 0x87, 0x50, 0xcb, 0x68,
	0xff, 0xe0, 0x64, 0xca, 0xb2, 0xd3, 0x29, 0xcb, 0x7e, 0x4f, 0x59, 0xf6, 0x75, 0xc6, 0x1a, 0xa7,
	0x33, 0xd6, 0xf8, 0x39, 0x63, 0x8d, 0xd7, 0xf7, 0x95, 0xf6, 0x47, 0xe3, 0x01, 0x2f, 0x70, 0x58,
	0x4d, 0x78, 0x80, 0x06, 0xaa, 0xaf, 0xb4, 0xb6, 0x9a, 0xf2, 0xe1, 0x6c, 0xa2, 0x3f, 0xb6, 0xe0,
	0x06, 0xab, 0xe1, 0x69, 0x3f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xbf, 0xf6, 0xfc, 0x83,
	0x03, 0x00, 0x00,
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
	// EpochInfos provide running epochInfos
	EpochInfos(ctx context.Context, in *QueryEpochsInfoRequest, opts ...grpc.CallOption) (*QueryEpochsInfoResponse, error)
	// CurrentEpoch provide current epoch of specified identifier
	CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EpochInfos(ctx context.Context, in *QueryEpochsInfoRequest, opts ...grpc.CallOption) (*QueryEpochsInfoResponse, error) {
	out := new(QueryEpochsInfoResponse)
	err := c.cc.Invoke(ctx, "/qom.epochs.v1.Query/EpochInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error) {
	out := new(QueryCurrentEpochResponse)
	err := c.cc.Invoke(ctx, "/qom.epochs.v1.Query/CurrentEpoch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// EpochInfos provide running epochInfos
	EpochInfos(context.Context, *QueryEpochsInfoRequest) (*QueryEpochsInfoResponse, error)
	// CurrentEpoch provide current epoch of specified identifier
	CurrentEpoch(context.Context, *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EpochInfos(ctx context.Context, req *QueryEpochsInfoRequest) (*QueryEpochsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochInfos not implemented")
}
func (*UnimplementedQueryServer) CurrentEpoch(ctx context.Context, req *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentEpoch not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EpochInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qom.epochs.v1.Query/EpochInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochInfos(ctx, req.(*QueryEpochsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CurrentEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qom.epochs.v1.Query/CurrentEpoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentEpoch(ctx, req.(*QueryCurrentEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qom.epochs.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EpochInfos",
			Handler:    _Query_EpochInfos_Handler,
		},
		{
			MethodName: "CurrentEpoch",
			Handler:    _Query_CurrentEpoch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qom/epochs/v1/query.proto",
}

func (m *QueryEpochsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochsInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryEpochsInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochsInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochsInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Epochs) > 0 {
		for iNdEx := len(m.Epochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryCurrentEpochRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrentEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x8
	}
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
func (m *QueryEpochsInfoRequest) Size() (n int) {
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

func (m *QueryEpochsInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for _, e := range m.Epochs {
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

func (m *QueryCurrentEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrentEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		n += 1 + sovQuery(uint64(m.CurrentEpoch))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEpochsInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEpochsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryEpochsInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEpochsInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochsInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
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
			m.Epochs = append(m.Epochs, EpochInfo{})
			if err := m.Epochs[len(m.Epochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCurrentEpochRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentEpochRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
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
			m.Identifier = string(dAtA[iNdEx:postIndex])
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
func (m *QueryCurrentEpochResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCurrentEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
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
