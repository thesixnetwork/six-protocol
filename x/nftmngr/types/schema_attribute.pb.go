// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/schema_attribute.proto

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

type SchemaAttribute struct {
	NftSchemaCode string                `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Name          string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DataType      string                `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	CurrentValue  *SchemaAttributeValue `protobuf:"bytes,4,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	Creator       string                `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SchemaAttribute) Reset()         { *m = SchemaAttribute{} }
func (m *SchemaAttribute) String() string { return proto.CompactTextString(m) }
func (*SchemaAttribute) ProtoMessage()    {}
func (*SchemaAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_8228233d6832e497, []int{0}
}

func (m *SchemaAttribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SchemaAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SchemaAttribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SchemaAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaAttribute.Merge(m, src)
}

func (m *SchemaAttribute) XXX_Size() int {
	return m.Size()
}

func (m *SchemaAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaAttribute proto.InternalMessageInfo

func (m *SchemaAttribute) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *SchemaAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaAttribute) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *SchemaAttribute) GetCurrentValue() *SchemaAttributeValue {
	if m != nil {
		return m.CurrentValue
	}
	return nil
}

func (m *SchemaAttribute) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type SchemaAttributeValue struct {
	// Types that are valid to be assigned to Value:
	//	*SchemaAttributeValue_NumberAttributeValue
	//	*SchemaAttributeValue_StringAttributeValue
	//	*SchemaAttributeValue_BooleanAttributeValue
	//	*SchemaAttributeValue_FloatAttributeValue
	Value isSchemaAttributeValue_Value `protobuf_oneof:"value"`
}

func (m *SchemaAttributeValue) Reset()         { *m = SchemaAttributeValue{} }
func (m *SchemaAttributeValue) String() string { return proto.CompactTextString(m) }
func (*SchemaAttributeValue) ProtoMessage()    {}
func (*SchemaAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8228233d6832e497, []int{1}
}

func (m *SchemaAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *SchemaAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SchemaAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *SchemaAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaAttributeValue.Merge(m, src)
}

func (m *SchemaAttributeValue) XXX_Size() int {
	return m.Size()
}

func (m *SchemaAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaAttributeValue proto.InternalMessageInfo

type isSchemaAttributeValue_Value interface {
	isSchemaAttributeValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SchemaAttributeValue_NumberAttributeValue struct {
	NumberAttributeValue *NumberAttributeValue `protobuf:"bytes,1,opt,name=number_attribute_value,json=numberAttributeValue,proto3,oneof" json:"number_attribute_value,omitempty"`
}
type SchemaAttributeValue_StringAttributeValue struct {
	StringAttributeValue *StringAttributeValue `protobuf:"bytes,2,opt,name=string_attribute_value,json=stringAttributeValue,proto3,oneof" json:"string_attribute_value,omitempty"`
}
type SchemaAttributeValue_BooleanAttributeValue struct {
	BooleanAttributeValue *BooleanAttributeValue `protobuf:"bytes,3,opt,name=boolean_attribute_value,json=booleanAttributeValue,proto3,oneof" json:"boolean_attribute_value,omitempty"`
}
type SchemaAttributeValue_FloatAttributeValue struct {
	FloatAttributeValue *FloatAttributeValue `protobuf:"bytes,4,opt,name=float_attribute_value,json=floatAttributeValue,proto3,oneof" json:"float_attribute_value,omitempty"`
}

func (*SchemaAttributeValue_NumberAttributeValue) isSchemaAttributeValue_Value()  {}
func (*SchemaAttributeValue_StringAttributeValue) isSchemaAttributeValue_Value()  {}
func (*SchemaAttributeValue_BooleanAttributeValue) isSchemaAttributeValue_Value() {}
func (*SchemaAttributeValue_FloatAttributeValue) isSchemaAttributeValue_Value()   {}

func (m *SchemaAttributeValue) GetValue() isSchemaAttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SchemaAttributeValue) GetNumberAttributeValue() *NumberAttributeValue {
	if x, ok := m.GetValue().(*SchemaAttributeValue_NumberAttributeValue); ok {
		return x.NumberAttributeValue
	}
	return nil
}

func (m *SchemaAttributeValue) GetStringAttributeValue() *StringAttributeValue {
	if x, ok := m.GetValue().(*SchemaAttributeValue_StringAttributeValue); ok {
		return x.StringAttributeValue
	}
	return nil
}

func (m *SchemaAttributeValue) GetBooleanAttributeValue() *BooleanAttributeValue {
	if x, ok := m.GetValue().(*SchemaAttributeValue_BooleanAttributeValue); ok {
		return x.BooleanAttributeValue
	}
	return nil
}

func (m *SchemaAttributeValue) GetFloatAttributeValue() *FloatAttributeValue {
	if x, ok := m.GetValue().(*SchemaAttributeValue_FloatAttributeValue); ok {
		return x.FloatAttributeValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SchemaAttributeValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SchemaAttributeValue_NumberAttributeValue)(nil),
		(*SchemaAttributeValue_StringAttributeValue)(nil),
		(*SchemaAttributeValue_BooleanAttributeValue)(nil),
		(*SchemaAttributeValue_FloatAttributeValue)(nil),
	}
}

func init() {
	proto.RegisterType((*SchemaAttribute)(nil), "thesixnetwork.sixprotocol.nftmngr.SchemaAttribute")
	proto.RegisterType((*SchemaAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.SchemaAttributeValue")
}

func init() { proto.RegisterFile("nftmngr/schema_attribute.proto", fileDescriptor_8228233d6832e497) }

var fileDescriptor_8228233d6832e497 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x4b, 0xe3, 0x40,
	0x18, 0xc7, 0x93, 0xbe, 0x6c, 0xb7, 0xb3, 0x5b, 0x16, 0x66, 0xdb, 0xdd, 0xa0, 0x10, 0xda, 0xe2,
	0xa1, 0x17, 0x13, 0x50, 0x7c, 0xb9, 0x5a, 0x41, 0x3c, 0x29, 0x54, 0xf1, 0x20, 0x42, 0x99, 0xa4,
	0x93, 0x36, 0x98, 0xcc, 0x94, 0xc9, 0x13, 0x4d, 0xbf, 0x85, 0x1f, 0xcb, 0x63, 0x8f, 0x1e, 0x4b,
	0xf3, 0x45, 0x24, 0xd3, 0x04, 0x69, 0x13, 0x30, 0xb7, 0xcc, 0xff, 0xe1, 0xf9, 0xfd, 0x9e, 0x99,
	0xcc, 0x20, 0x9d, 0x39, 0xe0, 0xb3, 0xa9, 0x30, 0x03, 0x7b, 0x46, 0x7d, 0x32, 0x26, 0x00, 0xc2,
	0xb5, 0x42, 0xa0, 0xc6, 0x5c, 0x70, 0xe0, 0xb8, 0x07, 0x33, 0x1a, 0xb8, 0x11, 0xa3, 0xf0, 0xca,
	0xc5, 0xb3, 0x11, 0xb8, 0x91, 0xcc, 0x6d, 0xee, 0x19, 0x69, 0xe7, 0x5e, 0x2f, 0x43, 0x30, 0x07,
	0xbe, 0xfa, 0xc7, 0x2f, 0xc4, 0x0b, 0x53, 0x4a, 0x7f, 0xa5, 0xa2, 0x3f, 0x77, 0x52, 0x70, 0x91,
	0xd5, 0xf1, 0x01, 0x6a, 0x31, 0x07, 0x36, 0xe9, 0x25, 0x9f, 0x50, 0x4d, 0xed, 0xaa, 0x83, 0xe6,
	0x68, 0x3b, 0xc4, 0x18, 0xd5, 0x18, 0xf1, 0xa9, 0x56, 0x91, 0x45, 0xf9, 0x8d, 0xf7, 0x51, 0x73,
	0x42, 0x80, 0x8c, 0x61, 0x31, 0xa7, 0x5a, 0x55, 0x16, 0x7e, 0x26, 0xc1, 0xfd, 0x62, 0x4e, 0xf1,
	0x13, 0x6a, 0xd9, 0xa1, 0x10, 0x94, 0xc1, 0x66, 0x02, 0xad, 0xd6, 0x55, 0x07, 0xbf, 0x8e, 0xce,
	0x8c, 0x6f, 0x37, 0x62, 0xec, 0x4c, 0xf8, 0x90, 0xb4, 0x8f, 0x7e, 0xa7, 0x34, 0xb9, 0xc2, 0x1a,
	0x6a, 0xd8, 0x82, 0x12, 0xe0, 0x42, 0xab, 0x4b, 0x71, 0xb6, 0xec, 0xc7, 0x55, 0xd4, 0x2e, 0x02,
	0x60, 0x8e, 0xfe, 0xb1, 0xd0, 0xb7, 0xa8, 0xd8, 0x3d, 0x1b, 0xb9, 0xe1, 0x72, 0x93, 0xdd, 0x48,
	0xc0, 0x36, 0xf8, 0x5a, 0x19, 0xb5, 0x59, 0x41, 0x9e, 0x08, 0x03, 0x10, 0x2e, 0x9b, 0xe6, 0x84,
	0x95, 0xf2, 0x47, 0x21, 0x01, 0x79, 0x61, 0x50, 0x90, 0x63, 0x81, 0xfe, 0x5b, 0x9c, 0x7b, 0x94,
	0xb0, 0x9c, 0xb1, 0x2a, 0x8d, 0xe7, 0x25, 0x8c, 0xc3, 0x0d, 0x21, 0xa7, 0xec, 0x58, 0x45, 0x05,
	0xec, 0xa1, 0x8e, 0xe3, 0x71, 0x92, 0xbb, 0x70, 0xe9, 0xef, 0x3e, 0x2d, 0x61, 0xbc, 0x4a, 0xfa,
	0x73, 0xbe, 0xbf, 0x4e, 0x3e, 0x1e, 0x36, 0x50, 0x5d, 0xd2, 0x87, 0xb7, 0xef, 0x6b, 0x5d, 0x5d,
	0xae, 0x75, 0x75, 0xb5, 0xd6, 0xd5, 0xb7, 0x58, 0x57, 0x96, 0xb1, 0xae, 0x7c, 0xc4, 0xba, 0xf2,
	0x78, 0x32, 0x75, 0x61, 0x16, 0x5a, 0x86, 0xcd, 0x7d, 0x73, 0xcb, 0x6d, 0x06, 0x6e, 0x74, 0x98,
	0xc9, 0xcd, 0xc8, 0xcc, 0x5e, 0x4b, 0x72, 0x7b, 0x03, 0xeb, 0x87, 0xac, 0x1c, 0x7f, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x46, 0x7a, 0x1d, 0xc7, 0x88, 0x03, 0x00, 0x00,
}

func (m *SchemaAttribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaAttribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSchemaAttribute(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CurrentValue != nil {
		{
			size, err := m.CurrentValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchemaAttribute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DataType) > 0 {
		i -= len(m.DataType)
		copy(dAtA[i:], m.DataType)
		i = encodeVarintSchemaAttribute(dAtA, i, uint64(len(m.DataType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchemaAttribute(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintSchemaAttribute(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SchemaAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SchemaAttributeValue_NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttributeValue_NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NumberAttributeValue != nil {
		{
			size, err := m.NumberAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchemaAttribute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SchemaAttributeValue_StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttributeValue_StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringAttributeValue != nil {
		{
			size, err := m.StringAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchemaAttribute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *SchemaAttributeValue_BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttributeValue_BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BooleanAttributeValue != nil {
		{
			size, err := m.BooleanAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchemaAttribute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func (m *SchemaAttributeValue_FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaAttributeValue_FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FloatAttributeValue != nil {
		{
			size, err := m.FloatAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchemaAttribute(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchemaAttribute(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchemaAttribute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *SchemaAttribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	l = len(m.DataType)
	if l > 0 {
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	if m.CurrentValue != nil {
		l = m.CurrentValue.Size()
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	return n
}

func (m *SchemaAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *SchemaAttributeValue_NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberAttributeValue != nil {
		l = m.NumberAttributeValue.Size()
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	return n
}

func (m *SchemaAttributeValue_StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringAttributeValue != nil {
		l = m.StringAttributeValue.Size()
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	return n
}

func (m *SchemaAttributeValue_BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BooleanAttributeValue != nil {
		l = m.BooleanAttributeValue.Size()
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	return n
}

func (m *SchemaAttributeValue_FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FloatAttributeValue != nil {
		l = m.FloatAttributeValue.Size()
		n += 1 + l + sovSchemaAttribute(uint64(l))
	}
	return n
}

func sovSchemaAttribute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozSchemaAttribute(x uint64) (n int) {
	return sovSchemaAttribute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *SchemaAttribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaAttribute
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
			return fmt.Errorf("proto: SchemaAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
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
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CurrentValue == nil {
				m.CurrentValue = &SchemaAttributeValue{}
			}
			if err := m.CurrentValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaAttribute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchemaAttribute
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

func (m *SchemaAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchemaAttribute
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
			return fmt.Errorf("proto: SchemaAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NumberAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &SchemaAttributeValue_NumberAttributeValue{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &SchemaAttributeValue_StringAttributeValue{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BooleanAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &SchemaAttributeValue_BooleanAttributeValue{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchemaAttribute
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
				return ErrInvalidLengthSchemaAttribute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchemaAttribute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FloatAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &SchemaAttributeValue_FloatAttributeValue{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchemaAttribute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchemaAttribute
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

func skipSchemaAttribute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchemaAttribute
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
					return 0, ErrIntOverflowSchemaAttribute
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
					return 0, ErrIntOverflowSchemaAttribute
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
				return 0, ErrInvalidLengthSchemaAttribute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchemaAttribute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchemaAttribute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchemaAttribute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchemaAttribute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchemaAttribute = fmt.Errorf("proto: unexpected end of group")
)
