// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v094/attribute_of_schema.proto

package v094

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

type AttributeOfSchema struct {
	NftSchemaCode    string             `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	SchemaAttributes []*SchemaAttribute `protobuf:"bytes,2,rep,name=schemaAttributes,proto3" json:"schemaAttributes,omitempty"`
}

func (m *AttributeOfSchema) Reset()         { *m = AttributeOfSchema{} }
func (m *AttributeOfSchema) String() string { return proto.CompactTextString(m) }
func (*AttributeOfSchema) ProtoMessage()    {}
func (*AttributeOfSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_585da7ba753e60bb, []int{0}
}
func (m *AttributeOfSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttributeOfSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttributeOfSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttributeOfSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeOfSchema.Merge(m, src)
}
func (m *AttributeOfSchema) XXX_Size() int {
	return m.Size()
}
func (m *AttributeOfSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeOfSchema.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeOfSchema proto.InternalMessageInfo

func (m *AttributeOfSchema) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *AttributeOfSchema) GetSchemaAttributes() []*SchemaAttribute {
	if m != nil {
		return m.SchemaAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*AttributeOfSchema)(nil), "thesixnetwork.sixprotocol.nftmngr.v094.AttributeOfSchema")
}

func init() {
	proto.RegisterFile("nftmngr/v094/attribute_of_schema.proto", fileDescriptor_585da7ba753e60bb)
}

var fileDescriptor_585da7ba753e60bb = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x2f, 0x33, 0xb0, 0x34, 0xd1, 0x4f, 0x2c, 0x29, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0x49, 0x8d, 0xcf, 0x4f, 0x8b, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x52, 0x2b, 0xc9, 0x48, 0x2d, 0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f,
	0xca, 0xd6, 0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27, 0xe7, 0xe7, 0xe8, 0x41, 0x4d, 0xd0, 0x03, 0x99,
	0x20, 0xa5, 0x8c, 0x62, 0x1e, 0xc4, 0x88, 0x78, 0xb8, 0xb1, 0x10, 0xc3, 0x94, 0xe6, 0x31, 0x72,
	0x09, 0x3a, 0xc2, 0xc4, 0xfc, 0xd3, 0x82, 0xc1, 0xaa, 0x84, 0x54, 0xb8, 0x78, 0xf3, 0xd2, 0x4a,
	0x20, 0x1c, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x54, 0x41, 0xa1,
	0x64, 0x2e, 0x01, 0x88, 0xa9, 0x70, 0x03, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xcc,
	0xf5, 0x88, 0x73, 0xa3, 0x5e, 0x30, 0xaa, 0xfe, 0x20, 0x0c, 0x03, 0x9d, 0x42, 0x4e, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2a, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x1f, 0xc5, 0x3a, 0xfd, 0xe2, 0xcc, 0x0a, 0x5d, 0x98, 0x7d, 0xfa, 0x15, 0xfa,
	0xb0, 0x70, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x06, 0x87, 0x46, 0x12, 0x1b, 0x58, 0xda, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0xf9, 0xbc, 0x30, 0x74, 0x01, 0x00, 0x00,
}

func (m *AttributeOfSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttributeOfSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttributeOfSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SchemaAttributes) > 0 {
		for iNdEx := len(m.SchemaAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SchemaAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAttributeOfSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintAttributeOfSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttributeOfSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttributeOfSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AttributeOfSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovAttributeOfSchema(uint64(l))
	}
	if len(m.SchemaAttributes) > 0 {
		for _, e := range m.SchemaAttributes {
			l = e.Size()
			n += 1 + l + sovAttributeOfSchema(uint64(l))
		}
	}
	return n
}

func sovAttributeOfSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttributeOfSchema(x uint64) (n int) {
	return sovAttributeOfSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AttributeOfSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttributeOfSchema
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
			return fmt.Errorf("proto: AttributeOfSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttributeOfSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeOfSchema
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
				return ErrInvalidLengthAttributeOfSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeOfSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttributeOfSchema
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
				return ErrInvalidLengthAttributeOfSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttributeOfSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaAttributes = append(m.SchemaAttributes, &SchemaAttribute{})
			if err := m.SchemaAttributes[len(m.SchemaAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttributeOfSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttributeOfSchema
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
func skipAttributeOfSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttributeOfSchema
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
					return 0, ErrIntOverflowAttributeOfSchema
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
					return 0, ErrIntOverflowAttributeOfSchema
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
				return 0, ErrInvalidLengthAttributeOfSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttributeOfSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttributeOfSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttributeOfSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttributeOfSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttributeOfSchema = fmt.Errorf("proto: unexpected end of group")
)
