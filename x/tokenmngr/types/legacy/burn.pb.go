// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmngr/v1/burn.proto

package legacy

import (
	fmt "fmt"
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
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Amount  uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Burn) Reset()         { *m = Burn{} }
func (m *Burn) String() string { return proto.CompactTextString(m) }
func (*Burn) ProtoMessage()    {}
func (*Burn) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebcc8fa7c9319f5, []int{0}
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

func (m *Burn) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Burn) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Burn)(nil), "thesixnetwork.sixprotocol.tokenmngr.Burn")
}

func init() { proto.RegisterFile("tokenmngr/v1/burn.proto", fileDescriptor_6ebcc8fa7c9319f5) }

var fileDescriptor_6ebcc8fa7c9319f5 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc9, 0xcf, 0x4e,
	0xcd, 0xcb, 0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2a, 0x2d, 0xca, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0xc1, 0xb5, 0x28, 0xc5,
	0x71, 0xb1, 0x38, 0x95, 0x16, 0xe5, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x04, 0x31, 0x65, 0xa6, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4,
	0x17, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x60, 0xed,
	0x12, 0xcc, 0x60, 0x71, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44,
	0x82, 0x05, 0x6c, 0x06, 0x94, 0xe7, 0x14, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x56, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x28, 0x2e,
	0xd5, 0x2f, 0xce, 0xac, 0xd0, 0x85, 0x39, 0x55, 0xbf, 0x42, 0x1f, 0xe1, 0xbf, 0x92, 0xca, 0x82,
	0xd4, 0x62, 0xfd, 0x32, 0xc3, 0x24, 0x36, 0xb0, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0x1e, 0xa3, 0x12, 0xfc, 0x00, 0x00, 0x00,
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
	if m.Amount != 0 {
		i = encodeVarintBurn(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintBurn(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
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
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovBurn(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovBurn(uint64(m.Amount))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
