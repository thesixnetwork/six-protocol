// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/opensea_display_option.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OpenseaDisplayOption struct {
	DisplayType string `protobuf:"bytes,1,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"`
	TraitType   string `protobuf:"bytes,2,opt,name=trait_type,json=traitType,proto3" json:"trait_type,omitempty"`
	MaxValue    uint64 `protobuf:"varint,3,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (m *OpenseaDisplayOption) Reset()         { *m = OpenseaDisplayOption{} }
func (m *OpenseaDisplayOption) String() string { return proto.CompactTextString(m) }
func (*OpenseaDisplayOption) ProtoMessage()    {}
func (*OpenseaDisplayOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a72f14b2fd8c4d9, []int{0}
}

func (m *OpenseaDisplayOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *OpenseaDisplayOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OpenseaDisplayOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *OpenseaDisplayOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenseaDisplayOption.Merge(m, src)
}

func (m *OpenseaDisplayOption) XXX_Size() int {
	return m.Size()
}

func (m *OpenseaDisplayOption) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenseaDisplayOption.DiscardUnknown(m)
}

var xxx_messageInfo_OpenseaDisplayOption proto.InternalMessageInfo

func (m *OpenseaDisplayOption) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func (m *OpenseaDisplayOption) GetTraitType() string {
	if m != nil {
		return m.TraitType
	}
	return ""
}

func (m *OpenseaDisplayOption) GetMaxValue() uint64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func init() {
	proto.RegisterType((*OpenseaDisplayOption)(nil), "thesixnetwork.sixprotocol.nftmngr.OpenseaDisplayOption")
}

func init() {
	proto.RegisterFile("nftmngr/opensea_display_option.proto", fileDescriptor_0a72f14b2fd8c4d9)
}

var fileDescriptor_0a72f14b2fd8c4d9 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0xcf, 0x2f, 0x48, 0xcd, 0x2b, 0x4e, 0x4d, 0x8c, 0x4f, 0xc9, 0x2c, 0x2e,
	0xc8, 0x49, 0xac, 0x8c, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6,
	0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0x41, 0xf5, 0x2b, 0x95, 0x72, 0x89, 0xf8,
	0x43, 0x8c, 0x70, 0x81, 0x98, 0xe0, 0x0f, 0x36, 0x40, 0x48, 0x91, 0x8b, 0x07, 0x66, 0x64, 0x49,
	0x65, 0x41, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x37, 0x54, 0x2c, 0xa4, 0xb2, 0x20,
	0x55, 0x48, 0x96, 0x8b, 0xab, 0xa4, 0x28, 0x31, 0xb3, 0x04, 0xa2, 0x80, 0x09, 0xac, 0x80, 0x13,
	0x2c, 0x02, 0x96, 0x96, 0xe6, 0xe2, 0xcc, 0x4d, 0xac, 0x88, 0x2f, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x56, 0x60, 0xd4, 0x60, 0x09, 0xe2, 0xc8, 0x4d, 0xac, 0x08, 0x03, 0xf1, 0x9d, 0xfc, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xc5, 0xf9, 0xfa, 0xc5, 0x99, 0x15, 0xba, 0x30, 0xf7, 0xeb,
	0x57, 0xe8, 0xc3, 0x42, 0x00, 0x64, 0x7b, 0x71, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x46, 0x5e, 0x13, 0x65, 0x19, 0x01, 0x00, 0x00,
}

func (m *OpenseaDisplayOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OpenseaDisplayOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OpenseaDisplayOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxValue != 0 {
		i = encodeVarintOpenseaDisplayOption(dAtA, i, uint64(m.MaxValue))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TraitType) > 0 {
		i -= len(m.TraitType)
		copy(dAtA[i:], m.TraitType)
		i = encodeVarintOpenseaDisplayOption(dAtA, i, uint64(len(m.TraitType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DisplayType) > 0 {
		i -= len(m.DisplayType)
		copy(dAtA[i:], m.DisplayType)
		i = encodeVarintOpenseaDisplayOption(dAtA, i, uint64(len(m.DisplayType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpenseaDisplayOption(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpenseaDisplayOption(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *OpenseaDisplayOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DisplayType)
	if l > 0 {
		n += 1 + l + sovOpenseaDisplayOption(uint64(l))
	}
	l = len(m.TraitType)
	if l > 0 {
		n += 1 + l + sovOpenseaDisplayOption(uint64(l))
	}
	if m.MaxValue != 0 {
		n += 1 + sovOpenseaDisplayOption(uint64(m.MaxValue))
	}
	return n
}

func sovOpenseaDisplayOption(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozOpenseaDisplayOption(x uint64) (n int) {
	return sovOpenseaDisplayOption(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *OpenseaDisplayOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpenseaDisplayOption
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
			return fmt.Errorf("proto: OpenseaDisplayOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OpenseaDisplayOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenseaDisplayOption
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
				return ErrInvalidLengthOpenseaDisplayOption
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpenseaDisplayOption
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenseaDisplayOption
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
				return ErrInvalidLengthOpenseaDisplayOption
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpenseaDisplayOption
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraitType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValue", wireType)
			}
			m.MaxValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpenseaDisplayOption
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOpenseaDisplayOption(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpenseaDisplayOption
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

func skipOpenseaDisplayOption(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpenseaDisplayOption
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
					return 0, ErrIntOverflowOpenseaDisplayOption
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
					return 0, ErrIntOverflowOpenseaDisplayOption
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
				return 0, ErrInvalidLengthOpenseaDisplayOption
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpenseaDisplayOption
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpenseaDisplayOption
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpenseaDisplayOption        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpenseaDisplayOption          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpenseaDisplayOption = fmt.Errorf("proto: unexpected end of group")
)
