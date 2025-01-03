// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/burn.proto

package types

import (
	fmt "fmt"
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

type Burn struct {
	Id      uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount  types.Coin `protobuf:"bytes,3,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *Burn) Reset()         { *m = Burn{} }
func (m *Burn) String() string { return proto.CompactTextString(m) }
func (*Burn) ProtoMessage()    {}
func (*Burn) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b4d0bda0c6c58ac, []int{0}
}
func (m *Burn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Burn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Burn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Burn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Burn.Merge(m, src)
}
func (m *Burn) XXX_Size() int {
	return m.Size()
}
func (m *Burn) XXX_DiscardUnknown() {
	xxx_messageInfo_Burn.DiscardUnknown(m)
}

var xxx_messageInfo_Burn proto.InternalMessageInfo

func (m *Burn) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Burn) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Burn) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Burn)(nil), "thesixnetwork.sixprotocol.tokenmngr.Burn")
}

func init() { proto.RegisterFile("tokenmngr/burn.proto", fileDescriptor_9b4d0bda0c6c58ac) }

var fileDescriptor_9b4d0bda0c6c58ac = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0x7c, 0x55, 0x3f, 0x11, 0x24, 0x86, 0xa8, 0x43, 0xe8, 0xe0, 0x46, 0x30, 0x90,
	0xa5, 0xbe, 0x2a, 0x0c, 0xec, 0xe1, 0x09, 0xc8, 0xc8, 0x16, 0x27, 0x56, 0x6a, 0x85, 0xf8, 0x56,
	0xb6, 0x03, 0xe1, 0x2d, 0x18, 0x78, 0x0a, 0x9e, 0xa4, 0x63, 0x47, 0x26, 0x40, 0xc9, 0x8b, 0xa0,
	0x26, 0x29, 0x7f, 0x26, 0xdb, 0xc7, 0xf7, 0xfc, 0xce, 0xd5, 0xf1, 0x66, 0x16, 0x4b, 0xa1, 0x2a,
	0x55, 0x68, 0xe0, 0xb5, 0x56, 0x6c, 0xa3, 0xd1, 0xa2, 0x7f, 0x6e, 0xd7, 0xc2, 0xc8, 0x46, 0x09,
	0xfb, 0x88, 0xba, 0x64, 0x46, 0x36, 0xbd, 0x9e, 0xe1, 0x3d, 0xfb, 0x9e, 0x9f, 0xcf, 0x0a, 0x2c,
	0xb0, 0xd7, 0x61, 0x7f, 0x1b, 0xac, 0x73, 0x9a, 0xa1, 0xa9, 0xd0, 0x00, 0x4f, 0x8d, 0x80, 0x87,
	0x15, 0x17, 0x36, 0x5d, 0x41, 0x86, 0x72, 0x44, 0x9f, 0xbd, 0x10, 0x6f, 0x12, 0xd7, 0x5a, 0xf9,
	0x27, 0x9e, 0x2b, 0xf3, 0x80, 0x84, 0x24, 0x9a, 0x24, 0xae, 0xcc, 0xfd, 0xc0, 0xfb, 0x9f, 0x69,
	0x91, 0x5a, 0xd4, 0x81, 0x1b, 0x92, 0xe8, 0x28, 0x39, 0x3c, 0x7d, 0xee, 0x4d, 0xd3, 0x0a, 0x6b,
	0x65, 0x83, 0x7f, 0x21, 0x89, 0x8e, 0x2f, 0x4f, 0xd9, 0x90, 0xc1, 0xf6, 0x19, 0x6c, 0xcc, 0x60,
	0x37, 0x28, 0x55, 0x0c, 0xdb, 0xf7, 0x85, 0xf3, 0xfa, 0xb1, 0xb8, 0x28, 0xa4, 0x5d, 0xd7, 0x9c,
	0x65, 0x58, 0xc1, 0xb8, 0xd0, 0x70, 0x2c, 0x4d, 0x5e, 0x82, 0x7d, 0xda, 0x08, 0xd3, 0x1b, 0x92,
	0x91, 0x1c, 0xdf, 0x6e, 0x5b, 0x4a, 0x76, 0x2d, 0x25, 0x9f, 0x2d, 0x25, 0xcf, 0x1d, 0x75, 0x76,
	0x1d, 0x75, 0xde, 0x3a, 0xea, 0xdc, 0x5d, 0xff, 0x42, 0xfd, 0xa9, 0x05, 0x8c, 0x6c, 0x96, 0x87,
	0x5e, 0xa0, 0x81, 0x9f, 0x26, 0x7b, 0x3e, 0x9f, 0xf6, 0x7f, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x24, 0x8a, 0x6d, 0x13, 0x63, 0x01, 0x00, 0x00,
}

func (m *Burn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Burn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Burn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBurn(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBurn(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBurn(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBurn(dAtA []byte, offset int, v uint64) int {
	offset -= sovBurn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Burn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBurn(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBurn(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovBurn(uint64(l))
	return n
}

func sovBurn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBurn(x uint64) (n int) {
	return sovBurn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Burn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBurn
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
			return fmt.Errorf("proto: Burn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Burn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurn
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
				return ErrInvalidLengthBurn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBurn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurn
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
				return ErrInvalidLengthBurn
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBurn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBurn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBurn
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
func skipBurn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBurn
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
					return 0, ErrIntOverflowBurn
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
					return 0, ErrIntOverflowBurn
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
				return 0, ErrInvalidLengthBurn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBurn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBurn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBurn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBurn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBurn = fmt.Errorf("proto: unexpected end of group")
)
