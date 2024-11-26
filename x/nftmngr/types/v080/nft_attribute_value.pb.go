// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/v080/nft_attribute_value.proto

package v080

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type NftAttributeValue struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*NftAttributeValue_NumberAttributeValue
	//	*NftAttributeValue_StringAttributeValue
	//	*NftAttributeValue_BooleanAttributeValue
	//	*NftAttributeValue_FloatAttributeValue
	Value               isNftAttributeValue_Value `protobuf_oneof:"value"`
	HiddenToMarketplace bool                      `protobuf:"varint,6,opt,name=hidden_to_marketplace,json=hiddenToMarketplace,proto3" json:"hidden_to_marketplace,omitempty"`
}

func (m *NftAttributeValue) Reset()         { *m = NftAttributeValue{} }
func (m *NftAttributeValue) String() string { return proto.CompactTextString(m) }
func (*NftAttributeValue) ProtoMessage()    {}
func (*NftAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f004df871d5ee2b0, []int{0}
}
func (m *NftAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftAttributeValue.Merge(m, src)
}
func (m *NftAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *NftAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NftAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_NftAttributeValue proto.InternalMessageInfo

type isNftAttributeValue_Value interface {
	isNftAttributeValue_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type NftAttributeValue_NumberAttributeValue struct {
	NumberAttributeValue *NumberAttributeValue `protobuf:"bytes,2,opt,name=number_attribute_value,json=numberAttributeValue,proto3,oneof" json:"number_attribute_value,omitempty"`
}
type NftAttributeValue_StringAttributeValue struct {
	StringAttributeValue *StringAttributeValue `protobuf:"bytes,3,opt,name=string_attribute_value,json=stringAttributeValue,proto3,oneof" json:"string_attribute_value,omitempty"`
}
type NftAttributeValue_BooleanAttributeValue struct {
	BooleanAttributeValue *BooleanAttributeValue `protobuf:"bytes,4,opt,name=boolean_attribute_value,json=booleanAttributeValue,proto3,oneof" json:"boolean_attribute_value,omitempty"`
}
type NftAttributeValue_FloatAttributeValue struct {
	FloatAttributeValue *FloatAttributeValue `protobuf:"bytes,5,opt,name=float_attribute_value,json=floatAttributeValue,proto3,oneof" json:"float_attribute_value,omitempty"`
}

func (*NftAttributeValue_NumberAttributeValue) isNftAttributeValue_Value()  {}
func (*NftAttributeValue_StringAttributeValue) isNftAttributeValue_Value()  {}
func (*NftAttributeValue_BooleanAttributeValue) isNftAttributeValue_Value() {}
func (*NftAttributeValue_FloatAttributeValue) isNftAttributeValue_Value()   {}

func (m *NftAttributeValue) GetValue() isNftAttributeValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NftAttributeValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NftAttributeValue) GetNumberAttributeValue() *NumberAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		return x.NumberAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetStringAttributeValue() *StringAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		return x.StringAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetBooleanAttributeValue() *BooleanAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		return x.BooleanAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetFloatAttributeValue() *FloatAttributeValue {
	if x, ok := m.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		return x.FloatAttributeValue
	}
	return nil
}

func (m *NftAttributeValue) GetHiddenToMarketplace() bool {
	if m != nil {
		return m.HiddenToMarketplace
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NftAttributeValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NftAttributeValue_NumberAttributeValue)(nil),
		(*NftAttributeValue_StringAttributeValue)(nil),
		(*NftAttributeValue_BooleanAttributeValue)(nil),
		(*NftAttributeValue_FloatAttributeValue)(nil),
	}
}

type NumberAttributeValue struct {
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NumberAttributeValue) Reset()         { *m = NumberAttributeValue{} }
func (m *NumberAttributeValue) String() string { return proto.CompactTextString(m) }
func (*NumberAttributeValue) ProtoMessage()    {}
func (*NumberAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f004df871d5ee2b0, []int{1}
}
func (m *NumberAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumberAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumberAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NumberAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberAttributeValue.Merge(m, src)
}
func (m *NumberAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *NumberAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_NumberAttributeValue proto.InternalMessageInfo

func (m *NumberAttributeValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringAttributeValue struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringAttributeValue) Reset()         { *m = StringAttributeValue{} }
func (m *StringAttributeValue) String() string { return proto.CompactTextString(m) }
func (*StringAttributeValue) ProtoMessage()    {}
func (*StringAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f004df871d5ee2b0, []int{2}
}
func (m *StringAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringAttributeValue.Merge(m, src)
}
func (m *StringAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *StringAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringAttributeValue proto.InternalMessageInfo

func (m *StringAttributeValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BooleanAttributeValue struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BooleanAttributeValue) Reset()         { *m = BooleanAttributeValue{} }
func (m *BooleanAttributeValue) String() string { return proto.CompactTextString(m) }
func (*BooleanAttributeValue) ProtoMessage()    {}
func (*BooleanAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f004df871d5ee2b0, []int{3}
}
func (m *BooleanAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BooleanAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BooleanAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BooleanAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanAttributeValue.Merge(m, src)
}
func (m *BooleanAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *BooleanAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanAttributeValue proto.InternalMessageInfo

func (m *BooleanAttributeValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type FloatAttributeValue struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *FloatAttributeValue) Reset()         { *m = FloatAttributeValue{} }
func (m *FloatAttributeValue) String() string { return proto.CompactTextString(m) }
func (*FloatAttributeValue) ProtoMessage()    {}
func (*FloatAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f004df871d5ee2b0, []int{4}
}
func (m *FloatAttributeValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FloatAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FloatAttributeValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FloatAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatAttributeValue.Merge(m, src)
}
func (m *FloatAttributeValue) XXX_Size() int {
	return m.Size()
}
func (m *FloatAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_FloatAttributeValue proto.InternalMessageInfo

func (m *FloatAttributeValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*NftAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.NftAttributeValue")
	proto.RegisterType((*NumberAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.NumberAttributeValue")
	proto.RegisterType((*StringAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.StringAttributeValue")
	proto.RegisterType((*BooleanAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.BooleanAttributeValue")
	proto.RegisterType((*FloatAttributeValue)(nil), "thesixnetwork.sixprotocol.nftmngr.v080.FloatAttributeValue")
}

func init() {
	proto.RegisterFile("nftmngr/v080/nft_attribute_value.proto", fileDescriptor_f004df871d5ee2b0)
}

var fileDescriptor_f004df871d5ee2b0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0xbd, 0x89, 0x21, 0xb0, 0xa9, 0xb2, 0xe0, 0xc4, 0x49, 0x61, 0x21, 0x0a, 0x84, 0x94,
	0xe0, 0x45, 0xa4, 0x89, 0xf2, 0xa7, 0x08, 0x45, 0x94, 0x26, 0x14, 0x0e, 0xba, 0xe2, 0x1a, 0x6b,
	0x0d, 0x6b, 0x63, 0x61, 0xef, 0x72, 0xf6, 0x1a, 0xcc, 0x5b, 0xdc, 0x03, 0xdd, 0x03, 0x5c, 0x49,
	0x79, 0xe5, 0x09, 0x5e, 0xe4, 0xe4, 0x05, 0x0b, 0xf9, 0xec, 0x82, 0xeb, 0xac, 0xf9, 0x3e, 0x7d,
	0xbf, 0xf1, 0xcc, 0x0e, 0xec, 0x31, 0x57, 0x84, 0xcc, 0x8b, 0xf0, 0x7a, 0xf8, 0x6d, 0x88, 0x99,
	0x2b, 0x6c, 0x22, 0x44, 0xe4, 0x3b, 0x89, 0xa0, 0xf6, 0x9a, 0x04, 0x09, 0x35, 0x57, 0x11, 0x17,
	0x1c, 0xf5, 0xc4, 0x82, 0xc6, 0x7e, 0xca, 0xa8, 0xd8, 0xf0, 0x68, 0x69, 0xc6, 0x7e, 0x2a, 0xeb,
	0x33, 0x1e, 0x98, 0xa7, 0x04, 0x33, 0x4b, 0xf8, 0xf4, 0xd1, 0xe3, 0xdc, 0x0b, 0x28, 0x96, 0xaa,
	0x93, 0xb8, 0x98, 0xb0, 0xed, 0x31, 0xa2, 0x7b, 0xa7, 0xc2, 0x77, 0x13, 0x57, 0xfc, 0xce, 0xf3,
	0xaf, 0xb2, 0x78, 0x84, 0xa0, 0xca, 0x48, 0x48, 0x75, 0xd0, 0x01, 0xfd, 0xa6, 0x25, 0xbf, 0x91,
	0x80, 0xef, 0x59, 0x12, 0x3a, 0x34, 0x7a, 0xde, 0x8c, 0xfe, 0xaa, 0x03, 0xfa, 0x6f, 0x47, 0x3f,
	0xcd, 0xcb, 0xba, 0x31, 0x27, 0x32, 0xa5, 0x48, 0xfc, 0xab, 0x58, 0x6d, 0x56, 0x51, 0xcf, 0xa8,
	0xb1, 0x88, 0x7c, 0xe6, 0x95, 0xa8, 0xaf, 0x5f, 0x46, 0xfd, 0x2f, 0x53, 0xca, 0xd4, 0xb8, 0xa2,
	0x8e, 0x36, 0xf0, 0x83, 0xc3, 0x79, 0x40, 0x09, 0x2b, 0x61, 0x55, 0x89, 0xfd, 0x75, 0x29, 0x76,
	0x7c, 0x8c, 0x29, 0x71, 0x35, 0xa7, 0x4a, 0x40, 0x37, 0x50, 0x73, 0x03, 0x4e, 0x4a, 0x0b, 0xd7,
	0x6b, 0x12, 0xfb, 0xe3, 0x52, 0xec, 0x9f, 0x2c, 0xa4, 0x04, 0x6d, 0xb9, 0xe5, 0x32, 0x1a, 0x41,
	0x6d, 0xe1, 0xcf, 0xe7, 0x94, 0xd9, 0x82, 0xdb, 0x21, 0x89, 0x96, 0x54, 0xac, 0x02, 0x32, 0xa3,
	0x7a, 0xbd, 0x03, 0xfa, 0x0d, 0xab, 0x75, 0x14, 0xa7, 0xfc, 0xdf, 0x59, 0x1a, 0xbf, 0x81, 0x35,
	0xd9, 0x56, 0xf7, 0x0b, 0x6c, 0x57, 0xad, 0x13, 0xb5, 0x4f, 0x06, 0xf9, 0x82, 0x54, 0xeb, 0xec,
	0xae, 0x5a, 0x43, 0xd1, 0xdd, 0xcc, 0xdd, 0x03, 0xa8, 0x55, 0x4e, 0xaf, 0x68, 0x6f, 0xe4, 0xf6,
	0xcf, 0xb0, 0x55, 0xf1, 0xd7, 0x45, 0x33, 0x38, 0x99, 0xc7, 0xd3, 0xfb, 0xbd, 0x01, 0x76, 0x7b,
	0x03, 0x3c, 0xee, 0x0d, 0x70, 0x7b, 0x30, 0x94, 0xdd, 0xc1, 0x50, 0x1e, 0x0e, 0x86, 0x72, 0xfd,
	0xdd, 0xf3, 0xc5, 0x22, 0x71, 0xcc, 0x19, 0x0f, 0x71, 0x61, 0xd8, 0x38, 0xf6, 0xd3, 0x41, 0x3e,
	0x6d, 0x9c, 0xe2, 0xfc, 0x46, 0xc5, 0x76, 0x45, 0x63, 0x79, 0xa9, 0x4e, 0x5d, 0xca, 0x5f, 0x9f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0xb3, 0x13, 0xbe, 0xc0, 0x03, 0x00, 0x00,
}

func (m *NftAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HiddenToMarketplace {
		i--
		if m.HiddenToMarketplace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NftAttributeValue_NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NumberAttributeValue != nil {
		{
			size, err := m.NumberAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StringAttributeValue != nil {
		{
			size, err := m.StringAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BooleanAttributeValue != nil {
		{
			size, err := m.BooleanAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *NftAttributeValue_FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftAttributeValue_FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FloatAttributeValue != nil {
		{
			size, err := m.FloatAttributeValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNftAttributeValue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *NumberAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumberAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NumberAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StringAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintNftAttributeValue(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BooleanAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BooleanAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BooleanAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value {
		i--
		if m.Value {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FloatAttributeValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FloatAttributeValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FloatAttributeValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftAttributeValue(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftAttributeValue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	if m.HiddenToMarketplace {
		n += 2
	}
	return n
}

func (m *NftAttributeValue_NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumberAttributeValue != nil {
		l = m.NumberAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringAttributeValue != nil {
		l = m.StringAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BooleanAttributeValue != nil {
		l = m.BooleanAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NftAttributeValue_FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FloatAttributeValue != nil {
		l = m.FloatAttributeValue.Size()
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}
func (m *NumberAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovNftAttributeValue(uint64(m.Value))
	}
	return n
}

func (m *StringAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovNftAttributeValue(uint64(l))
	}
	return n
}

func (m *BooleanAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value {
		n += 2
	}
	return n
}

func (m *FloatAttributeValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func sovNftAttributeValue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftAttributeValue(x uint64) (n int) {
	return sovNftAttributeValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: NftAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NumberAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_NumberAttributeValue{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_StringAttributeValue{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BooleanAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &BooleanAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_BooleanAttributeValue{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatAttributeValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FloatAttributeValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &NftAttributeValue_FloatAttributeValue{v}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HiddenToMarketplace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
			m.HiddenToMarketplace = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *NumberAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: NumberAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumberAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *StringAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: StringAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
				return ErrInvalidLengthNftAttributeValue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftAttributeValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *BooleanAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: BooleanAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BooleanAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftAttributeValue
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
			m.Value = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func (m *FloatAttributeValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftAttributeValue
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
			return fmt.Errorf("proto: FloatAttributeValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FloatAttributeValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipNftAttributeValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftAttributeValue
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
func skipNftAttributeValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftAttributeValue
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
					return 0, ErrIntOverflowNftAttributeValue
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
					return 0, ErrIntOverflowNftAttributeValue
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
				return 0, ErrInvalidLengthNftAttributeValue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftAttributeValue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftAttributeValue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftAttributeValue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftAttributeValue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftAttributeValue = fmt.Errorf("proto: unexpected end of group")
)
