// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: htlc/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryHTLCRequest is the request type for the Query/HTLC RPC method
type QueryHTLCRequest struct {
	// address is the address to query balances for
	HashLock string `protobuf:"bytes,1,opt,name=hash_lock,json=hashLock,proto3" json:"hash_lock,omitempty"`
}

func (m *QueryHTLCRequest) Reset()         { *m = QueryHTLCRequest{} }
func (m *QueryHTLCRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHTLCRequest) ProtoMessage()    {}
func (*QueryHTLCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a99e89fd1d8bb804, []int{0}
}
func (m *QueryHTLCRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHTLCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHTLCRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHTLCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHTLCRequest.Merge(m, src)
}
func (m *QueryHTLCRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHTLCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHTLCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHTLCRequest proto.InternalMessageInfo

func (m *QueryHTLCRequest) GetHashLock() string {
	if m != nil {
		return m.HashLock
	}
	return ""
}

// QueryBalanceResponse is the response type for the Query/HTLC RPC method
type QueryHTLCResponse struct {
	Htlc *HTLC `protobuf:"bytes,1,opt,name=htlc,proto3" json:"htlc,omitempty"`
}

func (m *QueryHTLCResponse) Reset()         { *m = QueryHTLCResponse{} }
func (m *QueryHTLCResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHTLCResponse) ProtoMessage()    {}
func (*QueryHTLCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a99e89fd1d8bb804, []int{1}
}
func (m *QueryHTLCResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHTLCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHTLCResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHTLCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHTLCResponse.Merge(m, src)
}
func (m *QueryHTLCResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHTLCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHTLCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHTLCResponse proto.InternalMessageInfo

func (m *QueryHTLCResponse) GetHtlc() *HTLC {
	if m != nil {
		return m.Htlc
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryHTLCRequest)(nil), "irismod.htlc.QueryHTLCRequest")
	proto.RegisterType((*QueryHTLCResponse)(nil), "irismod.htlc.QueryHTLCResponse")
}

func init() { proto.RegisterFile("htlc/query.proto", fileDescriptor_a99e89fd1d8bb804) }

var fileDescriptor_a99e89fd1d8bb804 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x28, 0xc9, 0x49,
	0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0x2c,
	0xca, 0x2c, 0xce, 0xcd, 0x4f, 0xd1, 0x03, 0xc9, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25,
	0xf4, 0x41, 0x2c, 0x88, 0x1a, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82,
	0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xa8, 0x2c, 0x3f,
	0xd8, 0x4c, 0x10, 0x01, 0x11, 0x50, 0xd2, 0xe7, 0x12, 0x08, 0x04, 0xd9, 0xe0, 0x11, 0xe2, 0xe3,
	0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xcd, 0xc5, 0x99, 0x91, 0x58, 0x9c, 0x11,
	0x9f, 0x93, 0x9f, 0x9c, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x01, 0x12, 0xf0, 0xc9,
	0x4f, 0xce, 0x56, 0xb2, 0xe6, 0x12, 0x44, 0xd2, 0x50, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4,
	0xc6, 0xc5, 0x02, 0x32, 0x13, 0xac, 0x98, 0xdb, 0x48, 0x48, 0x0f, 0xd9, 0x9d, 0x7a, 0x60, 0x95,
	0x60, 0x79, 0xa3, 0x4a, 0x2e, 0x56, 0xb0, 0x66, 0xa1, 0x02, 0x2e, 0x16, 0x90, 0xb0, 0x90, 0x1c,
	0xaa, 0x52, 0x74, 0xa7, 0x48, 0xc9, 0xe3, 0x94, 0x87, 0xd8, 0xac, 0xa4, 0xde, 0x74, 0xf9, 0xc9,
	0x64, 0x26, 0x45, 0x21, 0x79, 0x7d, 0xa8, 0x42, 0x7d, 0xb8, 0x0f, 0x8b, 0xf5, 0xab, 0xe1, 0xde,
	0xa8, 0x75, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xdd, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xb0, 0x21, 0x79, 0xa9, 0x25, 0x70, 0xc3,
	0x72, 0xf3, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0x21, 0x86, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1,
	0x81, 0x03, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x66, 0x76, 0xf1, 0xd9, 0x9f, 0x01, 0x00,
	0x00,
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
	// Balance queries the balance of a single coin for a single account
	HTLC(ctx context.Context, in *QueryHTLCRequest, opts ...grpc.CallOption) (*QueryHTLCResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) HTLC(ctx context.Context, in *QueryHTLCRequest, opts ...grpc.CallOption) (*QueryHTLCResponse, error) {
	out := new(QueryHTLCResponse)
	err := c.cc.Invoke(ctx, "/irismod.htlc.Query/HTLC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Balance queries the balance of a single coin for a single account
	HTLC(context.Context, *QueryHTLCRequest) (*QueryHTLCResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) HTLC(ctx context.Context, req *QueryHTLCRequest) (*QueryHTLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTLC not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_HTLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHTLCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HTLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.htlc.Query/HTLC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HTLC(ctx, req.(*QueryHTLCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.htlc.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HTLC",
			Handler:    _Query_HTLC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "htlc/query.proto",
}

func (m *QueryHTLCRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHTLCRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHTLCRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HashLock) > 0 {
		i -= len(m.HashLock)
		copy(dAtA[i:], m.HashLock)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.HashLock)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHTLCResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHTLCResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHTLCResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Htlc != nil {
		{
			size, err := m.Htlc.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryHTLCRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HashLock)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHTLCResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Htlc != nil {
		l = m.Htlc.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryHTLCRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHTLCRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHTLCRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashLock", wireType)
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
			m.HashLock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryHTLCResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHTLCResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHTLCResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Htlc", wireType)
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
			if m.Htlc == nil {
				m.Htlc = &HTLC{}
			}
			if err := m.Htlc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
