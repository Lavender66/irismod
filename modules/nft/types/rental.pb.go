// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/rental.proto

package types

import (
	fmt "fmt"
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

// RentalInfo defines a rental info
type RentalInfo struct {
	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	NftId   string `protobuf:"bytes,3,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Expires int64  `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (m *RentalInfo) Reset()         { *m = RentalInfo{} }
func (m *RentalInfo) String() string { return proto.CompactTextString(m) }
func (*RentalInfo) ProtoMessage()    {}
func (*RentalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e57beeaf512eaaa2, []int{0}
}
func (m *RentalInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RentalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RentalInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RentalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentalInfo.Merge(m, src)
}
func (m *RentalInfo) XXX_Size() int {
	return m.Size()
}
func (m *RentalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RentalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RentalInfo proto.InternalMessageInfo

func (m *RentalInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *RentalInfo) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *RentalInfo) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *RentalInfo) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*RentalInfo)(nil), "irismod.rental.RentalInfo")
}

func init() { proto.RegisterFile("nft/rental.proto", fileDescriptor_e57beeaf512eaaa2) }

var fileDescriptor_e57beeaf512eaaa2 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4b, 0x2b, 0xd1,
	0x2f, 0x4a, 0xcd, 0x2b, 0x49, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0x2c,
	0xca, 0x2c, 0xce, 0xcd, 0x4f, 0xd1, 0x83, 0x88, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xa5,
	0xf4, 0x41, 0x2c, 0x88, 0x2a, 0xa5, 0x1c, 0x2e, 0xae, 0x20, 0xb0, 0xbc, 0x67, 0x5e, 0x5a, 0xbe,
	0x90, 0x10, 0x17, 0x4b, 0x69, 0x71, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0x2d, 0x24, 0xc9, 0xc5, 0x91, 0x9c, 0x93, 0x58, 0x5c, 0x1c, 0x9f, 0x99, 0x22, 0xc1, 0x04, 0x16,
	0x67, 0x07, 0xf3, 0x3d, 0x53, 0x84, 0x44, 0xb9, 0xd8, 0xf2, 0xd2, 0x4a, 0x40, 0x12, 0xcc, 0x60,
	0x09, 0xd6, 0xbc, 0xb4, 0x12, 0xcf, 0x14, 0x21, 0x09, 0x2e, 0xf6, 0xd4, 0x8a, 0x82, 0xcc, 0xa2,
	0xd4, 0x62, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x18, 0xd7, 0xc9, 0xed, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x41, 0x0e, 0xcf, 0x4b, 0x2d, 0xd1, 0x87, 0x7a, 0x40, 0x3f, 0x37, 0x3f, 0xa5,
	0x34, 0x27, 0xb5, 0x58, 0x1f, 0xe4, 0xc5, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xe3,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf2, 0x50, 0x68, 0xf6, 0x00, 0x00, 0x00,
}

func (m *RentalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RentalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RentalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expires != 0 {
		i = encodeVarintRental(dAtA, i, uint64(m.Expires))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintRental(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintRental(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintRental(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRental(dAtA []byte, offset int, v uint64) int {
	offset -= sovRental(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RentalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovRental(uint64(m.Expires))
	}
	return n
}

func sovRental(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRental(x uint64) (n int) {
	return sovRental(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RentalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRental
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
			return fmt.Errorf("proto: RentalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RentalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRental(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRental
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
func skipRental(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
				return 0, ErrInvalidLengthRental
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRental
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRental
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRental        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRental          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRental = fmt.Errorf("proto: unexpected end of group")
)
