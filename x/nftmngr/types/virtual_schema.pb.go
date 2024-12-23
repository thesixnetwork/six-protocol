// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/virtual_schema.proto

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

type RegistryStatus int32

const (
	RegistryStatus_PENDING RegistryStatus = 0
	RegistryStatus_REJECT  RegistryStatus = 1
	RegistryStatus_ACCEPT  RegistryStatus = 2
)

var RegistryStatus_name = map[int32]string{
	0: "PENDING",
	1: "REJECT",
	2: "ACCEPT",
}

var RegistryStatus_value = map[string]int32{
	"PENDING": 0,
	"REJECT":  1,
	"ACCEPT":  2,
}

func (x RegistryStatus) String() string {
	return proto.EnumName(RegistryStatus_name, int32(x))
}

func (RegistryStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fd45310b3f69052, []int{0}
}

type VirtualSchemaProposal struct {
	Id                string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VirtualSchemaCode string                   `protobuf:"bytes,2,opt,name=virtualSchemaCode,proto3" json:"virtualSchemaCode,omitempty"`
	Registry          []*VirtualSchemaRegistry `protobuf:"bytes,3,rep,name=registry,proto3" json:"registry,omitempty"`
}

func (m *VirtualSchemaProposal) Reset()         { *m = VirtualSchemaProposal{} }
func (m *VirtualSchemaProposal) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaProposal) ProtoMessage()    {}
func (*VirtualSchemaProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd45310b3f69052, []int{0}
}

func (m *VirtualSchemaProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *VirtualSchemaProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *VirtualSchemaProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaProposal.Merge(m, src)
}

func (m *VirtualSchemaProposal) XXX_Size() int {
	return m.Size()
}

func (m *VirtualSchemaProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaProposal.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaProposal proto.InternalMessageInfo

func (m *VirtualSchemaProposal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualSchemaProposal) GetVirtualSchemaCode() string {
	if m != nil {
		return m.VirtualSchemaCode
	}
	return ""
}

func (m *VirtualSchemaProposal) GetRegistry() []*VirtualSchemaRegistry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type VirtualSchema struct {
	VirtualNftSchemaCode string                   `protobuf:"bytes,1,opt,name=virtualNftSchemaCode,proto3" json:"virtualNftSchemaCode,omitempty"`
	Registry             []*VirtualSchemaRegistry `protobuf:"bytes,2,rep,name=registry,proto3" json:"registry,omitempty"`
	Enable               bool                     `protobuf:"varint,3,opt,name=enable,proto3" json:"enable,omitempty"`
	ExpiredAtBlock       string                   `protobuf:"bytes,4,opt,name=expiredAtBlock,proto3" json:"expiredAtBlock,omitempty"`
}

func (m *VirtualSchema) Reset()         { *m = VirtualSchema{} }
func (m *VirtualSchema) String() string { return proto.CompactTextString(m) }
func (*VirtualSchema) ProtoMessage()    {}
func (*VirtualSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd45310b3f69052, []int{1}
}

func (m *VirtualSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *VirtualSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *VirtualSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchema.Merge(m, src)
}

func (m *VirtualSchema) XXX_Size() int {
	return m.Size()
}

func (m *VirtualSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchema.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchema proto.InternalMessageInfo

func (m *VirtualSchema) GetVirtualNftSchemaCode() string {
	if m != nil {
		return m.VirtualNftSchemaCode
	}
	return ""
}

func (m *VirtualSchema) GetRegistry() []*VirtualSchemaRegistry {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *VirtualSchema) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *VirtualSchema) GetExpiredAtBlock() string {
	if m != nil {
		return m.ExpiredAtBlock
	}
	return ""
}

type VirtualSchemaRegistry struct {
	NftSchemaCode    string         `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	SharedAttributes []string       `protobuf:"bytes,2,rep,name=sharedAttributes,proto3" json:"sharedAttributes,omitempty"`
	Status           RegistryStatus `protobuf:"varint,3,opt,name=status,proto3,enum=thesixnetwork.sixprotocol.nftmngr.RegistryStatus" json:"status,omitempty"`
}

func (m *VirtualSchemaRegistry) Reset()         { *m = VirtualSchemaRegistry{} }
func (m *VirtualSchemaRegistry) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaRegistry) ProtoMessage()    {}
func (*VirtualSchemaRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd45310b3f69052, []int{2}
}

func (m *VirtualSchemaRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *VirtualSchemaRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *VirtualSchemaRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaRegistry.Merge(m, src)
}

func (m *VirtualSchemaRegistry) XXX_Size() int {
	return m.Size()
}

func (m *VirtualSchemaRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaRegistry proto.InternalMessageInfo

func (m *VirtualSchemaRegistry) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *VirtualSchemaRegistry) GetSharedAttributes() []string {
	if m != nil {
		return m.SharedAttributes
	}
	return nil
}

func (m *VirtualSchemaRegistry) GetStatus() RegistryStatus {
	if m != nil {
		return m.Status
	}
	return RegistryStatus_PENDING
}

type VirtualSchemaRegistryRequest struct {
	NftSchemaCode    string   `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	SharedAttributes []string `protobuf:"bytes,2,rep,name=sharedAttributes,proto3" json:"sharedAttributes,omitempty"`
}

func (m *VirtualSchemaRegistryRequest) Reset()         { *m = VirtualSchemaRegistryRequest{} }
func (m *VirtualSchemaRegistryRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaRegistryRequest) ProtoMessage()    {}
func (*VirtualSchemaRegistryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd45310b3f69052, []int{3}
}

func (m *VirtualSchemaRegistryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *VirtualSchemaRegistryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaRegistryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *VirtualSchemaRegistryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaRegistryRequest.Merge(m, src)
}

func (m *VirtualSchemaRegistryRequest) XXX_Size() int {
	return m.Size()
}

func (m *VirtualSchemaRegistryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaRegistryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaRegistryRequest proto.InternalMessageInfo

func (m *VirtualSchemaRegistryRequest) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *VirtualSchemaRegistryRequest) GetSharedAttributes() []string {
	if m != nil {
		return m.SharedAttributes
	}
	return nil
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixprotocol.nftmngr.RegistryStatus", RegistryStatus_name, RegistryStatus_value)
	proto.RegisterType((*VirtualSchemaProposal)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaProposal")
	proto.RegisterType((*VirtualSchema)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchema")
	proto.RegisterType((*VirtualSchemaRegistry)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaRegistry")
	proto.RegisterType((*VirtualSchemaRegistryRequest)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaRegistryRequest")
}

func init() { proto.RegisterFile("nftmngr/virtual_schema.proto", fileDescriptor_1fd45310b3f69052) }

var fileDescriptor_1fd45310b3f69052 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6a, 0xd4, 0x50,
	0x14, 0xc6, 0x73, 0x33, 0x12, 0xdb, 0x53, 0x1a, 0xe2, 0x45, 0x25, 0x8b, 0x12, 0xc6, 0x20, 0x32,
	0x14, 0x4d, 0x70, 0xa4, 0xe0, 0xb6, 0x8d, 0x41, 0xea, 0x62, 0x1c, 0xd2, 0xc1, 0x85, 0x1b, 0xc9,
	0x9f, 0xdb, 0xc9, 0xa5, 0x99, 0xdc, 0x78, 0xef, 0x4d, 0x4d, 0xdf, 0xc2, 0xf7, 0x70, 0xeb, 0x43,
	0xb8, 0xec, 0x52, 0x77, 0x32, 0xf3, 0x22, 0xe2, 0x6d, 0x22, 0x4d, 0x3b, 0x60, 0x17, 0xb3, 0x4b,
	0xbe, 0xef, 0x7c, 0xe7, 0xfc, 0xce, 0x09, 0x81, 0xbd, 0xf2, 0x54, 0x2e, 0xca, 0x39, 0xf7, 0xcf,
	0x29, 0x97, 0x75, 0x5c, 0x7c, 0x12, 0x69, 0x4e, 0x16, 0xb1, 0x57, 0x71, 0x26, 0x19, 0x7e, 0x22,
	0x73, 0x22, 0x68, 0x53, 0x12, 0xf9, 0x85, 0xf1, 0x33, 0x4f, 0xd0, 0x46, 0xe9, 0x29, 0x2b, 0xbc,
	0x36, 0xe7, 0x7e, 0x43, 0xf0, 0xe8, 0xc3, 0x55, 0xf6, 0x44, 0x45, 0xa7, 0x9c, 0x55, 0x4c, 0xc4,
	0x05, 0x36, 0x41, 0xa7, 0x99, 0x8d, 0x86, 0x68, 0xb4, 0x1d, 0xe9, 0x34, 0xc3, 0xcf, 0xe1, 0xc1,
	0xf9, 0xf5, 0xc2, 0x80, 0x65, 0xc4, 0xd6, 0x95, 0x7d, 0xdb, 0xc0, 0x33, 0xd8, 0xe2, 0x64, 0x4e,
	0x85, 0xe4, 0x17, 0xf6, 0x60, 0x38, 0x18, 0xed, 0x8c, 0x5f, 0x7b, 0xff, 0xa5, 0xf1, 0x7a, 0x24,
	0x51, 0x9b, 0x8f, 0xfe, 0x75, 0x72, 0x7f, 0x21, 0xd8, 0xed, 0xd5, 0xe0, 0x31, 0x3c, 0x6c, 0x87,
	0x4f, 0x4e, 0xe5, 0x35, 0xb0, 0x2b, 0xee, 0xb5, 0x5e, 0x8f, 0x4d, 0xdf, 0x14, 0x1b, 0x7e, 0x0c,
	0x06, 0x29, 0xe3, 0xa4, 0x20, 0xf6, 0x60, 0x88, 0x46, 0x5b, 0x51, 0xfb, 0x86, 0x9f, 0x81, 0x49,
	0x9a, 0x8a, 0x72, 0x92, 0x1d, 0xca, 0xa3, 0x82, 0xa5, 0x67, 0xf6, 0x3d, 0xc5, 0x76, 0x43, 0x75,
	0xbf, 0xdf, 0xfc, 0x12, 0xdd, 0x0c, 0xfc, 0x14, 0x76, 0xcb, 0x35, 0xcb, 0xf5, 0x45, 0xbc, 0x0f,
	0x96, 0xc8, 0x63, 0xd5, 0x50, 0x72, 0x9a, 0xd4, 0x92, 0x08, 0xb5, 0xdd, 0x76, 0x74, 0x4b, 0xc7,
	0xc7, 0x60, 0x08, 0x19, 0xcb, 0x5a, 0x28, 0x56, 0x73, 0xfc, 0xf2, 0x0e, 0xfb, 0x77, 0x38, 0x27,
	0x2a, 0x18, 0xb5, 0x0d, 0xdc, 0x0a, 0xf6, 0xd6, 0x5f, 0x86, 0x7c, 0xae, 0x89, 0x90, 0x9b, 0x87,
	0xdf, 0x3f, 0x00, 0xb3, 0xcf, 0x82, 0x77, 0xe0, 0xfe, 0x34, 0x9c, 0xbc, 0x39, 0x9e, 0xbc, 0xb5,
	0x34, 0x0c, 0x60, 0x44, 0xe1, 0xbb, 0x30, 0x98, 0x59, 0xe8, 0xef, 0xf3, 0x61, 0x10, 0x84, 0xd3,
	0x99, 0xa5, 0x1f, 0xbd, 0xff, 0xb1, 0x74, 0xd0, 0xe5, 0xd2, 0x41, 0xbf, 0x97, 0x0e, 0xfa, 0xba,
	0x72, 0xb4, 0xcb, 0x95, 0xa3, 0xfd, 0x5c, 0x39, 0xda, 0xc7, 0x83, 0x39, 0x95, 0x79, 0x9d, 0x78,
	0x29, 0x5b, 0xf8, 0xbd, 0x3b, 0xf8, 0x82, 0x36, 0x2f, 0xba, 0x43, 0xf8, 0x8d, 0xdf, 0xfd, 0x6c,
	0xf2, 0xa2, 0x22, 0x22, 0x31, 0x94, 0xf3, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x66,
	0xa5, 0xa9, 0x84, 0x03, 0x00, 0x00,
}

func (m *VirtualSchemaProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVirtualSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VirtualSchemaCode) > 0 {
		i -= len(m.VirtualSchemaCode)
		copy(dAtA[i:], m.VirtualSchemaCode)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.VirtualSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExpiredAtBlock) > 0 {
		i -= len(m.ExpiredAtBlock)
		copy(dAtA[i:], m.ExpiredAtBlock)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.ExpiredAtBlock)))
		i--
		dAtA[i] = 0x22
	}
	if m.Enable {
		i--
		if m.Enable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVirtualSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.VirtualNftSchemaCode) > 0 {
		i -= len(m.VirtualNftSchemaCode)
		copy(dAtA[i:], m.VirtualNftSchemaCode)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.VirtualNftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchemaRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintVirtualSchema(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SharedAttributes) > 0 {
		for iNdEx := len(m.SharedAttributes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SharedAttributes[iNdEx])
			copy(dAtA[i:], m.SharedAttributes[iNdEx])
			i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.SharedAttributes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchemaRegistryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaRegistryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaRegistryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SharedAttributes) > 0 {
		for iNdEx := len(m.SharedAttributes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SharedAttributes[iNdEx])
			copy(dAtA[i:], m.SharedAttributes[iNdEx])
			i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.SharedAttributes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NftSchemaCode) > 0 {
		i -= len(m.NftSchemaCode)
		copy(dAtA[i:], m.NftSchemaCode)
		i = encodeVarintVirtualSchema(dAtA, i, uint64(len(m.NftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVirtualSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovVirtualSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *VirtualSchemaProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	l = len(m.VirtualSchemaCode)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovVirtualSchema(uint64(l))
		}
	}
	return n
}

func (m *VirtualSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VirtualNftSchemaCode)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovVirtualSchema(uint64(l))
		}
	}
	if m.Enable {
		n += 2
	}
	l = len(m.ExpiredAtBlock)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	return n
}

func (m *VirtualSchemaRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	if len(m.SharedAttributes) > 0 {
		for _, s := range m.SharedAttributes {
			l = len(s)
			n += 1 + l + sovVirtualSchema(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovVirtualSchema(uint64(m.Status))
	}
	return n
}

func (m *VirtualSchemaRegistryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovVirtualSchema(uint64(l))
	}
	if len(m.SharedAttributes) > 0 {
		for _, s := range m.SharedAttributes {
			l = len(s)
			n += 1 + l + sovVirtualSchema(uint64(l))
		}
	}
	return n
}

func sovVirtualSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozVirtualSchema(x uint64) (n int) {
	return sovVirtualSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *VirtualSchemaProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchemaProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, &VirtualSchemaRegistry{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVirtualSchema
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

func (m *VirtualSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualNftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualNftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, &VirtualSchemaRegistry{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
			m.Enable = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredAtBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpiredAtBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVirtualSchema
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

func (m *VirtualSchemaRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchemaRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharedAttributes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SharedAttributes = append(m.SharedAttributes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RegistryStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVirtualSchema
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

func (m *VirtualSchemaRegistryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVirtualSchema
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
			return fmt.Errorf("proto: VirtualSchemaRegistryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaRegistryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftSchemaCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharedAttributes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVirtualSchema
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
				return ErrInvalidLengthVirtualSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SharedAttributes = append(m.SharedAttributes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVirtualSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVirtualSchema
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

func skipVirtualSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVirtualSchema
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
					return 0, ErrIntOverflowVirtualSchema
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
					return 0, ErrIntOverflowVirtualSchema
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
				return 0, ErrInvalidLengthVirtualSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVirtualSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVirtualSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVirtualSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVirtualSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVirtualSchema = fmt.Errorf("proto: unexpected end of group")
)
