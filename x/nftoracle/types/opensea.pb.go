// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/opensea.proto

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

type Trait struct {
	TraitType   string `protobuf:"bytes,1,opt,name=trait_type,json=traitType,proto3" json:"trait_type,omitempty"`
	Value       string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	DisplayType string `protobuf:"bytes,3,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"`
	MaxValue    string `protobuf:"bytes,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (m *Trait) Reset()         { *m = Trait{} }
func (m *Trait) String() string { return proto.CompactTextString(m) }
func (*Trait) ProtoMessage()    {}
func (*Trait) Descriptor() ([]byte, []int) {
	return fileDescriptor_b592987820f26afc, []int{0}
}

func (m *Trait) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Trait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trait.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Trait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trait.Merge(m, src)
}

func (m *Trait) XXX_Size() int {
	return m.Size()
}

func (m *Trait) XXX_DiscardUnknown() {
	xxx_messageInfo_Trait.DiscardUnknown(m)
}

var xxx_messageInfo_Trait proto.InternalMessageInfo

func (m *Trait) GetTraitType() string {
	if m != nil {
		return m.TraitType
	}
	return ""
}

func (m *Trait) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Trait) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func (m *Trait) GetMaxValue() string {
	if m != nil {
		return m.MaxValue
	}
	return ""
}

func init() {
	proto.RegisterType((*Trait)(nil), "thesixnetwork.sixnft.nftoracle.Trait")
}

func init() { proto.RegisterFile("nftoracle/opensea.proto", fileDescriptor_b592987820f26afc) }

var fileDescriptor_b592987820f26afc = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4b, 0x2b, 0xc9,
	0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0xcf, 0x2f, 0x48, 0xcd, 0x2b, 0x4e, 0x4d, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2b, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x03, 0x31, 0xd3, 0x4a, 0xf4, 0xe0, 0xaa, 0x95, 0x6a, 0xb8, 0x58, 0x43, 0x8a,
	0x12, 0x33, 0x4b, 0x84, 0x64, 0xb9, 0xb8, 0x4a, 0x40, 0x8c, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x48, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x08, 0x17,
	0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x13, 0x58, 0x06, 0xc2, 0x11, 0x52, 0xe4, 0xe2, 0x49,
	0xc9, 0x2c, 0x2e, 0xc8, 0x49, 0xac, 0x84, 0x68, 0x63, 0x06, 0x4b, 0x72, 0x43, 0xc5, 0xc0, 0x1a,
	0xa5, 0xb9, 0x38, 0x73, 0x13, 0x2b, 0xe2, 0x21, 0x9a, 0x59, 0xc0, 0xf2, 0x1c, 0xb9, 0x89, 0x15,
	0x61, 0x20, 0xbe, 0x53, 0xe0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99,
	0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa3, 0x78, 0x41, 0xbf, 0x38,
	0xb3, 0x42, 0x17, 0xec, 0xb7, 0xe4, 0xfc, 0x1c, 0xfd, 0x0a, 0x7d, 0x84, 0xc7, 0x41, 0x2e, 0x28,
	0x4e, 0x62, 0x03, 0xcb, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x93, 0x79, 0xb8, 0x12,
	0x01, 0x00, 0x00,
}

func (m *Trait) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trait) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trait) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MaxValue) > 0 {
		i -= len(m.MaxValue)
		copy(dAtA[i:], m.MaxValue)
		i = encodeVarintOpensea(dAtA, i, uint64(len(m.MaxValue)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DisplayType) > 0 {
		i -= len(m.DisplayType)
		copy(dAtA[i:], m.DisplayType)
		i = encodeVarintOpensea(dAtA, i, uint64(len(m.DisplayType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintOpensea(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TraitType) > 0 {
		i -= len(m.TraitType)
		copy(dAtA[i:], m.TraitType)
		i = encodeVarintOpensea(dAtA, i, uint64(len(m.TraitType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpensea(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpensea(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Trait) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TraitType)
	if l > 0 {
		n += 1 + l + sovOpensea(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovOpensea(uint64(l))
	}
	l = len(m.DisplayType)
	if l > 0 {
		n += 1 + l + sovOpensea(uint64(l))
	}
	l = len(m.MaxValue)
	if l > 0 {
		n += 1 + l + sovOpensea(uint64(l))
	}
	return n
}

func sovOpensea(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozOpensea(x uint64) (n int) {
	return sovOpensea(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Trait) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpensea
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
			return fmt.Errorf("proto: Trait: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trait: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpensea
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
				return ErrInvalidLengthOpensea
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpensea
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraitType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpensea
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
				return ErrInvalidLengthOpensea
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpensea
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpensea
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
				return ErrInvalidLengthOpensea
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpensea
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpensea
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
				return ErrInvalidLengthOpensea
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpensea
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpensea(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpensea
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

func skipOpensea(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpensea
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
					return 0, ErrIntOverflowOpensea
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
					return 0, ErrIntOverflowOpensea
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
				return 0, ErrInvalidLengthOpensea
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpensea
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpensea
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpensea        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpensea          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpensea = fmt.Errorf("proto: unexpected end of group")
)
