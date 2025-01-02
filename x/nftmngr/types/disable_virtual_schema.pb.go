// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nftmngr/disable_virtual_schema.proto

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

type DisableVirtualSchemaProposal struct {
	Id                string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VirtualSchemaCode string                          `protobuf:"bytes,2,opt,name=virtualSchemaCode,proto3" json:"virtualSchemaCode,omitempty"`
	Registry          []*DisableVirtualSchemaRegistry `protobuf:"bytes,3,rep,name=registry,proto3" json:"registry,omitempty"`
	SubmitTime        time.Time                       `protobuf:"bytes,4,opt,name=submitTime,proto3,stdtime" json:"submitTime"`
	VotinStartTime    time.Time                       `protobuf:"bytes,5,opt,name=votinStartTime,proto3,stdtime" json:"votinStartTime"`
	VotingEndTime     time.Time                       `protobuf:"bytes,6,opt,name=votingEndTime,proto3,stdtime" json:"votingEndTime"`
}

func (m *DisableVirtualSchemaProposal) Reset()         { *m = DisableVirtualSchemaProposal{} }
func (m *DisableVirtualSchemaProposal) String() string { return proto.CompactTextString(m) }
func (*DisableVirtualSchemaProposal) ProtoMessage()    {}
func (*DisableVirtualSchemaProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{0}
}
func (m *DisableVirtualSchemaProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisableVirtualSchemaProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisableVirtualSchemaProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisableVirtualSchemaProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableVirtualSchemaProposal.Merge(m, src)
}
func (m *DisableVirtualSchemaProposal) XXX_Size() int {
	return m.Size()
}
func (m *DisableVirtualSchemaProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableVirtualSchemaProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DisableVirtualSchemaProposal proto.InternalMessageInfo

func (m *DisableVirtualSchemaProposal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DisableVirtualSchemaProposal) GetVirtualSchemaCode() string {
	if m != nil {
		return m.VirtualSchemaCode
	}
	return ""
}

func (m *DisableVirtualSchemaProposal) GetRegistry() []*DisableVirtualSchemaRegistry {
	if m != nil {
		return m.Registry
	}
	return nil
}

func (m *DisableVirtualSchemaProposal) GetSubmitTime() time.Time {
	if m != nil {
		return m.SubmitTime
	}
	return time.Time{}
}

func (m *DisableVirtualSchemaProposal) GetVotinStartTime() time.Time {
	if m != nil {
		return m.VotinStartTime
	}
	return time.Time{}
}

func (m *DisableVirtualSchemaProposal) GetVotingEndTime() time.Time {
	if m != nil {
		return m.VotingEndTime
	}
	return time.Time{}
}

type VirtualSchemaDisableRequest struct {
	VirtualNftSchemaCode string `protobuf:"bytes,1,opt,name=virtualNftSchemaCode,proto3" json:"virtualNftSchemaCode,omitempty"`
}

func (m *VirtualSchemaDisableRequest) Reset()         { *m = VirtualSchemaDisableRequest{} }
func (m *VirtualSchemaDisableRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualSchemaDisableRequest) ProtoMessage()    {}
func (*VirtualSchemaDisableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{1}
}
func (m *VirtualSchemaDisableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VirtualSchemaDisableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VirtualSchemaDisableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VirtualSchemaDisableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualSchemaDisableRequest.Merge(m, src)
}
func (m *VirtualSchemaDisableRequest) XXX_Size() int {
	return m.Size()
}
func (m *VirtualSchemaDisableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualSchemaDisableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualSchemaDisableRequest proto.InternalMessageInfo

func (m *VirtualSchemaDisableRequest) GetVirtualNftSchemaCode() string {
	if m != nil {
		return m.VirtualNftSchemaCode
	}
	return ""
}

type DisableVirtualSchemaRegistry struct {
	NftSchemaCode string         `protobuf:"bytes,1,opt,name=nftSchemaCode,proto3" json:"nftSchemaCode,omitempty"`
	Status        RegistryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=thesixnetwork.sixprotocol.nftmngr.RegistryStatus" json:"status,omitempty"`
}

func (m *DisableVirtualSchemaRegistry) Reset()         { *m = DisableVirtualSchemaRegistry{} }
func (m *DisableVirtualSchemaRegistry) String() string { return proto.CompactTextString(m) }
func (*DisableVirtualSchemaRegistry) ProtoMessage()    {}
func (*DisableVirtualSchemaRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2dd6dee3450cf76, []int{2}
}
func (m *DisableVirtualSchemaRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisableVirtualSchemaRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisableVirtualSchemaRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisableVirtualSchemaRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisableVirtualSchemaRegistry.Merge(m, src)
}
func (m *DisableVirtualSchemaRegistry) XXX_Size() int {
	return m.Size()
}
func (m *DisableVirtualSchemaRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DisableVirtualSchemaRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DisableVirtualSchemaRegistry proto.InternalMessageInfo

func (m *DisableVirtualSchemaRegistry) GetNftSchemaCode() string {
	if m != nil {
		return m.NftSchemaCode
	}
	return ""
}

func (m *DisableVirtualSchemaRegistry) GetStatus() RegistryStatus {
	if m != nil {
		return m.Status
	}
	return RegistryStatus_PENDING
}

func init() {
	proto.RegisterType((*DisableVirtualSchemaProposal)(nil), "thesixnetwork.sixprotocol.nftmngr.DisableVirtualSchemaProposal")
	proto.RegisterType((*VirtualSchemaDisableRequest)(nil), "thesixnetwork.sixprotocol.nftmngr.VirtualSchemaDisableRequest")
	proto.RegisterType((*DisableVirtualSchemaRegistry)(nil), "thesixnetwork.sixprotocol.nftmngr.DisableVirtualSchemaRegistry")
}

func init() {
	proto.RegisterFile("nftmngr/disable_virtual_schema.proto", fileDescriptor_d2dd6dee3450cf76)
}

var fileDescriptor_d2dd6dee3450cf76 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0x16, 0xaa, 0xe1, 0x69, 0x95, 0xb0, 0x76, 0x88, 0xca, 0x94, 0x96, 0x6a, 0x87,
	0x1e, 0xc0, 0x16, 0x41, 0x9c, 0x91, 0xc6, 0x38, 0x80, 0x10, 0x7f, 0xd2, 0x89, 0x03, 0x1c, 0xa6,
	0xa4, 0x71, 0x5d, 0x8b, 0x24, 0x0e, 0xf6, 0x9b, 0x92, 0x7d, 0x0a, 0x76, 0xe5, 0x1b, 0xed, 0xb8,
	0x23, 0x27, 0x40, 0xed, 0x17, 0x41, 0x38, 0x09, 0x6a, 0x20, 0x12, 0xf4, 0xe6, 0xbc, 0xaf, 0x9f,
	0x5f, 0xde, 0xf7, 0x79, 0x64, 0x7c, 0x9c, 0x2e, 0x20, 0x49, 0x85, 0x66, 0x91, 0x34, 0x41, 0x18,
	0xf3, 0xf3, 0x95, 0xd4, 0x90, 0x07, 0xf1, 0xb9, 0x99, 0x2f, 0x79, 0x12, 0xd0, 0x4c, 0x2b, 0x50,
	0xe4, 0x2e, 0x2c, 0xb9, 0x91, 0x45, 0xca, 0xe1, 0x93, 0xd2, 0x1f, 0xa8, 0x91, 0x85, 0xad, 0xcf,
	0x55, 0x4c, 0x2b, 0xfd, 0xf0, 0xa8, 0x06, 0xb5, 0x01, 0x86, 0x87, 0x42, 0x09, 0x65, 0x8f, 0xec,
	0xd7, 0xa9, 0xaa, 0x8e, 0x84, 0x52, 0x22, 0xe6, 0xcc, 0x7e, 0x85, 0xf9, 0x82, 0x81, 0x4c, 0xb8,
	0x81, 0x20, 0xc9, 0xca, 0x0b, 0x93, 0x2f, 0x3d, 0x7c, 0x74, 0x5a, 0x0e, 0xf6, 0xb6, 0xc4, 0xce,
	0x2c, 0xf5, 0xb5, 0x56, 0x99, 0x32, 0x41, 0x4c, 0x06, 0xb8, 0x2b, 0x23, 0x07, 0x8d, 0xd1, 0xf4,
	0x96, 0xdf, 0x95, 0x11, 0xb9, 0x87, 0x6f, 0xaf, 0xb6, 0x2f, 0x3e, 0x51, 0x11, 0x77, 0xba, 0xb6,
	0xfd, 0x77, 0x83, 0xbc, 0xc7, 0x7b, 0x9a, 0x0b, 0x69, 0x40, 0x5f, 0x38, 0xbd, 0x71, 0x6f, 0xba,
	0xef, 0x3d, 0xa6, 0xff, 0xdc, 0x94, 0xb6, 0x0d, 0xe4, 0x57, 0x18, 0xff, 0x37, 0x90, 0x9c, 0x62,
	0x6c, 0xf2, 0x30, 0x91, 0x70, 0x26, 0x13, 0xee, 0xdc, 0x18, 0xa3, 0xe9, 0xbe, 0x37, 0xa4, 0xe5,
	0xc6, 0xb4, 0xde, 0x98, 0x9e, 0xd5, 0x1b, 0x9f, 0xec, 0x5d, 0x7d, 0x1b, 0x75, 0x2e, 0xbf, 0x8f,
	0x90, 0xbf, 0xa5, 0x23, 0x2f, 0xf0, 0x60, 0xa5, 0x40, 0xa6, 0x33, 0x08, 0x74, 0x49, 0xba, 0xb9,
	0x03, 0xe9, 0x0f, 0x2d, 0x79, 0x8e, 0x0f, 0x6c, 0x45, 0x3c, 0x4d, 0x23, 0x0b, 0xeb, 0xef, 0x00,
	0x6b, 0x4a, 0x27, 0x6f, 0xf0, 0x9d, 0x86, 0x05, 0x95, 0x2d, 0x3e, 0xff, 0x98, 0x73, 0x03, 0xc4,
	0xc3, 0x87, 0x95, 0xe1, 0x2f, 0x17, 0xb0, 0x15, 0x46, 0x99, 0x55, 0x6b, 0x6f, 0xf2, 0x19, 0xb5,
	0xc7, 0x5d, 0xbb, 0x4b, 0x8e, 0xf1, 0x41, 0xda, 0x42, 0x6b, 0x16, 0xc9, 0x33, 0xdc, 0x37, 0x10,
	0x40, 0x6e, 0x6c, 0xf2, 0x03, 0xef, 0xc1, 0x7f, 0x84, 0x5a, 0xff, 0x62, 0x66, 0x85, 0x7e, 0x05,
	0x38, 0x79, 0x75, 0xb5, 0x76, 0xd1, 0xf5, 0xda, 0x45, 0x3f, 0xd6, 0x2e, 0xba, 0xdc, 0xb8, 0x9d,
	0xeb, 0x8d, 0xdb, 0xf9, 0xba, 0x71, 0x3b, 0xef, 0x1e, 0x09, 0x09, 0xcb, 0x3c, 0xa4, 0x73, 0x95,
	0xb0, 0x06, 0x9e, 0x19, 0x59, 0xdc, 0xaf, 0xf9, 0xac, 0x60, 0xf5, 0xbb, 0x80, 0x8b, 0x8c, 0x9b,
	0xb0, 0x6f, 0x3b, 0x0f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x31, 0xa0, 0xd6, 0x78, 0x03,
	0x00, 0x00,
}

func (m *DisableVirtualSchemaProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisableVirtualSchemaProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisableVirtualSchemaProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotingEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotinStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotinStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SubmitTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.Registry) > 0 {
		for iNdEx := len(m.Registry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VirtualSchemaCode) > 0 {
		i -= len(m.VirtualSchemaCode)
		copy(dAtA[i:], m.VirtualSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.VirtualSchemaCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VirtualSchemaDisableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualSchemaDisableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VirtualSchemaDisableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VirtualNftSchemaCode) > 0 {
		i -= len(m.VirtualNftSchemaCode)
		copy(dAtA[i:], m.VirtualNftSchemaCode)
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(len(m.VirtualNftSchemaCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DisableVirtualSchemaRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisableVirtualSchemaRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisableVirtualSchemaRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintDisableVirtualSchema(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
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
func (m *DisableVirtualSchemaProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	l = len(m.VirtualSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	if len(m.Registry) > 0 {
		for _, e := range m.Registry {
			l = e.Size()
			n += 1 + l + sovDisableVirtualSchema(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime)
	n += 1 + l + sovDisableVirtualSchema(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotinStartTime)
	n += 1 + l + sovDisableVirtualSchema(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime)
	n += 1 + l + sovDisableVirtualSchema(uint64(l))
	return n
}

func (m *VirtualSchemaDisableRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VirtualNftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	return n
}

func (m *DisableVirtualSchemaRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NftSchemaCode)
	if l > 0 {
		n += 1 + l + sovDisableVirtualSchema(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovDisableVirtualSchema(uint64(m.Status))
	}
	return n
}

func sovDisableVirtualSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDisableVirtualSchema(x uint64) (n int) {
	return sovDisableVirtualSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DisableVirtualSchemaProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DisableVirtualSchemaProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisableVirtualSchemaProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualSchemaCode", wireType)
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
			m.VirtualSchemaCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registry = append(m.Registry, &DisableVirtualSchemaRegistry{})
			if err := m.Registry[len(m.Registry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SubmitTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotinStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotinStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
				return ErrInvalidLengthDisableVirtualSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDisableVirtualSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotingEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *VirtualSchemaDisableRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VirtualSchemaDisableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualSchemaDisableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualNftSchemaCode", wireType)
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
			m.VirtualNftSchemaCode = string(dAtA[iNdEx:postIndex])
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
func (m *DisableVirtualSchemaRegistry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DisableVirtualSchemaRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisableVirtualSchemaRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDisableVirtualSchema
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
