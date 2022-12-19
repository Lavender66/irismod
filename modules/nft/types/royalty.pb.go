// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/royalty.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type RoyaltyInfo struct {
	Receiver        string                 `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	RoyaltyFraction cosmossdk_io_math.Uint `protobuf:"bytes,2,opt,name=royalty_fraction,json=royaltyFraction,proto3,customtype=cosmossdk.io/math.Uint" json:"royalty_fraction"`
}

func (m *RoyaltyInfo) Reset()         { *m = RoyaltyInfo{} }
func (m *RoyaltyInfo) String() string { return proto.CompactTextString(m) }
func (*RoyaltyInfo) ProtoMessage()    {}
func (*RoyaltyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c34aacedb01c5b, []int{0}
}
func (m *RoyaltyInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoyaltyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoyaltyInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoyaltyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoyaltyInfo.Merge(m, src)
}
func (m *RoyaltyInfo) XXX_Size() int {
	return m.Size()
}
func (m *RoyaltyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoyaltyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoyaltyInfo proto.InternalMessageInfo

type RoyaltyMetadata struct {
	Enabled            bool         `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DefaultRoyaltyInfo *RoyaltyInfo `protobuf:"bytes,2,opt,name=defaultRoyaltyInfo,proto3" json:"defaultRoyaltyInfo,omitempty"`
}

func (m *RoyaltyMetadata) Reset()         { *m = RoyaltyMetadata{} }
func (m *RoyaltyMetadata) String() string { return proto.CompactTextString(m) }
func (*RoyaltyMetadata) ProtoMessage()    {}
func (*RoyaltyMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c34aacedb01c5b, []int{1}
}
func (m *RoyaltyMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoyaltyMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoyaltyMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoyaltyMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoyaltyMetadata.Merge(m, src)
}
func (m *RoyaltyMetadata) XXX_Size() int {
	return m.Size()
}
func (m *RoyaltyMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RoyaltyMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RoyaltyMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RoyaltyInfo)(nil), "irismod.nft.royalty.RoyaltyInfo")
	proto.RegisterType((*RoyaltyMetadata)(nil), "irismod.nft.royalty.RoyaltyMetadata")
}

func init() { proto.RegisterFile("nft/royalty.proto", fileDescriptor_80c34aacedb01c5b) }

var fileDescriptor_80c34aacedb01c5b = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4a, 0x33, 0x41,
	0x14, 0x85, 0x77, 0x7f, 0x7e, 0x34, 0x4e, 0x8a, 0xe8, 0x2a, 0x12, 0x53, 0x4c, 0x42, 0x2a, 0x0b,
	0x99, 0x21, 0xda, 0x59, 0xa6, 0x10, 0x22, 0x08, 0xb2, 0x60, 0x63, 0x13, 0x66, 0x77, 0x67, 0x93,
	0xc1, 0xdd, 0xb9, 0x61, 0xe7, 0x26, 0x10, 0x5b, 0x5f, 0xc0, 0x47, 0xf0, 0x71, 0x52, 0xa6, 0x14,
	0x8b, 0xa0, 0x49, 0xe3, 0x63, 0xc8, 0xee, 0x8c, 0x92, 0x22, 0xdd, 0x5c, 0xbe, 0x73, 0xee, 0x39,
	0xdc, 0x21, 0x47, 0x3a, 0x45, 0x5e, 0xc0, 0x5c, 0x64, 0x38, 0x67, 0x93, 0x02, 0x10, 0x82, 0x63,
	0x55, 0x28, 0x93, 0x43, 0xc2, 0x74, 0x8a, 0xcc, 0xa1, 0xd6, 0xc9, 0x08, 0x46, 0x50, 0x71, 0x5e,
	0xbe, 0xac, 0xb4, 0x45, 0x63, 0x30, 0x39, 0x18, 0x1e, 0x09, 0x23, 0xf9, 0xac, 0x17, 0x49, 0x14,
	0x3d, 0x1e, 0x83, 0xd2, 0x8e, 0x9f, 0x59, 0x3e, 0xb4, 0x46, 0x3b, 0x58, 0xd4, 0x7d, 0x26, 0xf5,
	0xd0, 0xee, 0x1e, 0xe8, 0x14, 0x82, 0x16, 0xa9, 0x15, 0x32, 0x96, 0x6a, 0x26, 0x8b, 0xa6, 0xdf,
	0xf1, 0xcf, 0x0f, 0xc2, 0xbf, 0x39, 0x18, 0x90, 0x43, 0x57, 0x63, 0x98, 0x16, 0x22, 0x46, 0x05,
	0xba, 0xf9, 0xaf, 0xd4, 0xf4, 0xe9, 0x62, 0xd5, 0xf6, 0x3e, 0x56, 0xed, 0x53, 0xbb, 0xda, 0x24,
	0x4f, 0x4c, 0x01, 0xcf, 0x05, 0x8e, 0xd9, 0x83, 0xd2, 0x18, 0x36, 0x9c, 0xef, 0xc6, 0xd9, 0xae,
	0xff, 0x7f, 0xbf, 0xb5, 0xfd, 0xee, 0x8b, 0x4f, 0x1a, 0x2e, 0xfc, 0x4e, 0xa2, 0x48, 0x04, 0x8a,
	0xa0, 0x49, 0xf6, 0xa5, 0x16, 0x51, 0x26, 0x93, 0x2a, 0xbf, 0x16, 0xfe, 0x8e, 0xc1, 0x3d, 0x09,
	0x12, 0x99, 0x8a, 0x69, 0x86, 0x5b, 0x85, 0xab, 0x02, 0xf5, 0xcb, 0x0e, 0xdb, 0x71, 0x2c, 0xb6,
	0xa5, 0x0b, 0x77, 0x78, 0x6d, 0x8b, 0xfe, 0xed, 0xe2, 0x8b, 0x7a, 0x8b, 0x35, 0xf5, 0x97, 0x6b,
	0xea, 0x7f, 0xae, 0xa9, 0xff, 0xba, 0xa1, 0xde, 0x72, 0x43, 0xbd, 0xf7, 0x0d, 0xf5, 0x1e, 0x2f,
	0x46, 0x0a, 0xc7, 0xd3, 0x88, 0xc5, 0x90, 0xf3, 0x32, 0x43, 0x4b, 0xe4, 0x2e, 0x8b, 0xe7, 0x90,
	0x4c, 0x33, 0x69, 0x78, 0xf9, 0x77, 0x38, 0x9f, 0x48, 0x13, 0xed, 0x55, 0x47, 0xbd, 0xfa, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0xfc, 0x6c, 0xaa, 0xcf, 0x01, 0x00, 0x00,
}

func (this *RoyaltyInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoyaltyInfo)
	if !ok {
		that2, ok := that.(RoyaltyInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Receiver != that1.Receiver {
		return false
	}
	if !this.RoyaltyFraction.Equal(that1.RoyaltyFraction) {
		return false
	}
	return true
}
func (this *RoyaltyMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoyaltyMetadata)
	if !ok {
		that2, ok := that.(RoyaltyMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if !this.DefaultRoyaltyInfo.Equal(that1.DefaultRoyaltyInfo) {
		return false
	}
	return true
}
func (m *RoyaltyInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoyaltyInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoyaltyInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RoyaltyFraction.Size()
		i -= size
		if _, err := m.RoyaltyFraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRoyalty(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintRoyalty(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoyaltyMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoyaltyMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoyaltyMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultRoyaltyInfo != nil {
		{
			size, err := m.DefaultRoyaltyInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoyalty(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoyalty(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoyalty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoyaltyInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovRoyalty(uint64(l))
	}
	l = m.RoyaltyFraction.Size()
	n += 1 + l + sovRoyalty(uint64(l))
	return n
}

func (m *RoyaltyMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.DefaultRoyaltyInfo != nil {
		l = m.DefaultRoyaltyInfo.Size()
		n += 1 + l + sovRoyalty(uint64(l))
	}
	return n
}

func sovRoyalty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoyalty(x uint64) (n int) {
	return sovRoyalty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoyaltyInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoyalty
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
			return fmt.Errorf("proto: RoyaltyInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoyaltyInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoyalty
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
				return ErrInvalidLengthRoyalty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoyalty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoyaltyFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoyalty
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
				return ErrInvalidLengthRoyalty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoyalty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RoyaltyFraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoyalty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoyalty
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
func (m *RoyaltyMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoyalty
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
			return fmt.Errorf("proto: RoyaltyMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoyaltyMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoyalty
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultRoyaltyInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoyalty
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
				return ErrInvalidLengthRoyalty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoyalty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultRoyaltyInfo == nil {
				m.DefaultRoyaltyInfo = &RoyaltyInfo{}
			}
			if err := m.DefaultRoyaltyInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoyalty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoyalty
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
func skipRoyalty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoyalty
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
					return 0, ErrIntOverflowRoyalty
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
					return 0, ErrIntOverflowRoyalty
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
				return 0, ErrInvalidLengthRoyalty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoyalty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoyalty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoyalty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoyalty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoyalty = fmt.Errorf("proto: unexpected end of group")
)
