// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/disable_virtual_schema.proto

package types

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

type DisableVirtualSchema struct {
	NftSchemaCode string `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Creator       string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *DisableVirtualSchema) Reset()         { *m = DisableVirtualSchema{} }
func (m *DisableVirtualSchema) String() string { return proto.CompactTextString(m) }
func (*DisableVirtualSchema) ProtoMessage()    {}
func (*DisableVirtualSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{0}
}
func (m *DisableVirtualSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisableVirtualSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisableVirtualSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisableVirtualSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableVirtualSchema.Merge(m, src)
}
func (m *DisableVirtualSchema) XXX_Size() int {
	return m.Size()
}
func (m *DisableVirtualSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableVirtualSchema.DiscardUnknown(m)
}

var xxx_messageInfo_DisableVirtualSchema proto.InternalMessageInfo

func (m *DisableVirtualSchema) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *DisableVirtualSchema) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*DisableVirtualSchema)(nil), "thesixnetwork.sixprotocol.nftmngr.DisableVirtualSchema")
}

func init() {
	proto.RegisterFile("nftmngr/disable_virtual_schema.proto", fileDescriptor_d2dd6dee3450cf76)
}

var fileDescriptor_d2dd6dee3450cf76 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x4f, 0xc9, 0x2c, 0x4e, 0x4c, 0xca, 0x49, 0x8d, 0x2f, 0xcb, 0x2c, 0x2a,
	0x29, 0x4d, 0xcc, 0x89, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x2c, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6,
	0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0x41, 0xf5, 0x2b, 0x85, 0x71, 0x89, 0xb8,
	0x40, 0x8c, 0x08, 0x83, 0x98, 0x10, 0x0c, 0x36, 0x40, 0x48, 0x85, 0x8b, 0x37, 0x2f, 0xad, 0x04,
	0xc2, 0x71, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x15, 0x14, 0x92,
	0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x02, 0xcb, 0xc3, 0xb8, 0x4e,
	0xfe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe2, 0x3e, 0xfd, 0xe2, 0xcc, 0x0a, 0x5d, 0x98,
	0x03, 0xf5, 0x2b, 0xf4, 0x61, 0x5e, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xcb, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0x00, 0x9c, 0xa4, 0xfa, 0x00, 0x00, 0x00,
}

func (m *DisableVirtualSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisableVirtualSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisableVirtualSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDisableVirtualSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovDisableVirtualSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DisableVirtualSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	return n
}

func sovDisableVirtualSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDisableVirtualSchema(x uint64) (n int) {
	return sovDisableVirtualSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DisableVirtualSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDisableVirtualSchema
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
			return fmt.Errorf("proto: DisableVirtualSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisableVirtualSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDisableVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDisableVirtualSchema
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
func skipDisableVirtualSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDisableVirtualSchema
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
					return 0, ErrIntOverflowDisableVirtualSchema
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
					return 0, ErrIntOverflowDisableVirtualSchema
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
				return 0, ErrInvalidLengthDisableVirtualSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDisableVirtualSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDisableVirtualSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDisableVirtualSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDisableVirtualSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDisableVirtualSchema = fmt.Errorf("proto: unexpected end of group")
)
