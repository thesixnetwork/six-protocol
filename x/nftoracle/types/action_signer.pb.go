// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftoracle/action_signer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CreationFlow int32

const (
	CreationFlow_ORACLE         CreationFlow = 0
	CreationFlow_INTERNAL_OWNER CreationFlow = 1
)

var CreationFlow_name = map[int32]string{
	0: "ORACLE",
	1: "INTERNAL_OWNER",
}

var CreationFlow_value = map[string]int32{
	"ORACLE":         0,
	"INTERNAL_OWNER": 1,
}

func (x CreationFlow) String() string {
	return proto.EnumName(CreationFlow_name, int32(x))
}

func (CreationFlow) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb36fd886ff3e92e, []int{0}
}

type ActionSigner struct {
	ActorAddress string       `protobuf:"bytes,1,opt,name=actor_address,json=actorAddress,proto3" json:"actor_address,omitempty"`
	OwnerAddress string       `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	CreatedAt    time.Time    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ExpiredAt    time.Time    `protobuf:"bytes,4,opt,name=expired_at,json=expiredAt,proto3,stdtime" json:"expired_at"`
	Creator      string       `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	CreationFlow CreationFlow `protobuf:"varint,6,opt,name=creation_flow,json=creationFlow,proto3,enum=thesixnetwork.sixnft.nftoracle.CreationFlow" json:"creation_flow,omitempty"`
}

func (m *ActionSigner) Reset()         { *m = ActionSigner{} }
func (m *ActionSigner) String() string { return proto.CompactTextString(m) }
func (*ActionSigner) ProtoMessage()    {}
func (*ActionSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb36fd886ff3e92e, []int{0}
}
func (m *ActionSigner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionSigner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionSigner.Merge(m, src)
}
func (m *ActionSigner) XXX_Size() int {
	return m.Size()
}
func (m *ActionSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionSigner.DiscardUnknown(m)
}

var xxx_messageInfo_ActionSigner proto.InternalMessageInfo

func (m *ActionSigner) GetActorAddress() string {
	if m != nil {
		return m.ActorAddress
	}
	return ""
}

func (m *ActionSigner) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *ActionSigner) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *ActionSigner) GetExpiredAt() time.Time {
	if m != nil {
		return m.ExpiredAt
	}
	return time.Time{}
}

func (m *ActionSigner) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ActionSigner) GetCreationFlow() CreationFlow {
	if m != nil {
		return m.CreationFlow
	}
	return CreationFlow_ORACLE
}

type SetSignerSignature struct {
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *SetSignerSignature) Reset()         { *m = SetSignerSignature{} }
func (m *SetSignerSignature) String() string { return proto.CompactTextString(m) }
func (*SetSignerSignature) ProtoMessage()    {}
func (*SetSignerSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb36fd886ff3e92e, []int{1}
}
func (m *SetSignerSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetSignerSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetSignerSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetSignerSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSignerSignature.Merge(m, src)
}
func (m *SetSignerSignature) XXX_Size() int {
	return m.Size()
}
func (m *SetSignerSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSignerSignature.DiscardUnknown(m)
}

var xxx_messageInfo_SetSignerSignature proto.InternalMessageInfo

func (m *SetSignerSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *SetSignerSignature) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SetSignerParams struct {
	OwnerAddress string    `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ActorAddress string    `protobuf:"bytes,2,opt,name=actor_address,json=actorAddress,proto3" json:"actor_address,omitempty"`
	ExpiredAt    time.Time `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3,stdtime" json:"expired_at"`
}

func (m *SetSignerParams) Reset()         { *m = SetSignerParams{} }
func (m *SetSignerParams) String() string { return proto.CompactTextString(m) }
func (*SetSignerParams) ProtoMessage()    {}
func (*SetSignerParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb36fd886ff3e92e, []int{2}
}
func (m *SetSignerParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetSignerParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetSignerParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetSignerParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSignerParams.Merge(m, src)
}
func (m *SetSignerParams) XXX_Size() int {
	return m.Size()
}
func (m *SetSignerParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSignerParams.DiscardUnknown(m)
}

var xxx_messageInfo_SetSignerParams proto.InternalMessageInfo

func (m *SetSignerParams) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *SetSignerParams) GetActorAddress() string {
	if m != nil {
		return m.ActorAddress
	}
	return ""
}

func (m *SetSignerParams) GetExpiredAt() time.Time {
	if m != nil {
		return m.ExpiredAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("thesixnetwork.sixnft.nftoracle.CreationFlow", CreationFlow_name, CreationFlow_value)
	proto.RegisterType((*ActionSigner)(nil), "thesixnetwork.sixnft.nftoracle.ActionSigner")
	proto.RegisterType((*SetSignerSignature)(nil), "thesixnetwork.sixnft.nftoracle.SetSignerSignature")
	proto.RegisterType((*SetSignerParams)(nil), "thesixnetwork.sixnft.nftoracle.SetSignerParams")
}

func init() { proto.RegisterFile("nftoracle/action_signer.proto", fileDescriptor_cb36fd886ff3e92e) }

var fileDescriptor_cb36fd886ff3e92e = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0x52, 0x8d, 0x66, 0x4c, 0x6b, 0x19, 0x3c, 0x2c, 0x41, 0x37, 0x21, 0x5e, 0x82,
	0xe8, 0x0c, 0xd4, 0x83, 0xe7, 0x6d, 0x88, 0x20, 0x84, 0xd4, 0x6e, 0x0b, 0x82, 0x97, 0x30, 0xd9,
	0x4c, 0xb6, 0x8b, 0xbb, 0xfb, 0xc2, 0xcc, 0x0b, 0x89, 0xdf, 0xa2, 0x1f, 0xc1, 0x8f, 0xd3, 0x9b,
	0x3d, 0x7a, 0x52, 0x49, 0xbe, 0x88, 0xec, 0xec, 0x6e, 0xd3, 0x62, 0x10, 0xec, 0x25, 0xbc, 0xf7,
	0xf2, 0x7f, 0xff, 0xf9, 0xf3, 0x7b, 0x2c, 0x7d, 0x91, 0xcd, 0x10, 0xb4, 0x0c, 0x13, 0x25, 0x64,
	0x88, 0x31, 0x64, 0x63, 0x13, 0x47, 0x99, 0xd2, 0x7c, 0xae, 0x01, 0x81, 0x79, 0x78, 0xa1, 0x4c,
	0xbc, 0xca, 0x14, 0x2e, 0x41, 0x7f, 0xe1, 0x79, 0x39, 0x43, 0x7e, 0xb3, 0xd3, 0x7a, 0x16, 0x41,
	0x04, 0x56, 0x2a, 0xf2, 0xaa, 0xd8, 0x6a, 0xb5, 0x23, 0x80, 0x28, 0x51, 0xc2, 0x76, 0x93, 0xc5,
	0x4c, 0x60, 0x9c, 0x2a, 0x83, 0x32, 0x9d, 0x17, 0x82, 0xee, 0xf7, 0x1a, 0x6d, 0xfa, 0xf6, 0xb9,
	0x33, 0xfb, 0x1a, 0x7b, 0x49, 0xf7, 0x65, 0x88, 0xa0, 0xc7, 0x72, 0x3a, 0xd5, 0xca, 0x18, 0x97,
	0x74, 0x48, 0xaf, 0x11, 0x34, 0xed, 0xd0, 0x2f, 0x66, 0xb9, 0x08, 0x96, 0x99, 0xda, 0x8a, 0x6a,
	0x85, 0xc8, 0x0e, 0x2b, 0x51, 0x9f, 0xd2, 0x50, 0x2b, 0x89, 0x6a, 0x3a, 0x96, 0xe8, 0xee, 0x75,
	0x48, 0xef, 0xc9, 0x51, 0x8b, 0x17, 0x81, 0x78, 0x15, 0x88, 0x9f, 0x57, 0x81, 0x8e, 0x1f, 0x5f,
	0xfd, 0x6c, 0x3b, 0x97, 0xbf, 0xda, 0x24, 0x68, 0x94, 0x7b, 0x3e, 0xe6, 0x26, 0x6a, 0x35, 0x8f,
	0x75, 0x61, 0xf2, 0xe0, 0x7f, 0x4c, 0xca, 0x3d, 0x1f, 0x99, 0x4b, 0x1f, 0x59, 0x47, 0xd0, 0xee,
	0x43, 0x1b, 0xb4, 0x6a, 0xd9, 0x29, 0xdd, 0xb7, 0x65, 0x8e, 0x7b, 0x96, 0xc0, 0xd2, 0xad, 0x77,
	0x48, 0xef, 0xe0, 0xe8, 0x35, 0xff, 0x37, 0x6d, 0xde, 0x2f, 0x97, 0xde, 0x27, 0xb0, 0x0c, 0x9a,
	0xe1, 0xad, 0xae, 0x3b, 0xa4, 0xec, 0x4c, 0x61, 0x41, 0x33, 0xff, 0x95, 0xb8, 0xd0, 0x8a, 0x3d,
	0xa7, 0x0d, 0x53, 0x35, 0x25, 0xd2, 0xed, 0x20, 0x0f, 0x98, 0x2a, 0x63, 0x64, 0xa4, 0x4a, 0x92,
	0x55, 0xdb, 0xfd, 0x46, 0xe8, 0xd3, 0x1b, 0xbb, 0x8f, 0x52, 0xcb, 0x74, 0x07, 0x7d, 0xb2, 0x83,
	0xfe, 0x5f, 0x77, 0xac, 0xed, 0xb8, 0xe3, 0x5d, 0xba, 0x7b, 0xf7, 0xa2, 0xfb, 0x8a, 0xd3, 0xe6,
	0x6d, 0x1c, 0x8c, 0xd2, 0xfa, 0x49, 0xe0, 0xf7, 0x87, 0x83, 0x43, 0x87, 0x31, 0x7a, 0xf0, 0x61,
	0x74, 0x3e, 0x08, 0x46, 0xfe, 0x70, 0x7c, 0xf2, 0x69, 0x34, 0x08, 0x0e, 0xc9, 0xf1, 0xe9, 0xd5,
	0xda, 0x23, 0xd7, 0x6b, 0x8f, 0xfc, 0x5e, 0x7b, 0xe4, 0x72, 0xe3, 0x39, 0xd7, 0x1b, 0xcf, 0xf9,
	0xb1, 0xf1, 0x9c, 0xcf, 0xef, 0xa2, 0x18, 0x2f, 0x16, 0x13, 0x1e, 0x42, 0x2a, 0xee, 0x1c, 0x40,
	0x98, 0x78, 0xf5, 0xc6, 0xe6, 0x09, 0x21, 0x11, 0x2b, 0xb1, 0xfd, 0x54, 0xf0, 0xeb, 0x5c, 0x99,
	0x49, 0xdd, 0xfe, 0xf7, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xb2, 0xd8, 0x31, 0x44,
	0x03, 0x00, 0x00,
}

func (m *ActionSigner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionSigner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionSigner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreationFlow != 0 {
		i = encodeVarintActionSigner(dAtA, i, uint64(m.CreationFlow))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiredAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintActionSigner(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintActionSigner(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ActorAddress) > 0 {
		i -= len(m.ActorAddress)
		copy(dAtA[i:], m.ActorAddress)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.ActorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetSignerSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetSignerSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetSignerSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetSignerParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetSignerParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetSignerParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiredAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintActionSigner(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if len(m.ActorAddress) > 0 {
		i -= len(m.ActorAddress)
		copy(dAtA[i:], m.ActorAddress)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.ActorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintActionSigner(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActionSigner(dAtA []byte, offset int, v uint64) int {
	offset -= sovActionSigner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionSigner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ActorAddress)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovActionSigner(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt)
	n += 1 + l + sovActionSigner(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	if m.CreationFlow != 0 {
		n += 1 + sovActionSigner(uint64(m.CreationFlow))
	}
	return n
}

func (m *SetSignerSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	return n
}

func (m *SetSignerParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	l = len(m.ActorAddress)
	if l > 0 {
		n += 1 + l + sovActionSigner(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiredAt)
	n += 1 + l + sovActionSigner(uint64(l))
	return n
}

func sovActionSigner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActionSigner(x uint64) (n int) {
	return sovActionSigner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionSigner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionSigner
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
			return fmt.Errorf("proto: ActionSigner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionSigner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiredAt, dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationFlow", wireType)
			}
			m.CreationFlow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationFlow |= CreationFlow(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActionSigner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionSigner
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
func (m *SetSignerSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionSigner
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
			return fmt.Errorf("proto: SetSignerSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetSignerSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionSigner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionSigner
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
func (m *SetSignerParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActionSigner
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
			return fmt.Errorf("proto: SetSignerParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetSignerParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActionSigner
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
				return ErrInvalidLengthActionSigner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActionSigner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiredAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActionSigner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActionSigner
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
func skipActionSigner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActionSigner
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
					return 0, ErrIntOverflowActionSigner
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
					return 0, ErrIntOverflowActionSigner
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
				return 0, ErrInvalidLengthActionSigner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActionSigner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActionSigner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActionSigner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActionSigner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActionSigner = fmt.Errorf("proto: unexpected end of group")
)
