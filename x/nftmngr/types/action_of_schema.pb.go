// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/action_of_schema.proto

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

type ActionOfSchema struct {
	NftSchemaCode string `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Index         uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *ActionOfSchema) Reset()         { *m = ActionOfSchema{} }
func (m *ActionOfSchema) String() string { return proto.CompactTextString(m) }
func (*ActionOfSchema) ProtoMessage()    {}
func (*ActionOfSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_43d1ce2c7a79358b, []int{0}
}

func (m *ActionOfSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *ActionOfSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionOfSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *ActionOfSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionOfSchema.Merge(m, src)
}

func (m *ActionOfSchema) XXX_Size() int {
	return m.Size()
}

func (m *ActionOfSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionOfSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ActionOfSchema proto.InternalMessageInfo

func (m *ActionOfSchema) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *ActionOfSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionOfSchema) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*ActionOfSchema)(nil), "thesixnetwork.sixprotocol.nftmngr.ActionOfSchema")
}

func init() { proto.RegisterFile("nftmngr/action_of_schema.proto", fileDescriptor_43d1ce2c7a79358b) }

var fileDescriptor_43d1ce2c7a79358b = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4b, 0x2b, 0xc9,
	0xcd, 0x4b, 0x2f, 0xd2, 0x4f, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0x4f, 0x8b, 0x2f, 0x4e,
	0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0xc9, 0x48, 0x2d,
	0xce, 0xac, 0xc8, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0xce, 0xac, 0x00, 0x8b, 0x27,
	0xe7, 0xe7, 0xe8, 0x41, 0x75, 0x2a, 0x25, 0x70, 0xf1, 0x39, 0x82, 0x35, 0xfb, 0xa7, 0x05, 0x83,
	0xb5, 0x0a, 0xa9, 0x70, 0xf1, 0xe6, 0xa5, 0x95, 0x40, 0x38, 0xce, 0xf9, 0x29, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xa8, 0x82, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x4c, 0x60, 0x49, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x25, 0x08, 0xc2, 0x71, 0xf2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xd3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x14,
	0x97, 0xea, 0x17, 0x67, 0x56, 0xe8, 0xc2, 0x9c, 0xaa, 0x5f, 0xa1, 0x0f, 0xf3, 0x66, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x58, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x63, 0x46, 0x02,
	0x52, 0xfe, 0x00, 0x00, 0x00,
}

func (m *ActionOfSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionOfSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionOfSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintActionOfSchema(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintActionOfSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintActionOfSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionOfSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionOfSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *ActionOfSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovActionOfSchema(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovActionOfSchema(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovActionOfSchema(uint64(m.Index))
	}
	return n
}

func sovActionOfSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozActionOfSchema(x uint64) (n int) {
	return sovActionOfSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *ActionOfSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionOfSchema
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
			return fmt.Errorf("proto: ActionOfSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionOfSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionOfSchema
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
				return ErrInvalidLengthActionOfSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionOfSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionOfSchema
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
				return ErrInvalidLengthActionOfSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionOfSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionOfSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActionOfSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionOfSchema
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

func skipActionOfSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionOfSchema
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
					return 0, ErrIntOverflowActionOfSchema
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
					return 0, ErrIntOverflowActionOfSchema
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
				return 0, ErrInvalidLengthActionOfSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionOfSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionOfSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionOfSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionOfSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionOfSchema = fmt.Errorf("proto: unexpected end of group")
)
