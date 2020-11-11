// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random/random.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Random defines the feed standard
type Random struct {
	RequestTxHash string `protobuf:"bytes,1,opt,name=request_tx_hash,json=requestTxHash,proto3" json:"request_tx_hash,omitempty" yaml:"request_tx_hash"`
	Height        int64  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Value         string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Random) Reset()         { *m = Random{} }
func (m *Random) String() string { return proto.CompactTextString(m) }
func (*Random) ProtoMessage()    {}
func (*Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{0}
}
func (m *Random) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Random.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Random.Merge(m, src)
}
func (m *Random) XXX_Size() int {
	return m.Size()
}
func (m *Random) XXX_DiscardUnknown() {
	xxx_messageInfo_Random.DiscardUnknown(m)
}

var xxx_messageInfo_Random proto.InternalMessageInfo

func (m *Random) GetRequestTxHash() string {
	if m != nil {
		return m.RequestTxHash
	}
	return ""
}

func (m *Random) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Random) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Request defines the random request standard
type Request struct {
	Height           int64                                    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Consumer         string                                   `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	TxHash           string                                   `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty" yaml:"tx_hash"`
	Oracle           bool                                     `protobuf:"varint,4,opt,name=oracle,proto3" json:"oracle,omitempty"`
	ServiceFeeCap    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=service_fee_cap,json=serviceFeeCap,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"service_fee_cap" yaml:"service_fee_cap"`
	ServiceContextID string                                   `protobuf:"bytes,6,opt,name=service_context_id,json=serviceContextId,proto3" json:"service_context_id,omitempty" yaml:"service_context_id"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Request) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *Request) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Request) GetOracle() bool {
	if m != nil {
		return m.Oracle
	}
	return false
}

func (m *Request) GetServiceFeeCap() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ServiceFeeCap
	}
	return nil
}

func (m *Request) GetServiceContextID() string {
	if m != nil {
		return m.ServiceContextID
	}
	return ""
}

func init() {
	proto.RegisterType((*Random)(nil), "irismod.random.Random")
	proto.RegisterType((*Request)(nil), "irismod.random.Request")
}

func init() { proto.RegisterFile("random/random.proto", fileDescriptor_e5da2919a686585f) }

var fileDescriptor_e5da2919a686585f = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0xad, 0x57, 0x96, 0x6d, 0x46, 0xfb, 0x23, 0x33, 0x8d, 0xac, 0x87, 0xa4, 0xca, 0x29, 0x12,
	0x22, 0x51, 0xe1, 0xc6, 0x31, 0x45, 0x88, 0x71, 0x34, 0x9c, 0xb8, 0x44, 0xae, 0xf3, 0xa3, 0xb1,
	0x48, 0xe2, 0x10, 0x3b, 0x55, 0xc7, 0x77, 0x40, 0xe2, 0x73, 0xf0, 0x49, 0x76, 0xdc, 0x91, 0x53,
	0x40, 0xe9, 0x37, 0x28, 0x5f, 0x00, 0x35, 0xce, 0xba, 0x75, 0x27, 0xfb, 0xf9, 0xfd, 0xde, 0xd3,
	0xf3, 0x4f, 0x0f, 0x3f, 0xab, 0x58, 0x91, 0xc8, 0x3c, 0x34, 0x47, 0x50, 0x56, 0x52, 0x4b, 0x72,
	0x22, 0x2a, 0xa1, 0x72, 0x99, 0x04, 0xe6, 0x75, 0x74, 0x3e, 0x97, 0x73, 0xd9, 0x51, 0xe1, 0xe6,
	0x66, 0xa6, 0x46, 0xcf, 0xb9, 0x54, 0xb9, 0x54, 0xb1, 0x21, 0xb8, 0x14, 0x85, 0x21, 0xbc, 0xef,
	0xd8, 0xa2, 0x9d, 0x90, 0x44, 0xf8, 0xb4, 0x82, 0x6f, 0x35, 0x28, 0x1d, 0xeb, 0x65, 0x9c, 0x32,
	0x95, 0xda, 0x68, 0x8c, 0xfc, 0xa3, 0x68, 0xb4, 0x6e, 0xdc, 0x8b, 0x6b, 0x96, 0x67, 0x6f, 0xbc,
	0x47, 0x03, 0x1e, 0x3d, 0xee, 0x5f, 0x3e, 0x2d, 0xdf, 0x33, 0x95, 0x92, 0x0b, 0x6c, 0xa5, 0x20,
	0xe6, 0xa9, 0xb6, 0xf7, 0xc6, 0xc8, 0x1f, 0xd2, 0x1e, 0x91, 0x73, 0xbc, 0xbf, 0x60, 0x59, 0x0d,
	0xf6, 0x70, 0xe3, 0x48, 0x0d, 0xf0, 0xfe, 0xed, 0xe1, 0x03, 0x6a, 0xf4, 0x0f, 0x94, 0x68, 0x47,
	0x39, 0xc2, 0x87, 0x5c, 0x16, 0xaa, 0xce, 0xa1, 0xea, 0x3c, 0x8f, 0xe8, 0x16, 0x93, 0x17, 0xf8,
	0xe0, 0x2e, 0x69, 0xe7, 0x1b, 0x91, 0x75, 0xe3, 0x9e, 0x98, 0xa4, 0xdb, 0x84, 0x96, 0xde, 0x46,
	0x93, 0x15, 0xe3, 0x19, 0xd8, 0x4f, 0xc6, 0xc8, 0x3f, 0xa4, 0x3d, 0x22, 0x3f, 0x10, 0x3e, 0x55,
	0x50, 0x2d, 0x04, 0x87, 0xf8, 0x0b, 0x40, 0xcc, 0x59, 0x69, 0xef, 0x8f, 0x87, 0xfe, 0xd3, 0x57,
	0x97, 0x81, 0x59, 0x5a, 0x30, 0x63, 0x0a, 0x82, 0xc5, 0x64, 0x06, 0x9a, 0x4d, 0x82, 0xa9, 0x14,
	0x45, 0xf4, 0xe1, 0xa6, 0x71, 0x07, 0xf7, 0x6b, 0x79, 0xa4, 0xf7, 0x7e, 0xfd, 0x71, 0xfd, 0xb9,
	0xd0, 0x69, 0x3d, 0x0b, 0xb8, 0xcc, 0x43, 0x63, 0xd3, 0x1f, 0x2f, 0x55, 0xf2, 0x35, 0xd4, 0xd7,
	0x25, 0xa8, 0xce, 0x4a, 0xd1, 0xe3, 0x5e, 0xfd, 0x0e, 0x60, 0xca, 0x4a, 0x12, 0x63, 0x72, 0x67,
	0xc7, 0x65, 0xa1, 0x61, 0xa9, 0x63, 0x91, 0xd8, 0x56, 0xf7, 0xbf, 0x49, 0xdb, 0xb8, 0x67, 0x1f,
	0x0d, 0x3b, 0x35, 0xe4, 0xd5, 0xdb, 0x75, 0xe3, 0x5e, 0xee, 0xc6, 0xb8, 0xd7, 0x79, 0xf4, 0x4c,
	0xed, 0x8e, 0x27, 0xd1, 0xd5, 0x4d, 0xeb, 0xa0, 0xdb, 0xd6, 0x41, 0x7f, 0x5b, 0x07, 0xfd, 0x5c,
	0x39, 0x83, 0xdb, 0x95, 0x33, 0xf8, 0xbd, 0x72, 0x06, 0x9f, 0xc3, 0x07, 0x99, 0x37, 0xad, 0x2a,
	0x40, 0x87, 0x7d, 0xbb, 0xc2, 0x5c, 0x26, 0x75, 0x06, 0xaa, 0xef, 0x9e, 0xf9, 0xc0, 0xcc, 0xea,
	0x3a, 0xf4, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xd7, 0x7a, 0x19, 0x99, 0x02, 0x00,
	0x00,
}

func (m *Random) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Random) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Random) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintRandom(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RequestTxHash) > 0 {
		i -= len(m.RequestTxHash)
		copy(dAtA[i:], m.RequestTxHash)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.RequestTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceContextID) > 0 {
		i -= len(m.ServiceContextID)
		copy(dAtA[i:], m.ServiceContextID)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.ServiceContextID)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ServiceFeeCap) > 0 {
		for iNdEx := len(m.ServiceFeeCap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceFeeCap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRandom(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Oracle {
		i--
		if m.Oracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintRandom(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRandom(dAtA []byte, offset int, v uint64) int {
	offset -= sovRandom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Random) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestTxHash)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovRandom(uint64(m.Height))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovRandom(uint64(m.Height))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	if m.Oracle {
		n += 2
	}
	if len(m.ServiceFeeCap) > 0 {
		for _, e := range m.ServiceFeeCap {
			l = e.Size()
			n += 1 + l + sovRandom(uint64(l))
		}
	}
	l = len(m.ServiceContextID)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func sovRandom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRandom(x uint64) (n int) {
	return sovRandom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Random) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: Random: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Random: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Oracle = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFeeCap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceFeeCap = append(m.ServiceFeeCap, types.Coin{})
			if err := m.ServiceFeeCap[len(m.ServiceFeeCap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceContextID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceContextID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func skipRandom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
				return 0, ErrInvalidLengthRandom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRandom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRandom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRandom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRandom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRandom = fmt.Errorf("proto: unexpected end of group")
)
